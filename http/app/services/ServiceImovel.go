package services

import (
	"log"
	"servidorhttp/app/database"
	"servidorhttp/app/models"
	"servidorhttp/app/repositories"
)

type ServiceImovel struct {
	repo repositories.IImovelRepository
}

func NewServiceImovel() ServiceImovel {
	db, err := database.Connection()
	if err != nil {
		log.Panic(err)
	}
	repo := repositories.NewImovelRepositoy(db)
	return ServiceImovel{repo: repo}
}

func (r *ServiceImovel) All() []models.Imovel {
	return r.repo.GetAll()
}

func (r *ServiceImovel) Get(id uint) models.Imovel {
	return r.repo.GetById(id)
}
