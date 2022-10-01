package core

import (
	"main/models"
	"net/http"
)

type Executor struct {
	client *http.Client
	config *models.Config
}

func NewExecutor(config *models.Config) *Executor {
	return &Executor{
		client: &http.Client{},
		config: config,
	}
}
