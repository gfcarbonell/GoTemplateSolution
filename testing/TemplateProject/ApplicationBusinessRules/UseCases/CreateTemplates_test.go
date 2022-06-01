package UseCases

import (
	"TemplateSolution/src/Shared/Constants"
	"TemplateSolution/src/TemplateProject/ApplicationBusinessRules/DTOs"
	"TemplateSolution/src/TemplateProject/ApplicationBusinessRules/UseCases/CreateTemplateUseCase"
	"TemplateSolution/src/TemplateProject/EnterpriseBusinessRules/Entities"
	"TemplateSolution/src/TemplateProject/EnterpriseBusinessRules/Repositories/mocks"
	mockPresenters "TemplateSolution/src/TemplateProject/InterfaceAdapters/Presenters/mocks"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestCreateTemplateUseCase(t *testing.T) {
	testCases := []struct {
		Name                  string
		RequestData           DTOs.CreateTemplateDTO
		ResponseData          int64
		ExpectValueRepository *int64
		ExpectContent         int64
		ExpectedError         error
	}{
		{
			Name: "CreateTemplateSucess",
			RequestData: DTOs.CreateTemplateDTO{
				Name: "Gian",
			},
			ResponseData:          int64(1),
			ExpectValueRepository: func() *int64 { i := int64(1); return &i }(),
			ExpectContent:         int64(1),
			ExpectedError:         nil,
		},
		{
			Name: "CreateTemplateGenerateInvalidId",
			RequestData: DTOs.CreateTemplateDTO{
				Name: "Gian",
			},
			ResponseData:          int64(-1),
			ExpectValueRepository: func() *int64 { i := int64(-1); return &i }(),
			ExpectContent:         int64(1),
			ExpectedError:         Constants.ErrGenerateInvalidId,
		},
		{
			Name: "CreateTemplateGenerateBadRequest",
			RequestData: DTOs.CreateTemplateDTO{
				Name: "",
			},
			ResponseData:          int64(-1),
			ExpectValueRepository: func() *int64 { i := int64(-1); return &i }(),
			ExpectContent:         int64(1),
			ExpectedError:         Constants.ErrValidation,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			// Mock Presenters
			outputPort := mockPresenters.ICreateTemplatePresenter{}
			outputPort.On("Handle", tc.ResponseData, tc.ExpectedError)
			outputPort.On("GetContent").Return(tc.ExpectContent)
			outputPort.On("GetError").Return(tc.ExpectedError)

			// Instance Input Port
			var inputPort = &CreateTemplateUseCase.CreateTemplateInputPort{
				RequestData: tc.RequestData, OutputPort: &outputPort,
			}

			createTemplateRepository := &mockRepositories.ICreateTemplateRepository{}
			createTemplateRepository.On("Handle", Entities.TemplateEntity{Name: "Gian"}).Return(tc.ExpectValueRepository, tc.ExpectedError)

			// Mock Use Case
			createTemplateUseCase := CreateTemplateUseCase.NewCreateTemplateInteractor(createTemplateRepository)
			createTemplateUseCase.Handle(inputPort)

			if inputPort.GetOutputPort().GetError() == nil {
				assert.Nil(t, inputPort.GetOutputPort().GetError(), "Error should be nil")
				assert.Equal(t, tc.ResponseData, outputPort.GetContent(), "expected 1, got 1")
			} else if inputPort.GetOutputPort().GetError() == Constants.ErrGenerateInvalidId {
				assert.NotNil(t, inputPort.GetOutputPort().GetError(), "Error should not be nil")
				assert.NotEqual(t, tc.ResponseData, outputPort.GetContent(), "expected 1, got -1")
			} else {
				assert.NotSame(t, tc.RequestData, "Vlaidation")
			}
		})
	}
}
