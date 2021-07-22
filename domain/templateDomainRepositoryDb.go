package domain

import (
	"database/sql"

	"github.com/DTan13/template/errs"
)

type TemplateDomainRepositoryDb struct {
	client *sql.DB
}

func (d TemplateDomainRepositoryDb) GetAllDomains() ([]TemplateDomain, *errs.TemplateError) {
	// Operations on d
	return []TemplateDomain{}, nil
}

func NewTemplateDomainRepositoryDb(client *sql.DB) TemplateDomainRepositoryDb {
	return TemplateDomainRepositoryDb{client: client}
}
