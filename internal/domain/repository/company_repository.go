package repository

import (
	"api_company/internal/domain/entity"
	"context"
	"database/sql"
)

type CompanyRepository interface {
	GetAll(ctx context.Context) ([]entity.Company, error)
}

type companyRepository struct {
	db *sql.DB
}

func NewCompanyRepository(db *sql.DB) CompanyRepository {
	return &companyRepository{db: db}
}

func (r *companyRepository) GetAll(ctx context.Context) ([]entity.Company, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT ID, CNPJ, CompanyName, FantasyName FROM company where CNPJ is not null and FantasyName is not null and CompanyName is not null ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []entity.Company
	for rows.Next() {
		var company entity.Company
		if err := rows.Scan(&company.ID, &company.CNPJ, &company.CompanyName, &company.FantasyName); err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}
