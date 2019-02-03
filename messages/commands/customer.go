package commands

// Signup is a command requesting a new customer signup be performed.
type Signup struct {
	CustomerID string
	Name       string
	Nickname   string
}

// ChangeNickname is a command requesting an existing customer's nickname be
// changed.
type ChangeNickname struct {
	CustomerID  string
	NewNickname string
}
