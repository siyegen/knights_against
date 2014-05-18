package main

import (
	"errors"
	"fmt"
	"github.com/go-martini/martini"
	"log"
	"net"
	"net/http"
)

type GameServer struct {
	martini *martini.Martini
	router  martini.Router
}

func NewGameServer() *GameServer {
	g := &GameServer{
		martini: martini.New(),
	}
	g.router = martini.NewRouter()
	g.martini = martini.New()
	g.martini.Action(g.router.Handle)
	g.martini.MapTo(g.router, (*martini.Routes)(nil))
	return g
}

func (g *GameServer) Run() error {
	addr := ":9079"

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not start on %s. Err: %s", addr, err.Error()))
	}

	g.router.Get("/", func() (int, string) {
		return 200, "Serve me"
	})

	log.Printf("Starting Server at %s", addr)
	return http.Serve(l, g.martini)
}

func main() {
	fmt.Println("Hello Knights")

	gameServer := NewGameServer()

	log.Fatal(gameServer.Run())
}
