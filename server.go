package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nagahshi/gin-poc/controller"
	"github.com/nagahshi/gin-poc/middleware"
	"github.com/nagahshi/gin-poc/repository"
	"github.com/nagahshi/gin-poc/service"
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService       = service.NewVideoService(videoRepository)
	videoContoller  controller.VideoController = controller.New(videoService)

	loginService   service.LoginService      = service.NewLoginService()
	jwtService     service.JWTService        = service.NewJWTService()
	authController controller.AuthController = controller.NewAuthController(loginService, jwtService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery())
	server.Use(middleware.Logger())

	authRouter := server.Group("auth")
	{
		authRouter.POST("/login", func(ctx *gin.Context) {
			token := authController.Login(ctx)
			if token != "" {
				ctx.JSON(http.StatusOK, gin.H{
					"token": token,
				})
			} else {
				ctx.JSON(http.StatusUnauthorized, nil)
			}

		})
	}

	videoPrivateRouter := server.Group("videos", middleware.AuthorizeJWT())
	{
		videoPrivateRouter.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, videoContoller.FindAll())
		})

		videoPrivateRouter.POST("/", func(ctx *gin.Context) {
			err := videoContoller.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusCreated, gin.H{
					"message": "created",
				})
			}
		})

		videoPrivateRouter.PUT("/:id", func(ctx *gin.Context) {
			err := videoContoller.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusCreated, gin.H{
					"message": "updated",
				})
			}
		})

		videoPrivateRouter.DELETE("/:id", func(ctx *gin.Context) {
			err := videoContoller.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusCreated, gin.H{
					"message": "deleted",
				})
			}
		})
	}
	server.Run(":8002")
}
