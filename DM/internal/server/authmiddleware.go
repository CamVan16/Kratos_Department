package server

import (
	"DM/internal/biz"
	"context"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type contextKey string

const (
	UserContextKey contextKey = "user"
	reason         string     = "UNAUTHORIZED"
)

var (
	ErrMissingJwtToken = errors.Unauthorized(reason, "JWT token is missing")
	ErrTokenInvalid    = errors.Unauthorized(reason, "Token is invalid")
	ErrTokenParseFail  = errors.Unauthorized(reason, "Fail to parse JWT token")
	ErrAccessDenied    = errors.Unauthorized(reason, "access denied")
)

var ignoreOperations = []string{
	"/helloworld.v1.EmployeeService",
	"/helloworld.v1.SubDepartmentService",
	"/helloworld.v1.UserService",
}

func AuthMiddleware(allowedRoles ...string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				op := tr.Operation()
				//fmt.Println("Operation:", op)
				for _, e := range ignoreOperations {
					if strings.HasPrefix(op, e) {
						return handler(ctx, req)
					}
				}

				if ht, ok := tr.(*http.Transport); ok {
					tokenString := ht.Request().Header.Get("Authorization")
					if tokenString == "" {
						return nil, ErrMissingJwtToken
					}

					tokenString = strings.TrimPrefix(tokenString, "Bearer ")
					claims, err := biz.ParseJWT(tokenString)
					if err != nil {
						return nil, ErrTokenParseFail
					}

					hasAccess := false
					for _, role := range allowedRoles {
						if claims.Role == role {
							hasAccess = true
							break
						}
					}

					if !hasAccess {
						return nil, ErrAccessDenied
					}
					ctx = context.WithValue(ctx, UserContextKey, claims)
				}
			}

			return handler(ctx, req)
		}
	}
}
