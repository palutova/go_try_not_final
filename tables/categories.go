package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	editType "github.com/GoAdminGroup/go-admin/template/types/table"

	//"github.com/GoAdminGroup/go-admin/template/icon"
	//"github.com/GoAdminGroup/go-admin/template/types"
	//"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"

	//editType "github.com/GoAdminGroup/go-admin/template/types/table"
)


func GetCategoriesTable(ctx *context.Context) (CategoriesTable table.Table) {

	CategoriesTable = table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))
	info := CategoriesTable.GetInfo()
	// connect your custom connection
	// projectCategoriesTable = table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("mysql", "admin"))

	info.AddField("ID", "id", db.Int).FieldHide()
	info.AddField("Приоритет", "prior", db.Int).FieldSortable()
	info.AddField("Заголовок", "title", db.Varchar).FieldSortable()
	info.AddField("Slug", "slug", db.Varchar).FieldSortable()
	info.AddField("Дата обновления", "updated_at", db.Timestamp)
	info.AddField("Скрыт?", "hided", db.Boolean).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "true" { return "Да" }
		if model.Value == "false" { return "Нет" }
		return "???"
	}).FieldEditAble(editType.Switch).FieldEditOptions(types.FieldOptions{
		{Value: "false", Text: "Нет"},
		{Value: "true", Text: "Да"},
	})

	//info.AddButton("Articles", icon.Tv, action.PopUpWithIframe("/authors/list", "文章",
	//	action.IframeData{Src: "/admin/info/posts"}, "900px", "560px"))
	info.SetTable("categories").SetTitle("Категории").SetDescription("Категории проектов. Просмотр")

	formList := CategoriesTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()
	formList.AddField("Slug", "slug", db.Varchar, form.Text)
	formList.AddField("Заголовок", "title", db.Varchar, form.Text)
	formList.AddField("Приоритет", "prior", db.Int, form.Text)
	formList.AddField("Дата обновления", "updated_at", db.Timestamp, form.Text).FieldNotAllowEdit().FieldNotAllowAdd()
	formList.AddField("Категория скрыта?", "hided", db.Boolean, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "Нет", Value: "false"},
			{Text: "Да", Value: "true"},
		}).FieldDefault("true")
	formList.SetTable("categories").SetTitle("Категории проектов").SetDescription("Категории проектов. Редактирование")

	return
}

