package repositories

import "servidorhttp/app/models"

type IImovelRepository interface {
	GetAll() []models.Imovel
	Create(imovel *models.Imovel)
	GetById(id uint) models.Imovel
	Update(imovel *models.Imovel)
	Delete(id uint)
}
