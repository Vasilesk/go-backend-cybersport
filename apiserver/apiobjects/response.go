package apiobjects

// IResponse is an interface for server response
type IResponse interface {
	GetStatus() *string
	GetError() *string
	GetData() *map[string]interface{}
}

// BaseResponse is a type for server answer to user request
type BaseResponse struct {
	Data *map[string]interface{}
}

// GetStatus () of BaseResponse always returns "ok"
func (resp BaseResponse) GetStatus() *string {
	status := "ok"
	return &status
}

// GetError () of BaseResponse always returns nil
func (resp BaseResponse) GetError() *string {
	return nil
}

// GetData () of BaseResponse returns stored data
func (resp BaseResponse) GetData() *map[string]interface{} {
	return resp.Data
}

// ErrorResponse is a type for server answer to user request
type ErrorResponse struct {
	Error *string `json:"error,omitempty"`
}

// GetStatus () of ErrorResponse always returns "error"
func (resp ErrorResponse) GetStatus() *string {
	status := "error"
	return &status
}

// GetError () of ErrorResponse returns stored error
func (resp ErrorResponse) GetError() *string {
	return resp.Error
}

// GetData () of ErrorResponse always returns nil
func (resp ErrorResponse) GetData() *map[string]interface{} {
	return nil
}
