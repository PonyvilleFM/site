package mockdb

import (
	"crypto/md5"
	"fmt"
	"sync"

	"github.com/PonyvilleFM/site/db"
)

// UserDao stores and retrieves information on users.
type UserDao struct {
	sync.Mutex
	users map[string]*db.User
}

// NewUserDao is a convenience constructor.
func NewUserDao() *UserDao {
	return &UserDao{
		users: map[string]*db.User{},
	}
}

// Register will create a new user account with the given information
// and credentials or fail with an error.
func (ud *UserDao) Register(info db.User, password string) (*db.User, error) {
	ud.Lock()
	defer ud.Unlock()

	u := &info
	u.PasswordHash = hash(password)

	ud.users[info.Username] = u

	return u, nil
}

// Login authenticates a user with the given credentials or fails with an
// error.
func (ud *UserDao) Login(username, password string) (bool, *db.User, error) {
	ud.Lock()
	defer ud.Unlock()

	u, ok := ud.users[username]
	if !ok {
		return false, nil, db.ErrNoSuchUser
	}

	if hash(password) == u.PasswordHash {
		return true, u, nil
	}

	return false, nil, nil
}

// Info returns information about a user in the database.
func (ud *UserDao) Info(username string) (*db.User, error) {
	ud.Lock()
	defer ud.Unlock()

	u, ok := ud.users[username]
	if !ok {
		return nil, db.ErrNoSuchUser
	}

	return u, nil
}

// hash is a simple wrapper around the MD5 algorithm implementation in the
// Go standard library. It takes in data and a salt and returns the hashed
// representation.
func hash(data string) string {
	output := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", output)
}
