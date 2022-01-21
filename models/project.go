package models

import (
	"errors"
	"fmt"
	"github.com/fatih/structs" // для работы со структурами, читай документашку
	"html/template"
	"strconv"
	"time"
)
type Projects struct {
	ID         uint `gorm:"primary_key,column:title"`
	Title      string `gorm:"column:title"`
	Content    string `gorm:"column:content"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// получить первый элемент и вернуть его структурой описанной выше
func FirstProject() *Projects {
	s := new(Projects)
	Orm.First(s)
	return s
}

//получить элемент по id из параметра
func GetProjectByID(id int) (*Projects, error)  {
	s := new(Projects)
	Orm.First(s, id)
	if structs.IsZero(s) {
		fmt.Printf("\n++++++++ нет ни одного проекта +++++++++\n")
		return s, errors.New("Нет ни одного проекта")
	} else {
		return s, nil
	}
}

// получить только ID проекта
func (s *Projects) GETID() template.HTML {
	return template.HTML(strconv.Itoa(int(s.ID)))
}

// получить массив всех проектов
func GetAllProjects() (*[]Projects, error) {
	projectList := new([]Projects)
	Orm.Find(projectList)
	err := Orm.Find(projectList).Error
	if (len(*projectList) > 0) {
		return projectList, err
	} else {
		fmt.Printf("\n++++++++ нет ни одного проекта +++++++++\n")
		return projectList, errors.New("Нет ни одного проекта")
	}
}



