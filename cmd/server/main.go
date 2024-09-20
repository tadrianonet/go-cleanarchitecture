package main

import (
	_ "api_company/docs"
	"api_company/internal/domain/repository"
	"api_company/internal/infrastructure/db"
	"api_company/internal/interface/api"
	"api_company/internal/usecase"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	companyRepo := repository.NewCompanyRepository(db)
	companyUsecase := usecase.NewCompanyUsecase(companyRepo)
	companyHandler := api.NewCompanyHandler(companyUsecase)

	r := mux.NewRouter()
	r.HandleFunc("/company", companyHandler.GetAll).Methods("GET")
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Servidor rodando na porta 8080")
	http.ListenAndServe(":8080", r)
}
