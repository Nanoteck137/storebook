package apis

import (
	"net/http"
	"os"

	"github.com/nanoteck137/pyrin"
	"github.com/nanoteck137/storebook"
	"github.com/nanoteck137/storebook/core"
)

func RegisterHandlers(app core.App, router pyrin.Router) {
	g := router.Group("/api/v1")
	InstallSystemHandlers(app, g)
	InstallAuthHandlers(app, g)

	InstallCollectionHandlers(app, g)

	g = router.Group("/files")
	g.Register(
		pyrin.NormalHandler{
			Name:        "GetCollectionImage",
			Method:      http.MethodGet,
			Path:        "/collections/:id/images/:file",
			HandlerFunc: func(c pyrin.Context) error {
				id := c.Param("id")
				file := c.Param("file")

				dir := app.WorkDir().CollectionDirById(id)
				p := dir.Images()
				f := os.DirFS(p)

				return pyrin.ServeFile(c, f, file)
			},
		},
	)
}

func Server(app core.App) (*pyrin.Server, error) {
	s := pyrin.NewServer(&pyrin.ServerConfig{
		LogName: storebook.AppName,
		RegisterHandlers: func(router pyrin.Router) {
			RegisterHandlers(app, router)
		},
	})

	return s, nil
}
