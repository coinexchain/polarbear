package keybase

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/crypto/keys/hd"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
)

const (
	mnemonicEntropySize = 256
	defaultCoinType     = 688
)

//todo: add GetAddress with name
type KeyBase interface {
	CreateKey(name, password, bip39Passphrase string, account, index uint32) string
	DeleteKey(name, password string) string
	RecoverKey(name, mnemonic, password, bip39Passphrase string, account, index uint32) string
	AddKey(name, armor string) string
	ExportKey(name string) string
	ListKeys() string
	ResetPassword(name, password, newPassword string) string
	Sign(name, password, tx string) string
}

var _ KeyBase = DefaultKeyBase{}

type DefaultKeyBase struct {
	kb keys.Keybase
}

func NewDefaultKeyBase(root string) DefaultKeyBase {
	initDefaultKeyBaseConfig()
	return DefaultKeyBase{
		keys.New("keys", root),
	}
}

//todo: name repetition check
func (k DefaultKeyBase) CreateKey(name, password, bip39Passphrase string, account, index uint32) string {
	entropySeed, err := bip39.NewEntropy(mnemonicEntropySize)
	if err != nil {
		return ""
	}

	mnemonic, err := bip39.NewMnemonic(entropySeed[:])
	if err != nil {
		return ""
	}
	hdPath := hd.NewFundraiserParams(account, defaultCoinType, index)
	info, err := k.kb.Derive(name, mnemonic, bip39Passphrase, password, *hdPath)
	if err != nil {
		return ""
	}
	fmt.Println(info.GetAddress().String())
	return info.GetAddress().String() + "+" + mnemonic
}

func (k DefaultKeyBase) DeleteKey(name, password string) string {
	if err := k.kb.Delete(name, password, false); err != nil {
		return err.Error()
	}
	return ""
}

func (k DefaultKeyBase) RecoverKey(name, mnemonic, password, bip39Passphrase string, account, index uint32) string {
	info, err := k.kb.CreateAccount(name, mnemonic, bip39Passphrase, password, account, index)
	if err != nil {
		return ""
	}
	fmt.Println(info.GetAddress().String())
	return info.GetAddress().String()
}

func (k DefaultKeyBase) AddKey(name, armor string) string {
	if err := k.kb.Import(name, armor); err != nil {
		return err.Error()
	}
	return ""
}

func (k DefaultKeyBase) ExportKey(name string) string {
	armor, err := k.kb.Export(name)
	if err != nil {
		return ""
	}
	return armor
}

func (k DefaultKeyBase) ListKeys() string {
	_, err := k.kb.List()
	if err != nil {
		return err.Error()
	}
	//todo: make a json string show infos
	return ""
}

func (k DefaultKeyBase) ResetPassword(name, password, newPassword string) string {
	f := func() (string, error) { return newPassword, nil }
	if err := k.kb.Update(name, password, f); err != nil {
		return err.Error()
	}
	return ""
}

func (k DefaultKeyBase) Sign(name, password, tx string) string {
	sig, _, err := k.kb.Sign(name, password, []byte(tx))
	if err != nil {
		return ""
	}
	return string(sig)
}

func initDefaultKeyBaseConfig() {
	bench32MainPrefix := "coinex"
	bench32PrefixAccAddr := bench32MainPrefix
	// bench32PrefixAccPub defines the bench32 prefix of an account's public key
	bench32PrefixAccPub := bench32MainPrefix + types.PrefixPublic
	// bench32PrefixValAddr defines the bench32 prefix of a validator's operator address
	bench32PrefixValAddr := bench32MainPrefix + types.PrefixValidator + types.PrefixOperator
	// bench32PrefixValPub defines the bench32 prefix of a validator's operator public key
	bench32PrefixValPub := bench32MainPrefix + types.PrefixValidator + types.PrefixOperator + types.PrefixPublic
	// bench32PrefixConsAddr defines the bench32 prefix of a consensus node address
	bench32PrefixConsAddr := bench32MainPrefix + types.PrefixValidator + types.PrefixConsensus
	// bench32PrefixConsPub defines the bench32 prefix of a consensus node public key
	bench32PrefixConsPub := bench32MainPrefix + types.PrefixValidator + types.PrefixConsensus + types.PrefixPublic

	config := types.GetConfig()
	config.SetCoinType(defaultCoinType)
	config.SetBech32PrefixForAccount(bench32PrefixAccAddr, bench32PrefixAccPub)
	config.SetBech32PrefixForValidator(bench32PrefixValAddr, bench32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(bench32PrefixConsAddr, bench32PrefixConsPub)
	config.Seal()
}
