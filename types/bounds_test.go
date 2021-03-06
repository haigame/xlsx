package types_test

import (
	"encoding/xml"
	"github.com/plandem/xlsx/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBounds(t *testing.T) {
	type Entity struct {
		Attribute *types.Bounds `xml:"attribute,attr"`
	}

	//empty
	entity := Entity{Attribute: &types.Bounds{}}
	encoded, err := xml.Marshal(&entity)
	require.Empty(t, err)
	require.Equal(t, `<Entity></Entity>`, string(encoded))

	//encode
	b := types.BoundsFromIndexes(0, 0, 10, 10)
	entity = Entity{Attribute: &b}
	encoded, err = xml.Marshal(&entity)

	require.Empty(t, err)
	require.Equal(t, `<Entity attribute="A1:K11"></Entity>`, string(encoded))

	//decode
	var decoded Entity
	err = xml.Unmarshal(encoded, &decoded)
	require.Empty(t, err)

	require.Equal(t, entity, decoded)

	//methods
	require.Equal(t, types.Ref("A1:K11"), decoded.Attribute.ToRef())

	w, h := decoded.Attribute.Dimension()
	require.Equal(t, 11, w)
	require.Equal(t, 11, h)

	require.Equal(t, true, decoded.Attribute.Contains(0, 0))
	require.Equal(t, true, decoded.Attribute.ContainsRef("A1"))
	require.Equal(t, false, decoded.Attribute.Contains(12, 12))
	require.Equal(t, false, decoded.Attribute.ContainsRef("L12"))

	b1 := types.BoundsFromIndexes(10, 10, 0, 0)
	require.Equal(t, b, b1)
}
