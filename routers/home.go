package routers

type HomeRouter struct {
	baseRouter
}

func (r *HomeRouter) Get() {
	r.TplName = "index.tpl"
}
