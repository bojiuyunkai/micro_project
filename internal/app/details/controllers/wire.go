// +build wireinject

package controllers

import (
	"github.com/google/wire"
	"hz_project/internal/pkg/config"
	"hz_project/internal/pkg/database"
	"hz_project/internal/pkg/log"
	"hz_project/internal/app/details/services"
	"hz_project/internal/app/details/repositories"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	services.ProviderSet,
	//repositories.ProviderSet,
	ProviderSet,
)


func CreateDetailsController(cf string, sto repositories.DetailsRepository) (*DetailsController, error) {
	panic(wire.Build(testProviderSet))
}
