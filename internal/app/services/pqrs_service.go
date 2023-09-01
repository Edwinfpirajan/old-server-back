package services

import (
	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/entity"
)

func GetPqrsSacs() ([]entity.PqrsSac, error) {
	var pqrsSacs []entity.PqrsSac
	result := config.DB.Table("pqrs_sac").
		Select("*").
		Joins("INNER JOIN document_type on pqrs_sac.fk_document_type = document_type.id").
		Joins("INNER JOIN pqrs_type on pqrs_sac.fk_pqrs_type_id = pqrs_type.id").
		Scan(&pqrsSacs)

	if result.Error != nil {
		return nil, result.Error
	}

	return pqrsSacs, nil
}
