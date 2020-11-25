package service

import (
	"github.com/nagahshi/gin-poc/entity"
	"github.com/nagahshi/gin-poc/repository"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
}

type videoService struct {
	videoRepository repository.VideoRepository
}

func NewVideoService(repository repository.VideoRepository) VideoService {
	return &videoService{
		videoRepository: repository,
	}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videoRepository.Save(video)

	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videoRepository.Index()
}

func (service *videoService) Update(video entity.Video) {
	service.videoRepository.Update(video)
}

func (service *videoService) Delete(video entity.Video) {
	service.videoRepository.Delete(video)
}
