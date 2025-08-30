package config

import (
	"crypto/rand"
	"encoding/hex"
)
var JWTSecret string

func InitSecret(){
	bytes := make([]byte, 32)
	_,err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	JWTSecret= hex.EncodeToString(bytes)
}