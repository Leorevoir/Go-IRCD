package config

import "time"

const (
	DefaultPort         = 4242
	DefaultName         = "ircd.example.com"
	DefaultMaxClients   = 1000
	DefaultPingInterval = 90 * time.Second
	DefaultPingTimeout  = 30 * time.Second
)

type Ping struct {
	Interval time.Duration
	Timeout  time.Duration
}

type Config struct {
	Port       uint16
	ServerName string
	MaxClients uint
	Ping       Ping
	MOTD       []string
}

func DefaultConfig() *Config {
	return &Config{
		Port:       DefaultPort,
		ServerName: DefaultName,
		MaxClients: DefaultMaxClients,
		Ping: Ping{
			Interval: DefaultPingInterval,
			Timeout:  DefaultPingTimeout,
		},
		MOTD: []string{"Welcome to -GOLANG-SERVER!", "This is a sample message of the day."},
	}
}
