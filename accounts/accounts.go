// Package accounts handles logging in and registering users.
package accounts

type Account struct {
	Username string
	Password string
}

func (a *Account) Login() bool {
	// TODO: calls backend
	return true
}
