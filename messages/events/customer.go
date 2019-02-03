package events

// SignupCompleted is an event indicating that a new customer signup has been
// completed.
type SignupCompleted struct {
	CustomerID string
	Name       string
	Nickname   string
}

// NicknameChanged is an event indicating that an existing customer has had a
// nickname change.
type NicknameChanged struct {
	CustomerID  string
	NewNickname string
	OldNickname string
}
