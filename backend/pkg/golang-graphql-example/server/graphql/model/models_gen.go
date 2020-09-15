// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
)

type NewTodo struct {
	Text string `json:"text"`
}

// Pagination information
type PageInfo struct {
	// Has a next page ?
	HasNextPage bool `json:"hasNextPage"`
	// Has a previous page ?
	HasPreviousPage bool `json:"hasPreviousPage"`
	// Shortcut to first edge cursor in the result chunk
	StartCursor *string `json:"startCursor"`
	// Shortcut to last edge cursor in the result chunk
	EndCursor *string `json:"endCursor"`
}

type TodoConnection struct {
	Edges    []*TodoEdge `json:"edges"`
	PageInfo *PageInfo   `json:"pageInfo"`
}

type TodoEdge struct {
	Cursor string       `json:"cursor"`
	Node   *models.Todo `json:"node"`
}

type UpdateTodo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
