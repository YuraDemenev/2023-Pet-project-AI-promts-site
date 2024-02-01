package service

import (
	"site/pkg/cache"
	site "site/pkg/elements"
	"site/pkg/repository"
)

type Authorization interface {
	CreateUser(user site.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type Image interface {
	SaveImageLink(id int, imageLink string) error
	GetPromts(promt string) []string
	GenerateNewImageName() string
	SavePromtsToPromts(promts []string, imageLink string) error
	SavePromtsToConsideration(promts string, imageUrl string, userId int) error
}

type Pictures interface {
	GetUserName(id int) (string, error)
	GetNewImages(lastImageId int, postCache cache.CacheImages) (urlsHtml []string, imageNums []string, err error)
	GetImagePromts(imageUrl string, userId int) (string, error)
	SearchImages(promt string, lastImageId int) (urlsHtml []string, err error)
	AddLike(imageUrl string, userId int, countLike int) (string, error)
}

type Profile interface {
	GetNewImagesProfile(lastImageId int, userId int) (urlsHtmlImagesLike []string, urlsHtmlImagesUpload []string, err error)
	GetNewImages(lastImageId int, userId int, check string) (urlsHtml []string, err error)
	GetUserInfo(userId int) (countLikes int, countUploaded int, CountLikesOnUploaded int, err error)
	GetConsiderationImages(lastImageId int) (urlsHtml []string, err error)
	//For admin functions
	GetAdminInfo() (countImages int, err error)
	GetApproveButtons(imageUrl string) (HtmlStr string, err error)
	ConsiderImageAdmin(url string, status string) (promt string, err error)
	GetNewImagesAdmin(lastImageId int, userId int) (urlsHtml []string, err error)
}
type Service struct {
	Authorization
	Image
	Pictures
	Profile
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Image:         NewImageService(repos.Image),
		Pictures:      NewPicturesService(repos.Pictures),
		Profile:       NewProfileService(repos.Profile),
	}
}
