package repository

import (
	"site/pkg/cache"
	site "site/pkg/elements"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user site.User) (int, error)
	GetUser(username, password string) (site.User, error)
}

type Image interface {
	SaveLink(id int, imageLink string) error
	SavePromtsToPromts(promts []string, imageLink string) error
	SavePromtsToConsideration(promts string, imageUrl string, userId int) error
}

type Pictures interface {
	GetUserName(id int) (string, error)
	GetNewImages(lastImageId int, postCache cache.CacheImages) (urlsHtml []string, imageNums []string, err error)
	GetImagePromts(imageUrl string, userId int) (string, error)
	SearchImages(promtSlice string, lastImageId int) (urlsHtml []string, err error)
	AddLike(image_url string, userId int, countLike int) (string, error)
}

type Profile interface {
	GetNewImagesProfile(lastImageId int, userId int) (urlsHtmlImagesLike []string, urlsHtmlImagesUpload []string, err error)
	GetNewImages(lastImageId int, userId int, check string) (urlsHtml []string, err error)
	GetUserInfo(userId int) (countLikes int, countUploaded int, CountLikesOnUploaded int, err error)
	//For admin
	GetConsiderationImages(lastImageId int) (urlsHtml []string, err error)
	GetAdminInfo() (countImages int, err error)
	GetApproveButtons(imageUrl string) (HtmlStr string, err error)
	ConsiderImageAdmin(url string, status string) (promt string, err error)
	GetNewImagesAdmin(lastImageId int, userId int) (urlsHtml []string, err error)
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
