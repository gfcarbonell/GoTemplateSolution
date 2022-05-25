package SQLServer

import (
	"TemplateSolution/src/TemplateProject/EnterpriseBusinessRules/Entities"
	"TemplateSolution/src/TemplateProject/EnterpriseBusinessRules/Repositories"
	"database/sql"
	"errors"
	"fmt"
)

type CreateTemplateSQLServer struct {
	Repositories.ICreateTemplateRepository
	db *sql.DB
}

func NewCreateTemplateSQLServer(db *sql.DB) Repositories.ICreateTemplateRepository {
	return &CreateTemplateSQLServer{db: db}
}

func (iam *CreateTemplateSQLServer) Handle(entity *Entities.TemplateEntity) (*int64, error) {
	var id sql.NullInt64
	result := iam.db.QueryRow(`exec create_template @name=?`, entity.Name)

	if err := result.Scan(&id); err != nil {
		return nil, errors.New(err.Error())
	}

	if !id.Valid || id.Int64 <= 0 {
		return nil, errors.New("invalid ID")
	}

	fmt.Printf("Database Response as struct %+v\n", result)
	return &id.Int64, nil
}
