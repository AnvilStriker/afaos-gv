package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/AnvilStriker/afaos"
	"gopkg.in/yaml.v2"
)

var sideColors = map[afaos.Side]string{
	afaos.British: "red",
	afaos.French:  "blue",
}

var trLabels map[afaos.Transport]string = map[afaos.Transport]string{
	afaos.Bateau: "B",
	afaos.Ship:   "S",
	afaos.Wagon:  "W",
}

var trAttribs map[afaos.Transport]string = map[afaos.Transport]string{
	afaos.Bateau: `color="blue"`,
	afaos.Ship:   `color="seagreen",style="dashed"`,
	afaos.Wagon:  `color="tan"`,
}

var rsLabels map[afaos.Resource]string = map[afaos.Resource]string{
	afaos.Bateaux:  "B",
	afaos.Furs:     "F",
	afaos.Military: "X",
	afaos.Settlers: "T",
	afaos.Ships:    "S",
	afaos.Wagons:   "W",
}

var rsCanon []afaos.Resource = []afaos.Resource{
	afaos.Bateaux,
	afaos.Ships,
	afaos.Wagons,
	afaos.Furs,
	afaos.Military,
	afaos.Settlers,
}

type NodeLabel struct {
	Name      string
	VP        int
	Defense   int
	Resources string
	Start     bool
}

const (
	HasVP = 1 << iota
	HasDefense
	HasResources
)

var NameOnlyLabel = `[label=<<TABLE BORDER="%d" CELLBORDER="1" CELLSPACING="0" COLOR="%s">
                     <TR><TD>%s</TD></TR>
                    </TABLE>>]`


var NamePlusVPLabel = `[label=<<TABLE BORDER="%d" CELLBORDER="1" CELLSPACING="0" COLOR="%s">
                     <TR><TD>%s</TD><TD>%s</TD></TR>
                    </TABLE>>]`

var NameOverResourcesLabel = `[label=<<TABLE BORDER="%d" CELLBORDER="1" CELLSPACING="0" COLOR="%s">
                     <TR><TD>%s</TD></TR>
                     <TR><TD>%s</TD></TR>
                    </TABLE>>]`

var NameOverDefensePlusResourcesLabel = `[label=<<TABLE BORDER="%d" CELLBORDER="1" CELLSPACING="0" COLOR="%s">
                     <TR><TD colspan="2">%s</TD></TR>
                     <TR><TD>%s</TD><TD>%s</TD></TR>
                    </TABLE>>]`

var NamePlusVPOverResourcesLabel = `[label=<<TABLE BORDER="%d" CELLBORDER="1" CELLSPACING="0" COLOR="%s">
                     <TR><TD>%s</TD><TD>%s</TD></TR>
                     <TR><TD colspan="2">%s</TD></TR>
                    </TABLE>>]`

var NamePlusVPOverDefensePlusResourcesLabel = `[label=<<TABLE BORDER="%d" CELLBORDER="1" CELLSPACING="0" COLOR="%s">
                     <TR><TD colspan="2">%s</TD><TD>%s</TD></TR>
                     <TR><TD>%s</TD><TD colspan="2">%s</TD></TR>
                    </TABLE>>]`

func (nl NodeLabel) Render(side afaos.Side) string {
	/*
	top := []string{nl.Name}
	if nl.VP > 0 {
		top = append(top, strconv.Itoa(nl.VP))
	}

	topStr := strings.Join(top, " | ")

	bot := []string{}
	if nl.Defense > 0 {
		bot = append(bot, fmt.Sprintf("+%d", nl.Defense))
	}
	if nl.Resources != "" {
		bot = append(bot, nl.Resources)
	}

	botStr := strings.Join(bot, " | ")

	if botStr == "" {
		return fmt.Sprintf(OneRowLabel, topStr)
	}
	return fmt.Sprintf(TwoRowLabel, topStr, botStr)
	*/
	var vp, def, rs string
	name := nl.Name
	has := 0
	if nl.VP != 0 {
		vp = strconv.Itoa(nl.VP)
		has |= HasVP
	}
	if nl.Defense != 0 {
		def = "+" + strconv.Itoa(nl.Defense)
		has |= HasDefense
	}
	if nl.Resources != "" {
		rs = nl.Resources
		has |= HasResources
	}
	border := 0
	color := "black"
	if nl.Start {
		border = 1
		color = sideColors[side]
	}
	switch has {
	case HasVP:
		return fmt.Sprintf(NamePlusVPLabel, border, color, name, vp)
	case HasResources:
		return fmt.Sprintf(NameOverResourcesLabel, border, color, name, rs)
	case HasVP|HasResources:
		return fmt.Sprintf(NamePlusVPOverResourcesLabel, border, color, name, vp, rs)
	case HasDefense, HasDefense|HasResources:
		return fmt.Sprintf(NameOverDefensePlusResourcesLabel, border, color, name, def, rs)
	case HasVP|HasDefense, HasVP|HasDefense|HasResources:
		return fmt.Sprintf(NamePlusVPOverDefensePlusResourcesLabel, border, color, name, vp, def, rs)
	default:
		return fmt.Sprintf(NameOnlyLabel, border, color, name)
	}
}

