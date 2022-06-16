package models

import (
	"fmt"
	"strconv"

	"github.com/poncorobbin/go-simple-rest/pkg/db"
	_ "github.com/poncorobbin/go-simple-rest/pkg/db"
)

type Student struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (s *Student) Find() []Student {
	db, err := db.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer db.Close()

	rows, err := db.Query("select id, name, age from tb_student")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer rows.Close()

	var res []Student

	for rows.Next() {
		each := Student{}
		err = rows.Scan(&each.Id, &each.Name, &each.Age)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		res = append(res, each)
	}

	return res
}

func (s *Student) FindOne(id string) Student {
	db, err := db.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return Student{}
	}
	defer db.Close()

	student := Student{}
	err = db.QueryRow("select id, name, age from tb_student where id = ?", id).
		Scan(&student.Id, &student.Name, &student.Age)

	if err != nil {
		fmt.Println(err.Error())
		return Student{}
	}

	return student
}

func (s *Student) Save(body Student) Student {
	db, err := db.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return Student{}
	}
	defer db.Close()

	var student = body

	if student.Id != "" {
		_, err := db.Exec("update tb_student set name = ?,  age = ? where id = ?", student.Name, student.Age, student.Id)
		if err != nil {
			fmt.Println(err.Error())
			return Student{}
		}
	} else {
		id := strconv.Itoa(len(s.Find()) + 1)

		_, err := db.Exec("insert into tb_student values (?, ?, ?, ?)", id, student.Name, student.Age, student.Age)
		if err != nil {
			fmt.Println(err.Error())
			return Student{}
		}
	}

	return student
}
