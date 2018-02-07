package apiobjects

// Tournament is a type for a tournament
type Tournament struct {
	ID          *uint64 `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	LogoLink    *string `json:"logo_link"`
	IsActive    *bool   `json:"is_active"`
	GameID      *uint64 `json:"game_id"`
}
