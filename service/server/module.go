package server

import (
	"template/service/api"

	"github.com/gin-gonic/gin"
)

type ServerImpl interface {
	HelloWorld(c *gin.Context)
	NewReviewAPI(c *gin.Context)
}

type Server struct {
	api api.TemplateAPI
}

func NewServer() *Server {
	api := api.NewTemplateAPIImpl()
	return &Server{
		api: api,
	}
}

func NewServerImpl(r *gin.Engine) *gin.Engine {
	server := NewServer()
	r.POST("/postreview", server.NewReviewAPI)
	// r.GET("/getreview/:id",server.GetReviewByid)
	r.GET("/getreview", server.GetReviewAPI)
	r.DELETE("/delete/:id", server.DeleteReviewAPI)
	return r
}

var _ ServerImpl = &Server{}

func (s *Server) HelloWorld(c *gin.Context) {
	// Implement HelloWorld logic here
}
