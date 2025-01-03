package main

import (
	"fmt"
	"os"

	"github.com/BatmiBoom/BlogAggregator/internal/config"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = cfg.SetUser("lane")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(cfg)
}
