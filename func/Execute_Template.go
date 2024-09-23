package groupie

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

// ExecuteTemplate renders a template with the provided data and writes it to the response writer.
func ExecuteTemplate(temp *template.Template, s string, w http.ResponseWriter, info interface{}, status int) {
	var buf bytes.Buffer // Buffer to hold the rendered template content temporarily
	var err error
	if s == "alo" {
		err = temp.Execute(&buf, Result)
		if err != nil {
			fmt.Println(err)
			Error(w, http.StatusInternalServerError)
			return
		}
		temp.Execute(w, Result)

	} else if s == "err" {
		err = temp.Execute(&buf, info)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(status)
		temp.Execute(w, info)
	} else {
		//fmt.Println(Result.Tbn[status-1 : status])
		Sus := jdak{
			Tbn:  Result.Tbn[status-1 : status],
			Mapi: Result.Aloo.Index[status-1].Mapio,
		}
		fmt.Println(Sus.Mapi)
		err = temp.Execute(&buf, Sus)
		if err != nil {
			
			Error(w, http.StatusInternalServerError)
			return
		}
		temp.Execute(w, Sus)
	}
}
