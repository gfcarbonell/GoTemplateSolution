package GetTemplateByIdPort

import (
	"TemplateSolution/src/Shared/Common/Generic"
)

type IGetTemplateByIdInputPort interface {
	Generic.IInputPort[int64, IGetTemplateByIdOutputPort]
}
