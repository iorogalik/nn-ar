package requests

type OrganizationRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	City        string  `json:"city" validate:"required"`
	Address     string  `json:"address" validate:"required"`
	Lat         float64 `json:"lat" validate:"required"`
	Lon         float64 `json:"lon" validate:"required"`
}
