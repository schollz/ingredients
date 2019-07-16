package recipeingredients

import (
	"fmt"
	"path"
	"strings"
	"testing"

	log "github.com/schollz/logger"
	"github.com/stretchr/testify/assert"
)

func BenchmarkParse(b *testing.B) {
	log.SetLevel("error")
	fileToGet := ts[0].URL
	fileToGet = strings.TrimPrefix(fileToGet, "https://")
	if string(fileToGet[len(fileToGet)-1]) == "/" {
		fileToGet += "index.html"
	}
	fileToGet = path.Join("testing", "sites", fileToGet)
	r, err := NewFromFile(fileToGet)
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.Parse()
	}
}

func ExampleReadme() {
	r, _ := NewFromURL("https://joyfoodsunshine.com/the-most-amazing-chocolate-chip-cookies/")
	r.Parse()
	fmt.Println(r.PrintIngredientList())
	// Output:
	// 1 cup butter
	// 1 cup sugar
	// 1 cup brown sugar
	// 2 tsp vanilla
	// 2 whole eggs
	// 3 cups flour
	// 1 tsp baking soda
	// 1/2 tsp baking powder
	// 1 tsp salt
	// 2 cups chocolate chips
}

type URLIngredients struct {
	URL         string
	Ingredients []string
}

var ts = []URLIngredients{
	{
		"https://www.bonappetit.com/recipe/bas-best-chocolate-chip-cookies",
		[]string{"1 1/4 tsp kosher salt", "3/4 tsp baking soda", "3/4 cup butter", "1 cup brown sugar", "1/4 cup sugar", "1 whole egg", "2 whole egg yolks", "2 tsp vanilla", "6 oz chocolate chips"},
	},
	{
		"https://pinchofyum.com/the-best-soft-chocolate-chip-cookies",
		[]string{"8 tablespoon butter", "1/2 cup sugar", "1/4 cup brown sugar", "1 teaspoon vanilla", "1 whole egg", "1 1/2 cup flour", "1/2 teaspoon baking soda", "1/4 teaspoon salt", "3/4 cup chocolate chips"},
	},
	{
		"https://www.modernhoney.com/the-best-chocolate-chip-cookies/",
		[]string{"1 cup butter", "1 cup brown sugar", "1/2 cup sugar", "2 whole eggs", "2 teaspoons vanilla", "2 3/4 cups flour", "1 teaspoon cornstarch", "3/4 teaspoon baking soda", "3/4 teaspoon salt", "2 1/2 cups chocolate chips"},
	},
	{
		"https://laurenslatest.com/actually-perfect-chocolate-chip-cookies/",
		[]string{"3/4 cup butter", "1 cup brown sugar", "1/2 cup sugar", "1 whole egg yolk", "2 teaspoons vanilla", "2 cups flour", "1/2 teaspoon baking soda", "1/2 teaspoon salt", "1 cup chocolate chips", "1 cup chocolate chips"},
	},
	{
		"https://www.allrecipes.com/recipe/10813/best-chocolate-chip-cookies/",
		[]string{"1 cup butter", "1 cup sugar", "1 cup brown sugar", "2 whole eggs", "2 teaspoons vanilla", "1 teaspoon baking soda", "2 teaspoons water", "1/2 teaspoon salt", "3 cups flour", "2 cups chocolate chips", "1 cup walnuts"},
	},
	{
		"https://joyfoodsunshine.com/the-most-amazing-chocolate-chip-cookies/",
		[]string{"1 cup butter", "1 cup sugar", "1 cup brown sugar", "2 tsp vanilla", "2 whole eggs", "3 cups flour", "1 tsp baking soda", "1/2 tsp baking powder", "1 tsp salt", "2 cups chocolate chips"},
	},
	{
		"https://cakebycourtney.com/soft-chewy-chocolate-chip-cookies/",
		[]string{"1 cup butter", "1 cup sugar", "1 cup brown sugar", "2 whole eggs", "1 teaspoon vanilla", "3 1/2 cups flour", "1 teaspoon baking soda", "1 teaspoon baking powder", "1 teaspoon salt", "2 cups chocolate chips"},
	},
}

func TestTableChocolateChipCookies(t *testing.T) {
	log.SetLevel("trace")
	for _, t0 := range ts {
		fileToGet := t0.URL
		fileToGet = strings.TrimPrefix(fileToGet, "https://")
		if string(fileToGet[len(fileToGet)-1]) == "/" {
			fileToGet += "index.html"
		}
		fileToGet = path.Join("testing", "sites", fileToGet)
		r, err := NewFromFile(fileToGet)
		assert.Nil(t, err)
		err = r.Parse()
		assert.Nil(t, err)
		ingredients := make([]string, len(r.IngredientList().Ingredients))
		for i, ing := range r.IngredientList().Ingredients {
			ingredients[i] = fmt.Sprintf("%s %s %s", AmountToString(ing.Measure.Amount), ing.Measure.Name, ing.Name)
		}
		assert.Equal(t, t0.Ingredients, ingredients)
	}
}
