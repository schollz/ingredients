package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Create("corpus.go")
	check(err)
	defer f.Close()

	f.WriteString("package ingredients\n")

	var b []byte
	var pl pairList
	var i int

	// MAKE HERBS
	var herbList []string
	b, err = ioutil.ReadFile("corpus/herbs.json")
	if err != nil {
		panic(err)
	}
	if json.Unmarshal(b, &herbList) != nil {
		panic("could not unmarshal")
	}
	f.WriteString(`var herbMap = map[string]struct{}{` + "\n")
	pl = make(pairList, len(herbList))
	i = 0
	for _, k := range herbList {
		pl[i] = pair{k, len(k)}
		i++
	}
	sort.Slice(pl, func(i, j int) bool {
		return pl[i].Key < pl[j].Key
	})
	for _, p := range pl {
		f.WriteString(fmt.Sprintf(`"%s": {},`, p.Key) + "\n")
	}
	f.WriteString("}\n\n")

	// MAKE FRUITS
	var fruitList []string
	b, err = ioutil.ReadFile("corpus/fruits.json")
	if err != nil {
		panic(err)
	}
	if json.Unmarshal(b, &fruitList) != nil {
		panic("could not unmarshal")
	}
	f.WriteString(`var fruitMap = map[string]struct{}{` + "\n")
	pl = make(pairList, len(fruitList))
	i = 0
	for _, k := range fruitList {
		pl[i] = pair{k, len(k)}
		i++
	}
	sort.Slice(pl, func(i, j int) bool {
		return pl[i].Key < pl[j].Key
	})
	for _, p := range pl {
		f.WriteString(fmt.Sprintf(`"%s": {},`, p.Key) + "\n")
	}
	f.WriteString("}\n\n")

	// MAKE VEGETABLES
	var vegetableList []string
	b, err = ioutil.ReadFile("corpus/vegetables.json")
	if err != nil {
		panic(err)
	}
	if json.Unmarshal(b, &vegetableList) != nil {
		panic("could not unmarshal")
	}
	f.WriteString(`var vegetableMap = map[string]struct{}{` + "\n")
	pl = make(pairList, len(vegetableList))
	i = 0
	for _, k := range vegetableList {
		pl[i] = pair{k, len(k)}
		i++
	}
	sort.Slice(pl, func(i, j int) bool {
		return pl[i].Key < pl[j].Key
	})
	for _, p := range pl {
		f.WriteString(fmt.Sprintf(`"%s": {},`, p.Key) + "\n")
	}
	f.WriteString("}\n\n")

	// MAIN INGREDIENT LIST
	// sort the ingredient corpus by the length of each term
	// and then by alphabetizing
	b, err = ioutil.ReadFile("corpus/ingredients.txt")
	corpusIngredients := strings.Split(string(b), "\n")
	ingredientSizes := make(map[string]int)
	for _, ing := range corpusIngredients {
		ing = strings.TrimSpace(ing)
		if len(ing) == 0 {
			continue
		}
		ingredientSizes[ing] = len(ing)
		ingredientSizes[ing+"s"] = len(ing) + 1
	}
	for _, ing := range fruitList {
		ing = strings.TrimSpace(ing)
		if len(ing) == 0 {
			continue
		}
		ingredientSizes[ing] = len(ing)
		ingredientSizes[ing+"s"] = len(ing) + 1
	}
	for _, ing := range herbList {
		ing = strings.TrimSpace(ing)
		if len(ing) == 0 {
			continue
		}
		ingredientSizes[ing] = len(ing)
		ingredientSizes[ing+"s"] = len(ing) + 1
	}
	for _, ing := range vegetableList {
		ing = strings.TrimSpace(ing)
		if len(ing) == 0 {
			continue
		}
		ingredientSizes[ing] = len(ing)
		ingredientSizes[ing+"s"] = len(ing) + 1
	}

	pl = make(pairList, len(ingredientSizes))
	i = 0
	for k, v := range ingredientSizes {
		pl[i] = pair{k, v}
		i++
	}
	sort.Slice(pl, func(i, j int) bool {
		if pl[i].Value == pl[j].Value {
			return pl[i].Key < pl[j].Key
		}
		return pl[i].Value > pl[j].Value
	})

	corpusIngredients = make([]string, len(pl))
	for i, p := range pl {
		corpusIngredients[i] = " " + strings.TrimSpace(p.Key) + " "
	}
	f.WriteString(`var corpusIngredients = []string{"` + strings.Join(corpusIngredients, `"`+",\n"+`"`) + `"}` + "\n")
	f.Sync()

	// MAIN MEASURE LIST
	pl = make(pairList, len(corpusMeasuresMap))
	i = 0
	for k := range corpusMeasuresMap {
		pl[i] = pair{k, len(k)}
		i++
	}
	sort.Slice(pl, func(i, j int) bool {
		if pl[i].Value == pl[j].Value {
			return pl[i].Key < pl[j].Key
		}
		return pl[i].Value > pl[j].Value
	})
	corpusMeasures := make([]string, len(pl))
	for i, p := range pl {
		corpusMeasures[i] = " " + p.Key + " "
	}
	f.WriteString(`var corpusMeasures = []string{"` + strings.Join(corpusMeasures, `"`+",\n"+`"`) + `"}` + "\n")
	f.Sync()

	// MAKE NUMBERS
	b, err = ioutil.ReadFile("corpus/numbers.txt")
	corpusNumbers := strings.Split(string(b), "\n")
	for v := range corpusFractionNumberMap {
		corpusNumbers = append(corpusNumbers, v)
	}
	for i, c := range corpusNumbers {
		corpusNumbers[i] = " " + strings.TrimSpace(c) + " "
	}
	f.WriteString(`var corpusNumbers = []string{"` + strings.Join(corpusNumbers, `"`+",\n"+`"`) + `"}` + "\n")
	f.Sync()

	// MAKE DIRECTIONS CORPUS
	b, err = ioutil.ReadFile("corpus/directions_pos.txt")
	corpusDirections := strings.Fields(string(b))
	corpusDirectionsMap := make(map[string]struct{})
	for _, c := range corpusDirections {
		corpusDirectionsMap[" "+strings.ToLower(c)+" "] = struct{}{}
	}
	pl = make(pairList, len(corpusDirectionsMap))
	i = 0
	for k := range corpusDirectionsMap {
		pl[i] = pair{k, len(k)}
		i++
	}
	sort.Slice(pl, func(i, j int) bool {
		return pl[i].Key < pl[j].Key
	})
	corpusDirections = make([]string, len(pl))
	for i, p := range pl {
		corpusDirections[i] = p.Key
	}
	f.WriteString(`var corpusDirections = []string{"` + strings.Join(corpusDirections, `"`+",\n"+`"`) + `"}` + "\n")
	f.Sync()

	// MAKE DIRECTIONS NEG CORPUS
	b, err = ioutil.ReadFile("corpus/directions_neg.txt")
	corpusDirections = strings.Fields(string(b))
	corpusDirectionsMap = make(map[string]struct{})
	for _, c := range corpusDirections {
		corpusDirectionsMap[" "+strings.ToLower(c)+" "] = struct{}{}
	}
	pl = make(pairList, len(corpusDirectionsMap))
	i = 0
	for k := range corpusDirectionsMap {
		pl[i] = pair{k, len(k)}
		i++
	}
	sort.Slice(pl, func(i, j int) bool {
		return pl[i].Key < pl[j].Key
	})
	corpusDirections = make([]string, len(pl))
	for i, p := range pl {
		corpusDirections[i] = p.Key
	}
	f.WriteString(`var corpusDirectionsNeg = []string{"` + strings.Join(corpusDirections, `"`+",\n"+`"`) + `"}` + "\n")
	f.Sync()

	f.WriteString(`type fractionNumber struct {
		fractionString string
		value          float64
	}
	`)
	f.WriteString(`var corpusFractionNumberMap = map[string]fractionNumber{` + "\n")
	pl = make(pairList, len(corpusFractionNumberMap))
	i = 0
	for k := range corpusFractionNumberMap {
		pl[i] = pair{k, len(k)}
		i++
	}
	sort.Slice(pl, func(i, j int) bool {
		return pl[i].Key < pl[j].Key
	})
	for _, p := range pl {
		f.WriteString(fmt.Sprintf(`"%s": fractionNumber{"%s",%10.10f},`, p.Key, corpusFractionNumberMap[p.Key].fractionString, corpusFractionNumberMap[p.Key].value) + "\n")
	}
	f.WriteString("}\n\n")

	f.WriteString(`var corpusMeasuresMap = map[string]string{` + "\n")
	pl = make(pairList, len(corpusMeasuresMap))
	i = 0
	for k := range corpusMeasuresMap {
		pl[i] = pair{k, len(k)}
		i++
	}
	sort.Slice(pl, func(i, j int) bool {
		return pl[i].Key < pl[j].Key
	})
	for _, p := range pl {
		f.WriteString(fmt.Sprintf(`"%s": "%s",`, p.Key, corpusMeasuresMap[p.Key]) + "\n")
	}
	f.WriteString("}\n\n")

	// MAKE DENSITIES
	var densities map[string]float64
	b, err = ioutil.ReadFile("corpus/densities.json")
	if err != nil {
		panic(err)
	}
	if json.Unmarshal(b, &densities) != nil {
		panic("could not unmarshal")
	}
	f.WriteString(`var densities = map[string]float64{` + "\n")
	pl = make(pairList, len(densities))
	i = 0
	for k := range densities {
		pl[i] = pair{k, len(k)}
		i++
	}
	sort.Slice(pl, func(i, j int) bool {
		return pl[i].Key < pl[j].Key
	})
	for _, p := range pl {
		f.WriteString(fmt.Sprintf(`"%s": %2.10f,`, p.Key, densities[p.Key]) + "\n")
	}
	f.WriteString("}\n\n")

}

