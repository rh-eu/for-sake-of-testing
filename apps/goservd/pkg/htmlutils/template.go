package htmlutils

import (
	"html/template"
	"log"
	"net/http"
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

// Render ...
func (g *TemplateGroup) Render(w http.ResponseWriter, name string, context interface{}) {

	log.Print("Rendering ...")
	log.Printf("Render . context.addr: %v", &context)
	log.Printf("Render . Context ... %v  .... index: %v", context, name)

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}

	log.Printf("Data ... %v", data)

	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message and use
	// the http.Error() function to send a generic 500 Internal Server Error
	// response to the user.
	ts, err := template.ParseFiles("./sitedata/templates/index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	log.Printf("template ts ... %v", ts)

	// We then use the Execute() method on the template set to write the template
	// content as the response body. The last parameter to Execute() represents any
	// dynamic data that we want to pass in, which for now we'll leave as nil.
	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

}
