package client

//import "go-telnet"

func connect(provider string) {
	switch provider {
	case "telnet":
		//var caller telnet.Caller = telnet.StandardCaller
		//telnet.DialToAndCallTLS("localhost:5555", caller, nil)
	case "http":

	default:
		panic("unknown provider: " + provider)
	}
}

