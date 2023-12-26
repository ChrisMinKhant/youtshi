package model

type SuccessResponse struct {
	Status  int
	Message string
	Data    any
}

func (response *SuccessResponse) GetStatus() int {
	return response.Status
}

func (response *SuccessResponse) GetMessage() string {
	return response.Message
}

func (response *SuccessResponse) GetData() any {
	return response.Data
}

func (response *SuccessResponse) SetStatus(status int) {
	response.Status = status
}

func (response *SuccessResponse) SetMessage(message string) {
	response.Message = message
}

func (response *SuccessResponse) SetData(data any) {
	response.Data = data
}

func (response *SuccessResponse) Get() []any {
	return []any{response.Status, response.Message, response.Data}
}
