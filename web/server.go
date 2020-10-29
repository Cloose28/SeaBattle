package web

import (
	"SeaBattle/battle"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ginServer struct {
	model  *battle.SeaBattleGame
	engine *gin.Engine
}

func NewGinServer(model *battle.SeaBattleGame) *ginServer {
	server := &ginServer{
		model: model,
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.POST("/create-matrix", server.createMatrix)
	router.POST("/ship", server.ship)
	router.POST("/shot", server.shot)
	router.POST("/clear", server.clear)
	router.GET("/state", server.state)

	server.engine = router
	return server
}

func (s *ginServer) Run(port int) error {
	return s.engine.Run(fmt.Sprintf(":%d", port))
}

type matrixRequest struct {
	Size int `json:"range"`
}

func (s *ginServer) createMatrix(context *gin.Context) {
	log.Println("create matrix request")
	defer log.Println("create matrix proceed")

	var params matrixRequest
	if err := context.BindJSON(&params); err != nil {
		context.Status(http.StatusBadRequest)
		return
	}
	err := s.model.CreateGame(params.Size)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}
	context.Status(http.StatusOK)
}


type shipRequest struct {
	Coordinates string `json:"Coordinates"`
}
func (s *ginServer) ship(context *gin.Context) {
	log.Println("ship request")
	defer log.Println("ship proceed")

	var request shipRequest
	if err := context.BindJSON(&request); err != nil {
		context.Status(http.StatusBadRequest)
		return
	}

	err := s.model.InitShips(request.Coordinates)
	if err != nil {
		context.Status(http.StatusBadRequest)
		return
	}
	context.Status(http.StatusOK)
}
func (s *ginServer) shot(context *gin.Context) {

}
func (s *ginServer) clear(context *gin.Context) {

}

func (s *ginServer) state(context *gin.Context) {

}
