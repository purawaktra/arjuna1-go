package modules

import "github.com/gin-gonic/gin"

type Arjuna1Router struct {
	engine *gin.Engine
	rh     Arjuna1RequestHandler
}

func CreateArjuna1Router(engine *gin.Engine, rh Arjuna1RequestHandler) Arjuna1Router {
	return Arjuna1Router{
		engine: engine,
		rh:     rh,
	}
}
