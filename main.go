package main

import (
	"fmt"
	mysql "mysql-replica/adapter"
)

func main() {
	db := mysql.NewDatabase()

	fmt.Println(db.ReadDB)
	fmt.Println(db.WriteDB)
}
