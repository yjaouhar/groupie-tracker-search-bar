package groupie

func Uni() {
	p := make(map[string]bool)
	for _, v := range Result.aloo.Index {
		for mk := range v.Mapio {
			if !p["mk"] {
				p[mk] = true
			}
		}
	}
	Fu["relation"] = p
	p=make(map[string]bool)
	for _, v := range Result.Tbn {
		for _, mk := range v.Members {
			if !p[mk] {
				p[mk] = true
			}
		}
	}
	Fu["member"] = p
}
