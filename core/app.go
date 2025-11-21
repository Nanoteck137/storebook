package core

import (
	"github.com/nanoteck137/pyrin/trail"
	"github.com/nanoteck137/storebook/config"
	"github.com/nanoteck137/storebook/database"
	"github.com/nanoteck137/storebook/types"
)

// Inspiration from Pocketbase: https://github.com/pocketbase/pocketbase
// File: https://github.com/pocketbase/pocketbase/blob/master/core/app.go
type App interface {
	Logger() *trail.Logger

	DB() *database.Database
	Config() *config.Config

	WorkDir() types.WorkDir

	Bootstrap() error
}
