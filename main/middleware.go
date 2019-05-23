package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
)

const (
	notAuthorized, serverName, basic = "Not Authorized for this request.", "Authenticator", "Basic"
)

// Only allow authenticated requests. This currently only supports
// HTTP Basic. Reject any requests that do do contain a valid
// Authorization header username and password, Base62 encoded.
func (s *server) authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a := r.Header.Get("Authorization")
		if len(a) == 0 {
			fail(w, notAuthorized, 401)
		} else {
			u, p, err := s.extractCredentials(r)
			if err != nil {
				fail(w, err.Error(), 401)
			} else {
				_, err := getByUsernameAndPassword(u, p, true, s)
				if err != nil {
					fail(w, notAuthorized, 401)
				} else {
					next.ServeHTTP(w, r)
				}
			}
		}
	})
}

// Only admin users can:
// 1. CRUD any admin user.
// 2. CRUD any application.
// 3. List all applications.

// A non-admin API user can:
// 1. Only CRUD their own application.
// 2. Only CRUD users that they created.
func (s *server) scopeAuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, _ := s.extractCredentials(r)
		cu, _ := getByUsernameAndPassword(username, password, true, s)
		uri := r.RequestURI
		if strings.HasSuffix(uri, su) {
			if !cu.Admin {
				fail(w, notAuthorized, 401)
				return
			}
		}

		if strings.Contains(uri, app) {
			u, _ := url.Parse(r.RequestURI)
			id := path.Base(u.Path)
			if len(id) == 0 && r.Method != "POST" {
				fail(w, notAuthorized, 401)
				return
			} else {
				app, _ := getApplicationById(id, s)
				if !cu.Admin && app.CreatedBy != cu.Id {
					fail(w, notAuthorized, 401)
					return
				}
			}
		}

		if strings.Contains(uri, u) {
			u, _ := url.Parse(r.RequestURI)
			id := path.Base(u.Path)
			if len(id) == 0 && r.Method != "POST" {
				fail(w, notAuthorized, 401)
				return
			}

			var (
				newUser User
				err     error
			)

			if r.Method == "POST" || r.Method == "PUT" {
				decoder := json.NewDecoder(r.Body)
				err = decoder.Decode(&newUser)
			} else {
				newUser, err = getById(id, s)
			}

			if err != nil {
				fail(w, "Failed to parse user.", 200)
				return
			}

			users, _ := listUsersByCreatedId(cu.Id, s)
			if users == nil {
				fail(w, "Could not find user.", 200)
				return
			}

			creator := false
			for _, u := range users {
				if u.Id == newUser.Id {
					creator = true
					break
				}
			}

			if !cu.Admin && (!creator || newUser.Admin) {
				fail(w, notAuthorized, 401)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func (s *server) extractCredentials(r *http.Request) (string, string, error) {
	a := r.Header.Get("Authorization")
	p := strings.Split(a, " ")
	if len(p) < 2 || !strings.EqualFold(p[0], basic) {
		return "", "", errors.New(notAuthorized)
	}
	d, _ := b64.StdEncoding.DecodeString(p[1])
	slice := strings.Split(string(d), ":")
	if len(slice) < 2 {
		return "", "", errors.New(notAuthorized)
	}
	return slice[0], slice[1], nil
}

func (s *server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.String()
		log.Printf("%s %s\n", r.Method, url)
		next.ServeHTTP(w, r)
	})
}
