package app

import (
	"fmt"
	"log"
	"net/http"
        "io/ioutil" 

	"github.com/julienschmidt/httprouter"
	"github.com/rh/for-sake-of-testing/apps/goservd/pkg/htmlutils"
        "github.com/rh/for-sake-of-testing/apps/goservd/pkg/sitedata"
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

func faviconHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
     //http.ServeFile(w, r, "/static/images/favicon.ico")

     file, err := sitedata.Assets.Open("images/favicon.ico")
     if err != nil {
             log.Fatal(err)
     }

     content, err := ioutil.ReadAll(file)
     if err != nil {
             log.Fatal(err)
     }
     w.Write([]byte(string(content)))
     defer file.Close()

}

// NewApp ...
func NewApp() *App {
	k := &App{
		r: httprouter.New(),
	}

	router := k.r

        router.GET("/favicon.ico", faviconHandler)

	rootHandler := k.getRoothandler()
	router.GET("/", rootHandler)

	fs1 := http.FileServer(http.Dir("./sitedata/static/"))
	log.Printf("Filesystem fs1 ... %v", fs1)

	fs2 := http.FileServer(sitedata.Assets)
	log.Printf("Filesystem fs2 ... %v", fs2)

	router.Handler("GET", "/static/*filepath", fs2)
	router.Handler("GET", "/built/*filepath", fs2)

	//router.Handler("GET", "/static/*filepath", http.StripPrefix("/static/", http.FileServer(http.Dir("./sitedata/static/"))))
	//router.Handler("GET", "/built/*filepath", http.StripPrefix("/built/", http.FileServer(http.Dir("./sitedata/built/"))))

	return k
}