func printRanks(ord [][]string) {
	for i, row := range ord {
		items := make([]string, 0)
		for _, s := range row {
			items = append(items, fmt.Sprintf(`"%s";`, s))
		}
		rank := "same"
		switch i {
		case 0:
			rank = "min"
		case len(ord) - 1:
			rank = "max"
		default:
			;
		}
		
		fmt.Printf(`    {rank="%s"; rankdir=LR; %s}`+"\n", rank, strings.Join(items, " "))
	}
	fmt.Println()
}
		 
func printNodes(side afaos.Side, locs []*afaos.Location) {
	for _, loc := range locs {
		nl := NodeLabel{Name: loc.Name, VP: loc.VP, Defense: loc.Defense, Start: loc.Start}

		rm := map[afaos.Resource]string{}
		for _, res := range loc.Resources {
			rm[res] = rsLabels[res]
		}
		rs := []string{}
		for _, res := range rsCanon {
			if s, ok := rm[res]; ok {
				rs = append(rs, s)
			}
		}
		if loc.Money > 0 {
			rs = append(rs, strconv.Itoa(loc.Money))
		}
		nl.Resources = strings.Join(rs, " ")
		fmt.Printf(`    "%s" %s`+"\n", loc.Name, nl.Render(side))
	}
	fmt.Println()
}

func printLinks(network afaos.Connections, ordMap map[string]int) {
	done := make(map[afaos.LocPair]bool)
	for locPair, tr := range network {
		if done[locPair] {
			continue
		}
		done[locPair] = true
		f := `    "%s" -> "%s" [%s]`
		orig, dest := locPair.Orig, locPair.Dest

		// see if there's a link in the reverse direction
		revLocPair := afaos.LocPair{Orig: locPair.Dest, Dest: locPair.Orig}
		if network[revLocPair] == tr {
			done[revLocPair] = true
			f = `    "%s" -> "%s" [%s,dir="both"]`
			// for bidirectional links, let ordMap determine order
			if ordMap[orig] > ordMap[dest] {
				orig, dest = revLocPair.Orig, revLocPair.Dest
			}
		}
		fmt.Printf(f+";\n", orig, dest, trAttribs[tr])
	}
}

func printDigraph(name string, side afaos.Side, ord [][]string, locs []*afaos.Location, network afaos.Connections, ordMap map[string]int) {
	fmt.Printf("digraph %s {\n", name)
	fmt.Println("    graph [splines=true,nodesep=0.4];")
	fmt.Println("    node [shape=plaintext];")
	fmt.Println()

	printRanks(ord)
	printNodes(side, locs)
	printLinks(network, ordMap)

	fmt.Println("}")
}

func main() {
	/*
	side := afaos.British
	b, err := ioutil.ReadFile("/Users/mikec/Development/Graphviz/Play/afaos-br.dat")
	*/
	side := afaos.French
	b, err := ioutil.ReadFile("/Users/mikec/Development/Graphviz/Play/afaos-fr.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	locs := make([]*afaos.Location, 0, 10)
	err = yaml.Unmarshal(b, &locs)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err = ioutil.ReadFile("/Users/mikec/Development/Graphviz/Play/afaos-nodes.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	ord := make([][]string, 0)
	err = yaml.Unmarshal(b, &ord)
	if err != nil {
		fmt.Println(err)
		return
	}
	ord = afaos.FilterOrd(ord, locs) // Eliminate from ord any node not named in the locs data
	if side == afaos.French {
		ord = afaos.InvertOrd(ord)  // Reverse the layout order for the "other" side
	}
	ordMap := afaos.MakeOrdMap(ord)
	network := afaos.MakeNetwork(locs)

	printDigraph("B", side, ord, locs, network, ordMap)
}
