package Presenters

type CreateTemplatePresenter struct {
	ICreateTemplatePresenter
	content int64
	err     error
}

func NewCreateTemplatePresenter() ICreateTemplatePresenter {
	return &CreateTemplatePresenter{}
}

func (iam *CreateTemplatePresenter) Handle(response int64, err error) {
	iam.content = response
	iam.err = err
}

func (iam *CreateTemplatePresenter) GetContent() int64 {
	return iam.content
}

func (iam *CreateTemplatePresenter) GetError() error {
	return iam.err
}
