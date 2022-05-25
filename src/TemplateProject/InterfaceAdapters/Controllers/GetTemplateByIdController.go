package Controllers

import (
	"TemplateSolution/src/Shared/Models"
	"TemplateSolution/src/Shared/Utils"
	"TemplateSolution/src/TemplateProject/ApplicationBusinessRules/UseCases/GetTemplateByIdUseCase"
	"TemplateSolution/src/TemplateProject/InterfaceAdapters/Presenters"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type GetTemplateByIdController struct {
	IGetTemplateByIdController
	GetTemplateByIdInteractor GetTemplateByIdUseCase.IGetTemplateByIdInteractor
}

// NewGetTemplateByIdController - Injection Dependency
func NewGetTemplateByIdController(getTemplateByIdUseCase GetTemplateByIdUseCase.IGetTemplateByIdInteractor) IGetTemplateByIdController {
	return &GetTemplateByIdController{GetTemplateByIdInteractor: getTemplateByIdUseCase}
}

// Handle godoc
// @Summary get a template by ID
// @ID get-template-by-id
// @Tags root
// @Accept json
// @Produce json
// @Success 200 {ReadTemplateDTO} content
// @Param id path int true "id"
// @Router /api/templates/{id} [get]
func (iam *GetTemplateByIdController) Handle(ctx echo.Context) error {
	var outputPort = Presenters.NewGetTemplateByIdPresenter()

	requestData, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if !errors.Is(err, nil) && requestData < 0 {
		return Utils.JSONHandleError(http.StatusBadRequest, Models.ErrorModel{Code: -1, Message: "Invalid Id"}, ctx)
	}

	Utils.InfoLogger.Println("Request: ", requestData)

	var input = GetTemplateByIdUseCase.GetTemplateByIdInputPort{
		RequestData: requestData, OutputPort: outputPort,
	}

	iam.GetTemplateByIdInteractor.Handle(&input)

	if !errors.Is(outputPort.GetError(), nil) {
		return Utils.JSONHandleError(http.StatusInternalServerError, Models.ErrorModel{Code: -1, Message: outputPort.GetError().Error()}, ctx)
	}

	Utils.InfoLogger.Println(outputPort.GetContent())

	return ctx.JSON(http.StatusCreated, outputPort.GetContent())
}
