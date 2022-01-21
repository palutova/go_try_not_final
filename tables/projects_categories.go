package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	//"github.com/GoAdminGroup/go-admin/template/icon"
	//"github.com/GoAdminGroup/go-admin/template/types"
	//"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"

	//editType "github.com/GoAdminGroup/go-admin/template/types/table"
)


func GetProjectsCategoriesTable(ctx *context.Context) (ProjectsCategoriesTable table.Table) {

	ProjectsCategoriesTable = table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))
	info := ProjectsCategoriesTable.GetInfo()
	// connect your custom connection
	// projectCategoriesTable = table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("mysql", "admin"))

	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("ID проекта", "projects_id", db.Int).FieldSortable()
	info.AddField("ID категории", "categories_id", db.Int).FieldSortable()

	//info.AddButton("Articles", icon.Tv, action.PopUpWithIframe("/authors/list", "文章",
	//	action.IframeData{Src: "/admin/info/posts"}, "900px", "560px"))
	info.SetTable("projects_categories").SetTitle("Категории").SetDescription("Категории проектов. Просмотр")

	formList := ProjectsCategoriesTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()
	formList.AddField("ID проекта", "projects_id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()
	formList.AddField("ID категории", "categories_id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()
	formList.SetTable("projects_categories").SetTitle("Категории и Проекты").SetDescription("Категории и Проекты, Таблица для связи. Редактирование")

	return
}

