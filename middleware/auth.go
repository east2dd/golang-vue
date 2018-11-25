package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/xyingsoft/golang-vue/models"
	u "github.com/xyingsoft/golang-vue/utils"
)

var JwtAuthentication = func(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{}
		requestPath := r.URL.Path

		for _, value := range notAuth {
			match, _ := regexp.MatchString(value, requestPath)
			if match {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			response = u.Message(false, "Missing auth token")
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response, http.StatusForbidden)
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response, http.StatusForbidden)
			return
		}

		tokenPart := splitted[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("TOKEN_PASSWORD")), nil
		})

		if err != nil { //Malformed token, returns with http code 403 as usual
			response = u.Message(false, "Malformed authentication token")
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response, http.StatusForbidden)
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			response = u.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response, http.StatusForbidden)
			return
		}

		fmt.Sprintf("User %", tk.Username)
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
}
