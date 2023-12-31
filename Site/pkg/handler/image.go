package handler

import (
	"bytes"
	"fmt"
	"html/template"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
)

func (h *Handler) uploadImagePost(c *gin.Context) {
	//Get Promts
	promts := c.Request.PostFormValue("inputPromts")
	if len(promts) >= 1024 {
		generateErrorAller(http.StatusBadRequest, "To big promt", "Please write promt smaller you have more than 1024 symbols", nil, *&c)
		return
	}
	//Split string promt to slice
	promtsSlice := h.service.Image.GetPromts(promts)
	if len(promtsSlice) == 0 {
		generateErrorAller(http.StatusBadRequest, "No promts", "Please write promts for your image", nil, *&c)
		return
	}

	//Get file
	file, err := c.FormFile("image")
	if err != nil {
		generateErrorAller(http.StatusBadRequest, "No image", "Please choose image", err, *&c)
		return
	}
	//Recreate name
	file.Filename = h.service.Image.GenerateNewImageName()

	//Save file
	err = c.SaveUploadedFile(file, "../static/images/highQuality/"+file.Filename)
	if err != nil {
		generateErrorAller(http.StatusInternalServerError, "Upload image failed", "Please try again", err, *&c)
		return
	}

	//Save low quality
	fileContainer, err := file.Open()
	if err != nil {
		generateErrorAller(http.StatusInternalServerError, "Upload image failed", "Please try again", err, *&c)
		return
	}
	byteContainer, err := io.ReadAll(fileContainer)
	if err != nil {
		generateErrorAller(http.StatusInternalServerError, "Upload image failed", "Please try again", err, *&c)
		return
	}

	image, _, err := image.Decode(bytes.NewReader(byteContainer))
	if err != nil {
		generateErrorAller(http.StatusInternalServerError, "Upload image failed", "Please try again", err, *&c)
		return
	}
	out, err := os.Create("../static/images/lowQuality/" + file.Filename)

	fileLowQuality := resize.Resize(350, 0, image, resize.Lanczos2)
	opt := jpeg.Options{
		Quality: 90,
	}
	err = jpeg.Encode(out, fileLowQuality, &opt)

	//Save 20px (For lazy loading in html)
	out, err = os.Create("../static/images/20pxImage/" + file.Filename)
	fileLowQuality = resize.Resize(20, 0, image, resize.Lanczos2)
	opt = jpeg.Options{
		Quality: 0,
	}
	err = jpeg.Encode(out, fileLowQuality, &opt)

	//Get user id
	userId, ok := getUserId(c)
	if ok != nil {
		generateErrorAller(http.StatusBadGateway, "Server fail", ok.Error(), err, *&c)
		return
	}

	//Save link to database
	err = h.service.Image.SaveImageLink(userId, file.Filename)
	if err != nil {
		generateErrorAller(http.StatusInternalServerError, "Server fail", "Please try again", err, *&c)
		return
	}

	err = h.service.Image.SavePromts(promtsSlice, file.Filename)
	if err != nil {
		generateErrorAller(http.StatusInternalServerError, "Server fail", "Please try again", err, *&c)
		return
	}

	//Success message
	htmlStr := "<div class='alert alert-success' role='alert'>Succes upload image</div>"
	tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl.Execute(c.Writer, nil)

	htmlStr = fmt.Sprintf("<img src='%s' />", "../static/images/lowQuality/"+file.Filename)
	tmpl, _ = template.New("t").Parse(htmlStr)
	tmpl.Execute(c.Writer, nil)

	return
}

func (h *Handler) uploadImageGet(c *gin.Context) {
	//Get user user name
	idData, _ := c.Get(userCtx)
	id, check := idData.(int)
	if check == false {
		generateErrorAller(http.StatusBadGateway, "Server fail", "please try again", nil, *&c)
		return
	}
	userName, err := h.service.Pictures.GetUserName(id)
	if err != nil {
		generateErrorAller(http.StatusBadGateway, "Server fail", "please try again", nil, *&c)
		return
	}
	tmpl, _ := template.ParseFiles("../templates/uploadImage.html")
	data := map[string]string{
		"Username": userName,
	}

	tmpl.Execute(c.Writer, data)

	return
}
