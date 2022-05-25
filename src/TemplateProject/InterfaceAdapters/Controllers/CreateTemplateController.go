package Controllers

import (
	"TemplateSolution/src/Shared/Models"
	"TemplateSolution/src/Shared/Utils"
	"TemplateSolution/src/TemplateProject/ApplicationBusinessRules/DTOs"
	"TemplateSolution/src/TemplateProject/ApplicationBusinessRules/UseCases/CreateTemplateUseCase"
	"TemplateSolution/src/TemplateProject/InterfaceAdapters/Presenters"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CreateTemplateController struct {
	ICreateTemplateController
	CreateTemplateInteractor CreateTemplateUseCase.ICreateTemplateInteractor
}

// NewCreateTemplateController - Injection Dependency
func NewCreateTemplateController(createTemplateUseCase CreateTemplateUseCase.ICreateTemplateInteractor) ICreateTemplateController {
	return &CreateTemplateController{CreateTemplateInteractor: createTemplateUseCase}
}

// Handle godoc
// @Summary Add a new item to the template list
// @ID create-template-id
// @Tags root
// @Accept json
// @Produce json
// @Success 201 {int} id
// @Param data body DTOs.CreateTemplateDTO true "data"
// @Router /api/templates/ [post]
func (iam *CreateTemplateController) Handle(ctx echo.Context) error {
	var requestData DTOs.CreateTemplateDTO
	var outputPort = Presenters.NewCreateTemplatePresenter()

	if err := ctx.Bind(&requestData); !errors.Is(err, nil) {
		return Utils.JSONHandleError(http.StatusBadRequest, Models.ErrorModel{Code: -1, Message: "Bad request."}, ctx)
	}

	Utils.InfoLogger.Println("Request: ", requestData)

	var input = CreateTemplateUseCase.CreateTemplateInputPort{
		RequestData: requestData, OutputPort: outputPort,
	}

	iam.CreateTemplateInteractor.Handle(&input)

	if !errors.Is(outputPort.GetError(), nil) {
		return Utils.JSONHandleError(http.StatusInternalServerError, Models.ErrorModel{Code: -1, Message: outputPort.GetError().Error()}, ctx)
	}

	Utils.InfoLogger.Println(outputPort.GetContent())
	return ctx.JSON(http.StatusCreated, outputPort.GetContent())
}
