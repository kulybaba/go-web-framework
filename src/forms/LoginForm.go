package forms

import (
	"strings"

	"github.com/asaskevich/govalidator"
)

type Login struct {
	Email    string `valid:"required, email"`
	Password string `valid:"required, length(8|20)"`
}

func (form *Login) Validate() []string {
	if _, err := govalidator.ValidateStruct(form); err != nil {
		return strings.Split(err.Error(), ";")
	}
	return nil
}
