package routes

import "github.com/olendr/boilerplate/controllers"

var pc = &controllers.PostController{}

func init() {
	Add("GET", "/post", pc.Index)
}
