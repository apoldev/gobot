package server

import (
	"github.com/apoldev/gobot/store"
	"github.com/apoldev/gobot/telegram"
	"github.com/gin-gonic/gin"
)

type Server struct {
	r     *gin.Engine
	Store *store.Store
	Bot   *telegram.Bot
}

func New() *Server {

	r := gin.Default()

	return &Server{
		r: r,
	}
}

func (s *Server) Configure() {

	s.configureStore()

	s.configureRouters()

	s.configureBot()

}

func (s *Server) Start() {

	go s.Bot.StartGetUpdates()

	s.r.Run(":8081")

}

func (s *Server) configureStore() {

	s.Store = store.NewStore()

}

func (s *Server) configureBot() {

	s.Bot = telegram.NewBot()
	s.Bot.Store = s.Store

}
