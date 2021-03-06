package authorization

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
)

//go:generate mockgen -destination=./mocks/mock_Service.go -package=mocks github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/authorization Service
type Service interface {
	// Http middleware
	Middleware() gin.HandlerFunc
	// Check if it is authorized
	IsAuthorized(ctx context.Context, action, resource string) (bool, error)
	// Check authorized and fail if not authorized
	CheckAuthorized(ctx context.Context, action, resource string) error
}

func NewService(cfgManager config.Manager) Service {
	return &service{
		cfgManager: cfgManager,
	}
}
