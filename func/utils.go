package groupie

var (
//	Artist      Artists
	// Location    Locations
	// ConcertDate ConcertDates
	// Relation    Relations
	Cards       Card
	Fetched     bool
	Id          int
	Isfetched   bool
	//Mok         blasa
	Result      jdak
	Fu = make( map[string]interface{})
	
)

type Err struct {
	Status  int
	Message string
}



// type blasa struct {
// 	Index []struct {
// 		// ID        int      `json:"id"`
// 		Locations []string `json:"locations"`
// 		// Dates     string   `json:"dates"`
// 	} `json:"index"`
// }

type Card struct {
	Art struct {
		ID           int      `json:"id"`
		Image        string   `json:"image"`
		Name         string   `json:"name"`
		Members      []string `json:"members"`
		CreationDate int      `json:"creationDate"`
		FirstAlbum   string   `json:"firstAlbum"`
		Loco         []string
	}
	Rela map[string]interface{}
}

const Url = "https://groupietrackers.herokuapp.com/api/"

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////:



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
	
	Mp map[string]interface{}
}

