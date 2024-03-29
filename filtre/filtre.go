package filtre

type Filters struct {
	CreationDateMin float64
	CreationDateMax float64
	FirstAlbumMin   float64
	FirstAlbumMax   float64
	NumMembers      int
	Locations       []string
}
