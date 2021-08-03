package main

import (
	"github.com/rachadiannovansyah/go-echo-boilerplate-ipj/migrations/list"

	gm "github.com/ShkrutDenis/go-migrations"
	gmStore "github.com/ShkrutDenis/go-migrations/store"
)

func main() {
	gm.Run(getMigrationsList())
}

func getMigrationsList() []gmStore.Migratable {
	return []gmStore.Migratable{
		&list.CreateUserTable{},
		&list.CreatePostTable{},
		&list.UpdateUserTable{},
		&list.CreateCategoryTable{},
	}
}
