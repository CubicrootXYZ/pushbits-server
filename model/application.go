package model

// Application holds information like the name, the token, and the associated user of an application.
type Application struct {
	ID       uint   `gorm:"AUTO_INCREMENT;primary_key" json:"id"`
	Token    string `gorm:"type:string;size:64;unique" json:"token"`
	UserID   uint   `json:"-"`
	Name     string `gorm:"type:string" json:"name"`
	MatrixID string `gorm:"type:string" json:"-"`
}

// CreateApplication is used to process queries for creating applications.
type CreateApplication struct {
	Name string `form:"name" query:"name" json:"name" binding:"required"`
}

// UpdateApplication is used to process queries for updating applications.
type UpdateApplication struct {
	Name string `json:"name"`
}
