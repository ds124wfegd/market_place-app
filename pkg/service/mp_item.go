package service

import (
	"github.com/ds124wfegd/mp-app"
	"github.com/ds124wfegd/mp-app/pkg/repository"
)

type MpItemService struct {
	repo     repository.MpItem
	listRepo repository.MpList
}

func NewMpItemService(repo repository.MpItem, listRepo repository.MpList) *MpItemService {
	return &MpItemService{repo: repo, listRepo: listRepo}
}

func (s *MpItemService) Create(userId, listId int, item mp.MpItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		// list does not exists or does not belongs to user
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *MpItemService) GetAll(userId, listId int) ([]mp.MpItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *MpItemService) GetById(userId, itemId int) (mp.MpItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *MpItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *MpItemService) Update(userId, itemId int, input mp.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}
