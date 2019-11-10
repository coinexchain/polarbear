package keybase

import (
	"fmt"
	"github.com/cosmos/go-bip39"
	"github.com/magiconair/properties/assert"
	assert2 "github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var testKeyBase = NewDefaultKeyBase("./tmp")

func getRandAddress() (string, string, error) {
	entropy, err := bip39.NewEntropy(mnemonicEntropySize)
	if err != nil {
		return "", "", err
	}
	return GetAddressFromEntropy(entropy)
}

func TestDefaultKeyBase(t *testing.T) {

	name := "default"
	password := "12345678"
	newPassword := "11223344"
	bip39Passphrase := "11111111"
	chainId := "coinexdex-test1"
	var address string
	var mnemonic string
	var account uint32 = 0
	var index uint32 = 0

	defer os.RemoveAll("./tmp")
	mnemonic = testKeyBase.CreateKey(name, password, bip39Passphrase, account, index)
	info := strings.Split(mnemonic, "+")
	assert.Equal(t, len(info), 2)
	mnemonic = info[1]
	address = info[0]
	fmt.Println(address)
	assert2.NotNil(t, mnemonic)
	t.Log("create key pass")

	mnemonic2 := testKeyBase.CreateKey(name, password, bip39Passphrase, account, index)
	assert.Equal(t, mnemonic2, errPrefix + errCreate + "key with same name is already exist")
	t.Log("repeat create key pass")

	res := testKeyBase.DeleteKey(name, password)
	assert.Equal(t, res, "")
	t.Log("delete key pass")

	res = testKeyBase.RecoverKey(name, mnemonic, password, bip39Passphrase, account, index)
	assert.Equal(t, res, address)
	t.Log("recover key pass")

	keys := testKeyBase.ListKeys()
	assert2.Contains(t, keys, name)
	t.Log("list keys pass")

	res = testKeyBase.ResetPassword(name, password, newPassword)
	assert.Equal(t, res, "")
	t.Log("reset password pass")

	armor := testKeyBase.ExportKey(name, newPassword, newPassword)
	assert2.NotEqual(t, "", armor)
	//fmt.Println(armor)
	t.Log("export keys pass")

	res = testKeyBase.AddKey(name, armor, newPassword)
	assert.Equal(t, res, errPrefix+errAdd+"Cannot overwrite key "+name)
	res = testKeyBase.DeleteKey(name, newPassword)
	assert.Equal(t, res, "")
	res = testKeyBase.AddKey(name, armor, newPassword)
	assert.Equal(t, res, "")
	addr := testKeyBase.GetAddress(name)
	assert.Equal(t, addr, address)
	t.Log("add key pass")
	t.Log("get address pass")

	unsignedFmtStr := "{\"account_number\":\"0\"," +
		"\"chain_id\":\"coinexdex-test1\"," +
		"\"fee\":" +
		"{\"amount\":[{\"amount\":\"200000\",\"denom\":\"cet\"}]," +
		"\"gas\":\"6000\"}," +
		"\"memo\":\"\"," +
		"\"msgs\":[{" +
		"\"type\":\"bankx/MsgSend\"," +
		"\"value\":{" +
		"\"amount\":[{\"amount\":\"1000000\",\"denom\":\"cet\"}]," +
		"\"from_address\":\"%s\"," +
		"\"to_address\":\"coinex1rd3tgkzd8q8akaug53hnqwhr378xfeljchmzls\"," +
		"\"unlock_time\":\"0\"}}]," +
		"\"sequence\":\"2\"}"
	unsignedStr := fmt.Sprintf(unsignedFmtStr, address)
	signer := testKeyBase.GetSigner(unsignedStr)
	assert.Equal(t, signer, name)
	t.Log("getSigner pass")

	res = testKeyBase.Sign(name, newPassword, unsignedStr)
	assert2.NotEqual(t, "", res)
	t.Log("sign pass")

	stdTxFmtStr := "{\"type\":\"cosmos-sdk/StdTx\"," +
		"\"value\":{" +
		"\"msg\":[{\"type\":\"bankx/MsgSetMemoRequired\",\"value\":{\"address\":\"%s\",\"required\":true}}]," +
		"\"fee\":{\"amount\":[{\"denom\":\"cet\",\"amount\":\"2000000000\"}],\"gas\":\"6000000\"}," +
		"\"signatures\":null," +
		"\"memo\":\"Sent with example memo\"}}"
	address = "coinex1222kjfxeprsgaywfsuewlswrwmx3724etsepzc"
	stdTx := fmt.Sprintf(stdTxFmtStr, address)
	fmt.Println(stdTx)
	res = testKeyBase.SignStdTx(name, newPassword, stdTx,
		chainId, 0, 1)
	assert2.NotEqual(t, "", res)
	//fmt.Println(res)

	res = testKeyBase.SignAndBuildBroadcast(name, newPassword, stdTx, chainId, "sync", 0, 1)
	assert2.NotEqual(t, "", res)
	//fmt.Println(res)

	for i := 0; i < 10; i++ {
		addr0, mnemonic, err := getRandAddress()
		assert.Equal(t, nil, err)
		name = fmt.Sprintf("user%d", i)
		addr := testKeyBase.RecoverKey(name, mnemonic, "password", "", 0, 0)
		assert.Equal(t, addr0, addr)
	}
}
