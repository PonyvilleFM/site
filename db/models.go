// Package db contains models and DAO interfaces that packages
// db/postgresdb and db/mockdb will implement.
//
// This is intended to be the primary way that the underlying
// application stores and fetches data.
//
// DAO: https://en.wikipedia.org/wiki/Data_access_object
package db

// User is a logical "user" account created for administrators or DJ's to
// log in and make API calls.
type User struct {
	Username      string `db:"username"`
	Email         string `db:"email"`
	PasswordHash  string `db:"password_hash"`
	IsAdmin       bool   `db:"is_admin"`
	IsDJ          bool   `db:"is_dj"`
	TwitterHandle string `db:"twitter_handle"`
}
