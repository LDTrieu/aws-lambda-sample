package model

type FaError struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	LangCode string `json:"langCode"`
	Err      error  `json:"-"`
}

func NewFaErr(code int, message string, err error) *FaError {
	return &FaError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func (ins *FaError) Error() string {
	return ins.Err.Error()
}

func (ins *FaError) ToString() string {
	return ins.Message
}

func (ins *FaError) Update(code int, err error) {
	ins.Code = code
	ins.Message = err.Error()
	ins.Err = err
}
