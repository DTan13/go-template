package domain

import (
	"github.com/DTan13/template/dto"
	"github.com/DTan13/template/errs"
)

// TemplateDomain is domain
type TemplateDomain struct {
	id   string
	Name string
}

type TemplateDomainRepository interface {
	GetAllDomains() ([]TemplateDomain, *errs.TemplateError)
}

func (d TemplateDomain) ToTemplateDomainDTO() dto.TemplateDomain {
	return dto.TemplateDomain{
		Name: d.Name,
	}
}
