package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/bayu-wibowo-mf/personal-go-graphql/graph/model"
	"github.com/go-pg/pg"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
	DB    *pg.DB
}
