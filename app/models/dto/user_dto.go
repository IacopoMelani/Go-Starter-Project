package dto

import (
	"errors"

	"gopkg.in/guregu/null.v4"
)

// UserDTO - Defines user dto
type UserDTO struct {
	ID       int64       `json:"id"`
	Name     null.String `json:"name"`
	Lastname null.String `json:"lastname"`
	Gender   null.String `json:"gender"`
}

// MapUserDTO - Maps dto from a Mappable interface
func MapUserDTO(m Mappable) UserDTO {
	return m.Map().(UserDTO)
}

// Validate - Valididates the user dto
func (u UserDTO) Validate() (bool, []error) {

	valid := true
	errs := []error{}

	if u.Name.IsZero() {
		valid = false
		errs = append(errs, errors.New("missing name"))
	}

	if u.Lastname.IsZero() {
		valid = false
		errs = append(errs, errors.New("missing lastname"))
	}

	if u.Gender.IsZero() {
		valid = false
		errs = append(errs, errors.New("gender missing"))
	}

	return valid, errs
}
