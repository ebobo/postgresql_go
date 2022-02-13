package main

import (
	"fmt"

	"github.com/ebobo/postgresql_go/pkg/store"
)

func main() {

	fmt.Println("Hello, Qi")

	//runSql()

	store.RunSqlx()
}
