package database

import (
	"context"
	"lore-keeper-be/internal/types"
)

type Database interface {
	InitDB(ctx context.Context) error

	AddCharacter(ctx context.Context, character *types.Character) error
	DeleteCharacter(ctx context.Context, name string) error
	UpdateCharacter(ctx context.Context, character *types.Character) error
	GetCharacter(ctx context.Context, name string) (*types.Character, error)

	AddWorld(ctx context.Context, world *types.World) error
	DeleteWorld(ctx context.Context, name string) error
	UpdateWorld(ctx context.Context, world *types.World) error
	GetWorld(ctx context.Context, name string) (*types.World, error)

	AddCity(ctx context.Context, city *types.City) error
	DeleteCity(ctx context.Context, name string) error
	UpdateCity(ctx context.Context, city *types.City) error
	GetCity(ctx context.Context, name string) (*types.City, error)

	AddFaction(ctx context.Context, faction *types.Faction) error
	DeleteFaction(ctx context.Context, name string) error
	UpdateFaction(ctx context.Context, faction *types.Faction) error
	GetFaction(ctx context.Context, name string) (*types.Faction, error)

	Shutdown() error
}
