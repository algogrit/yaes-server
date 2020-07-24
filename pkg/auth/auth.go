package auth

//go:generate mockgen -destination mock_$GOFILE -source=$GOFILE -package auth

import (
	"github.com/justinas/alice"
)

// Auth package is to provide a middleware capable of providing authentication
type Auth interface {
	Middleware() alice.Chain
}
