package model

type Error struct {
	errorCode    string
	status       int
	errorMessage string
}

func NewError() *Error {
	return &Error{}
}

func (errorStruct *Error) GetErrorCode() string {
	return errorStruct.errorCode
}

func (errorStruct *Error) GetStatus() int {
	return errorStruct.status
}

func (errorStruct *Error) GetErrorMessage() string {
	return errorStruct.errorMessage
}

func (errorStruct *Error) SetErrorCode(errorCode string) {
	errorStruct.errorCode = errorCode
}

func (errorStruct *Error) SetStatus(status int) {
	errorStruct.status = status
}

func (errorStruct *Error) SetErrorMessage(errorMessage string) {
	errorStruct.errorMessage = errorMessage
}

func (errorStruct *Error) Get() []any {
	return []any{errorStruct.errorCode, errorStruct.status, errorStruct.errorMessage}
}

func (errorStruct *Error) Set(errorCode string, errorStatus int, errorMessage string) *Error {
	return &Error{
		errorCode:    errorCode,
		status:       errorStatus,
		errorMessage: errorMessage,
	}
}
