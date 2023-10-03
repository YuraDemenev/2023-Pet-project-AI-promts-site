package repository

import (
	site "site/pkg/elements"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user site.User) (int, error)
	GetUser(username, password string) (site.User, error)
}

type Image interface {
	SaveLink(id int, imageLink string) error
	SavePromts(promts []string, imageLink string) error
}

type Pictures interface {
	GetUserName(id int) (string, error)
	GetNewImages(lastImageId int) (urlsHtml []string, err error)
	GetImagePromts(imageUrl string, userId int) (string, error)
	SearchImages(promtSlice string, lastImageId int) (urlsHtml []string, err error)
	AddLike(image_url string, userId int, countLike int) (string, error)
}

type Profile interface {
	GetNewImagesProfile(lastImageId int, userId int) (urlsHtmlImagesLike []string, urlsHtmlImagesUpload []string, err error)
	GetNewImages(lastImageId int, userId int, check string) (urlsHtml []string, err error)
	GetUserInfo(userId int) (countLikes int, countUploaded int, CountLikesOnUploaded int, err error)
}

type Repository struct {
	Authorization
	Image
	Pictures
	Profile
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Image:         NewImagePostgres(db),
		Pictures:      NewPicturesPostgres(db),
		Profile:       NewProfilePostgres(db),
	}
}
