package domain

import (
	"github.com/DTan13/template/errs"
)

type TemplateDomainStub struct {
	domains []TemplateDomain
}

func (s TemplateDomainStub) GetAllDomains() ([]TemplateDomain, *errs.TemplateError) {
	return s.domains, nil
}

func NewTemplateDomainStud() TemplateDomainStub {
	return TemplateDomainStub{domains: []TemplateDomain{
		{
			id:   "1",
			Name: "Template Domain 1",
		}, {
			id:   "2",
			Name: "Template Domain 2",
		}, {
			id:   "3",
			Name: "Template Domain 3",
		}, {
			id:   "4",
			Name: "Template Domain 4",
		}, {
			id:   "5",
			Name: "Template Domain 5",
		},
	}}
}
