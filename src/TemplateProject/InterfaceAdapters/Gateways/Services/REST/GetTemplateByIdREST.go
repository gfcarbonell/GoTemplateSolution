package REST

import (
	"TemplateSolution/src/Shared/Common/Constants"
	"TemplateSolution/src/Shared/Utils"
	"TemplateSolution/src/TemplateProject/EnterpriseBusinessRules/Models"
	"TemplateSolution/src/TemplateProject/EnterpriseBusinessRules/Proxies"
	"encoding/json"
	"fmt"
)

type GetTemplateByIdREST struct {
	Proxies.IGetTemplateByIdProxy
}

func NewGetTemplateByIdREST() Proxies.IGetTemplateByIdProxy {
	return &GetTemplateByIdREST{}
}

func (iam *GetTemplateByIdREST) Handle(input int64) (*Models.TemplateModel, error) {
	url := fmt.Sprintf("%s/%d", Utils.SETTING.URL.Templates, input)
	var bodyBytes, _, err = Utils.HttpClient(Constants.MethodGet, url, nil, nil)
	var responseObject Models.TemplateModel

	if err != nil {
		return nil, err
	}

	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("API Response as struct %+v\n", responseObject)

	return &responseObject, nil
}
