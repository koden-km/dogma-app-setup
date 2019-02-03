package main

import (
	"context"
	"fmt"

	"github.com/dogmatiq/dogma"
	dogmatiqapp "github.com/koden-km/dogma-app-setup"
	"github.com/koden-km/dogma-app-setup/messages/commands"

	"github.com/dogmatiq/dogmatest/engine"
	"github.com/dogmatiq/enginekit/config"
)

func main() {
	fmt.Println("Dogmatiq app...")

	app := &dogmatiqapp.App{}

	cfg, err := config.NewApplicationConfig(app)
	if err != nil {
		panic(err)
	}

	en, err := engine.New(cfg)
	if err != nil {
		panic(err)
	}

	messages := []dogma.Message{
		commands.Signup{
			CustomerID: "cust001",
			Name:       "Not Seven",
			Nickname:   "dude",
		},
		commands.ChangeNickname{
			CustomerID:  "cust001",
			NewNickname: "dudeguy",
		},
	}

	for _, m := range messages {
		err := en.Dispatch(
			context.Background(),
			m,
			// engine.EnableProjections(false),
		)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Done.")
}
