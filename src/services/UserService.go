package services

import (
	"fmt"
	"time"
	"net/http"
	"golang.org/x/crypto/bcrypt"

	"github.com/petrokulybaba/go-basic-framework/configs"
	"github.com/petrokulybaba/go-basic-framework/src/forms"
	"github.com/petrokulybaba/go-basic-framework/src/models"
	"github.com/petrokulybaba/go-basic-framework/src/repositories"
)

func Login(w http.ResponseWriter, r *http.Request, loginForm forms.Login) {
	if err := loginForm.Validate(); err != nil {
		RenderTemplate(w, configs.Routes["login"]["name"], map[string]interface{}{
			"errors": err,
		})
	} else {
		var errors []string
		userRepository := repositories.UserRepository{}
		user, err := userRepository.FindOneBy(models.User{Email: loginForm.Email})
		if err != nil {
			errors = append(errors, "Wrong email")
		}

		if !CheckPasswordHash(loginForm.Password, user.Password) {
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

func Registration(w http.ResponseWriter, r *http.Request, registrationForm forms.Registration) {
	if err := registrationForm.Validate(); err != nil {
		RenderTemplate(w, configs.Routes["registration"]["name"], map[string]interface{}{
			"errors": err,
		})
	} else {
		userRepository := repositories.UserRepository{}
		if _, err := userRepository.FindOneBy(models.User{Email: registrationForm.Email}); err == nil {
			RenderTemplate(w, configs.Routes["registration"]["name"], map[string]interface{}{
				"errors": []string{fmt.Sprintf("Email %s already taken", registrationForm.Email)},
			})
			return
		}

		hash, err := HashPassword(registrationForm.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		user := models.User{
			FirstName: registrationForm.FirstName,
			LastName:  registrationForm.LastName,
			Email:     registrationForm.Email,
			Password:  hash,
			Created:   time.Now(),
			Updated:   time.Now(),
		}

		err = userRepository.Create(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
