package extractrecipe

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIngredientsInString(t *testing.T) {
	fmt.Println(strings.Index("1/2 cup chocolate chips (melted", "chocolate chips"))
	g := GetIngredientsInString("1/2 cup chocolate chips (melted")
	fmt.Println(g)
	assert.Equal(t, "chocolate chips", "hi")
}
