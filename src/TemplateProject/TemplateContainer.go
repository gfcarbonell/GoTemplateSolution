package TemplateProject

import (
	"TemplateSolution/src/TemplateProject/ApplicationBusinessRules/Ports/CreateTemplatePort"
	"TemplateSolution/src/TemplateProject/ApplicationBusinessRules/Ports/GetTemplateByIdPort"
	"TemplateSolution/src/TemplateProject/ApplicationBusinessRules/UseCases/CreateTemplateUseCase"
	"TemplateSolution/src/TemplateProject/ApplicationBusinessRules/UseCases/GetTemplateByIdUseCase"
	"TemplateSolution/src/TemplateProject/EnterpriseBusinessRules/Proxies"
	"TemplateSolution/src/TemplateProject/EnterpriseBusinessRules/Repositories"
	"TemplateSolution/src/TemplateProject/InterfaceAdapters/Controllers"
	"TemplateSolution/src/TemplateProject/InterfaceAdapters/Gateways/Persistences/SQLServer"
	"TemplateSolution/src/TemplateProject/InterfaceAdapters/Gateways/Services/REST"
	"TemplateSolution/src/TemplateProject/InterfaceAdapters/Presenters"
	"database/sql"
)

type ITemplateContainer interface {
	NewTemplateController() Controllers.TemplateController
}

type TemplateContainer struct {
	Db *sql.DB
}

// NewTemplateContainer - Injection Dependency
func NewTemplateContainer(db *sql.DB) ITemplateContainer {
	return &TemplateContainer{Db: db}
}

func (iam *TemplateContainer) NewTemplateController() Controllers.TemplateController {
	return Controllers.TemplateController{
		CreateTemplateController:  iam.NewCreateTemplateController(),
		GetTemplateByIdController: iam.NewGetTemplateByIdController(),
	}
}

func (iam *TemplateContainer) NewCreateTemplateController() Controllers.ICreateTemplateController {
	return Controllers.NewCreateTemplateController(iam.NewCreateTemplateInteractor())
}

func (iam *TemplateContainer) NewCreateTemplateInteractor() CreateTemplateUseCase.ICreateTemplateInteractor {
	return CreateTemplateUseCase.NewCreateTemplateInteractor(iam.NewCreateTemplateSQLServer())
}

func (iam *TemplateContainer) NewCreateTemplateInputPort() CreateTemplatePort.ICreateTemplateInputPort {
	return CreateTemplateUseCase.NewCreateTemplateInputPort(iam.NewCreateTemplateInputPort().GetRequestData(), iam.NewCreateTemplateInputPort().GetOutputPort())
}

func (iam *TemplateContainer) NewCreateTemplatePresenter() Presenters.ICreateTemplatePresenter {
	return Presenters.NewCreateTemplatePresenter()
}

func (iam *TemplateContainer) NewCreateTemplateSQLServer() Repositories.ICreateTemplateRepository {
	return SQLServer.NewCreateTemplateSQLServer(iam.Db)
}

func (iam *TemplateContainer) NewGetTemplateByIdController() Controllers.IGetTemplateByIdController {
	return Controllers.NewGetTemplateByIdController(iam.NewGetTemplateByIdInteractor())
}

func (iam *TemplateContainer) NewGetTemplateByIdInteractor() GetTemplateByIdUseCase.IGetTemplateByIdInteractor {
	return GetTemplateByIdUseCase.NewGetTemplateByIdInteractor(iam.NewGetTemplateByIdREST(), iam.NewGetTemplateByIdSQLServer())
}

func (iam *TemplateContainer) NewGetTemplateByIdInputPort() GetTemplateByIdPort.IGetTemplateByIdInputPort {
	return GetTemplateByIdUseCase.NewGetTemplateByIdInputPort(iam.NewGetTemplateByIdInputPort().GetRequestData(), iam.NewGetTemplateByIdInputPort().GetOutputPort())
}

func (iam *TemplateContainer) NewGetTemplateByIdPresenter() Presenters.IGetTemplateByIdPresenter {
	return Presenters.NewGetTemplateByIdPresenter()
}

func (iam *TemplateContainer) NewGetTemplateByIdREST() Proxies.IGetTemplateByIdProxy {
	return REST.NewGetTemplateByIdREST()
}

func (iam *TemplateContainer) NewGetTemplateByIdSQLServer() Repositories.IGetTemplateByIdRepository {
	return SQLServer.NewGetTemplateByIdSQLServer(iam.Db)
}
