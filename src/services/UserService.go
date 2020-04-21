package services

import (
	"net/http"
	"golang.org/x/crypto/bcrypt"

	"github.com/petrokulybaba/go-basic-framework/configs"
	"github.com/petrokulybaba/go-basic-framework/src/forms"
	"github.com/petrokulybaba/go-basic-framework/src/models"
	"github.com/petrokulybaba/go-basic-framework/src/repositories"
)

func Login(w http.ResponseWriter, r *http.Request, login forms.Login) {
	if err := login.Validate(); err != nil {
		RenderTemplate(w, configs.Routes["login"]["name"], map[string]interface{}{
			"errors": err,
		})
	} else {
		var errors []string
		userRepository := repositories.UserRepository{}
		user, err := userRepository.FindOneBy(models.User{Email: login.Email})
		if err != nil {
			errors = append(errors, "Wrong email")
		}

		if !CheckPasswordHash(login.Password, user.Password) {
			errors = append(errors, "Wrong password")
		}

		if errors != nil {
			RenderTemplate(w, configs.Routes["login"]["name"], map[string]interface{}{
				"errors": errors,
			})
		}

		http.Redirect(w, r, configs.Routes["index"]["path"], http.StatusFound)
	}
}

func HashPassword(password string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(res), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
