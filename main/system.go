package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

func (s *server) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var user User
		err := decoder.Decode(&user)
		if err != nil {
			fail(w, "Invalid credentials.", 401)
			return
		}

		_, err = getByUsernameAndPassword(user.Username, user.Password, false, s)
		if err != nil {
			fail(w, "Invalid credentials.", 401)
		} else {
			fail(w, "Success.", 200)
		}
	}
}

func (s *server) createUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var user User
		err := decoder.Decode(&user)
		check(err)

		if user.Application == 0 && !user.Api {
			fail(w, "User must belong to an application.", 200)
			return
		}

		_, err = getApplicationById(strconv.Itoa(user.Application), s)
		if err != nil {
			fail(w, "Could not find application.", 200)
			return
		}

		_, err = getByUsernameOrEmailAndApplication(user.Username, user.Email, user.Application, s)
		if err == nil {
			fail(w, "User already exists.", 200)
			return
		} else {
			username, password, _ := s.extractCredentials(r)
			currentUser, _ := getByUsernameAndPassword(username, password, true, s)

			if user.Api && (!currentUser.Api && !currentUser.Admin) {
				fail(w, notAuthorized, 401)
				return
			}

			var err error
			if currentUser.Admin && user.Admin {
				_, err = create(user, s, true, currentUser)
			} else {
				_, err = create(user, s, false, currentUser)
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(200)

			newUser, _ := getByUsernameOrEmailAndApplication(user.Username, user.Email, user.Application, s)
			u, _ := json.Marshal(newUser)
			_, err = w.Write([]byte(string(u)))
			check(err)
		}
	}
}

func (s *server) getUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, _ := url.Parse(r.RequestURI)
		id := path.Base(u.Path)
		user, err := getById(id, s)
		if err != nil {
			fail(w, "Could not find user.", 200)
		} else {
			us, _ := json.Marshal(user)
			_, err := w.Write([]byte(string(us)))
			check(err)
		}
	}
}

func (s *server) listUsers(api bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := listAllUsers(api, s)
		if err != nil {
			fail(w, err.Error(), 200)
		} else {
			m, _ := json.Marshal(users)
			_, err = w.Write([]byte(string(m)))
			check(err)
		}
	}
}

func (s *server) listApplications() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, _ := s.extractCredentials(r)
		cu, _ := getByUsernameAndPassword(username, password, true, s)

		var (
			apps []Application
			err  error
		)

		if cu.Admin {
			apps, err = listAllApplications(s)
		} else {
			apps, err = listApplications(cu.Id, s)
		}

		if err != nil {
			fail(w, err.Error(), 200)
		} else {
			m, _ := json.Marshal(apps)
			_, err = w.Write([]byte(string(m)))
			check(err)
		}
	}
}

func (s *server) listApplicationUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, _ := url.Parse(r.RequestURI)
		id := path.Base(u.Path)
		_, err := getApplicationById(id, s)
		if err != nil {
			fail(w, "Could not find application.", 200)
		} else {
			users, _ := listApplicationUsers(id, s)
			us, _ := json.Marshal(users)
			_, err := w.Write([]byte(string(us)))
			check(err)
		}
	}
}

func (s *server) createApplication() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var app Application
		err := decoder.Decode(&app)
		check(err)

		_, err = getApplicationByNameOrUrl(app.Name, app.Url, s)
		if err == nil {
			fail(w, "Application already exists.", 200)
		} else {
			username, password, _ := s.extractCredentials(r)
			currentUser, _ := getByUsernameAndPassword(username, password, true, s)
			_, err = createApplication(app, s, currentUser)
			check(err)

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(200)
			newApp, _ := getApplicationByName(app.Name, s)
			u, _ := json.Marshal(newApp)
			_, err := w.Write([]byte(string(u)))
			check(err)
		}
	}
}

func (s *server) getApplication() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, _ := url.Parse(r.RequestURI)
		id := path.Base(u.Path)
		app, err := getApplicationById(id, s)
		if err != nil {
			fail(w, "Could not find application.", 200)
		} else {
			us, _ := json.Marshal(app)
			_, err := w.Write([]byte(string(us)))
			check(err)
		}
	}
}

func (s *server) updateApplication() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, _ := url.Parse(r.RequestURI)
		id := path.Base(u.Path)
		selected, err := getApplicationById(id, s)
		if err != nil {
			fail(w, "Could not find application.", 200)
		} else {
			decoder := json.NewDecoder(r.Body)
			var newApp Application
			err := decoder.Decode(&newApp)
			check(err)
			_, current := currentUser(r, s)
			selected = bindApplication(selected, newApp, current)
			updateApplication(selected, s)
			su, _ := json.Marshal(selected)
			_, err = w.Write([]byte(string(su)))
			check(err)
		}
	}
}

func (s *server) updateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, _ := url.Parse(r.RequestURI)
		id := path.Base(u.Path)
		selected, err := getById(id, s)
		if err != nil {
			fail(w, "Could not find user.", 200)
		} else {
			decoder := json.NewDecoder(r.Body)
			var newUser User
			err := decoder.Decode(&newUser)
			check(err)
			_, err = getByUsernameOrEmail(newUser.Username, newUser.Email, true, s)
			if err == nil {
				fail(w, "User already exists.", 200)
			} else {
				_, current := currentUser(r, s)
				selected = bind(selected, newUser, current)
				update(selected, s)
				su, _ := json.Marshal(selected)
				_, err = w.Write([]byte(string(su)))
				check(err)
			}
		}
	}
}

func currentUser(r *http.Request, s *server) (error, User) {
	username, password, _ := s.extractCredentials(r)
	current, _ := getByUsernameAndPassword(username, password, true, s)
	return nil, current
}

func bindApplication(e Application, n Application, current User) Application {
	e.Name = n.Name
	e.Description = n.Description
	e.Url = n.Url
	e.Active = n.Active
	e.UpdatedBy = current.Id
	return e
}

func bind(e User, n User, current User) User {
	e.Email = n.Email
	e.Username = n.Username
	e.Password = n.Password
	e.Active = n.Active
	e.UpdatedBy = current.Id
	return e
}

func (s *server) deleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, _ := url.Parse(r.RequestURI)
		id := path.Base(u.Path)
		selected, err := getById(id, s)
		if err != nil {
			fail(w, "Could not find user.", 200)
		} else {
			if count(s) == 1 {
				fail(w, "Cannot delete last API user.", 200)
			} else {
				check(err)
				selected.Deleted = true
				selected.Active = false
				update(selected, s)
				w.WriteHeader(200)
			}
		}
	}
}

func (s *server) deleteApplication() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, _ := url.Parse(r.RequestURI)
		id := path.Base(u.Path)
		selected, err := getApplicationById(id, s)
		if err != nil {
			fail(w, "Could not find application.", 200)
		} else {
			count := countApplicationUsers(id, s)
			if count == 1 {
				fail(w, "Cannot deleted application. Delete its remaining users first.", 200)
			} else {
				check(err)
				selected.Deleted = true
				selected.Active = false
				updateApplication(selected, s)
				w.WriteHeader(200)
			}
		}
	}
}
