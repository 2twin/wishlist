package server

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"
	"wishes/internal/app"
)

type Server struct {
	app        *app.App
	httpServer *http.Server
}

func NewServer() *Server {
	return &Server{
		app: app.NewApp(),
		httpServer: &http.Server{
			Addr: ":8081",
		},
	}
}

func (s *Server) InitRoutes() {
	mux := http.NewServeMux()
	s.httpServer.Handler = mux

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("frontend/index.html", "frontend/wishlist.html")
		if err != nil {
			log.Fatal(err)
		}

		err = t.Execute(w, s.app.Users)
		if err != nil {
			log.Fatal(err)
		}
	})

	fs := http.FileServer(http.Dir("./assets"))
	mux.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/add_user", s.LogMiddleware(s.AddUser))
	mux.HandleFunc("/remove_user", s.LogMiddleware(s.RemoveUser))
	mux.HandleFunc("/add_wish", s.LogMiddleware(s.AddWish))
	mux.HandleFunc("/remove_wish", s.LogMiddleware(s.RemoveWish))
	mux.HandleFunc("/edit_wish", s.LogMiddleware(s.EditWish))
	mux.HandleFunc("/toggle_wish_status", s.LogMiddleware(s.ToggleWishStatus))
	mux.HandleFunc("/get_wishes", s.LogMiddleware(s.GetWishes))
}

func (s *Server) ListenAndServe() error {
	s.InitRoutes()

	log.Printf("Starting server on port %s", s.httpServer.Addr)
	err := s.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Shutdown() error {
	log.Println("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}
