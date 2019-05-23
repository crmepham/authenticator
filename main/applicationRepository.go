package main

import (
	"database/sql"
	"errors"
	"strconv"
)

func getApplicationByNameOrUrl(n string, u string, s *server) (Application, error) {
	args := []interface{}{n, u}
	return getApplication(s, "select * from application where (name = ? or url = ?) and active = true limit 1", args...)
}

func getApplicationByName(n string, s *server) (Application, error) {
	args := []interface{}{n}
	return getApplication(s, "select * from application where name = ? and active = true limit 1", args...)
}

func getApplicationById(id string, s *server) (Application, error) {
	i, err := strconv.Atoi(id)
	check(err)
	return getApplication(s, "select * from application where id = ? and active = true and deleted = false", i)
}

func getApplication(s *server, q string, args ...interface{}) (Application, error) {
	db, err := sql.Open(s.db.driver, s.db.buildUrl())
	check(err)

	defer db.Close()
	row := db.QueryRow(q, args...)
	var a Application
	err = row.Scan(&a.Id, &a.Name, &a.Description, &a.Url, &a.Active, &a.Deleted, &a.Created, &a.CreatedBy, &a.Updated, &a.UpdatedBy)
	if err != nil {
		return a, errors.New("No application found.")
	} else {
		return a, nil
	}
}

func countUserApplications(id int, s *server) int {
	db, err := sql.Open(s.db.driver, s.db.buildUrl())
	check(err)

	defer db.Close()

	var count int
	_ = db.QueryRow("select count(id) from application where deleted = false and created_by = ?", id).Scan(&count)
	return count
}

func listApplicationUsers(id string, s *server) ([]User, error) {
	i, err := strconv.Atoi(id)
	check(err)

	db, err := sql.Open(s.db.driver, s.db.buildUrl())
	check(err)

	defer db.Close()

	rows, err := db.Query("select * from user where active = true and api = false and deleted = false and application_id = ? order by id asc", i)
	if err != nil {
		return nil, errors.New(err.Error())
	} else {
		return parse(rows), nil
	}
}

func listApplications(id int, s *server) ([]Application, error) {
	args := []interface{}{id}
	return listApps(s, "select * from application where active = true and deleted = false and created_by = ? order by id asc", args...)
}

func listAllApplications(s *server) ([]Application, error) {
	args := []interface{}{}
	return listApps(s, "select * from application where active = true and deleted = false  order by id asc", args...)
}

func listApps(s *server, q string, args ...interface{}) ([]Application, error) {
	db, err := sql.Open(s.db.driver, s.db.buildUrl())
	check(err)

	defer db.Close()

	rows, err := db.Query(q, args...)
	if err != nil {
		return nil, errors.New(err.Error())
	} else {
		return parseApplications(rows), nil
	}
}

func parseApplications(rows *sql.Rows) []Application {
	var applications []Application
	for rows.Next() {
		var id int
		var name string
		var description string
		var url string
		var active bool
		var deleted bool
		var created string
		var createdBy int
		var updated string
		var updatedBy int
		err := rows.Scan(&id, &name, &description, &url, &active, &deleted, &created, &createdBy, &updated, &updatedBy)
		check(err)
		applications = append(applications, Application{id, name, description, url, active, deleted,
			created, createdBy, updated, updatedBy})
	}
	return applications
}

func countApplicationUsers(id string, s *server) int {
	i, err := strconv.Atoi(id)
	check(err)

	db, err := sql.Open(s.db.driver, s.db.buildUrl())
	check(err)

	defer db.Close()

	var count int
	_ = db.QueryRow("select count(id) from user where deleted = false and active = true and application_id = ?", i).Scan(&count)
	return count
}

func createApplication(a Application, s *server, c User) (Application, error) {
	db, err := sql.Open(s.db.driver, s.db.buildUrl())
	check(err)

	defer db.Close()

	ins, err := db.Query("insert into application (name, description, url, active, created_by, updated_by) values (?,?,?,?,?,?)",
		a.Name, a.Description, a.Url, true, c.Id, c.Id)

	defer ins.Close()
	check(err)

	if err != nil {
		return a, err
	} else {
		return a, nil
	}
}

func updateApplication(a Application, s *server) {
	db, err := sql.Open(s.db.driver, s.db.buildUrl())
	check(err)

	defer db.Close()

	insert, err := db.Prepare("update application set name = ?, description = ?, url = ?, active = ?, deleted = ?, updated_by = ?, updated = CURRENT_TIMESTAMP where id = ?")
	check(err)
	_, err = insert.Exec(&a.Name, &a.Description, &a.Url, &a.Active, &a.Deleted, &a.UpdatedBy, &a.Id)
	check(err)
}
