package web

import (
	"SeaBattle/battle"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ginServer struct {
	model *battle.SeaBattleGame
	engine *gin.Engine
}


func NewGinServer(model *battle.SeaBattleGame) *ginServer {
	server := &ginServer{
		model: model,
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.GET("/create-matrix", server.createMatrix)
	router.GET("/ship", server.ship)
	router.GET("/shot", server.shot)
	router.GET("/clear", server.clear)
	router.GET("/state", server.state)

	server.engine = router
	return server
}

func (s *ginServer) Run(port int) error {
	return s.engine.Run(fmt.Sprintf(":%d", port))
}

func (s *ginServer) createMatrix(context *gin.Context) {

}

func (s *ginServer) ship(context *gin.Context) {

}
func (s *ginServer) shot(context *gin.Context) {

}
func (s *ginServer) clear(context *gin.Context) {

}

func (s *ginServer) state(context *gin.Context) {

}



