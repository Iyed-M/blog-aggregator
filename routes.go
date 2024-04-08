package main

func (s *server) addRoutes() {
	v1 := s.app.Group("/v1")
	v1.Get("/readiness", hanldeReadiness)
	v1.Get("/err", handleErr)
	v1.Post("/users", s.handleCreateUser)
}
