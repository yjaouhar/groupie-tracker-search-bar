package groupie

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func Search(w http.ResponseWriter, r *http.Request) {
	x := r.URL.Query()
	sl := []int{}
	
	fmt.Println(x["art"])
	for i, v := range Artist {
		
		// sls := []interface{}{}
		// sls = append(sls, v.Members,v.Loco)
		if strings.Contains(strings.ToLower(v.Name), strings.ToLower(x["art"][0])) {
			sl = append(sl, i+1)
			continue
		} else if strings.Contains(strings.ToLower(v.FirstAlbum), strings.ToLower(x["art"][0]))  {
			sl = append(sl, i+1)
			continue
		}else if  strings.Contains(strings.ToLower(strconv.Itoa(v.CreationDate)), strings.ToLower(x["art"][0])){
			sl = append(sl, i+1)
			continue
		}else {
			for _, mem := range v.Members {
				if strings.Contains(strings.ToLower(mem), strings.ToLower(x["art"][0])) {
					sl = append(sl, i+1)
					break
				}
			}
			if len(sl) != 0 {
				for _, lo := range v.Loco {
					if strings.Contains(strings.ToLower(lo), strings.ToLower(x["art"][0])) {
						sl = append(sl, i+1)
						break
					}
				}
			}
		}
	}
	fmt.Println(sl)
}
