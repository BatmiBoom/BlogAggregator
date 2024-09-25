package main

import (
	"fmt"
	"os"

	"github.com/BatmiBoom/BlogAggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(cfg)

	err = cfg.SetUser("Batmiboom")
	if err != nil {
		fmt.Println(err)
	}
	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("ERROR : %v", err)
		os.Exit(1)
	}

	fmt.Println(cfg)
}
