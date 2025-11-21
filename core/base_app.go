package core

import (
	"os"

	"github.com/nanoteck137/pyrin/trail"
	"github.com/nanoteck137/storebook"
	"github.com/nanoteck137/storebook/config"
	"github.com/nanoteck137/storebook/database"
	"github.com/nanoteck137/storebook/types"
)

var _ App = (*BaseApp)(nil)

type BaseApp struct {
	logger          *trail.Logger
	db              *database.Database
	config          *config.Config
}

func (app *BaseApp) Logger() *trail.Logger {
	return app.logger
}

func (app *BaseApp) DB() *database.Database {
	return app.db
}

func (app *BaseApp) Config() *config.Config {
	return app.config
}

func (app *BaseApp) WorkDir() types.WorkDir {
	return app.config.WorkDir()
}

func (app *BaseApp) Bootstrap() error {
	var err error

	workDir := app.config.WorkDir()

	dirs := []string{
		workDir.CollectionsDir(),
	}

	for _, dir := range dirs {
		err = os.Mkdir(dir, 0755)
		if err != nil && !os.IsExist(err) {
			return err
		}
	}

	app.db, err = database.Open(workDir.DatabaseFile())
	if err != nil {
		return err
	}

	if app.config.RunMigrations {
		err = app.db.RunMigrateUp()
		if err != nil {
			return err
		}
	}

	return nil
}

func NewBaseApp(config *config.Config) *BaseApp {
	return &BaseApp{
		logger: storebook.DefaultLogger(),
		config: config,
	}
}
