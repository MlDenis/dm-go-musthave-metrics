package environment

import (
	"flag"
)

const (
	ServerDefaultHostAdressURL = "localhost:8080"

	ServerEnvHostAdressURL = "ADDRESS"
)

type ServerConfig struct {
	HostAdress string
}

func NewServerConfig() ServerConfig {

	hostAdressFlag := flag.String("a", ServerDefaultHostAdressURL, "HOST_ADRESS")

	flag.Parse()

	sc := ServerConfig{
		HostAdress: GetEnvString(ServerEnvHostAdressURL, *hostAdressFlag),
	}

	return sc
}
