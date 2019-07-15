package recipeingredients

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

// func BenchmarkParse(b *testing.B) {
// 	log.SetLevel("error")
// 	for n := 0; n < b.N; n++ {
// 		r, _ := NewFromFile("testing/sites/lasagna.html")
// 		r.Parse()
// 	}
// }

// // TODO: test non recipes
// // func TestParseNonRecipes(t *testing.T) {
// // 	files := []string{
// // 		"testing/sites/spain-europe-middle-class.html",
// // 	}
// // 	for _, f := range files {
// // 		log.Infof("working on %s", f)
// // 		r, err := NewFromFile(f)
// // 		assert.Nil(t, err)
// // 		err = r.Parse()
// // 		assert.NotNil(t, err)
// // 	}
// // }

// func ExampleBananaBread1() {
// 	log.SetLevel("info")
// 	r, _ := NewFromURL("https://thesaltymarshmallow.com/best-banana-bread-recipe/")
// 	r.Parse()
// 	fmt.Println(r.PrintIngredientList())
// 	// Output:
// 	// 1 whole butter
// 	// 3 whole bananas
// 	// 2 whole eggs
// 	// 1 teaspoon vanilla
// 	// 2 cups flour
// 	// 1 cup granulated sugar
// 	// 1 teaspoon baking soda
// 	// 1/2 teaspoon salt
// 	// 1/2 teaspoon cinnamon
// }

// func ExampleChocolateChip1() {
// 	log.SetLevel("info")
// 	r, _ := NewFromURL("https://joyfoodsunshine.com/the-most-amazing-chocolate-chip-cookies/")
// 	r.Parse()
// 	fmt.Println(r.PrintIngredientList())
// 	// Output:
// 	// 1 cup butter
// 	// 1 cup white sugar
// 	// 1 cup brown sugar
// 	// 2 tsp vanilla
// 	// 2 whole eggs
// 	// 3 cups flour
// 	// 1 tsp baking soda
// 	// 1/2 tsp baking powder
// 	// 1 tsp salt
// 	// 2 cups chocolate chips
// }

// func ExampleChocolateChip2() {
// 	log.SetLevel("info")
// 	urlToParse := "https://www.allrecipes.com/recipe/10813/best-chocolate-chip-cookies/"
// 	fileToGet := urlToParse
// 	fileToGet = strings.TrimPrefix(fileToGet, "https://")
// 	if string(fileToGet[len(fileToGet)-1]) == "/" {
// 		fileToGet += "index.html"
// 	}
// 	fileToGet = path.Join("testing", "sites", fileToGet)
// 	r, err := NewFromFile(fileToGet)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	err = r.Parse()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(r.IngredientList())
// 	// Output:
// 	// 	1 cup butter
// 	// 1 cup white sugar
// 	// 1 cup brown sugar (packed)
// 	// 2 whole eggs
// 	// 2 teaspoons vanilla
// 	// 1 teaspoon baking soda
// 	// 2 teaspoons water (hot)
// 	// 1/2 teaspoon salt
// 	// 3 cups flour (all purpose)
// 	// 2 cups chocolate chips (semisweet)
// 	// 1 cup walnuts (chopped)
// }

// type URLIngredients struct {
// 	URL         string
// 	Ingredients string
// }

// func TestTableChocolateChipCookies(t *testing.T) {
// 	log.SetLevel("info")
// 	ts := []URLIngredients{
// 		{
// 			"https://www.bonappetit.com/recipe/bas-best-chocolate-chip-cookies",
// 			`1 1/4 tsp kosher salt (morton)
// 3/4 tsp baking soda
// 3/4 cup butter (unsalted)
// 1 cup brown sugar (dark)
// 1/4 cup granulated sugar
// 1 whole egg
// 2 whole egg yolks
// 2 tsp vanilla
// 6 oz chocolate chips (coarsely chopped or semisweet)`,
// 		},
// 	}
// 	for _, t0 := range ts {
// 		fileToGet := t0.URL
// 		fileToGet = strings.TrimPrefix(fileToGet, "https://")
// 		if string(fileToGet[len(fileToGet)-1]) == "/" {
// 			fileToGet += "index.html"
// 		}
// 		fileToGet = path.Join("testing", "sites", fileToGet)
// 		r, err := NewFromFile(fileToGet)
// 		assert.Nil(t, err)
// 		err = r.Parse()
// 		assert.Nil(t, err)
// 		assert.Equal(t, t0.Ingredients, strings.TrimSpace(fmt.Sprint(r.IngredientList())))
// 	}
// }

