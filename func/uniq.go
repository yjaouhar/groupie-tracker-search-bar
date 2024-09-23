package groupie

import (
	"strconv"
)

func Uni() {
	p := make(map[string]bool)
	for _, v := range Result.Aloo.Index {
		for mk := range v.Mapio {
			if !p["mk"] {
				p[mk] = true
			}
		}
	}
	Fu["location"] = p
	p = make(map[string]bool)
	for _, v := range Result.Tbn {
		for _, mk := range v.Members {
			if !p[mk] {
				p[mk] = true
			}
		}

	}
	Fu["member"] = p
	p = make(map[string]bool)
	for _, v := range Result.Tbn {
		if !p[strconv.Itoa(v.CreationDate)] {
			p[strconv.Itoa(v.CreationDate)] = true
		}

	}
	Fu["creation date"] = p
	p = make(map[string]bool)
	for _, v := range Result.Tbn {

		if !p[v.FirstAlbum] {
			p[v.FirstAlbum] = true
		}

	}
	Fu["First Album"] = p

}
