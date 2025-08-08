package main

import (
	"fmt"
	"github.com/Leorevoir/Go-IRCD/internal/config"
)

func main() {
	c := config.DefaultConfig()

	fmt.Printf("Server name: %s\n", c.ServerName)
	fmt.Printf("Port: %d\n", c.Port)
}
