package routers

import (
	"net/http"
	Controller "servidorhttp/app/controllers"
)

var controllerImovel = Controller.NewControllerImovel()

var listRouters = []Router{
	{
		Uri:      "/imovels",
		Method:   http.MethodGet,
		function: controllerImovel.All,
		Auth:     false,
	},
	{
		Uri:      "/imovel/{id}",
		Method:   http.MethodGet,
		function: controllerImovel.Get,
		Auth:     false,
	},
}
