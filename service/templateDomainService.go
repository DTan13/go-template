package service

import (
	"errors"

	"github.com/DTan13/template/domain"
	"github.com/DTan13/template/dto"
	"github.com/DTan13/template/errs"
)

type TemplateDomainService interface {
	GetAllDomains() ([]dto.TemplateDomain, error)
}

type DefaultTemplateDomain struct {
	repo domain.TemplateDomainRepository
}

func (d DefaultTemplateDomain) GetAllDomains() ([]dto.TemplateDomain, error) {
	r, err := d.repo.GetAllDomains()
	if err != errs.TemplateErrorNil {
		return nil, errors.New(err.Error())
	} else {
		var result []dto.TemplateDomain

		for _, d := range r {
			result = append(result, d.ToTemplateDomainDTO())
		}

		return result, nil
	}
}

func NewTemplateDomain(repo domain.TemplateDomainRepository) DefaultTemplateDomain {
	return DefaultTemplateDomain{repo: repo}
}
