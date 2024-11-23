package repositories

import (
	"log"
	"servidorhttp/app/models"

	"gorm.io/gorm"
)

type ImovelRepository struct {
	gorm.Model
	db *gorm.DB
}

func NewImovelRepositoy(db *gorm.DB) IImovelRepository {
	return &ImovelRepository{db: db}
}

func (r *ImovelRepository) Create(imovel *models.Imovel) {
	r.db.Create(&imovel)
}

func (r *ImovelRepository) Update(imovel *models.Imovel) {
	r.db.Save(&imovel)
}

func (r *ImovelRepository) Delete(id uint) {
	var imovel models.Imovel
	if err := r.db.Delete(&imovel, id).Error; err != nil {
		log.Panic(err)
	}
}

func (r *ImovelRepository) GetById(id uint) models.Imovel {
	var imovel models.Imovel
	if err := r.db.Find(&imovel, id).Error; err != nil {
		log.Panic(err)
	}
	return imovel
}

func (r *ImovelRepository) GetAll() []models.Imovel {
	var imoveis []models.Imovel
	if err := r.db.Find(&imoveis).Error; err != nil {
		log.Panic(err)
	}
	return imoveis
}
