package customer

import (
	"github.com/dogmatiq/dogma"
	"github.com/koden-km/dogma-app-setup/messages/commands"
	"github.com/koden-km/dogma-app-setup/messages/events"
)

// customer is an aggregate root for a customer.
type customer struct {
	// Nickname is the current customer nickname.
	Nickname string
}

func (c *customer) ApplyEvent(m dogma.Message) {
	switch x := m.(type) {
	case events.SignupCompleted:
		c.Nickname = x.Nickname
	case events.NicknameChanged:
		c.Nickname = x.NewNickname
	}
}

// Aggregate implements the business logic for a customer.
type Aggregate struct{}

// New returns a new Aggregate instance.
func (Aggregate) New() dogma.AggregateRoot {
	return &customer{}
}

// Configure configures the behavior of the engine as it relates to this
// handler.
func (Aggregate) Configure(c dogma.AggregateConfigurer) {
	c.Name("customer")
	c.RouteCommandType(commands.Signup{})
	c.RouteCommandType(commands.ChangeNickname{})
}

// RouteCommandToInstance returns the ID of the aggregate instance that is
// targetted by m.
func (Aggregate) RouteCommandToInstance(m dogma.Message) string {
	switch x := m.(type) {
	case commands.Signup:
		return x.CustomerID
	case commands.ChangeNickname:
		return x.CustomerID
	default:
		panic(dogma.UnexpectedMessage)
	}
}

// HandleCommand handles a command message that has been routed to this
// handler.
func (Aggregate) HandleCommand(s dogma.AggregateCommandScope, m dogma.Message) {
	switch x := m.(type) {
	case commands.Signup:
		signup(s, x)
	case commands.ChangeNickname:
		changeNickname(s, x)
	default:
		panic(dogma.UnexpectedMessage)
	}
}

func signup(s dogma.AggregateCommandScope, m commands.Signup) {
	if !s.Create() {
		s.Log("customer has already signed up")
		return
	}

	s.RecordEvent(events.SignupCompleted{
		CustomerID: m.CustomerID,
		Name:       m.Name,
		Nickname:   m.Nickname,
	})
}

func changeNickname(s dogma.AggregateCommandScope, m commands.ChangeNickname) {
	c := s.Root().(*customer)

	if c.Nickname != m.NewNickname {
		s.RecordEvent(events.NicknameChanged{
			CustomerID:  m.CustomerID,
			NewNickname: m.NewNickname,
			OldNickname: c.Nickname,
		})
	}
}
