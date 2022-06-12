package controllers

import (
	"myblog/models"
	"myblog/util"
)

type BlogController struct {
	baseController
}

func (c *BlogController) list() {
	var (
		page     int
		pagesize int = 6
		offset   int
		list     []*models.Post
		hosts    []*models.Post
		cateId   int
		keyword  string
	)

	categorys := []*models.Category{}
	c.o.QueryTable(new(models.Category).TableName()).All(&categorys)
	c.Data["cates"] = categorys
	if page, _ = c.GetInt("page"); page < 1 {
		page = 1
	}
	offset = (page - 1) * pagesize
	query := c.o.QueryTable(new(models.Post).TableName())

	if c.actionName == "resource" {
		query = query.Filter("tpyes", 0)
	} else {
		query = query.Filter("tpye", 1)
	}

	if cateId, _ = c.GetInt("cate_id"); cateId != 0 {
		query = query.Filter("category_id", cateId)
	}
	keyword = c.Input().Get("keyword")
	if keyword != "" {
		query = query.Filter("title_contains", keyword)
	}
	query.OrderBy("-views").Limit(10, 0).All(&hosts)
	if c.actionName == "home" {
		query = query.Filter("is_top", 1)
	}
	count, _ := query.Count()
	c.Data["count"] = count
	query.OrderBy("-created").Limit(pagesize, offset).All(&list)

	c.Data["list"] = list
	c.Data["pagebar"] = util.NewPager(page, int(count), pagesize, "/"+c.actionName, true).ToString()
	c.Data["hosts"] = hosts
}

func (c *BlogController) Home() {
	config := models.Config{Name: "start"}
	c.o.Read(&config, "name")
	var notices []*models.Post
	c.o.QueryTable(new(models.Post).TableName()).Filter("category_id", 2).All(&notices)
	c.Data["notices"] = notices

	if config.Value != "1" {
		c.Ctx.WriteString("系统维护")
		return
	}

	c.list()
	c.TplName = c.controllerName + "/home.html"
}

func (c *BlogController) Article() {

}
