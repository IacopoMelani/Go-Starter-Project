package dto

// DTO - Defines an interface for DTOs
type DTO interface {
	Validate() (bool, []error)
}

// Mappable - Defines an interface for a model that want translate himself to a dto
type Mappable interface {
	Map() DTO
}
