package SQLServer

import (
	"TemplateSolution/src/TemplateProject/EnterpriseBusinessRules/Entities"
	"TemplateSolution/src/TemplateProject/EnterpriseBusinessRules/Repositories"
	"database/sql"
	"errors"
	"fmt"
)

type GetTemplateByIdSQLServer struct {
	Repositories.IGetTemplateByIdRepository
	db *sql.DB
}

func NewGetTemplateByIdSQLServer(db *sql.DB) Repositories.IGetTemplateByIdRepository {
	return &GetTemplateByIdSQLServer{db: db}
}

func (iam *GetTemplateByIdSQLServer) Handle(entity int64) (*Entities.TemplateEntity, error) {
	var entitySQL Entities.TemplateEntity
	result := iam.db.QueryRow(`exec get_template_by_id @id=?`, entity)

	if err := result.Scan(&entitySQL.Id, &entitySQL.Name); err != nil {
		return nil, errors.New(err.Error())
	}

	fmt.Printf("Database Response as struct %+v\n", result)
	return &entitySQL, nil
}
