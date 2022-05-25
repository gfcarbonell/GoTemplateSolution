package CreateTemplatePort

import "TemplateSolution/src/Shared/Common/Generic"

type ICreateTemplateOutputPort interface {
	Generic.IOutputPort[int64, error]
}
