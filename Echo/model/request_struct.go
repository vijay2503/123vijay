package model

type Login struct {
	Name   string `json:"name"`
	RollNo string `json:"roll_no"`
}
type StudentDetails struct {
	Name       string `json:"name" validate:"required" gorm:"type:varchar(250);column:name"`
	RollNo     int    `json:"roll_no" validate:"required" gorm:"type:varchar(250);column:roll_no"`
	Department string `json:"department" validate:"required" gorm:"type:varchar(250);column:department"`
	Email      string `json:"email,omitempty" validate:"omitempty,email" gorm:"type:varchar(250);column:email"`
	Phone      string `json:"phone,omitempty" validate:"required,min=10,max=10,numeric" gorm:"type:varchar(250);column:phone"`
	Country    string `json:"country,omitempty" gorm:"type:varchar(250);column:country"`
	State      string `json:"state,omitempty" gorm:"type:varchar(250);column:state"`
	City       string `json:"city,omitempty" gorm:"type:varchar(250);column:city"`
}
