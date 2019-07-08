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
	for n := 0; n < b.N; n++ {
		r, _ := NewFromFile("testing/sites/lasagna.html")
		r.Parse()
	}
}

// TODO: test non recipes
// func TestParseNonRecipes(t *testing.T) {
// 	files := []string{
// 		"testing/sites/spain-europe-middle-class.html",
// 	}
// 	for _, f := range files {
// 		log.Infof("working on %s", f)
// 		r, err := NewFromFile(f)
// 		assert.Nil(t, err)
// 		err = r.Parse()
// 		assert.NotNil(t, err)
// 	}
// }

func ExampleBananaBread1() {
	log.SetLevel("info")
	r, _ := NewFromURL("https://thesaltymarshmallow.com/best-banana-bread-recipe/")
	r.Parse()
	fmt.Println(r.PrintIngredientList())
	// Output:
	// 1 whole butter
	// 3 whole bananas
	// 2 whole eggs
	// 1 teaspoon vanilla
	// 2 cups flour
	// 1 cup granulated sugar
	// 1 teaspoon baking soda
	// 1/2 teaspoon salt
	// 1/2 teaspoon cinnamon
}

func ExampleChocolateChip1() {
	log.SetLevel("info")
	r, _ := NewFromURL("https://joyfoodsunshine.com/the-most-amazing-chocolate-chip-cookies/")
	r.Parse()
	fmt.Println(r.PrintIngredientList())
	// Output:
	// 1 cup butter
	// 1 cup white sugar
	// 1 cup brown sugar
	// 2 tsp vanilla
	// 2 whole eggs
	// 3 cups flour
	// 1 tsp baking soda
	// 1/2 tsp baking powder
	// 1 tsp salt
	// 2 cups chocolate chips
}

func ExampleChocolateChip2() {
	log.SetLevel("info")
	urlToParse := "https://www.allrecipes.com/recipe/10813/best-chocolate-chip-cookies/"
	fileToGet := urlToParse
	fileToGet = strings.TrimPrefix(fileToGet, "https://")
	if string(fileToGet[len(fileToGet)-1]) == "/" {
		fileToGet += "index.html"
	}
	fileToGet = path.Join("testing", "sites", fileToGet)
	r, err := NewFromFile(fileToGet)
	if err != nil {
		fmt.Println(err)
	}
	err = r.Parse()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.IngredientList())
	// Output:
	// 	1 cup butter
	// 1 cup white sugar
	// 1 cup brown sugar (packed)
	// 2 whole eggs
	// 2 teaspoons vanilla
	// 1 teaspoon baking soda
	// 2 teaspoons water (hot)
	// 1/2 teaspoon salt
	// 3 cups flour (all purpose)
	// 2 cups chocolate chips (semisweet)
	// 1 cup walnuts (chopped)
}

type URLIngredients struct {
	URL         string
	Ingredients string
}

func TestTableChocolateChipCookies(t *testing.T) {
	log.SetLevel("info")
	ts := []URLIngredients{
		{"https://www.bonappetit.com/recipe/bas-best-chocolate-chip-cookies", `1 1/4 tsp kosher salt (morton)
3/4 tsp baking soda
3/4 cup butter (unsalted)
1 cup brown sugar (dark)
1/4 cup granulated sugar
1 whole egg
2 whole egg yolks
2 tsp vanilla
6 oz chocolate chips (coarsely chopped or semisweet)`},
	}
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
		assert.Equal(t, t0.Ingredients, strings.TrimSpace(fmt.Sprint(r.IngredientList())))
	}
}

func ExampleChocolateChip3() {
	log.SetLevel("info")
	urlToParse := "https://www.bonappetit.com/recipe/bas-best-chocolate-chip-cookies"
	fileToGet := urlToParse
	fileToGet = strings.TrimPrefix(fileToGet, "https://")
	if string(fileToGet[len(fileToGet)-1]) == "/" {
		fileToGet += "index.html"
	}
	fileToGet = path.Join("testing", "sites", fileToGet)
	r, err := NewFromFile(fileToGet)
	if err != nil {
		fmt.Println(err)
	}
	err = r.Parse()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.IngredientList())
	// Output:
	// 1 1/4 tsp kosher salt (morton)
	// 3/4 tsp baking soda
	// 3/4 cup butter (unsalted)
	// 1 cup brown sugar (dark)
	// 1/4 cup granulated sugar
	// 1 whole egg
	// 2 whole egg yolks
	// 2 tsp vanilla
	// 6 oz chocolate chips (coarsely chopped or semisweet)
}

func ExampleChocolateChip4() {
	log.SetLevel("info")
	urlToParse := "https://pinchofyum.com/the-best-soft-chocolate-chip-cookies"
	fileToGet := urlToParse
	fileToGet = strings.TrimPrefix(fileToGet, "https://")
	if string(fileToGet[len(fileToGet)-1]) == "/" {
		fileToGet += "index.html"
	}
	fileToGet = path.Join("testing", "sites", fileToGet)
	r, err := NewFromFile(fileToGet)
	if err != nil {
		fmt.Println(err)
	}
	err = r.Parse()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.IngredientList())
	// Output:
	// 8 tablespoon butter (s of salted)
	// 1/2 cup white sugar
	// 1/4 cup brown sugar (packed light)
	// 1 teaspoon vanilla
	// 1 whole egg
	// 1 1/2 cup flour (s all purpose)
	// 1/2 teaspoon baking soda
	// 1/4 teaspoon salt
	// 3/4 cup chocolate chips

}

func ExampleChocolateChip5() {
	log.SetLevel("info")
	urlToParse := "https://www.modernhoney.com/the-best-chocolate-chip-cookies/"
	fileToGet := urlToParse
	fileToGet = strings.TrimPrefix(fileToGet, "https://")
	if string(fileToGet[len(fileToGet)-1]) == "/" {
		fileToGet += "index.html"
	}
	fileToGet = path.Join("testing", "sites", fileToGet)
	r, err := NewFromFile(fileToGet)
	if err != nil {
		fmt.Println(err)
	}
	err = r.Parse()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.IngredientList())
	// Output:
	// 1 cup butter (cold)
	// 1 cup brown sugar
	// 1/2 cup sugar (2 tablespoons)
	// 2 whole eggs
	// 2 teaspoons vanilla
	// 2 3/4 cups flour
	// 1 teaspoon cornstarch
	// 3/4 teaspoon baking soda
	// 3/4 teaspoon salt
	// 2 1/2 cups chocolate chips
}

func ExampleChocolateChip6() {
	log.SetLevel("info")
	urlToParse := "https://laurenslatest.com/actually-perfect-chocolate-chip-cookies/"
	fileToGet := urlToParse
	fileToGet = strings.TrimPrefix(fileToGet, "https://")
	if string(fileToGet[len(fileToGet)-1]) == "/" {
		fileToGet += "index.html"
	}
	fileToGet = path.Join("testing", "sites", fileToGet)
	r, err := NewFromFile(fileToGet)
	if err != nil {
		fmt.Println(err)
	}
	err = r.Parse()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.IngredientList())
	// Output:
	// 3/4 cup butter
	// 1 cup brown sugar (packed)
	// 1/2 cup granulated sugar
	// 1 whole egg yolk
	// 2 teaspoons vanilla
	// 2 cups flour (all purpose)
	// 1/2 teaspoon baking soda
	// 1/2 teaspoon salt
	// 1 cup chocolate chips (semi sweet)
	// 1 cup chocolate chips (milk)
}
