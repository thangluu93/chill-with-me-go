package core

import (
	"main/models"
	"net/http"
)

type Executor struct {
	client *http.Client
	config *models.Config
}
