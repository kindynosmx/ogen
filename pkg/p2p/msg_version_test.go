package p2p_test

import (
	"github.com/kindynosmx/ogen/pkg/p2p"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMsgVersion(t *testing.T) {
	f := fuzz.New().NilChance(0)
	v := new(p2p.MsgVersion)
	f.Fuzz(v)

	ser, err := v.Marshal()
	assert.NoError(t, err)

	desc := new(p2p.MsgVersion)
	err = desc.Unmarshal(ser)
	assert.NoError(t, err)

	assert.Equal(t, v, desc)

	assert.Equal(t, p2p.MsgVersionCmd, v.Command())
	assert.Equal(t, uint64(240), v.MaxPayloadLength())

}
