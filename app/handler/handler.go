package handler

import (
	"github.com/fthyilmz/workshop-go.git/app/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	Connection *mongo.Client
}

type DashboardHandler struct {
}

type LoginHandler struct {
	LoginService utils.LoginService
	JWtService   utils.JWTService
}
