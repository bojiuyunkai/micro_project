// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package services

import (
	"github.com/google/wire"
	"hz_project/internal/app/reviews/repositories"
	"hz_project/internal/pkg/config"
	"hz_project/internal/pkg/database"
	"hz_project/internal/pkg/log"
)

// Injectors from wire.go:

func CreateReviewsService(cf string, sto repositories.ReviewsRepository) (ReviewsService, error) {
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
	reviewsService := NewReviewService(logger, sto)
	return reviewsService, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, database.ProviderSet, ProviderSet)
