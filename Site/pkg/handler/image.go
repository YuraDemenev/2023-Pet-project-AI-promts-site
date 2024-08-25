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
	"sync"

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
	//promtsSlice := h.service.Image.GetPromts(promts)
	if len(promts) == 0 {
		generateErrorAller(http.StatusBadRequest, "No promts", "Please write promts for your image", nil, c)
		return
	}

	//Get file
	file, err := c.FormFile("image")
	if err != nil {
		generateErrorAller(http.StatusBadRequest, "No image", "Please choose image", err, c)
		return
	}
	//Recreate name
	file.Filename = h.service.Image.GenerateNewImageName()

	//Get user id
	userId, ok := getUserId(c)
	if ok != nil {
		generateErrorAller(http.StatusBadGateway, "Server fail", ok.Error(), err, c)
		return
	}

	err = h.service.Image.SavePromtsToConsideration(promts, file.Filename, userId)
	if err != nil {
		generateErrorAller(http.StatusInternalServerError, "Server fail", "Please try again", err, c)
		return
	}

	//Save file
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = c.SaveUploadedFile(file, "../images/highQuality/"+file.Filename)
		if err != nil {
			generateErrorAller(http.StatusInternalServerError, "Upload image failed", "Please try again", err, c)
			return
		}
	}()

	//Save low quality
	fileContainer, err := file.Open()
	if err != nil {
		generateErrorAller(http.StatusInternalServerError, "Upload image failed", "Please try again", err, c)
		return
	}

	byteContainer, err := io.ReadAll(fileContainer)
	if err != nil {
		generateErrorAller(http.StatusInternalServerError, "Upload image failed", "Please try again", err, c)
		return
	}

	image, _, err := image.Decode(bytes.NewReader(byteContainer))
	if err != nil {
		generateErrorAller(http.StatusInternalServerError, "Upload image failed", "Please try again", err, c)
		return
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		out, err := os.Create("../images/lowQuality/" + file.Filename)

		if err != nil {
			generateErrorAller(http.StatusInternalServerError, "Upload image failed", "Please try again", err, c)
			return
		}

		fileLowQuality := resize.Resize(350, 0, image, resize.Lanczos2)
		opt := jpeg.Options{
			Quality: 90,
		}

		err = jpeg.Encode(out, fileLowQuality, &opt)
		if err != nil {
			generateErrorAller(http.StatusInternalServerError, "Upload image failed", "Please try again", err, c)
			return
		}
	}()

	//Save 20px (For lazy loading in html)

	wg.Add(1)
	go func() {
		defer wg.Done()

		out, err := os.Create("../images/20pxImage/" + file.Filename)
		if err != nil {
			generateErrorAller(http.StatusInternalServerError, "Upload image failed", "Please try again", err, c)
			return
		}

		fileLowQuality := resize.Resize(20, 0, image, resize.Lanczos2)
		opt := jpeg.Options{
			Quality: 0,
		}

		err = jpeg.Encode(out, fileLowQuality, &opt)
		if err != nil {
			generateErrorAller(http.StatusInternalServerError, "Upload image failed", "Please try again", err, c)
			return
		}
	}()

	wg.Wait()

	//Success message
	htmlStr := "<div class='alert alert-success' role='alert'>Succes upload image. Your image is under consideration</div>"
	tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl.Execute(c.Writer, nil)

	htmlStr = fmt.Sprintf("<img src='%s' />", "../images/lowQuality/"+file.Filename)
	tmpl, _ = template.New("t").Parse(htmlStr)
	tmpl.Execute(c.Writer, nil)

}

func (h *Handler) uploadImageGet(c *gin.Context) {
	//Get user user name
	idData, _ := c.Get(userCtx)
	id, check := idData.(int)
	if !check {
		generateErrorAller(http.StatusBadGateway, "Server fail", "please try again", nil, c)
		return
	}
	userName, err := h.service.Pictures.GetUserName(id)
	if err != nil {
		generateErrorAller(http.StatusBadGateway, "Server fail", "please try again", nil, c)
		return
	}
	tmpl, _ := template.ParseFiles("../templates/uploadImage.html")
	data := map[string]string{
		"Username": userName,
		"URL":      h.url,
	}

	tmpl.Execute(c.Writer, data)

}
