package apis

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nanoteck137/pyrin"
	"github.com/nanoteck137/storebook"
	"github.com/nanoteck137/storebook/core"
	"github.com/nanoteck137/storebook/utils"
)

// TODO(patrik): Remove
var logger = storebook.DefaultLogger()

func LoggedIn(app core.App, c pyrin.Context) error {
	passwordHeader := c.Request().Header.Get("X-Password")
	if passwordHeader != "" {
		if app.Config().Password == passwordHeader {
			return nil
		}
	}

	authHeader := c.Request().Header.Get("Authorization")
	tokenString := utils.ParseAuthHeader(authHeader)
	if tokenString == "" {
		return InvalidAuth("invalid authorization header")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(app.Config().JwtSecret), nil
	})
	if err != nil {
		// TODO(patrik): Handle error better
		return InvalidAuth("invalid authorization token")
	}

	jwtValidator := jwt.NewValidator(jwt.WithIssuedAt())

	if _, ok := token.Claims.(jwt.MapClaims); ok {
		if err := jwtValidator.Validate(token.Claims); err != nil {
			return InvalidAuth("invalid authorization token")
		}

		return nil
	}

	return InvalidAuth("invalid authorization token")
}

func ConvertURL(c pyrin.Context, path string) string {
	host := c.Request().Host

	scheme := "http"

	h := c.Request().Header.Get("X-Forwarded-Proto")
	if h != "" {
		scheme = h
	}

	return fmt.Sprintf("%s://%s%s", scheme, host, path)
}
