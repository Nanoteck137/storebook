package apis

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nanoteck137/pyrin"
	"github.com/nanoteck137/storebook/core"
	"github.com/nanoteck137/validate"
)

type Signin struct {
	Token string `json:"token"`
}

type SigninBody struct {
	Password string `json:"password"`
}

func (b SigninBody) Validate() error {
	return validate.ValidateStruct(&b,
		validate.Field(&b.Password, validate.Required),
	)
}

func InstallAuthHandlers(app core.App, group pyrin.Group) {
	group.Register(
		pyrin.ApiHandler{
			Name:         "Signin",
			Path:         "/auth/signin",
			Method:       http.MethodPost,
			ResponseType: Signin{},
			BodyType:     SigninBody{},
			HandlerFunc: func(c pyrin.Context) (any, error) {
				body, err := pyrin.Body[SigninBody](c)
				if err != nil {
					return nil, err
				}

				if app.Config().Password != body.Password {
					return nil, InvalidCredentials()
				}

				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"iat": time.Now().Unix(),
					// "exp":    time.Now().Add(1000 * time.Second).Unix(),
				})

				tokenString, err := token.SignedString(([]byte)(app.Config().JwtSecret))
				if err != nil {
					return nil, err
				}

				return Signin{
					Token: tokenString,
				}, nil
			},
		},
	)
}
