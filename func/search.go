package groupie

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func Search(w http.ResponseWriter, r *http.Request) {
	Result.Mok = nil
	//found := false
	x := r.URL.Query().Get("art")
	if x == "" {
		Error(w, 400)
		return
	}
	fmt.Println(x)
	for i, v := range Result.Tbn {
		if strings.Contains(strings.ToLower(v.Name), strings.ToLower(x)) {
			Result.Mok = append(Result.Mok, Result.Tbn[i])
			fmt.Println(v.Name)
			continue
		} else if strings.Contains(strings.ToLower(v.FirstAlbum), strings.ToLower(x)) && len(x)!= 4 {
			Result.Mok = append(Result.Mok, Result.Tbn[i])
			fmt.Println(v.FirstAlbum)
			continue
		} else if strings.Contains(strings.ToLower(strconv.Itoa(v.CreationDate)), strings.ToLower(x))  && len(x)== 4 {
			Result.Mok = append(Result.Mok, Result.Tbn[i])
			fmt.Println(v.CreationDate)
			continue
		} else {
			found := false
			for _, mem := range v.Members {
				if strings.Contains(strings.ToLower(mem), strings.ToLower(x)) {
					Result.Mok = append(Result.Mok, Result.Tbn[i])
					found = true
					break
				}
			}
			if !found {
				for v, _ := range Result.Aloo.Index[i].Mapio {
					if strings.Contains(strings.ToLower(v), strings.ToLower(x)) {
						Result.Mok = append(Result.Mok, Result.Tbn[i])
						break
					}
				}
			}

		}
	}
	// if !found {
	// 	for i, lo := range Result.Aloo.Index {
	// 		for v, _ := range lo.Mapio {
	// 			if strings.Contains(strings.ToLower(v), strings.ToLower(x)) {
	// 				Result.Mok = append(Result.Mok, Result.Tbn[i])
	// 				break
	// 			}
	// 		}

	// 	}
	// }
	if len(Result.Mok) == 0 {
		Error(w, 404)
		return
	}
	temp, err := template.ParseFiles("template/index.html")
	if err != nil {
		Error(w, http.StatusInternalServerError)
		return
	}
	// Execute the parsed template and write it to the response writer.
	ExecuteTemplate(temp, "alo", w, nil, 0)
}
