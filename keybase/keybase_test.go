package keybase

import (
	"github.com/magiconair/properties/assert"
	assert2 "github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var testKeyBase = NewDefaultKeyBase("./tmp")

func TestDefaultKeyBase(t *testing.T) {

	name := "default"
	password := "12345678"
	newPassword := "11223344"
	bip39Passphrase := "11111111"
	var mnemonic string
	var account uint32 = 0
	var index uint32 = 0

	mnemonic = testKeyBase.CreateKey(name, password, bip39Passphrase, account, index)
	info := strings.Split(mnemonic, "+")
	assert.Equal(t, len(info), 2)
	mnemonic = info[1]
	assert2.NotNil(t, mnemonic)

	res := testKeyBase.DeleteKey(name, password)
	assert.Equal(t, res, "")

	res = testKeyBase.RecoverKey(name, mnemonic, password, bip39Passphrase, account, index)
	assert.Equal(t, res, info[0])

	//keys := testKeyBase.ListKeys()

	res = testKeyBase.ResetPassword(name, password, newPassword)
	assert.Equal(t, res, "")

	mnemonic = "enlist shoe journey effort unfair scout layer affair arrow twice happy ready horn buyer loan deposit merge fancy panda gospel pole type essence side"
	addr := testKeyBase.RecoverKey("alice", mnemonic, "password", "", 0, 0)
	assert.Equal(t, "coinex1000ujfjr5tj4nac33mr7t76y2zvmzdmmpwfnx7", addr)

	_ = os.RemoveAll("./tmp")
}
