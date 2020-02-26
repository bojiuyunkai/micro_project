// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package controllers

import (
	"github.com/google/wire"
	"hz_project/internal/app/details/repositories"
	"hz_project/internal/app/details/services"
	"hz_project/internal/pkg/config"
	"hz_project/internal/pkg/database"
	"hz_project/internal/pkg/log"
)

// Injectors from wire.go:

func CreateDetailsController(cf string, sto repositories.DetailsRepository) (*DetailsController, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.New(options)
	if err != nil {
		return nil, err
	}
	detailsService := services.NewDetailService(logger, sto)
	detailsController := NewDetailsController(logger, detailsService)
	return detailsController, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, database.ProviderSet, services.ProviderSet, ProviderSet)