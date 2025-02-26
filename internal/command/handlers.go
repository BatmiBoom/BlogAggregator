package command

import "fmt"

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Arguments) < 1 || len(cmd.Arguments) > 3 {
		return fmt.Errorf("ERROR: login expects one argument - USERNAME")
	}

	if cmd.Name != "login" {
		return fmt.Errorf(
			"ERROR: the handler in use 'handlerLogin' is not the right for this command : %s",
			cmd.Name,
		)
	}

	err := s.Cfg.SetUser(cmd.Arguments[1])
	if err != nil {
		return err
	}

	fmt.Println("The login was succesful, user set correctly")

	return nil
}
