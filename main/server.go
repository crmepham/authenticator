package main

import (
	"bufio"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"os"
	"strings"
)

type server struct {
	router *mux.Router
	config map[string]string
	db     dbDriver
}

const (
	p, app, u, su, id, a = "/api/v1/", "application/", "user/", "system/user", "{1-9}", "application.properties"
)

func (s *server) init() {
	s.router = mux.NewRouter()
	s.loadConfig()
	s.loadDBDriver()
	s.registerMiddleWare()
	s.registerHandlers()
}

func (db dbDriver) buildUrl() string {
	return db.username + ":" + db.password + "@tcp(" + db.url + ":" + db.port + ")/" + db.name
}

func (s *server) registerHandlers() {
	s.router.HandleFunc(p+"applications", s.listApplications()).Methods("GET")
	s.router.HandleFunc(p+"application", s.createApplication()).Methods("POST")
	s.router.HandleFunc(p+app+id, s.getApplication()).Methods("GET")
	s.router.HandleFunc(p+app+id, s.updateApplication()).Methods("PUT")
	s.router.HandleFunc(p+app+id, s.deleteApplication()).Methods("DELETE")
	s.router.HandleFunc(p+app+"users/"+id, s.listApplicationUsers()).Methods("GET")
	s.router.HandleFunc(p+"user", s.createUser()).Methods("POST")
	s.router.HandleFunc(p+"system/users", s.listUsers(true)).Methods("GET")
	s.router.HandleFunc(p+"users", s.listUsers(false)).Methods("GET")
	s.router.HandleFunc(p+u+id, s.getUser()).Methods("GET")
	s.router.HandleFunc(p+u+id, s.deleteUser()).Methods("DELETE")
	s.router.HandleFunc(p+u+id, s.updateUser()).Methods("PUT")
	s.router.HandleFunc(p+"login", s.login()).Methods("POST")
}

func (s *server) registerMiddleWare() {
	s.router.Use(s.loggingMiddleware)
	s.router.Use(s.authenticationMiddleware)
	s.router.Use(s.scopeAuthenticationMiddleware)
}

func (s *server) loadConfig() {
	file, err := os.Open("../" + a)
	check(err)
	defer file.Close()

	reader := bufio.NewReader(file)

	s.config = make(map[string]string)
	for {
		line, err := reader.ReadString('\n')
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				s.config[key] = value
			}
		}
		if err == io.EOF {
			break
		}
		check(err)
	}
}

func (s *server) loadDBDriver() {
	s.db.password = s.config["database.password"]
	s.db.name = s.config["database.name"]
	s.db.username = s.config["database.username"]
	s.db.driver = s.config["database.driver"]
	s.db.port = s.config["database.port"]
	s.db.url = s.config["database.url"]
}
