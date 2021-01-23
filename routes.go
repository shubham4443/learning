package main

import "learning/handlers"

func (s *Server) routes() {
	s.router.Methods("GET").Path("/").HandlerFunc(handlers.GetAllKeys(s.store))
	s.router.Methods("POST").Path("/").HandlerFunc(handlers.SetKeys(s.store))
	s.router.Methods("DELETE").Path("/").HandlerFunc(handlers.DeleteAllKeys(s.store))
	s.router.Methods("GET").Path("/{key}").HandlerFunc(handlers.GetSingleKey(s.store))
	s.router.Methods("DELETE").Path("/{key}").HandlerFunc(handlers.DeleteSingleKey(s.store))
}
