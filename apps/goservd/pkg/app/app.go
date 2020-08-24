package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rh/for-sake-of-testing/apps/goservd/pkg/htmlutils"
)

// App ...
type App struct {
	r  *httprouter.Router
	tg *htmlutils.TemplateGroup
}

type pageContext struct {
}

func (k *App) getPageContext(r *http.Request, urlBase string) *pageContext {
	c := &pageContext{}

	return c
}

func (k *App) getRoothandler() httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Printf("Hello World!\n")
		// w.Write([]byte("Hello World!"))
		urlBase := "/"
		k.tg.Render(w, "index.html", k.getPageContext(r, urlBase))
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
