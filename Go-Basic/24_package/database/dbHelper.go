package database

var connection string

// package initalization
// will tell go that if the package is used, it will run this function first
func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
