package apiobjects

// Player is a type for player
type Player struct {
	ID          *uint64  `json:"id"`
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	LogoLink    *string  `json:"logo_link"`
	Rating      *float64 `json:"rating"`
}
