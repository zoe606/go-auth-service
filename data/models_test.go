package data

import (
	"database/sql"
	"time"
)

type PostgresTestRepository struct {
	Conn *sql.Conn
}

func NewPostgresTestRepository(conn *sql.Conn) *PostgresTestRepository {
	return &PostgresTestRepository{Conn: conn}
}

// GetAll returns a slice of all users, sorted by last name
func (u *PostgresRepository) GetAll() ([]*User, error) {
	users := []*User{}
	return users, nil
}

// GetByEmail returns one user by email
func (u *PostgresRepository) GetByEmail(email string) (*User, error) {
	user := User{
		ID:        1,
		FirstName: "First",
		LastName:  "last",
		Email:     "dummy@gmail.com",
		Password:  "somepassword",
		Active:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return &user, nil
}

// GetOne returns one user by id
func (u *PostgresRepository) GetOne(id int) (*User, error) {
	user := User{
		ID:        1,
		FirstName: "First",
		LastName:  "last",
		Email:     "dummy@gmail.com",
		Password:  "somepassword",
		Active:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return &user, nil
}

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *PostgresRepository) Update(user User) error {
	return nil
}

// Delete deletes one user from the database, by User.ID
func (u *PostgresRepository) Delete(id int) error {
	return nil
}

// DeleteByID deletes one user from the database, by ID
func (u *PostgresRepository) DeleteByID(id int) error {
	return nil
}

// Insert inserts a new user into the database, and returns the ID of the newly inserted row
func (u *PostgresRepository) Insert(user User) (int, error) {
	newID := 2
	return newID, nil
}

// ResetPassword is the method we will use to change a user's password.
func (u *PostgresRepository) ResetPassword(password string, user User) error {
	return nil
}

// PasswordMatches uses Go's bcrypt package to compare a user supplied password
// with the hash we have stored for a given user in the database. If the password
// and hash match, we return true; otherwise, we return false.
func (u *PostgresRepository) PasswordMatches(plainText string, user User) (bool, error) {
	return true, nil
}
