package ingredients

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIngredientsInString(t *testing.T) {
	g := GetIngredientsInString("* 1 1/2 cups (255g) chocolate chips (semi-sweet or milk)")
	assert.Equal(t, "chocolate chips", g[0].Word)
}
