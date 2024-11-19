package routers

import (
	"net/http"
	Controller "servidorhttp/app/controllers"
)

var listRouters = []Router{
	{
		Uri:      "/home",
		Method:   http.MethodGet,
		function: (Controller.ControllerHome{}).Welcome,
		Auth:     false,
	},
}
