package services

import (
	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/config"
)

type PayrollService struct{}

func NewPayrollService() *PayrollService {
	return &PayrollService{}
}

func (s *PayrollService) GetAllPayRollService() (models.Headquarters, error) {
	var headquarters models.Headquarters
	err := config.DB.Model(&headquarters).Find(&headquarters).Error
	if err != nil {
		return nil, err
	}
	return headquarters, nil
}
