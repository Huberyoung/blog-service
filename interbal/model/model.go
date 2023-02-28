package model

type Model struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	CreatedOn  uint   `json:"created_on"`
	CreatedBy  string `json:"created_by"`
	ModifiedOn uint   `json:"modified_on"`
	ModifiedBy string `json:"modified_by"`
	DeleteOn   uint   `json:"delete_on"`
	IsDel      uint8  `json:"is_del"`
}
