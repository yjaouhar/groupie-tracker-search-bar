package groupie

import (
	"html/template"
	"net/http"
)

// Gethandel handles HTTP requests for the home page ("/")
func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		Error(w, http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		Error(w, http.StatusNotFound)
		return
	}

	// Check if data has been fetched and is available for rendering
	if !Isfetched {
		Error(w, http.StatusInternalServerError)
		return
	}

	temp, err := template.ParseFiles("template/index.html")
	if err != nil {
		Error(w, http.StatusInternalServerError)
		return
	}
	// Execute the parsed template and write it to the response writer.
	ExecuteTemplate(temp, "artist", w, nil, 0)
}
