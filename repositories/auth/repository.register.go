package repositories

import (
	"authservice/models"
	"authservice/schemas"
	"net/http"

	"gorm.io/gorm"
)

type RepositoryRegister interface {
	RegisterRepository(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError)
}

type repositoryRegister struct {
	db *gorm.DB
}

func NewRepositoryRegister(db *gorm.DB) *repositoryRegister {
	return &repositoryRegister{db: db}
}

func (r *repositoryRegister) RegisterRepository(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError) {
	var user models.ModelAuth

	db := r.db.Model(&user)

	checkUserAccount := db.Debug().First(&user, "email = ?", input.Email)
	if checkUserAccount.RowsAffected > 0 {
		return &user, schemas.SchemaDatabaseError{
			Code: http.StatusConflict,
			Type: "error_01",
		}
	}

	user.Fullname = input.Fullname
	user.Email = input.Email
	user.Password = input.Password

	addNewUser := db.Debug().Create(&user).Commit()
	if addNewUser.RowsAffected < 1 {
		return &user, schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
	}

	return &user, schemas.SchemaDatabaseError{}
}
