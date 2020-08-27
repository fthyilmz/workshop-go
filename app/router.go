package router

import (
	"github.com/gin-gonic/gin"

	"github.com/fthyilmz/workshop-go.git/app/handler"
	"github.com/fthyilmz/workshop-go.git/app/middleware"
	"github.com/fthyilmz/workshop-go.git/app/utils"
	"github.com/fthyilmz/workshop-go.git/config"
	_ "github.com/fthyilmz/workshop-go.git/docs"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	"github.com/swaggo/gin-swagger"        // gin-swagger middleware
)

func New() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	client := config.GetMongodbClient()

	var loginService = utils.LoginService{Client: client}
	var jwtService = utils.JWTAuthService()

	h := handler.Handler{Connection: client}
	l := handler.LoginHandler{LoginService: loginService, JWtService: jwtService}
	d := handler.DashboardHandler{}

	r.POST("/v1/login", l.Login)

	authGroup := r.Group("/v1")

	authGroup.Use(middleware.AuthorizeJWT())

	authGroup.GET("/furniture", h.All)
	authGroup.GET("/furniture/:id", h.Get)
	authGroup.POST("/furniture", h.Store)
	authGroup.PUT("/furniture/:id", h.Update)
	authGroup.DELETE("/furniture/:id", h.Delete)

	authGroup.GET("/room", h.GetRoomList)
	authGroup.GET("/total/furniture", d.TotalFurniture)
	authGroup.GET("/total/furniture/:apartmentId", d.TotalFurnitureOfApartment)

	return r
}
