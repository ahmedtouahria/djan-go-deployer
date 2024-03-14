package deployer

import (
	"database/sql"
	"fmt"
	"os"
)

// Database represents a database configuration.
type Database struct {
	name     string
	password string
	username string
}

// DatabaseBuilder is a builder for constructing a Database.
type DatabaseBuilder struct {
	database Database
}

// NewDatabaseBuilder creates a new instance of DatabaseBuilder.
func NewDatabaseBuilder() *DatabaseBuilder {
	return &DatabaseBuilder{}
}

// WithName sets the name for the Database.
func (b *DatabaseBuilder) WithName(name string) *DatabaseBuilder {
	b.database.name = name
	return b
}

// WithPassword sets the password for the Database.
func (b *DatabaseBuilder) WithPassword(password string) *DatabaseBuilder {
	b.database.password = password
	return b
}

// WithUsername sets the username for the Database.
func (b *DatabaseBuilder) WithUsername(username string) *DatabaseBuilder {
	b.database.username = username
	return b
}

// Build constructs and returns the final Database.
func (b *DatabaseBuilder) Build() Database {
	return b.database
}

const (
	Host     = "localhost"
	Port     = 5432
	User     = "your_username"
	Password = "your_password"
	Dbname   = "your_database"
)

func createDatabase() error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", Host, Port, User, Password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	// Connect to the default PostgreSQL database to create a new database
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", Dbname))
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(os.Stdout, "Database '%s' created successfully.\n", []any{Dbname}...)
	if err != nil {
		return err
	}
	return nil
}
