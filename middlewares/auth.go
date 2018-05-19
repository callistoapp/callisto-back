package middlewares

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"callisto/models"
	"github.com/gorilla/context"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.RequestURI != "/graphql" {
			next.ServeHTTP(w, r)
			return
		}
		cookie, err := r.Cookie("connect.sid")
		if err != nil || cookie == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		client := &http.Client{}
		if req, err := http.NewRequest("GET", "http://auth:3001/profile", nil); err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		} else {
			req.AddCookie(cookie)
			resp, err := client.Do(req)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			body, _ := ioutil.ReadAll(resp.Body)
			var loggedUser models.AuthenticatedUser
			json.Unmarshal([]byte(body), &loggedUser)
			context.Set(r,"loggedUser", loggedUser)
			next.ServeHTTP(w, r)
		}
	})
}