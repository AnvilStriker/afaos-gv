package afaos

const (
	Albany       = "Albany"
	Baltimore    = "Baltimore"
	Boston       = "Boston"
	Deerfield    = "Deerfield"
	Halifax      = "Halifax"
	NewHaven     = "New Haven"
	NewYork      = "New York"
	Pemaquid     = "Penaquid"
	Philadelphia = "Philadelphia"
	PortRoyal    = "Port Royal"
	StMarys      = "St. Mary's"
)

var BritishStart []*Location = []*Location{
	&Location{
		Name:  NewYork,
		Start: true,
		Destinations: map[Transport][]string{
			Bateau: []string{
				Albany,
			},
			Ship: []string{
				Philadelphia,
				NewHaven,
				Boston,
			},
		},
		Resources: []Resource{
			Ships,
			Military,
			Settlers,
		},
		Money: 3,
	},
	&Location{
		Name:  NewHaven,
		Start: true,
		Destinations: map[Transport][]string{
			Bateau: []string{
				Deerfield,
			},
			Ship: []string{
				Pemaquid,
				NewYork,
				Boston,
			},
		},
		Resources: []Resource{
			Ships,
		},
		Money: 2,
	},
	&Location{
		Name:  Boston,
		Start: true,
		Destinations: map[Transport][]string{
			Ship: []string{
				NewYork,
				NewHaven,
				Pemaquid,
				PortRoyal,
				Halifax,
			},
		},
		Resources: []Resource{
			Ships,
			Military,
			Settlers,
		},
		Money: 3,
	},
}
