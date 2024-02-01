package handler

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Function that start when user go to /profile
func (h *Handler) profileGet(c *gin.Context) {
	idData, _ := c.Get(userCtx)
	id, check := idData.(int)
	if !check {
		generateErrorAller(http.StatusBadGateway, "Server fail", "please try again", nil, *&c)
		return
	}

	//If you a anonym redirect to home
	if id == anonymId {
		c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/pictures/")
	}

	// if we admin we get another profile
	if id == adminId {
		userName, err := h.service.Pictures.GetUserName(id)
		if err != nil {
			generateErrorAller(http.StatusBadGateway, "Server fail", "please try again", err, *&c)
			return
		}

		//Get Urls
		urlsHtml, err := h.service.Profile.GetConsiderationImages(0)
		if err != nil {
			generateErrorAller(http.StatusInternalServerError, "Server fail", "please try again", err, *&c)
			return
		}
		//Convert urls to html
		HTMLImages := make([]template.HTML, len(urlsHtml))
		for i := range HTMLImages {
			HTMLImages[i] = template.HTML(urlsHtml[i])
		}

		countImages, err := h.service.Profile.GetAdminInfo()
		if err != nil {
			generateErrorAller(http.StatusInternalServerError, "Server fail", "please try again", err, *&c)
			return
		}

		data := map[string]interface{}{
			"Username":       userName,
			"CountImages":    countImages,
			"UrlsForApprove": HTMLImages,
		}

		tmpl, _ := template.ParseFiles("../templates/profileAdmin.html")
		tmpl.Execute(c.Writer, data)

	} else {

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
	}
}

// Function that start when user scroll down
func (h *Handler) profilePost(c *gin.Context) {
	var urls []string
	//Get user id
	idData, _ := c.Get(userCtx)
	id, checkBool := idData.(int)
	if !checkBool {
		err := errors.New("Can not convert id to int")
		generateErrorAller(http.StatusBadGateway, "Server fail", "please try again", err, *&c)
		return
	}

	//If you a anonym redirect to home
	if id == anonymId {
		c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/pictures/")
	}

	if id == adminId {
		//Get last image id
		lastImageId := c.Request.Header.Get("lastImageId")
		lastImageIdInt, err := strconv.Atoi(lastImageId)
		if err != nil {
			generateErrorAller(http.StatusBadGateway, "Server fail", "please try again", err, *&c)
			return
		}

		urls, err = h.service.Profile.GetNewImagesAdmin(lastImageIdInt, id)
		if err != nil {
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

	} else {

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
}

// Func for admin. for approve or not images
func (h *Handler) profileConsider(c *gin.Context) {

	//Get headers
	status := c.Request.Header.Get("status")
	url := c.Request.Header.Get("url")
	promt, err := h.service.Profile.ConsiderImageAdmin(url, status)

	if err != nil {
		generateErrorAller(http.StatusBadGateway, "Server fail", "please try again", err, *&c)
		return
	}

	//Split promt and save it to table promts
	//For search
	if status == "Yes" {
		promts := h.service.GetPromts(promt)
		h.service.Image.SavePromtsToPromts(promts, url)
	}

	htmlStr := fmt.Sprintf("<div class='alert alert-success' role='alert'>%s!</div>", "Succes")
	tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl.Execute(c.Writer, nil)
}
