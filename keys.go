package checkers

import "cosmossdk.io/collections"

const (
	ModuleName     = "checkers"
	MaxIndexLength = 256
)

var (
	ParamsKey      = collections.NewPrefix("Params")
	StoredGamesKey = collections.NewPrefix("StoredGames/value/")
)
