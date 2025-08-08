package config

import "time"

/*
* default configuration values for testing purposes
 */
const (
	DefaultPort         = 4242
	DefaultName         = "ircd.example.com"
	DefaultMaxClients   = 1000
	DefaultPingInterval = 90 * time.Second
	DefaultPingTimeout  = 30 * time.Second
)

/*
* ping configuration
 */
type Ping struct {
	Interval time.Duration
	Timeout  time.Duration
}

/*
* server configuration
 */
type Config struct {
	Port       uint16
	ServerName string
	MaxClients uint
	Ping       Ping
	MOTD       []string
}

/*
* default configuration values for testing purposes
 */
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
