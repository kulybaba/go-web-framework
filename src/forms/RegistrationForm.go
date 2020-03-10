package forms

import (
	"strings"

	"github.com/asaskevich/govalidator"
)

type Registration struct {
	FirstName      string `valid:"required, length(2|30)"`
	LastName       string `valid:"required, length(2|30)"`
	Email          string `valid:"required, email"`
	Password       string `valid:"required, length(8|20)"`
	RepeatPassword string `valid:"required, length(8|20)"`
}

func (form *Registration) Validate() []string {
	if _, err := govalidator.ValidateStruct(form); err != nil {
		return strings.Split(err.Error(), ";")
	}
	return nil
}
