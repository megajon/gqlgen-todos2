package graph

import (
	"github.com/megajon/gqlgen-todos2/postgres"
)

type Resolver struct {
	UsersRepo postgres.UsersRepo
	TodosRepo postgres.TodosRepo
}
