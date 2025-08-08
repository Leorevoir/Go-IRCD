package server

import (
	"encoding/json"
	"github.com/Leorevoir/Go-IRCD/pkg/error"
	"os"
)

const (
	DefaultPingFrequency  = 30
	DefaultPingMaxLatency = 5
	DefaultSSLPort        = 6697
	DefaultPort           = 4242
)

type SSLCertificate struct {
	KeyFile  string
	CertFile string
}

type ConfigPing struct {
	Frequency  int
	MaxLatency int
}

type Config struct {
	Name           string
	Network        string
	Port           uint16
	SSLPort        uint16
	MOTD           string
	Ping           ConfigPing
	SSLCertificate SSLCertificate
}

func FromJSON(path string) Config {
	var config Config

	file, err := os.Open(path)
	error.Assert(err, "Failed to open config file")
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	error.Assert(err, "Failed to decode config file")

	return setDefaultConfig(config)
}

func setDefaultConfig(config Config) Config {
	if config.Ping.Frequency == 0 {
		config.Ping.Frequency = DefaultPingFrequency
	}
	if config.Ping.MaxLatency == 0 {
		config.Ping.MaxLatency = DefaultPingMaxLatency
	}
	if config.Port == 0 {
		config.Port = DefaultPort
	}
	if config.SSLPort == 0 {
		config.SSLPort = DefaultSSLPort
	}
	return config
}
