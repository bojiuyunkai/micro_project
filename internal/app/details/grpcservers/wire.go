// +build wireinject

package grpcservers

import (
	"github.com/google/wire"
	"hz_project/internal/pkg/config"
	"hz_project/internal/pkg/database"
	"hz_project/internal/pkg/log"
	"hz_project/internal/app/details/services"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateDetailsServer(cf string, service services.DetailsService) (*DetailsServer, error) {
	panic(wire.Build(testProviderSet))
}