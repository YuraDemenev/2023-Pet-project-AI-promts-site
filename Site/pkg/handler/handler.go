package handler

import (
	"net/http"
	"site/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	//router.LoadHTMLFiles("../templates/registration.html", "../templates/authorization.html", "../templates/pictures.html")
	router.LoadHTMLGlob("../templates/*")
	fs := http.FileSystem(http.Dir("../static"))
	router.StaticFS("static/", fs)

	auth := router.Group("/auth")
	{
		auth.GET("/sign-up", h.registration)
		auth.POST("/sign-up", h.signUp)

		auth.GET("/sign-in", h.authorization)
		auth.POST("/sign-in", h.signIn)
	}
	pictures := router.Group("/pictures", h.userIdentity)
	{
		//For home page
		pictures.GET("/", h.watchPictureGet)
		pictures.POST("/", h.watchPicturePost)

		//For load picture
		pictures.GET("/uploadImage", h.uploadImageGet)
		pictures.POST("/uploadImage", h.uploadImagePost)

		//For function when user click on image
		pictures.POST("/:info", h.showPromts)

		//For search
		pictures.GET("/search", h.searchGet)
		pictures.POST("/search", h.searchPost)

		//For add or delete like
		pictures.POST("/addLike", h.addLikePost)

		//For watch profile
		pictures.GET("/profile", h.profileGet)
		pictures.POST("/profile", h.profilePost)

	}

	return router
}
