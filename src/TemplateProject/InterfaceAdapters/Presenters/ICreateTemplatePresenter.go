package Presenters

import (
	"TemplateSolution/src/Shared/Common/Generic"
)

type ICreateTemplatePresenter interface {
	Generic.IPresenter[int64, int64, error]
}
