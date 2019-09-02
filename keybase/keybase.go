package keybase

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/crypto/keys/hd"
	"github.com/cosmos/go-bip39"
)

const (
	mnemonicEntropySize = 256
	defaultCoinType     = 688
)

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

var _ KeyBase = DefaultKeyBase(nil)

type DefaultKeyBase struct {
	kb keys.Keybase
}

func NewDefaultKeyBase(root string) DefaultKeyBase {
	return DefaultKeyBase{
		keys.New("db", root),
	}
}
func (k DefaultKeyBase) CreateKey(name, password, bip39Passphrase string, account, index uint32) string {
	entropySeed, err := bip39.NewEntropy(mnemonicEntropySize)
	if err != nil {
		return err.Error()
	}

	mnemonic, err := bip39.NewMnemonic(entropySeed[:])
	if err != nil {
		return err.Error()
	}
	hdPath := hd.NewFundraiserParams(account, defaultCoinType, index)
	info, err := k.kb.Derive(name, mnemonic, bip39Passphrase, password, *hdPath)
	if err != nil {
		return err.Error()
	}
	fmt.Println(info.GetAddress().String())
	return ""
}

func (k DefaultKeyBase) DeleteKey(name, password string) string {
	return k.kb.Delete(name, password, false).Error()
}

func (k DefaultKeyBase) RecoverKey(name, mnemonic, password, bip39Passphrase string, account, index uint32) string {
	info, err := k.kb.CreateAccount(name, mnemonic, bip39Passphrase, password, account, index)
	if err != nil {
		return err.Error()
	}
	fmt.Println(info.GetAddress().String())
	return ""
}

func (k DefaultKeyBase) AddKey(name, armor string) string {
	return k.kb.Import(name, armor).Error()
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
		return ""
	}
	//todo: make a json string show infos
	return ""
}

func (k DefaultKeyBase) ResetPassword(name, password, newPassword string) string {
	f := func() (string, error) { return newPassword, nil }
	return k.kb.Update(name, password, f).Error()
}

func (k DefaultKeyBase) Sign(name, password, tx string) string {
	sig, _, err := k.kb.Sign(name, password, []byte(tx))
	if err != nil {
		return ""
	}
	return string(sig)
}
