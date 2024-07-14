package database

var connection string

// otomatis dieksekusi
func init() {
	connection = "MYSQL"
}

func GetDatabase() string {
	return connection
}
