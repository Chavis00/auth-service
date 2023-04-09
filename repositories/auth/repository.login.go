package repositories

import (
	"authservice/models"
	"authservice/pkg"
	"authservice/schemas"
	"net/http"

	"gorm.io/gorm"
)

type RepositoryLogin interface {
	LoginRepository(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError)
}

type repositoryLogin struct {
	db *gorm.DB
}

func NewRepositoryLogin(db *gorm.DB) *repositoryLogin {
	return &repositoryLogin{db: db}
}

func (r *repositoryLogin) LoginRepository(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError) {
	var user models.ModelAuth
	db := r.db.Model(&user)

	user.Email = input.Email
	user.Password = input.Password
	checkUserAccount := db.Debug().First(&user, "email = ?", input.Email)

	if checkUserAccount.RowsAffected < 1 {
		return &user, schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
	}

	if !user.Active {
		return &user, schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
	}

	comparePassword := pkg.ComparePassword(user.Password, input.Password)

	if comparePassword != nil {
		return &user, schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_03",
		}
	}

	return &user, schemas.SchemaDatabaseError{}
}
