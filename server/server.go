package server

import (
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

type MsFormServer struct {
	router *gin.Engine
	port   string
}

func NewMsFormServer(port string) *MsFormServer {
	result := new(MsFormServer)

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	result.port = port
	result.router = gin.New()
	result.router.Use(gin.Logger())

	result.router.GET("/ping", Ping)
	result.router.GET("/form", Form)

	return result
}

func (srv *MsFormServer) Run() {
	srv.router.Run(":" + srv.port)
}

func (srv *MsFormServer) GET(endPoint string, handler MsFormServerHandler) {
	srv.router.GET(endPoint, gin.HandlerFunc(handler))
}

func (srv *MsFormServer) ServeHTTP(res *httptest.ResponseRecorder, req *http.Request) {
	srv.router.ServeHTTP(res, req)
}
