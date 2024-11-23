package models

type Imovel struct {
	Id       int     `json:"id" gorm:"column:id"`
	Tipo     string  `json:"tipo" gorm:"column:tipo"`
	Preco    float32 `json:"preco" gorm:"column:preco"`
	Endereco string  `json:"endereco" gorm:"column:endereco"`
	Status   string  `json:"status" gorm:"column:status"`
}

func (r Imovel) TableName() string {
	return "imovel"
}
