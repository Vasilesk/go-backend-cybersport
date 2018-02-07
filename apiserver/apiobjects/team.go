package apiobjects

// Team is a type for a team
type Team struct {
	ID          *uint64  `json:"id"`
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	LogoLink    *string  `json:"logo_link"`
	Rating      *float64 `json:"rating"`
}
