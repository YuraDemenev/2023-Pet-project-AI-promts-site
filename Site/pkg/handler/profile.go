package handler

import (
	"errors"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Function that start when user go to /profile
func (h *Handler) profileGet(c *gin.Context) {
	idData, _ := c.Get(userCtx)
	id, check := idData.(int)
	if check == false {
		generateErrorAller(http.StatusBadGateway, "Server fail", "please try again", nil, *&c)
		return
	}

	//If you a anonym redirect to home
	if id == anonymId {
		c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/pictures/")
	}

	userName, err := h.service.Pictures.GetUserName(id)
	if err != nil {
		generateErrorAller(http.StatusBadGateway, "Server fail", "please try again", err, *&c)
		return
	}

	//Get Urls
	urslImagesLike, urslImagesUpload, err := h.service.Profile.GetNewImagesProfile(0, id)
	if err != nil {
		generateErrorAller(http.StatusInternalServerError, "Server fail", "please try again", err, *&c)
		return
	}

	//Create a html of div imagesLike
	HTMLImagesLike := make([]template.HTML, len(urslImagesLike))
	for i := range HTMLImagesLike {
		HTMLImagesLike[i] = template.HTML(urslImagesLike[i])
	}

	//Create a html of div imagesUpload
	HTMLImagesUpload := make([]template.HTML, len(urslImagesUpload))
	for i := range urslImagesUpload {
		HTMLImagesUpload[i] = template.HTML(urslImagesUpload[i])
	}

	countLikes, countUploaded, countLikesOnUploaded, err := h.service.GetUserInfo(id)
	if err != nil {
		generateErrorAller(http.StatusInternalServerError, "Server fail", "please try again", err, *&c)
		return
	}

	data := map[string]interface{}{
		"Username":             userName,
		"UrlsLike":             HTMLImagesLike,
		"UrlsUpload":           HTMLImagesUpload,
		"CountLikes":           countLikes,
		"CountUploaded":        countUploaded,
		"CountLikesOnUploaded": countLikesOnUploaded,
	}

	tmpl, _ := template.ParseFiles("../templates/profile.html")
	tmpl.Execute(c.Writer, data)
	return
}

// Function that start when user scroll down
func (h *Handler) profilePost(c *gin.Context) {
	var urls []string
	//Get user id
	idData, _ := c.Get(userCtx)
	id, checkBool := idData.(int)
	if checkBool == false {
		err := errors.New("Can not convert id to int")
		generateErrorAller(http.StatusBadGateway, "Server fail", "please try again", err, *&c)
		return
	}

	//If you a anonym redirect to home
	if id == anonymId {
		c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/pictures/")
	}

	//Get last image id
	lastImageId := c.Request.Header.Get("lastImageId")
	lastImageIdInt, err := strconv.Atoi(lastImageId)
	if err != nil {
		generateErrorAller(http.StatusBadGateway, "Server fail", "please try again", err, *&c)
		return
	}

	//We need to know what picture send(like or upload) and
	//for recognize it we get header promt
	check := c.Request.Header.Get("promt")
	if check == "like" || check == "upload" {
		urls, err = h.service.Profile.GetNewImages(lastImageIdInt, id, check)
		if err != nil {
			generateErrorAller(http.StatusInternalServerError, "Server error", "Please try again", err, *&c)
			return
		}

	} else {
		err := errors.New("check not like and not upload")
		generateErrorAller(http.StatusInternalServerError, "Server error", "Please try again", err, *&c)
		return
	}

	for _, htmlStr := range urls {
		if htmlStr == "" {
			break
		}
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(c.Writer, nil)
	}

	return
}
