package main

import (
	"database/sql"
	"errors"
	"strconv"
)

func getByUsernameAndPassword(u string, p string, api bool, s *server) (User, error) {
	args := []interface{}{u, p, api}
	return get(s, "select * from user where username = ? and password = ? and active = true and api = ? limit 1", args...)
}

func getByUsernameAndPasswordAndApplication(u string, p string, id int, api bool, s *server) (User, error) {
	args := []interface{}{u, p, id, api}
	return get(s, "select * from user where username = ? and password = ? and application_id = ? and active = true and api = ? limit 1", args...)
}

func getByUsernameOrEmail(u string, e string, api bool, s *server) (User, error) {
	args := []interface{}{u, e, api}
	return get(s, "select * from user where (username = ? or email = ?) and active = true and api = ? limit 1", args...)
}

func getByUsernameOrEmailAndApplication(u string, e string, id int, s *server) (User, error) {
	args := []interface{}{u, e, id}
	return get(s, "select * from user where (username = ? or email = ?) and application_id = ? and active = true limit 1", args...)
}

func getById(id string, s *server) (User, error) {
	i, err := strconv.Atoi(id)
	check(err)
	return get(s, "select * from user where id = ? and deleted = false", i)
}

func get(s *server, q string, args ...interface{}) (User, error) {
	db, err := sql.Open(s.db.driver, s.db.buildUrl())
	check(err)

	defer db.Close()
	row := db.QueryRow(q, args...)
	var user User
	err = row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Application, &user.Api, &user.Active, &user.Admin, &user.Deleted,
		&user.Created, &user.CreatedBy, &user.Updated, &user.UpdatedBy)
	if err != nil {
		return user, errors.New("No user found.")
	} else {
		return user, nil
	}
}

func listUsersByCreatedId(id int, s *server) ([]User, error) {
	args := []interface{}{id}
	return list(s, "select * from user where active = true and deleted = false and created_by = ? order by id asc", args...)
}

func listAllUsers(api bool, s *server) ([]User, error) {
	args := []interface{}{api}
	return list(s, "select * from user where active = true and api = ? and deleted = false order by id asc", args...)
}

func list(s *server, q string, args ...interface{}) ([]User, error) {
	db, err := sql.Open(s.db.driver, s.db.buildUrl())
	check(err)

	defer db.Close()

	rows, err := db.Query(q, args...)
	if err != nil {
		return nil, errors.New(err.Error())
	} else {
		return parse(rows), nil
	}
}

func parse(rows *sql.Rows) []User {
	var users []User
	for rows.Next() {
		var id int
		var username string
		var password string
		var email string
		var application int
		var api bool
		var active bool
		var admin bool
		var deleted bool
		var created string
		var createdBy int
		var updated string
		var updatedBy int
		err := rows.Scan(&id, &username, &password, &email, &application, &api, &active, &admin, &deleted, &created, &createdBy, &updated, &updatedBy)
		check(err)
		users = append(users, User{id, username, password, email, application, api, active, admin, deleted, created, createdBy, updated, updatedBy})
	}
	return users
}

func count(s *server) int {
	db, err := sql.Open(s.db.driver, s.db.buildUrl())
	check(err)

	defer db.Close()

	var count int
	_ = db.QueryRow("select count(id) from user where deleted = false and api = true and active = true").Scan(&count)
	return count
}

func create(u User, s *server, admin bool, c User) (User, error) {
	db, err := sql.Open(s.db.driver, s.db.buildUrl())
	check(err)

	defer db.Close()

	ins, err := db.Query("insert into user (username, password, application_id, email, api, active, admin, created_by, updated_by) values (?,?,?,?,?,?,?,?,?)",
		u.Username, u.Password, u.Application, u.Email, u.Api, true, admin, c.Id, c.Id)

	defer ins.Close()
	check(err)

	if err != nil {
		return u, err
	} else {
		return u, nil
	}
}

func update(u User, s *server) {
	db, err := sql.Open(s.db.driver, s.db.buildUrl())
	check(err)

	defer db.Close()

	insert, err := db.Prepare("update user set email = ?, username = ?, password = ?, active = ?, deleted = ?, updated_by = ?, updated = CURRENT_TIMESTAMP where id = ?")
	check(err)
	_, err = insert.Exec(&u.Email, &u.Username, &u.Password, &u.Active, &u.Deleted, &u.UpdatedBy, &u.Id)
	check(err)
}
