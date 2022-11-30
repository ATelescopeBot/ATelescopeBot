package main

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Envvars struct {
	Token string
}

func (env *Envvars) Init() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	env.Token = os.Getenv("TOKEN")

	if env.Token == "" {
		return errors.New("TOKEN environment variable can't be empty. Edit .env file.")
	}

	return nil
}