// func ExampleChocolateChip3() {
// 	log.SetLevel("info")
// 	urlToParse := "https://www.bonappetit.com/recipe/bas-best-chocolate-chip-cookies"
// 	fileToGet := urlToParse
// 	fileToGet = strings.TrimPrefix(fileToGet, "https://")
// 	if string(fileToGet[len(fileToGet)-1]) == "/" {
// 		fileToGet += "index.html"
// 	}
// 	fileToGet = path.Join("testing", "sites", fileToGet)
// 	r, err := NewFromFile(fileToGet)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	err = r.Parse()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(r.IngredientList())
// 	// Output:
// 	// 1 1/4 tsp kosher salt (morton)
// 	// 3/4 tsp baking soda
// 	// 3/4 cup butter (unsalted)
// 	// 1 cup brown sugar (dark)
// 	// 1/4 cup granulated sugar
// 	// 1 whole egg
// 	// 2 whole egg yolks
// 	// 2 tsp vanilla
// 	// 6 oz chocolate chips (coarsely chopped or semisweet)
// }

// func ExampleChocolateChip4() {
// 	log.SetLevel("info")
// 	urlToParse := "https://pinchofyum.com/the-best-soft-chocolate-chip-cookies"
// 	fileToGet := urlToParse
// 	fileToGet = strings.TrimPrefix(fileToGet, "https://")
// 	if string(fileToGet[len(fileToGet)-1]) == "/" {
// 		fileToGet += "index.html"
// 	}
// 	fileToGet = path.Join("testing", "sites", fileToGet)
// 	r, err := NewFromFile(fileToGet)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	err = r.Parse()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(r.IngredientList())
// 	// Output:
// 	// 8 tablespoon butter (s of salted)
// 	// 1/2 cup white sugar
// 	// 1/4 cup brown sugar (packed light)
// 	// 1 teaspoon vanilla
// 	// 1 whole egg
// 	// 1 1/2 cup flour (s all purpose)
// 	// 1/2 teaspoon baking soda
// 	// 1/4 teaspoon salt
// 	// 3/4 cup chocolate chips

// }

// func ExampleChocolateChip5() {
// 	log.SetLevel("info")
// 	urlToParse := "https://www.modernhoney.com/the-best-chocolate-chip-cookies/"
// 	fileToGet := urlToParse
// 	fileToGet = strings.TrimPrefix(fileToGet, "https://")
// 	if string(fileToGet[len(fileToGet)-1]) == "/" {
// 		fileToGet += "index.html"
// 	}
// 	fileToGet = path.Join("testing", "sites", fileToGet)
// 	r, err := NewFromFile(fileToGet)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	err = r.Parse()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(r.IngredientList())
// 	// Output:
// 	// 1 cup butter (cold)
// 	// 1 cup brown sugar
// 	// 1/2 cup sugar (2 tablespoons)
// 	// 2 whole eggs
// 	// 2 teaspoons vanilla
// 	// 2 3/4 cups flour
// 	// 1 teaspoon cornstarch
// 	// 3/4 teaspoon baking soda
// 	// 3/4 teaspoon salt
// 	// 2 1/2 cups chocolate chips
// }

// func ExampleChocolateChip6() {
// 	log.SetLevel("info")
// 	urlToParse := "https://laurenslatest.com/actually-perfect-chocolate-chip-cookies/"
// 	fileToGet := urlToParse
// 	fileToGet = strings.TrimPrefix(fileToGet, "https://")
// 	if string(fileToGet[len(fileToGet)-1]) == "/" {
// 		fileToGet += "index.html"
// 	}
// 	fileToGet = path.Join("testing", "sites", fileToGet)
// 	r, err := NewFromFile(fileToGet)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	err = r.Parse()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(r.IngredientList())
// 	// Output:
// 	// 3/4 cup butter
// 	// 1 cup brown sugar (packed)
// 	// 1/2 cup granulated sugar
// 	// 1 whole egg yolk
// 	// 2 teaspoons vanilla
// 	// 2 cups flour (all purpose)
// 	// 1/2 teaspoon baking soda
// 	// 1/2 teaspoon salt
// 	// 1 cup chocolate chips (semi sweet)
// 	// 1 cup chocolate chips (milk)
// }

