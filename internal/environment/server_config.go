package environment

import (
	"flag"
)

const (
	serverDefaultHostAdressURL = "localhost:8080"

	serverEnvHostAdressURL = "ADDRESS"
)

type ServerConfig struct {
	HostAdress string
}

func NewServerConfig() ServerConfig {

	hostAdressFlag := flag.String("a", serverDefaultHostAdressURL, "HOST_ADRESS")

	flag.Parse()

	sc := ServerConfig{
		HostAdress: GetEnvString(serverEnvHostAdressURL, *hostAdressFlag),
	}

	return sc
}
