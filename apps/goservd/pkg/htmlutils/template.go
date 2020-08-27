package htmlutils

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rh/for-sake-of-testing/apps/goservd/pkg/sitedata"
	"github.com/shurcooL/httpfs/html/vfstemplate"
)

// Todo ...
type Todo struct {
	Title string
	Done  bool
}

// TodoPageData ...
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

// TemplateGroup ...
type TemplateGroup struct {
}

// FuncMap ...
func FuncMap() template.FuncMap {
	return template.FuncMap{
		"jsonstring": JSONString,
	}
}

// JSONString ...
func JSONString(v interface{}) (template.JS, error) {
	a, err := json.Marshal(v)
	if err != nil {
		return template.JS(""), err
	}
	return template.JS(a), nil
}

//// LoadFilesInDir ...
//func LoadFilesInDir(dir string) (map[string]string, error) {
//	dirData := map[string]string{}
//
//	log.Printf("Template Directory ... %v", dir)
//
//	files, err := AssetDir(dir)
//	if err != nil {
//		return dirData, errors.Wrapf(err, "Could not load bindata dir %v", dir)
//	}
//
//	return dirData, nil
//}

// Render ...
func (g *TemplateGroup) Render(w http.ResponseWriter, name string, context interface{}) {

	log.Print("Rendering ...")
	log.Printf("Render . context.addr: %v", &context)
	log.Printf("Render . Context ... %v  .... index: %v", context, name)

	//t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	//err = t.ExecuteTemplate(w, "T", "<script>alert('you have been pwned')</script>")

	//tData, err1 := LoadFilesInDir("templates")
	//if err1 != nil {
	//	log.Printf("Error loading template files: %v", err1)
	//}

	//log.Printf("tData ... %v", tData)

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}

	log.Printf("Data ... %v", data)

	file, err := sitedata.Assets.Open("templates/hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("The content of file hello.txt is:", string(content))
	defer file.Close()

	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message and use
	// the http.Error() function to send a generic 500 Internal Server Error
	// response to the user.
	// t, err := template.ParseFiles("./sitedata/templates/index.html")

	ts := template.New("index.html").Funcs(FuncMap())

	ts, err = vfstemplate.ParseFiles(sitedata.Assets, ts, "index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// We then use the Execute() method on the template set to write the template
	// content as the response body. The last parameter to Execute() represents any
	// dynamic data that we want to pass in, which for now we'll leave as nil.
	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

}
