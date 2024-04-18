package entity

type Example struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Field1 string `json:"field_1"`
	Field2 int    `json:"field_2"`
	Field3 bool   `json:"field_3"`
}


type ExamplePostRequest struct {
	Field1 string `json:"field_1" binding:"required"`
	Field2 int    `json:"field_2" binding:"required"`
	Field3 bool   `json:"field_3" binding:"required"`
}