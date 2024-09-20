package usecase

import (
	"api_company/internal/domain/entity"
	"api_company/internal/domain/repository"
	"context"
)

type CompanyUsecase interface {
	GetAllCompanies(ctx context.Context) ([]entity.Company, error)
}

type companyUsecase struct {
	repo repository.CompanyRepository
}

func NewCompanyUsecase(repo repository.CompanyRepository) CompanyUsecase {
	return &companyUsecase{repo: repo}
}

func (u *companyUsecase) GetAllCompanies(ctx context.Context) ([]entity.Company, error) {
	return u.repo.GetAll(ctx)
}
