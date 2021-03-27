package helpers

import (
	"fmt"
)

func PgURI(Host, Port, User, Password, DBName string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", User, Password, Host, Port, DBName)
}