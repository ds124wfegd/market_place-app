package service

import (
	"github.com/ds124wfegd/mp-app"
	"github.com/ds124wfegd/mp-app/pkg/repository"
)

type MpListService struct {
	repo repository.MpList
}

func NewMpListService(repo repository.MpList) *MpListService {
	return &MpListService{repo: repo}
}

func (s *MpListService) Create(userId int, list mp.MpList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *MpListService) GetAll(userId int) ([]mp.MpList, error) {
	return s.repo.GetAll(userId)
}

func (s *MpListService) GetById(userId, listId int) (mp.MpList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *MpListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *MpListService) Update(userId, listId int, input mp.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, listId, input)
}
