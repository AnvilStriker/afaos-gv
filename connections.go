package afaos

type LocPair struct {
	Orig string
	Dest string
}

type Connections map[LocPair]Transport

func MakeNetwork(locs []*Location) Connections {
	conns := make(Connections)
	for _, orig := range locs {
		for transport, destLocs := range orig.Destinations {
			for _, dest := range destLocs {
				conns[LocPair{Orig: orig.Name, Dest: dest}] = transport
			}
		}
	}
	return conns
}

func MakeOrdMap(names [][]string) map[string]int {
	m := make(map[string]int)
	c := 0
	for _, row := range names {
		for _, name := range row {
			m[name] = c
			c++
		}
	}
	return m
}

func FilterOrd(names [][]string, locs []*Location) [][]string {
	locMap := make(map[string]bool)
	for _, loc := range locs {
		locMap[loc.Name] = true
	}
	for r, row := range names {
		filtered := make([]string, 0 , len(row))
		for _, name := range row {
			if locMap[name] {
				filtered = append(filtered, name)
			}
		}
		names[r] = filtered
	}
	return names
}

func InvertOrd(names [][]string) [][]string {
	for top, bot := 0, len(names)-1; top < bot; top, bot = top+1, bot-1 {
		names[top], names[bot] = names[bot], names[top]
	}
	for _, row := range names {
		for left, right := 0, len(row)-1; left < right; left, right = left+1, right-1 {
			row[left], row[right] = row[right], row[left]
		}
	}
	return names
}	
