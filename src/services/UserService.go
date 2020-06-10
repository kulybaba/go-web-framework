package services

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/petrokulybaba/go-web-framework/configs"
	"github.com/petrokulybaba/go-web-framework/src/forms"
	"github.com/petrokulybaba/go-web-framework/src/models"
	"github.com/petrokulybaba/go-web-framework/src/repositories"
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

		token := GenerateToken()
		http.SetCookie(w, &http.Cookie{
			Name:     configs.SESSION_COOKIE_NAME,
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 24 * 30),
			HttpOnly: true,
		})

		err = RedisSet(token, user.ID)
		if err != nil {
			log.Fatal(err)
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

		token := GenerateToken()
		http.SetCookie(w, &http.Cookie{
			Name:     configs.SESSION_COOKIE_NAME,
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 24 * 30),
			HttpOnly: true,
		})

		err = RedisSet(token, user.ID)
		if err != nil {
			log.Fatal(err)
		}

		http.Redirect(w, r, configs.Routes["index"]["path"], http.StatusFound)
	}
}

func Logout(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     configs.SESSION_COOKIE_NAME,
		Expires:  time.Now(),
		MaxAge:   -1,
		HttpOnly: true,
	})

	err := RedisDelete(configs.SESSION_COOKIE_NAME)
	if err != nil {
		log.Fatal(err)
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

func GenerateToken() string {
	uuid, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	return uuid.String()
}
