package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	err error
)

type Task struct {
	ID          int
	Title       string
	StartDate   time.Time
	DueDate     time.Time
	Status      bool
	Priority    bool
	Description string
	CreatedAt   time.Time
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "example"
	dbName := "gotest"
	dbHost := "localhost"
	dbPort := "3306"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Repository interface {
	Find(id int) (Task, error)
	Add(task Task) (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: dbConn()}
}

func (r repository) Find(id int) (Task, error) {
	var task Task

	rows, _ := r.db.Query("SELECT * FROM tasks WHERE id=?", id)
	defer rows.Close()

	//
	if !rows.Next() {
		return task, errors.New("record not found")
	}

	for rows.Next() {
		err := rows.Scan(&task.ID, &task.Title, &task.StartDate, &task.DueDate, &task.Status, &task.Priority, &task.Description, &task.CreatedAt)
		if err != nil {
			return task, errors.New("record not found")
		}
	}

	if err = rows.Err(); err != nil {
		return task, err
	}

	return task, nil
}

func (r repository) Add(task Task) (int64, error) {

	stmt, err := r.db.Prepare("INSERT INTO tasks (title,start_date,due_date,status,priority,description,created_at) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(task.Title, task.StartDate, task.DueDate, task.Status, task.Priority, task.Description, task.CreatedAt)
	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastID, nil
}

func main() {
	db := dbConn()
	myDB := NewRepository(db)

	task := Task{
		Title:       "React Native",
		StartDate:   time.Now(),
		DueDate:     time.Now(),
		Status:      true,
		Priority:    true,
		Description: "React Native is a JavaScript framework for writing real, natively rendering mobile applications for iOS and Android.",
		CreatedAt:   time.Now(),
	}

	lastID, err := myDB.Add(task)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(lastID)

}
