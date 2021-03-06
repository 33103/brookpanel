package controllers

type IndexController struct {
	BaseController
}

// Index
func (c *IndexController) Index() {

	c.Data["title"] = "首页-" + c.appname

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "index/css.html"
	c.LayoutSections["footerjs"] = "index/js.html"
	c.setTpl("index/index.html", "shared/public_page.html")

}
