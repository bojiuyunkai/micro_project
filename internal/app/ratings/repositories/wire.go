// +build wireinject

package repositories

import (
	"github.com/google/wire"
	"hz_project/internal/pkg/config"
	"hz_project/internal/pkg/database"
	"hz_project/internal/pkg/log"
)



var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateRatingRepository(f string) (RatingsRepository, error) {
	panic(wire.Build(testProviderSet))
}

