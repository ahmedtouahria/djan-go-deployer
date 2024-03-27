package deployer

import (
	"fmt"
	"os"
	"os/exec"
)

// Database represents a database configuration.
type Database struct {
	Name     string
	Password string
	Username string
	Host     string
	Port     string
}

// Error implements error.
func (Database) Error() string {
	panic("unimplemented")
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
	b.database.Name = name
	return b
}

// WithPassword sets the password for the Database.
func (b *DatabaseBuilder) WithPassword(password string) *DatabaseBuilder {
	b.database.Password = password
	return b
}

// WithUsername sets the username for the Database.
func (b *DatabaseBuilder) WithUsername(username string) *DatabaseBuilder {
	b.database.Username = username
	return b
}

// WithHost sets the host for the Database.
func (b *DatabaseBuilder) WithHost(host string) *DatabaseBuilder {
	b.database.Host = host
	return b
}

// WithPort sets the port for the Database.
func (b *DatabaseBuilder) WithPort(port string) *DatabaseBuilder {
	b.database.Port = port
	return b
}
// Build constructs and returns the final Database.
func (b *DatabaseBuilder) Build() error {
	// Get or create the database
	err := CreatePostgresScript(b.database.Name, b.database.Username, b.database.Password)
	if err != nil {
		fmt.Printf("Error getting or creating database: %v\n", err)
		return err
	}
	return nil
}


func CreatePostgresScript(dbName, dbUser, dbPassword string) error {
	// Generate the bash script content
	scriptContent := fmt.Sprintf(`#!/bin/bash

DB_NAME="%s"
DB_USER="%s"
DB_PASSWORD="%s"

# Run psql commands to create the database and user
sudo -u postgres psql -c "CREATE DATABASE $DB_NAME;"
sudo -u postgres psql -c "CREATE USER $DB_USER WITH PASSWORD '$DB_PASSWORD';"
sudo -u postgres psql -c "ALTER ROLE $DB_USER SET client_encoding TO 'utf8';"
sudo -u postgres psql -c "ALTER ROLE $DB_USER SET default_transaction_isolation TO 'read committed';"
sudo -u postgres psql -c "ALTER ROLE $DB_USER SET timezone TO 'UTC';"
sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER;"

echo "Database '$DB_NAME' and user '$DB_USER' created successfully."
`, dbName, dbUser, dbPassword)

	// Write the script content to a temporary file
	tmpfile, err := os.CreateTemp("", "create_postgres_script_*.sh")
	if err != nil {
		return err
	}
	defer os.Remove(tmpfile.Name()) // Clean up the temporary file
	if _, err := tmpfile.WriteString(scriptContent); err != nil {
		return err
	}

	// Close the temporary file
	if err := tmpfile.Close(); err != nil {
		return err
	}

	// Make the temporary file executable
	if err := os.Chmod(tmpfile.Name(), 0755); err != nil {
		return err
	}

	// Execute the bash script
	cmd := exec.Command("bash", tmpfile.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
return nil
}