func TestNew(t *testing.T) {
	var urlToParse string
	urlToParse = "https://www.bonappetit.com/recipe/bas-best-chocolate-chip-cookies"
	urlToParse = "https://thesaltymarshmallow.com/best-banana-bread-recipe/"
	urlToParse = "https://pinchofyum.com/the-best-soft-chocolate-chip-cookies"
	urlToParse = "https://laurenslatest.com/actually-perfect-chocolate-chip-cookies/"
	urlToParse = "https://www.allrecipes.com/recipe/10813/best-chocolate-chip-cookies/"
	fileToGet := urlToParse
	fileToGet = strings.TrimPrefix(fileToGet, "https://")
	if string(fileToGet[len(fileToGet)-1]) == "/" {
		fileToGet += "index.html"
	}
	fileToGet = path.Join("testing", "sites", fileToGet)

	b, _ := ioutil.ReadFile(fileToGet)
	findIngredientsFromHTML(b)
	assert.Nil(t, nil)
}

func findIngredientsFromHTML(b []byte) {
	doc, err := html.Parse(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	var f func(n *html.Node, ingredientChildren *[]string) (s string)
	f = func(n *html.Node, ingredientChildren *[]string) (s string) {
		childrenText := []string{}
		// fmt.Printf("%+v\n", n)
		score := 0
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			var childText string
			childText = f(c, ingredientChildren)
			if childText != "" {
				childrenText = append(childrenText, childText)
				score += scoreLine(childText)
			}
		}
		if score > 5 && len(childrenText) < 15 && len(childrenText) > 1 {
			*ingredientChildren = append(*ingredientChildren, childrenText...)
			for _, child := range childrenText {
				fmt.Printf("[%s]\n", child)
			}
		}
		if len(childrenText) > 0 {
			// fmt.Println(childrenText)
			s = strings.Join(childrenText, " ")
		} else if n.DataAtom == 0 && strings.TrimSpace(n.Data) != "" {
			s = strings.TrimSpace(n.Data)
		}
		return s
	}
	var ingredientChildren []string
	f(doc, &ingredientChildren)
	fmt.Println(ingredientChildren)
}

func scoreLine(line string) (score int) {
	line = SanitizeLine(line)
	i := 0
	lineInfos := make([]LineInfo, 1)
	lineInfos[i].Line = line
	lineInfos[i].IngredientsInString = GetIngredientsInString(line)
	if len(lineInfos[i].LineOriginal) > 50 {
		return
	}
	if len(lineInfos[i].IngredientsInString) == 2 && len(lineInfos[i].IngredientsInString[1].Word) > len(lineInfos[i].IngredientsInString[0].Word) {
		lineInfos[i].IngredientsInString[0] = lineInfos[i].IngredientsInString[1]
	}
	lineInfos[i].AmountInString = GetNumbersInString(line)
	lineInfos[i].MeasureInString = GetMeasuresInString(line)

	// does it contain an ingredient?
	if len(lineInfos[i].IngredientsInString) > 0 {
		score++
	}

	// disfavor containing multiple ingredients
	if len(lineInfos[i].IngredientsInString) > 1 {
		score = score - len(lineInfos[i].IngredientsInString) + 1
	}

	// does it contain an amount?
	if len(lineInfos[i].AmountInString) > 0 {
		score++
	}
	// does it contain a measure (cups, tsps)?
	if len(lineInfos[i].MeasureInString) > 0 {
		score++
	}
	// does the ingredient come after the measure?
	if len(lineInfos[i].IngredientsInString) > 0 && len(lineInfos[i].MeasureInString) > 0 && lineInfos[i].IngredientsInString[0].Position > lineInfos[i].MeasureInString[0].Position {
		score++
	}
	// does the ingredient come after the amount?
	if len(lineInfos[i].IngredientsInString) > 0 && len(lineInfos[i].AmountInString) > 0 && lineInfos[i].IngredientsInString[0].Position > lineInfos[i].AmountInString[0].Position {
		score++
	}
	// does the measure come after the amount?
	if len(lineInfos[i].MeasureInString) > 0 && len(lineInfos[i].AmountInString) > 0 && lineInfos[i].MeasureInString[0].Position > lineInfos[i].AmountInString[0].Position {
		score++
	}

	// disfavor lots of puncuation
	puncuation := []string{".", ",", "!", "?"}
	for _, punc := range puncuation {
		if strings.Count(lineInfos[i].LineOriginal, punc) > 1 {
			score--
		}
	}

	// disfavor long lines
	if len(line) > 50 {
		score = score - (len(line) - 50)
	}

	// does it start with a list indicator (* or -)?
	fields := strings.Fields(line)
	if len(fields) > 0 && (fields[0] == "*" || fields[0] == "-") {
		score++
	}
	// if only one thing is right, its wrong
	if score == 1 {
		score = 0.0
	}
	return
}
