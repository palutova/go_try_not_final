package tables

import (
	"bin/models"
	"fmt"
	//"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	editType "github.com/GoAdminGroup/go-admin/template/types/table"
	template2 "html/template"
	//"time"
	//"errors"

	//"github.com/GoAdminGroup/go-admin/template"
	//"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

// GetPostsTable return the model of table posts.
func GetProjectsTable(ctx *context.Context) (projectsTable table.Table) {
	projectsTable = table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := projectsTable.GetInfo()
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Путь", "slug", db.Varchar)
	info.AddField("Заголовок", "title", db.Varchar)
	info.AddField("Дата публикации", "published_at", db.Varchar)

	//
	//info.AddField("Role Name", "role_name", db.Varchar).FieldJoin(types.Join{
	//	Table: "role",         // table name which you want to join
	//	Field: "id",           // table field name of your own
	//	JoinField: "user_id",  // table field name of the table which you want to join
	//})
	info.AddField("Категории", "categories_id", db.Int).FieldDisplay(func(value types.FieldModel) interface{} {
		return template.Default().
			Link().
			SetURL("/info/categories/detail?__goadmin_detail_pk=" + value.Value).
			SetContent(template.HTML(value.Value)).
			OpenInNewTab().
			SetTabTitle(template.HTML("categories Detail(" + value.Value + ")")).
			GetContent()
	})
	//

	//	добавляет фильтр
	//FieldFilterable(types.FilterType{FormType: form.DatetimeRange})
	//
	//info.AddField("Ссылка на сайт", "siteurl", db.Varchar)
	info.AddField("Скрыт?", "hided", db.Boolean).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "true" { return "Да" }
		if model.Value == "false" { return "Нет" }
		return "???"
	}).FieldEditAble(editType.Switch).FieldEditOptions(types.FieldOptions{
		{Value: "false", Text: "Нет"},
		{Value: "true", Text: "Да"},
	})
	// для сортировки, там некрасивое окно поиска появляется, нужно ли оно или нет.
	//}).FieldFilterable(types.FilterType{FormType: form.SelectSingle}).FieldFilterOptions(types.FieldOptions{
	//	{Value: "false", Text: "Нет"},
	//	{Value: "true", Text: "Да"},
	//})
	// надо ли все эти поля отображать в админке или нет ?
	//info.AddField("Горизонтальная тумба?", "wide_block", db.Boolean)
	info.AddField("Изображение в шапке", "headimg", db.Varchar).FieldDisplay(func(model types.FieldModel) interface{} {
		//fmt.Printf(" \n +++++ headimg is ++++++ \n ", model.Row["headimg"].(string))
		if (model.Row["headimg"].(string)  != "") {
			return template.Default().Image().
				SetSrc(template2.HTML("/uploads/" + model.Row["headimg"].(string))).
				SetHeight("auto").SetWidth("120").WithModal().GetContent()
		} else {
			return "Не задано"
		}

	})
	//info.AddField("Изображение в шапке", "headimg", db.Varchar)

	//info.AddField("Превью", "preview", db.Varchar).FieldImage("50", "auto", "/uploads/") // если без всплывабщего окна

	info.AddField("Превью", "preview", db.Varchar).FieldDisplay(func(model types.FieldModel) interface{} {
		return template.Default().Image().
			SetSrc(template2.HTML("/uploads/" + model.Row["preview"].(string))).
			SetHeight("auto").SetWidth("120").WithModal().GetContent()
	})

	// надо ли все эти поля отображать в админке или нет ?
	//info.AddField("Второе превью", "second_preview", db.Varchar)
	//info.AddField("Видео превью", "video_preview", db.Varchar)
	//info.AddField("Повторять видео превью?", "video_preview_loop", db.Boolean)

	// projects_project_categories
	//info.AddField("Категории", "ID", db.Varchar).FieldJoin(types.Join{
	//	Field:     "author_id",
	//	JoinField: "id",
	//	Table:     "authors",
	//}).FieldHide()
	// example from tutorial
	//info.AddField("Role Name", "role_name", db.Varchar).FieldJoin(types.Join{
	//	Table: "role",         // table name which you want to join
	//	Field: "id",           // table field name of your own
	//	JoinField: "user_id",  // table field name of the table which you want to join
	//})
	//

	info.SetTable("projects").SetTitle("Проекты").SetDescription("страница просмотра всех проектов")

	formList := projectsTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()
	formList.AddField("Путь", "slug", db.Varchar, form.Text)
	formList.AddField("Заголовок", "title", db.Varchar, form.Text)
	formList.AddField("Ссылка на сайт", "siteurl", db.Varchar, form.Text)
	formList.AddField("Контент", "content", db.Varchar, form.RichText).FieldEnableFileUpload()
	formList.AddField("Дата", "published_at", db.Varchar, form.Datetime)
	//formList.AddField("Категории 1", "category_id", db.Int, form.Select).FieldOptions(types.FieldOptions{
	//	{
	//		Text: "beer",
	//		Value: "beer1",
	//	}, {
	//		Text: "juice",
	//		Value: "juice",
	//	}, {
	//		Text: "water",
	//		Value: "water",
	//	}, {
	//		Text: "red bull",
	//		Value: "red bull",
	//	},
	//}).
	//	// returns a []string  here, the corresponding value is ultimately responds field values of this column, the corresponding value is displayed when edit form
	//	FieldDisplay(func(model types.FieldModel) interface{} {
	//		return []string{"beer1", "juice"}
	//	});
	formList.AddField("Категории", "category_id", db.Int, form.Select).

		FieldOptionInitFn(func(model types.FieldModel) types.FieldOptions {
			categories := models.GetCategoriesByProjectIDForSelect(1)
			fmt.Printf("%q", categories)
			//return types.FieldOptions {
			//	categories,
			//}
			//if categories, err := models.GetAllCategories(); err == nil {
			//	fmt.Printf("%q", categories)
			//	//return types.FieldOptions {
			//	//	for i := 0; i < len(categories); i++ {
			//	//		cat := categories[i]
			//	//		{ Text: cat.Title, Value: cat.ID },
			//	//	}
			//	//}
			//} else {
			//	println("ERROR(table projects): %d", err.Error())
			//}
			//return types.FieldOptions {
			//	{ Text: "test" , Value: "0" },
			//	{ Text: "test 2" , Value: "1", Selected: true },
			//	{ Text: "test 2" , Value: "2", Selected: true },
			//	{ Text: "test 3" , Value: "3" },
			//}
			//a, erra := json.Marshal(categories)
			//return types.FieldOptions.Marshal(a)
			return types.FieldOptions(categories)

		})
	//.FieldDisplayInitFn(func(model types.FieldModel) types.FieldDisplay {
	//		if selectedCategories, err := models.GetCategoriesByProjectID(4); err == nil {
	//			fmt.Printf("%q", selectedCategories)
				//return types.FieldOptions {
				//	for i := 0; i < len(categories); i++ {
				//		cat := categories[i]
				//		{ Text: cat.Title, Value: cat.ID },
				//	}
				//}
			//} else {
				// Если топика нет, прервём с ошибкой
				//println("ERROR(table projects): %d", err.Error())
			//}
			//return []string{"1", "2"}
		//})

		//FieldDisplay(func(model types.FieldModel) interface{} {
		//	//return template.Default().Image().
		//	//	SetSrc(template2.HTML("/uploads/" + model.Row["preview"].(string))).
		//	//	SetHeight("auto").SetWidth("120").WithModal().GetContent()
		//	model.Row["project_categories_id"]
		//	{Text: "Нет", Value: "false"},
		//})
		//FieldOptions(types.FieldOptions{
		//	{Text: "Нет", Value: "false"},
		//	{Text: "Да", Value: "true"},
		//})

	//).FieldDisplay(func(value types.FieldModel) interface{} {
	//	return template.Default().
	//		Link().
	//		SetURL("/info/project_categories/detail?__goadmin_detail_pk=" + value.Value).
	//		SetContent(template.HTML(value.Value)).
	//		OpenInNewTab().
	//		SetTabTitle(template.HTML("project_categories Detail(" + value.Value + ")")).
	//		GetContent()
	//})
	//formList.AddField("Дата", "updated_at", db.Varchar, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()
	//formList.AddField("Дата", "created_at", db.Varchar, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()

	formList.AddField("Проект скрыт?", "hided", db.Boolean, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "Нет", Value: "false"},
			{Text: "Да", Value: "true"},
		}).FieldDefault("true")

	formList.AddField("Превью", "preview", db.Varchar, form.File)
	formList.AddField("Второе превью", "second_preview", db.Varchar, form.File)

	formList.AddField("Горизонтальная тумба?", "wide_block", db.Boolean, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "Нет", Value: "false"},
			{Text: "Да", Value: "true"},
		}).FieldDefault("false")
	formList.AddField("Видео превью", "video_preview", db.Varchar, form.File)
	formList.AddField("Повторять видео превью?", "video_preview_loop", db.Boolean, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "Нет", Value: "false"},
			{Text: "Да", Value: "true"},
		}).FieldDefault("false")

	formList.AddField("Изображение в шапке", "headimg", db.Varchar, form.File)
	//
	formList.EnableAjax("Проект успешно сохранен", "Что-то пошло не так, проверьте данные")
	//
	formList.SetTable("projects").SetTitle("Проекты").SetDescription("раздел проекты редактирование")

	return
}

