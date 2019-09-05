package wallet

import (
	 "github.com/coinexchain/polarbear/keybase"
)

var Api Wallet


type Wallet struct {
	keybase.KeyBase
}

func Init(root string) {
	Api.KeyBase = keybase.NewDefaultKeyBase(root)
}

func CreateKey(name, password, bip39Passphrase string, account, index uint32) string {
	return Api.CreateKey(name, password, bip39Passphrase, account, index)
}

func DeleteKey(name, password string) string {
	return Api.DeleteKey(name, password)
}

func RecoverKey(name, mnemonic, password, bip39Passphrase string, account, index uint32) string {
	return Api.RecoverKey(name, mnemonic, password, bip39Passphrase, account, index)
}

func AddKey(name, armor string) string {
	return Api.AddKey(name, armor)
}

func ExportKey(name string) string {
	return Api.ExportKey(name)
}

func ListKeys() string {
	return Api.ListKeys()
}

func ResetPassword(name, password, newPassword string) string {
	return Api.ResetPassword(name, password, newPassword)
}

func Sign(name, password, tx string) string {
	return Api.Sign(name, password, tx)
}
