package recipeingredients

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIngredientsInString(t *testing.T) {
	fmt.Println(strings.Index("1/2 cup chocolate chips (melted", "chocolate chips"))
	g := GetIngredientsInString("* 1 1/2 cups (255g) chocolate chips (semi-sweet or milk)")
	fmt.Println(g)
	assert.Equal(t, "chocolate chips", "chocolate chips")
}
