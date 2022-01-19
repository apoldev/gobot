package server

func (s *Server) configureRouters() {

	s.r.GET("/", s.HandlerDefault())

}
