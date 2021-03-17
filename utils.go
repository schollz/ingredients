package ingredients

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"

	log "github.com/schollz/logger"
)

// ConvertStringToNumber
func ConvertStringToNumber(s string) float64 {
	switch s {
	case "½":
		return 0.5
	case "¼":
		return 0.25
	case "¾":
		return 0.75
	case "⅛":
		return 1.0 / 8
	case "⅜":
		return 3.0 / 8
	case "⅝":
		return 5.0 / 8
	case "⅞":
		return 7.0 / 8
	case "⅔":
		return 2.0 / 3
	case "⅓":
		return 1.0 / 3
	}
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

func AmountToString(amount float64) string {
	r, _ := parseDecimal(fmt.Sprintf("%2.10f", amount))
	rationalFraction := float64(r.n) / float64(r.d)
	if rationalFraction > 0 {
		bestFractionDiff := 1e9
		bestFraction := 0.0
		var fractions = map[float64]string{
			0:       "",
			1:       "",
			1.0 / 2: "1/2",
			1.0 / 3: "1/3",
			2.0 / 3: "2/3",
			1.0 / 6: "1/6",
			1.0 / 8: "1/8",
			3.0 / 8: "3/8",
			5.0 / 8: "5/8",
			7.0 / 8: "7/8",
			1.0 / 4: "1/4",
			3.0 / 4: "3/4",
		}
		for f := range fractions {
			currentDiff := math.Abs(f - rationalFraction)
			if currentDiff < bestFractionDiff {
				bestFraction = f
				bestFractionDiff = currentDiff
			}
		}
		if fractions[bestFraction] == "" {
			return strconv.FormatInt(int64(math.Round(amount)), 10)
		}
		if r.i > 0 {
			return strconv.FormatInt(r.i, 10) + " " + fractions[bestFraction]
		} else {
			return fractions[bestFraction]
		}
	}
	return strconv.FormatInt(r.i, 10)
}

// A rational number r is expressed as the fraction p/q of two integers:
// r = p/q = (d*i+n)/d.
type rational struct {
	i int64 // integer
	n int64 // fraction numerator
	d int64 // fraction denominator
}

func gcd(x, y int64) int64 {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func parseDecimal(s string) (r rational, err error) {
	sign := int64(1)
	if strings.HasPrefix(s, "-") {
		sign = -1
	}
	p := strings.IndexByte(s, '.')
	if p < 0 {
		p = len(s)
	}
	if i := s[:p]; len(i) > 0 {
		if i != "+" && i != "-" {
			r.i, err = strconv.ParseInt(i, 10, 64)
			if err != nil {
				return rational{}, err
			}
		}
	}
	if p >= len(s) {
		p = len(s) - 1
	}
	if f := s[p+1:]; len(f) > 0 {
		n, err := strconv.ParseUint(f, 10, 64)
		if err != nil {
			return rational{}, err
		}
		d := math.Pow10(len(f))
		if math.Log2(d) > 63 {
			err = fmt.Errorf(
				"ParseDecimal: parsing %q: value out of range", f,
			)
			return rational{}, err
		}
		r.n = int64(n)
		r.d = int64(d)
		if g := gcd(r.n, r.d); g != 0 {
			r.n /= g
			r.d /= g
		}
		r.n *= sign
	}
	return r, nil
}

// GetIngredientsInString returns the word positions of the ingredients
func GetIngredientsInString(s string) (wordPositions []WordPosition) {
	return getWordPositions(s, corpusIngredients)
}

// GetNumbersInString returns the word positions of the numbers in the ingredient string
func GetNumbersInString(s string) (wordPositions []WordPosition) {
	return getWordPositions(s, corpusNumbers)
}

// GetMeasuresInString returns the word positions of the measures in a ingredient string
func GetMeasuresInString(s string) (wordPositions []WordPosition) {
	return getWordPositions(s, corpusMeasures)
}

// WordPosition shows a word and its position
// Note: the position is memory-dependent as it will
// be the position after the last deleted word
type WordPosition struct {
	Word     string
	Position int
}

func getWordPositions(s string, corpus []string) (wordPositions []WordPosition) {
	wordPositions = []WordPosition{}
	for _, ing := range corpus {
		pos := strings.Index(s, ing)
		if pos > -1 {
			s = strings.Replace(s, ing, strings.Repeat(" ", utf8.RuneCountInString(ing)), 1)
			ing = strings.TrimSpace(ing)
			wordPositions = append(wordPositions, WordPosition{ing, pos})
		}
	}
	sort.Slice(wordPositions, func(i, j int) bool {
		return wordPositions[i].Position < wordPositions[j].Position
	})
	return
}

// getOtherInBetweenPositions returns the word positions comment string in the ingredients
func getOtherInBetweenPositions(s string, pos1, pos2 WordPosition) (other string) {
	if pos1.Position > pos2.Position {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			log.Error(s, pos1, pos2)
			log.Error(r)
		}
	}()
	other = s[pos1.Position+len(pos1.Word)+1 : pos2.Position]
	other = strings.TrimSpace(other)
	return
}

// SanitizeLine removes parentheses, trims the line, converts to lower case,
// replaces fractions with unicode and then does special conversion for ingredients (like eggs).
func SanitizeLine(s string) string {
	s = strings.ToLower(s)
	s = strings.Replace(s, "⁄", "/", -1)
	s = strings.Replace(s, " / ", "/", -1)

	// special cases
	s = strings.Replace(s, "butter milk", "buttermilk", -1)
	s = strings.Replace(s, "bicarbonate of soda", "baking soda", -1)
	s = strings.Replace(s, "soda bicarbonate", "baking soda", -1)

	// remove parentheses
	re := regexp.MustCompile(`(?s)\((.*)\)`)
	for _, m := range re.FindAllStringSubmatch(s, -1) {
		s = strings.Replace(s, m[0], " ", 1)
	}

	s = " " + strings.TrimSpace(s) + " "

	// replace unicode fractions with fractions
	for v := range corpusFractionNumberMap {
		s = strings.Replace(s, v, " "+corpusFractionNumberMap[v].fractionString+" ", -1)
	}

	// remove non-alphanumeric
	reg, _ := regexp.Compile("[^a-zA-Z0-9/]+")
	s = reg.ReplaceAllString(s, " ")

	// replace fractions with unicode fractions
	for v := range corpusFractionNumberMap {
		s = strings.Replace(s, corpusFractionNumberMap[v].fractionString, " "+v+" ", -1)
	}

	s = strings.Replace(s, " one ", " 1 ", -1)

	return s
}

var gramConversions = map[string]float64{
	"ounce": 28.3495,
	"gram":  1,
	"pound": 453.592,
}

var conversionToCup = map[string]float64{
	"tbl":        0.0625,
	"tsp":        0.020833,
	"cup":        1.0,
	"pint":       2.0,
	"quart":      4.0,
	"gallon":     16.0,
	"milliliter": 0.00423,
	"can":        1.75,
}
var ingredientToCups = map[string]float64{
	"eggs":    0.25,
	"egg":     0.25,
	"garlic":  0.0280833,
	"chicken": 3,
	"celery":  0.5,
	"onion":   1,
	"carrot":  1,
	"butter":  0.5,
}

func cupsToOther(cups float64, ingredient string) (amount float64, measure string) {
	if _, ok := ingredientToCups[ingredient]; ok {
		measure = "whole"
		amount = math.Round(cups / ingredientToCups[ingredient])
		return
	}
	if cups > 0.125 {
		amount = cups
		measure = "cup"
	} else if cups > 0.020833*3 {
		amount = cups * 16
		measure = "tbl"
	} else {
		amount = cups * 48
		measure = "tsp"
	}
	if math.IsInf(amount, 0) {
		amount = 0
	}

	return
}

// normalizeIngredient will try to normalize the ingredient to 1 cup
func normalizeIngredient(ingredient, measure string, amount float64) (cups float64, err error) {
	// convert measure to standard measure
	newMeasure, ok := corpusMeasuresMap[measure]
	if !ok && measure != "whole" {
		err = fmt.Errorf("could not find '%s'", measure)
		return
	}
	measure = newMeasure
	if _, ok := ingredientToCups[ingredient]; ok && measure == "" {
		// special ingredients
		cups = amount * ingredientToCups[ingredient]
	} else if _, ok := conversionToCup[measure]; ok {
		// check if it has a standard volume measurement
		cups = float64(amount) * conversionToCup[measure]
	} else if _, ok := gramConversions[measure]; ok {
		// check if it has a standard weight measurement
		var density float64
		density, ok = densities[ingredient]
		if !ok {
			density = 200 // grams / cup
		}
		cups = amount * gramConversions[measure] / density
	} else {
		if _, ok := fruitMap[ingredient]; ok {
			cups = 1 * amount
		} else if _, ok := vegetableMap[ingredient]; ok {
			cups = 1 * amount
		} else if _, ok := herbMap[ingredient]; ok {
			cups = 0.0208333 * amount
		} else {
			err = errors.New("could not convert weight or volume")
		}
	}
	return
}

func determineMeasurementsFromCups(cups float64) (amount float64, measure string, amountString string, err error) {
	if cups > 0.125 {
		amount = cups
		measure = "cup"
	} else if cups > 0.020833*3 {
		amount = cups * 16
		measure = "tablespoon"
	} else {
		amount = cups * 48
		measure = "teaspoon"
	}
	amountString = AmountToString(amount)
	if math.IsInf(amount, 0) {
		amount = 0
	}
	if math.IsInf(cups, 0) {
		cups = 0
	}
	return
}
