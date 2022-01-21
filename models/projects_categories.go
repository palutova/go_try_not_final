package models

import (
	//"errors"
	"encoding/json"
	"fmt"
	"github.com/GoAdminGroup/go-admin/template/types"

	//"github.com/fatih/structs"
	"html/template"
)

type Project struct {
	//gorm.Model
	ID        uint `gorm:"column:id"`
}

type ProjectsCategories struct {
	//gorm.Model
	ID        uint      `gorm:"column:id"`
	ProjectsId uint     `gorm:"column:projects_id"`
	CategoriesId uint   `gorm:"column:categories_id"`
}

type Category struct {
	//gorm.Model
	ID        uint      `gorm:"column:id"`
	Title     string    `gorm:"column:title"`
	Slug      string    `gorm:"column:slug"`
	Prior     int       `gorm:"column:prior"`
	Hided  	  bool      `gorm:"column:hided"`
	Projects []Project  `gorm:"many2many:projects_categories;"`
}
type CategorySelected struct {
	Text     string `json:"Text"`
	Value    string `json:"Value"`
	Selected bool   `json:"Selected"`
}

type FieldOption struct {
	Text          string            `json:"text"`
	Value         string            `json:"value"`
	TextHTML      template.HTML     `json:"-"`
	Selected      bool              `json:"-"`
	SelectedLabel template.HTML     `json:"-"`
	Extra         map[string]string `json:"-"`
}
type FieldOptions []FieldOption

func GetCategoriesByProjectID(id int) (*[]Category) {
	categoryList := new([]Category)
	// https://golang.hotexamples.com/ru/examples/github.com.jinzhu.gorm/DB/Preload/golang-db-preload-method-examples.html
	// реализует запрос
	// SELECT * FROM categories INNER JOIN projects_categories ON projects_categories.categories_id = categories.id WHERE projects_categories.projects_id = 1;
	Orm = Orm.Joins("join projects_categories on projects_categories.categories_id = categories.id ")
	Orm = Orm.Where("projects_categories.projects_id = ?", id)
	Orm = Orm.Select("categories.id, categories.title")
	println(" количество записей выборки в моделе projects_categories %d", Orm.Find(categoryList).RowsAffected)
	if err = Orm.Find(categoryList).Error; err != nil {
		println(" ошибка в моделе projects_categories %d", err.Error())
		return nil
	}
	if selectedCount := Orm.Find(categoryList).RowsAffected; selectedCount == 0 {
		println(" 0 Категорий у проекта в моделе projects_categories %d", selectedCount)
		return nil
	}
	return categoryList
}
//func GetCategoriesByProjectIDForSelect(id int) ([]CategorySelected) {
func GetCategoriesByProjectIDForSelect(id int) (value types.FieldOptions) {
	allCategories := GetAllCategories()
	//for i := 0; i < len(allCategories); i++ {
	//	cat := allCategories[i]
	//	{ Text: cat.Title, Value: cat.ID, Selected: false },
	//}
	fmt.Printf("\n  категории %q \n", allCategories)

	selectedCategories := GetCategoriesByProjectID(id)
	fmt.Printf("\n Выбранные категории %q \n", selectedCategories)
	b, err := json.Marshal(selectedCategories)

	// чтобы сформировать json ответ из структуры
	if err != nil { fmt.Println(err) } else { fmt.Println(string(b)) }
	//

	var CategorySelectedList = types.FieldOptions{
		{ Text: "test" , Value: "0", Selected: false },
		{ Text: "test 2" , Value: "1", Selected: true },
		{ Text: "test 2" , Value: "2", Selected: false },
		{ Text: "test 3" , Value: "3", Selected: false },
	}
	//a, erra := json.Marshal(CategorySelectedList)
	//fmt.Println(" ========== ")
	//if err != nil {
	//	fmt.Println(erra)
	//	return nil
	//} else {
	//	fmt.Println(string(a))
		//return a
	//}
	fmt.Println(" ========== ")
	//return types.FieldOptions(a)
	return CategorySelectedList
}

func GetAllCategories() (*[]Category) {
	categoriesList := new([]Category)
	Orm.Find(categoriesList)
	err := Orm.Find(categoriesList).Error
	if err != nil {
		println("Нет ни одной категории")
		return nil
	}
	if (len(*categoriesList) > 0) {
		return categoriesList
	} else {
		println("Нет ни одной категории")
		return categoriesList
	}
}