type pair struct {
	Key   string
	Value int
}

type pairList []pair

func (p pairList) Len() int           { return len(p) }
func (p pairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p pairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type fractionNumber struct {
	fractionString string
	value          float64
}

var corpusFractionNumberMap = map[string]fractionNumber{
	"½": fractionNumber{"1/2", 1.0 / 2},
	"¼": fractionNumber{"1/4", 1.0 / 4},
	"¾": fractionNumber{"3/4", 3.0 / 4},
	"⅛": fractionNumber{"1/8", 1.0 / 8},
	"⅜": fractionNumber{"3/8", 3.0 / 8},
	"⅝": fractionNumber{"5/8", 5.0 / 8},
	"⅞": fractionNumber{"7/8", 7.0 / 8},
	"⅔": fractionNumber{"2/3", 2.0 / 3},
	"⅓": fractionNumber{"1/3", 1.0 / 3},
}

var corpusMeasuresMap = map[string]string{
	"tablespoon":  "tbl",
	"tablespoons": "tbl",
	"tbl":         "tbl",
	"tbsp":        "tbl",
	"tblsp":       "tbl",
	"tbsps":       "tbl",
	"tbls":        "tbl",
	"tbs":         "tbl",
	"teaspoons":   "tsp",
	"teaspoon":    "tsp",
	"tsp":         "tsp",
	"t":           "tsp",
	"tsps":        "tsp",
	"cups":        "cup",
	"cup":         "cup",
	"c":           "cup",
	"ounces":      "ounce",
	"ounce":       "ounce",
	"oz":          "ounce",
	"grams":       "gram",
	"g":           "gram",
	"gram":        "gram",
	"milliliter":  "milliliter",
	"ml":          "milliliter",
	"pint":        "pint",
	"pints":       "pint",
	"quart":       "quart",
	"quarts":      "quart",
	"pound":       "pound",
	"pounds":      "pound",
	"cans":        "can",
	"canned":      "can",
	"can":         "can",
}
