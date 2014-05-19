package main

import (
	"errors"
	"fmt"
	"github.com/go-martini/martini"
	"log"
	"net"
	"net/http"
)

type KAServer struct {
	martini *martini.Martini
	router  martini.Router
}

func NewKAServer() *KAServer {
	k := &KAServer{
		martini: martini.New(),
	}
	k.router = martini.NewRouter()
	k.martini = martini.New()
	k.martini.Action(k.router.Handle)
	k.martini.MapTo(k.router, (*martini.Routes)(nil))
	return k
}

func (k *KAServer) Run() error {
	addr := ":9079"

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not start on %s. Err: %s", addr, err.Error()))
	}

	k.router.Get("/", func() (int, string) {
		return 200, "Serve me"
	})
	k.router.Get("/game", k.StartGame)

	log.Printf("Starting Server at %s", addr)
	return http.Serve(l, k.martini)
}

func (k *KAServer) StartGame() {
	// get input (size, characters)
	// create game, load level
	// start game tick
	game := CreateGame(5, 5, 0, 0)
	fmt.Println(game.level)
}

func main() {
	fmt.Println("Hello Knights")

	app := NewKAServer()

	log.Fatal(app.Run())
}
