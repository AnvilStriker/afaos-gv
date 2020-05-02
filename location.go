package afaos

type Side int

const (
	British Side = iota
	French
)

type Transport string

const (
	Bateau  = "Bateau"
	Ship    = "Ship"
	Wagon   = "Wagon"
	NoRoute = "(no route)"
)

type Resource string

const (
	Bateaux  = "Bateaux"
	Furs     = "Furs"
	Military = "Military"
	Settlers = "Settlers"
	Ships    = "Ships"
	Wagons   = "Wagons"
)

type Location struct {
	Name         string
	Start        bool
	Destinations map[Transport][]string
	Resources    []Resource
	Money        int
        Defense      int
	VP           int
}
