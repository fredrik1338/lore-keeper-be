package database

import (
	"context"
	"lore-keeper-be/internal/types"
)

type Database interface {
	InitDB(ctx context.Context) error

	AddCharacter(ctx context.Context, character *types.Character) error
	ListCharacters(ctx context.Context) ([]string, error)
	DeleteCharacter(ctx context.Context, name string) error
	UpdateCharacter(ctx context.Context, character *types.Character) error
	GetCharacter(ctx context.Context, name string) (*types.Character, error)

	AddWorld(ctx context.Context, world *types.World) error
	ListWorlds(ctx context.Context) ([]string, error)
	DeleteWorld(ctx context.Context, name string) error
	UpdateWorld(ctx context.Context, world *types.World) error
	GetWorld(ctx context.Context, name string) (*types.World, error)

	AddCity(ctx context.Context, city *types.City) error
	ListCities(ctx context.Context) ([]string, error)
	DeleteCity(ctx context.Context, name string) error
	UpdateCity(ctx context.Context, city *types.City) error
	GetCity(ctx context.Context, name string) (*types.City, error)

	AddFaction(ctx context.Context, faction *types.Faction) error
	ListFactions(ctx context.Context) ([]string, error)
	DeleteFaction(ctx context.Context, name string) error
	UpdateFaction(ctx context.Context, faction *types.Faction) error
	GetFaction(ctx context.Context, name string) (*types.Faction, error)

	AddUniverse(ctx context.Context, universe *types.Universe) error
	ListUniverses(ctx context.Context) ([]string, error)
	DeleteUniverse(ctx context.Context, name string) error
	UpdateUniverse(ctx context.Context, universe *types.Universe) error
	GetUniverse(ctx context.Context, name string) (*types.Universe, error)

	Shutdown() error
}
