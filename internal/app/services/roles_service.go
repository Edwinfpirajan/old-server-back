package services

import (
	"errors"

	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/config"
	"gorm.io/gorm"
)

type RolesService struct{}

func NewRolesService() *RolesService {
	return &RolesService{}
}

func (s *RolesService) GetAllRoles() ([]models.Roles, error) {
	var roles []models.Roles
	err := config.DB.Model(&roles).Find(&roles).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, errors.New("No se encontraron roles")
	}
	return roles, nil
}
