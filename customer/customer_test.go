package customer_test

import (
	"testing"

	. "github.com/dogmatiq/dogmatest/assert"
	"github.com/koden-km/dogma-app-setup/internal/testrunner"
	"github.com/koden-km/dogma-app-setup/messages/commands"
	"github.com/koden-km/dogma-app-setup/messages/events"
)

func TestCustomer_Signup(t *testing.T) {
	t.Run(
		"it signs up the customer",
		func(t *testing.T) {
			testrunner.Runner.
				Begin(t).
				ExecuteCommand(
					commands.Signup{
						CustomerID: "C007",
						Name:       "Dude Guy",
						Nickname:   "dudeguy",
					},
					EventRecorded(
						events.SignupCompleted{
							CustomerID: "C007",
							Name:       "Dude Guy",
							Nickname:   "dudeguy",
						},
					),
				)
		},
	)

	t.Run(
		"it does not signup a customer that is already signed up",
		func(t *testing.T) {
			cmd := commands.Signup{
				CustomerID: "C007",
				Name:       "Dude Guy",
				Nickname:   "dudeguy",
			}

			testrunner.Runner.
				Begin(t).
				Prepare(cmd).
				ExecuteCommand(
					cmd,
					NoneOf(
						EventTypeRecorded(events.SignupCompleted{}),
					),
				)
		},
	)
}

func TestCustomer_ChangeNickname(t *testing.T) {
	t.Run(
		"it changes the nickname of the customer",
		func(t *testing.T) {
			testrunner.Runner.
				Begin(t).
				Prepare(commands.Signup{
					CustomerID: "C007",
					Name:       "Dude Guy",
					Nickname:   "dudeguy",
				}).
				ExecuteCommand(
					commands.ChangeNickname{
						CustomerID:  "C007",
						NewNickname: "lucky7",
					},
					EventRecorded(
						events.NicknameChanged{
							CustomerID:  "C007",
							NewNickname: "lucky7",
							OldNickname: "dudeguy",
						},
					),
				)
		},
	)

	t.Run(
		"it does not change nickname if it is the same",
		func(t *testing.T) {
			testrunner.Runner.
				Begin(t).
				Prepare(commands.Signup{
					CustomerID: "C007",
					Name:       "Dude Guy",
					Nickname:   "dudeguy",
				}).
				ExecuteCommand(
					commands.ChangeNickname{
						CustomerID:  "C007",
						NewNickname: "dudeguy",
					},
					NoneOf(
						EventTypeRecorded(events.NicknameChanged{}),
					),
				)
		},
	)
}
