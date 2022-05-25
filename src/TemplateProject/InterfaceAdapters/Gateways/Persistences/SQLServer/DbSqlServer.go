package SQLServer

import (
	"TemplateSolution/src/Shared/Utils"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

type IDbSQL[T any] interface {
	open() T
	close()
	getConnectionString() string
}

type DbSQLServer struct {
	IDbSQL[*sql.DB]
	Db *sql.DB
}

func (iam *DbSQLServer) Open() (*sql.DB, error) {
	db, err := sql.Open("mssql", iam.GetConnectionString())

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (iam *DbSQLServer) Close() {
	defer iam.Db.Close()
}

func (iam *DbSQLServer) GetConnectionString() string {
	return fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		Utils.SETTING.Database.Host,
		Utils.SETTING.Database.Username,
		Utils.SETTING.Database.Password,
		Utils.SETTING.Database.Port,
		Utils.SETTING.Database.DatabaseName)
}
