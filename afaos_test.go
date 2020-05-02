package afaos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAfaos(t *testing.T) {
	conns := MakeNetwork(BritishStart)

	assert.Equal(t, 13, len(conns))
	assert.Equal(t, Ship, conns[LocPair{Orig:NewYork, Dest:NewHaven}])
	assert.Equal(t, Ship, conns[LocPair{Orig:NewHaven, Dest:NewYork}])
	assert.Equal(t, Bateau, conns[LocPair{Orig:NewHaven, Dest:Deerfield}])
	assert.Equal(t, NoRoute, conns[LocPair{Orig:Deerfield, Dest:NewHaven}])
}
