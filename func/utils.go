package groupie

var (
	Isfetched bool
	Result    jdak
	Fu        = make(map[string]interface{})
)

const Url = "https://groupietrackers.herokuapp.com/api/"

type jdak struct {
	Tbn []struct {
		ID           int      `json:"id"`
		Image        string   `json:"image"`
		Name         string   `json:"name"`
		Members      []string `json:"members"`
		CreationDate int      `json:"creationDate"`
		FirstAlbum   string   `json:"firstAlbum"`
	}
	Mok []struct {
		ID           int      `json:"id"`
		Image        string   `json:"image"`
		Name         string   `json:"name"`
		Members      []string `json:"members"`
		CreationDate int      `json:"creationDate"`
		FirstAlbum   string   `json:"firstAlbum"`
	}

	Aloo struct {
		Index []struct {
			Mapio map[string][]string `json:"datesLocations"`
		} `json:"index"`
	}
	Mapi map[string][]string

	Mp map[string]interface{}
}
type Err struct {
	Status  int
	Message string
}
