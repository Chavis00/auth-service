package repositories

import (
	"authservice/models"
	"authservice/schemas"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type RepositoryActivation interface {
	ActivationRepository(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError)
}

type repositoryActivation struct {
	db *gorm.DB
}

func NewRepositoryActivation(db *gorm.DB) *repositoryActivation {
	return &repositoryActivation{db: db}
}

func (r *repositoryActivation) ActivationRepository(input *schemas.SchemaAuth) (*models.ModelAuth, schemas.SchemaDatabaseError) {

	var user models.ModelAuth
	db := r.db.Model(&user)

	user.Email = input.Email

	checkUserAccount := db.Debug().First(&user, "email = ?", input.Email)

	if checkUserAccount.RowsAffected < 1 {
		return &user, schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
	}

	db.Debug().First(&user, "active = ?", input.Active)

	if user.Active {
		return &user, schemas.SchemaDatabaseError{
			Code: http.StatusBadRequest,
			Type: "error_02",
		}
	}

	user.Active = input.Active
	user.UpdatedAt = time.Now().Local()
	updateActivation := db.Debug().Where("email = ?", input.Email).Updates(user)

	if updateActivation.RowsAffected < 1 {
		return &user, schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_03",
		}
	}

	return &user, schemas.SchemaDatabaseError{}
}
