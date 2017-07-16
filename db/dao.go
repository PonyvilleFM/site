package db

// UserDAOer is a common interface for all of the database calls involving
// user accounts (logging in, registration, etc.)
type UserDAOer interface {
	// Register will create a new user account with the given credentials
	// or fail with a given error.
	Register(username, email, password string) (*User, error)

	// Login returns a User instance if the given username and password
	// are the same as the ones in the database.
	Login(username, password string) (ok bool, user *User, err error)

	// Info returns information about a user in the database.
	Info(username string) (*User, error)
}
