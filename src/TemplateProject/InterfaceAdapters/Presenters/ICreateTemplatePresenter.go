package Presenters

import (
	"TemplateSolution/src/Shared/Common/Generic"
)

//go:generate mockery --name=ICreateTemplatePresenter --filename=CreateTemplatePresenterMock.go --outpkg=mockPresenters
type ICreateTemplatePresenter interface {
	Generic.IPresenter[int64, int64, error]
}
