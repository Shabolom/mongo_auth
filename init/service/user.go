package service

import (
	"net/http"
	"test_hh_1/init/domain"
	"test_hh_1/init/models"
	"test_hh_1/init/repository"
	"test_hh_1/tools"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

var userRepo = repository.NewUserRepo()

func (ua *UserService) GiveToken(refToken string, user models.User) (error, int) {
	userEntity := domain.RefreshToken{
		UserID:   user.ID.String(),
		RefToken: refToken,
	}

	err := userRepo.GiveToken(userEntity)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusCreated
}

func (ua *UserService) RefreshToken(oldRefToken, newRef string) tools.BaseResult {

	err := userRepo.RefreshToken(oldRefToken, newRef)
	if err != nil {
		return tools.BaseResult{
			Err:    err,
			Status: http.StatusBadRequest,
		}
	}

	return tools.BaseResult{
		Status: http.StatusOK,
		Result: "токены обновлены",
	}
}
