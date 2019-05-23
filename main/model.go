package main

type Error struct {
	Message string `json:"message"`
}

type dbDriver struct {
	driver   string
	username string
	password string
	url      string
	port     string
	name     string
}

type User struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Application int    `json:"application_id"`
	Api         bool   `json:"api"`
	Active      bool   `json:"active"`
	Admin       bool   `json:"admin"`
	Deleted     bool   `json:"deleted"`
	Created     string `json:"created"`
	CreatedBy   int    `json:"created_by"`
	Updated     string `json:"updated"`
	UpdatedBy   int    `json:"updated_by"`
}

type Application struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Active      bool   `json:"active"`
	Deleted     bool   `json:"deleted"`
	Created     string `json:"created"`
	CreatedBy   int    `json:"created_by"`
	Updated     string `json:"updated"`
	UpdatedBy   int    `json:"updated_by"`
}
