package secp256k1

import (
	"testing"

	"github.com/tendermint/tendermint/crypto/secp256k1"
	"github.com/stretchr/testify/require"
)

func TestEcdh(t *testing.T) {
	aPrivKey := secp256k1.GenPrivKeySecp256k1([]byte("alice-haha"))
	aPubKey := aPrivKey.PubKey().(secp256k1.PubKeySecp256k1)
	bPrivKey := secp256k1.GenPrivKeySecp256k1([]byte("bob-hehe"))
	bPubKey := bPrivKey.PubKey().(secp256k1.PubKeySecp256k1)
	secret1, err := Ecdh(bPubKey[:], aPrivKey[:])
	require.Equal(t, nil, err)
	secret2, err := Ecdh(aPubKey[:], bPrivKey[:])
	require.Equal(t, nil, err)
	require.Equal(t, secret1, secret2)
}
