package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

var dbConnection Repository

const (
	host     = "localhost"
	port     = 5432
	user     = "carreira"
	password = "123456"
	dbname   = "crud"
)

func connect() *Repository {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	error := db.Ping()
	if error != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return &Repository{db: db}

}

func (db Repository) insert(p person) (int64, error) {
	var id int64

	err := db.db.QueryRow("INSERT INTO person(id, name, cpf, age) VALUES(nextval('personseq'), $1, $2, $3) RETURNING id", p.Name, p.Cpf, p.Age).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("Person: %v", err)
	}

	return id, nil

}

func (db Repository) update(p person) (bool, error) {

	result, err := db.db.Exec("UPDATE person SET name = $1, cpf = $2, age = $3 WHERE id = $4;", p.Name, p.Cpf, p.Age, p.ID)

	if err != nil {
		panic(err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("Update Person: %v", err)
	}
	return count == 1, err

}

func (db Repository) query(id int64) (person, error) {

	var p person
	row := db.db.QueryRow("SELECT * FROM person WHERE id = $1;", id)
	if err := row.Scan(&p.ID, &p.Name, &p.Cpf, &p.Age); err != nil {
		if err == sql.ErrNoRows {
			panic(err)
		}
		panic(err)
	}

	return p, nil

}

func (db Repository) queryAll() ([]person, error) {

	var persons []person
	rows, err := db.db.Query("SELECT * FROM person;")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var p person
		if err := rows.Scan(&p.ID, &p.Name, &p.Cpf, &p.Age); err != nil {
			panic(err)
		}
		persons = append(persons, p)
	}

	return persons, nil

}

func (db Repository) delete(id int64) (int64, error) {
	result, err := db.db.Exec("DELETE FROM person WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("Update Person: %v", err)
	}

	return count, nil

}
