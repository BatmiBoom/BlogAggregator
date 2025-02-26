package command

import (
	"github.com/BatmiBoom/BlogAggregator/internal/config"
	"github.com/BatmiBoom/BlogAggregator/internal/database"
)

type State struct {
	Db  *database.Queries
	Cfg *config.Config
}

type Command struct {
	Name      string
	Arguments []string
}

type Commands struct {
	Cmds map[string]func(*State, Command) error
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Cmds[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {
	err := c.Cmds[cmd.Name](s, cmd)
	if err != nil {
		return err
	}

	return nil
}
