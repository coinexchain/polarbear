package cip0013

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"

	"github.com/coinexchain/polarbear/secp256k1"
)

/*
https://en.bitcoin.it/wiki/Protocol_documentation#Variable_length_integer
< 0xFD		1	uint8_t
<= 0xFFFF	3	0xFD followed by the length as uint16_t
<= 0xFFFF FFFF	5	0xFE followed by the length as uint32_t
-		9	0xFF followed by the length as uint64_t
*/

func int2bytes(i int) []byte {
	var buf [9]byte
	if i < 0xFD {
		return []byte{byte(i)}
	} else if i < 0xFFFF {
		buf[0] = 0xFD
		binary.LittleEndian.PutUint16(buf[1:3], uint16(i))
		return buf[:3]
	} else if i < 0xFFFFFFFF {
		buf[0] = 0xFE
		binary.LittleEndian.PutUint32(buf[1:5], uint32(i))
		return buf[:5]
	} else {
		buf[0] = 0xFF
		binary.LittleEndian.PutUint32(buf[1:9], uint64(i))
		return buf[:]
	}
}


func SignOffChainMessage(msg []byte, seckey []byte) (string, error) {
	prefix := "\u0018CoinEx Token Signed Message:\n"
	lenBytes := int2bytes(len(msg))
	hashIn := append([]byte(prefix), lenBytes...)
	hashIn := append(hashIn, msg...)
	hash := sha256.Sum256(hashIn)

	resBytes, err := secp256k1.Sign(hash[:], seckey)
	if err != nil {
		return "", err
	}

	var buf [65]byte
	buf[0] = resBytes[64]+31
	copy(buf[1:], resBytes[:64])
	return base64.StdEncoding.EncodeToString(buf[:]), nil
}



