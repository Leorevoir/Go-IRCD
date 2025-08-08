package config

import "time"

const (
	IRCD_DefaultPort         = 4242
	IRCD_DefaultName         = "IRCD-GOLANG-SERVER"
	IRCD_DefaultMaxClients   = 1000
	IRCD_DefaultPingInterval = 90 * time.Second
	IRCD_DefaultPingTimeout  = 30 * time.Second
)

type IRCDPing struct {
	Interval time.Duration
	Timeout  time.Duration
}

type IRCDConfig struct {
	Port       uint16
	ServerName string
	MaxClients uint
	Ping       IRCDPing
	MOTD       []string
}

func DefaultConfig() *IRCDConfig {
	return &IRCDConfig{
		Port:       IRCD_DefaultPort,
		ServerName: IRCD_DefaultName,
		MaxClients: IRCD_DefaultMaxClients,
		Ping: IRCDPing{
			Interval: IRCD_DefaultPingInterval,
			Timeout:  IRCD_DefaultPingTimeout,
		},
		MOTD: []string{"Welcome to IRCD-GOLANG-SERVER!", "This is a sample message of the day."},
	}
}
