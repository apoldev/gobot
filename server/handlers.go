package server

import "github.com/gin-gonic/gin"

func (s *Server) HandlerDefault() func(c *gin.Context) {

	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"app":     "gobot",
			"version": "1.0",
		})
	}
}
