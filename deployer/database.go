package deployer

// Database represents a database configuration.
type Database struct{
	name string
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

