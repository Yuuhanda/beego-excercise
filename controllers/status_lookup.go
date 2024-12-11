package controllers

import (
    "myproject/services"
)

type StatusLookupController struct {
    statusService *services.StatusLookupService
}

func NewStatusLookupController() *StatusLookupController {
    return &StatusLookupController{
        statusService: services.NewStatusLookupService(),
    }
}
