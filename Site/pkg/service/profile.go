package service

import (
	"site/pkg/repository"
)

type ProfileService struct {
	repo repository.Profile
}

func NewProfileService(repo repository.Profile) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) GetNewImagesProfile(lastImageId int, userId int) (urlsHtmlImagesLike []string, urlsHtmlImagesUpload []string, err error) {
	return s.repo.GetNewImagesProfile(lastImageId, userId)
}

func (s *ProfileService) GetNewImages(lastImageId int, userId int, check string) (urlsHtml []string, err error) {
	return s.repo.GetNewImages(lastImageId, userId, check)
}

func (s *ProfileService) GetUserInfo(userId int) (countLikes int, countUploaded int, CountLikesOnUploaded int, err error) {
	return s.repo.GetUserInfo(userId)
}
func (s *ProfileService) GetConsiderationImages(lastImageId int) (urlsHtml []string, err error) {
	return s.repo.GetConsiderationImages(lastImageId)
}

func (s *ProfileService) GetAdminInfo() (countImages int, err error) {
	return s.repo.GetAdminInfo()
}

func (s *ProfileService) GetApproveButtons(imageUrl string) (HtmlStr string, err error) {
	return s.repo.GetApproveButtons(imageUrl)
}

func (s *ProfileService) ConsiderImageAdmin(url string, status string) (promt string, err error) {
	return s.repo.ConsiderImageAdmin(url, status)
}

func (s *ProfileService) GetNewImagesAdmin(lastImageId int, userId int) (urlsHtml []string, err error) {
	return s.repo.GetNewImagesAdmin(lastImageId, userId)
}
