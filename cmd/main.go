package main

import (
	"github.com/ebobo/postgresql_go/pkg/store"
	"github.com/ebobo/utilities_go/pkg/greeting"
)

func main() {

	greeting.HelloQi()

	//runSql()

	store.RunSqlx()
}
