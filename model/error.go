package model

type Error struct {
	ErrorCode    string
	Status       int
	ErrorMessage string
}

func (errorStruct *Error) GetErrorCode() string {
	return errorStruct.ErrorCode
}

func (errorStruct *Error) GetStatus() int {
	return errorStruct.Status
}

func (errorStruct *Error) GetErrorMessage() string {
	return errorStruct.ErrorMessage
}

func (errorStruct *Error) SetErrorCode(errorCode string) {
	errorStruct.ErrorCode = errorCode
}

func (errorStruct *Error) SetStatus(status int) {
	errorStruct.Status = status
}

func (errorStruct *Error) SetErrorMessage(errorMessage string) {
	errorStruct.ErrorMessage = errorMessage
}

func (errorStruct *Error) Get() []any {
	return []any{errorStruct.ErrorCode, errorStruct.Status, errorStruct.ErrorMessage}
}
