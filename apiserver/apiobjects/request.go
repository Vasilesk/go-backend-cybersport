package apiobjects

// BaseRequest is a type for user request to the server
type BaseRequest struct {
	V       *float64  `json:"v"`
	ID      *uint64   `json:"id"`
	Limit   *uint64   `json:"limit"`
	Offset  *uint64   `json:"offset"`
	Players *[]Player `json:"players"`
	Teams   *[]Team   `json:"teams"`
}
