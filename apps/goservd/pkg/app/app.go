package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// App ...
type App struct {
	r *httprouter.Router
}

func (k *App) getRoothandler() httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Printf("Hello World!\n")

	})
}

// Run ...
func (k *App) Run() {
	log.Printf("goservd is up and running")
	log.Fatal(http.ListenAndServe(":5051", k.r))
}

// NewApp ...
func NewApp() *App {
	k := &App{
		r: httprouter.New(),
	}

	router := k.r

	rootHandler := k.getRoothandler()
	router.GET("/hello", rootHandler)

	return k
}
