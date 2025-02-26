package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/BatmiBoom/BlogAggregator/internal/command"
	"github.com/BatmiBoom/BlogAggregator/internal/config"
	"github.com/BatmiBoom/BlogAggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Not enough arguments were provided")
		os.Exit(1)
	}

	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.Db_url)

	dbQueries := database.New(db)

	state := command.State{
		Db:  dbQueries,
		Cfg: &cfg,
	}

	login_cmd := command.Command{
		Name:      "login",
		Arguments: args[1:],
	}

	cmds := command.Commands{
		Cmds: map[string]func(*command.State, command.Command) error{},
	}

	cmds.Register("login", command.HandlerLogin)
	err = cmds.Run(&state, login_cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
