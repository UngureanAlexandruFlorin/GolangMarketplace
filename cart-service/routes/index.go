package routes

import (
	"net/http"

	"local.com/golangMarketplace/cartService/controllers"
	"local.com/golangMarketplace/cartService/middlewares"
)

var RoutesRouter *Router

func InitMounts(path string) {
	var create http.Handler = http.HandlerFunc(controllers.Create)
	var read http.Handler = http.HandlerFunc(controllers.Read)
	var update http.Handler = http.HandlerFunc(controllers.Update)
	var delete http.Handler = http.HandlerFunc(controllers.Delete)

	RoutesRouter.HandleFunc(path+"/create", create)
	RoutesRouter.HandleFunc(path+"/read", middlewares.AuthMiddleware(read))
	RoutesRouter.HandleFunc(path+"/update", middlewares.AuthMiddleware(update))
	RoutesRouter.HandleFunc(path+"/delete", middlewares.AuthMiddleware(delete))

}

func init() {
	RoutesRouter = new(Router)
	RoutesRouter.Exists = true
	RoutesRouter.MountRoute = InitMounts

}

func GET() {

}

func POST() {

}

func DELETE() {

}

func PUT() {

}
