package storage

import (
	"database/sql"
	"fmt"
	"github.com/dannyvelas/examplego_api/apierror"
	"github.com/dannyvelas/examplego_api/models"
)

type AdminRepo struct {
	database Database
}

func NewAdminRepo(database Database) AdminRepo {
	return AdminRepo{database: database}
}

func (adminRepo AdminRepo) GetOne(id string) (models.Admin, error) {
	const query = `SELECT id, password FROM admins WHERE email = $1`

	var admin models.Admin
	err := adminRepo.database.driver.QueryRow(query, id).
		Scan(&admin.Id, &admin.Password)
	if err == sql.ErrNoRows {
		return models.Admin{}, fmt.Errorf("admin_repo: GetOne: %w", apierror.NotFound)
	} else if err != nil {
		return models.Admin{}, fmt.Errorf("admin_repo: GetOne: %v", WrapSQLError(ErrQueryScanOneRow, err))
	}

	return admin, nil
}
