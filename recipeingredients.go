package recipeingredients

//go:generate go run corpus/main.go
//go:generate gofmt -s -w corpus.go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strings"

	"github.com/jaytaylor/html2text"
	log "github.com/schollz/logger"
)

// Recipe contains the info for the file and the lines
type Recipe struct {
	FileName    string                        `json:"filename"`
	FileContent string                        `json:"file_content"`
	Lines       []LineInfo                    `json:"lines"`
	Directions  []string                      `json:"directions"`
	Ingredients []Ingredient                  `json:"ingredients"`
	Ratios      map[string]map[string]float64 `json:"ratios"`
}

// LineInfo has all the information for the parsing of a given line
type LineInfo struct {
	LineOriginal        string
	Line                string         `json:",omitempty"`
	IngredientsInString []WordPosition `json:",omitempty"`
	AmountInString      []WordPosition `json:",omitempty"`
	MeasureInString     []WordPosition `json:",omitempty"`
	Ingredient          Ingredient     `json:",omitempty"`
}

// Ingredient is the basic struct for ingredients
type Ingredient struct {
	Name      string  `json:"name,omitempty"`
	Comment   string  `json:"comment,omitempty"`
	Measure   Measure `json:"measure,omitempty"`
	Frequency float64 `json:"frequency,omitempty`
}

// Measure includes the amount, name and the cups for conversions
type Measure struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
	Cups   float64 `json:"cups"`
	Weight float64 `json:"weight"`
}

// IngredientList is a list of ingredients
type IngredientList struct {
	Ingredients []Ingredient `json:"ingredients"`
}

func (il IngredientList) String() string {
	s := ""
	for _, ing := range il.Ingredients {
		s += fmt.Sprintf("%s %s %s", AmountToString(ing.Measure.Amount), ing.Measure.Name, ing.Name)
		if ing.Comment != "" {
			s += " (" + ing.Comment + ")"
		}
		s += "\n"
	}
	return s
}

// Save saves the recipe to a file
func (r *Recipe) Save(fname string) (err error) {
	b, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		return
	}
	err = ioutil.WriteFile(fname, b, 0644)
	return
}

// Load will load a recipe file
func Load(fname string) (r *Recipe, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return
	}
	r = new(Recipe)
	err = json.Unmarshal(b, r)
	return
}

// NewFromFile generates a new parser from a file
func NewFromFile(fname string) (r *Recipe, err error) {
	r = &Recipe{FileName: fname}
	_, err = os.Stat(fname)
	return
}

// NewFromURL generates a new parser from a url
func NewFromURL(url string) (r *Recipe, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	r = &Recipe{FileName: url}
	r.FileContent, err = html2text.FromString(string(html), html2text.Options{PrettyTables: false, OmitLinks: true})

	return
}

// Parse is the main parser for a given recipe.
// It looks for the following
// - Contains number
// - Contains mass/volume
// - Contains ingredient
// - Number occurs before ingredient
// - Number occurs before mass/volume
// - Number of ingredients is 1
// - Percent of other words is less than 50%
// - Part of list (contains - or *)
func (r *Recipe) Parse() (rerr error) {
	if r == nil {
		r = &Recipe{}
	}
	if r.FileContent == "" && r.FileName != "" {
		var bFile []byte
		bFile, rerr = ioutil.ReadFile(r.FileName)
		if rerr != nil {
			return
		}
		r.FileContent, rerr = html2text.FromString(string(bFile), html2text.Options{PrettyTables: false, OmitLinks: true})
		if rerr != nil {
			return
		}
	}

	lines := strings.Split(r.FileContent, "\n")
	scores := make([]float64, len(lines))
	lineInfos := make([]LineInfo, len(lines))
	i := -1
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		i++
		lineInfos[i].LineOriginal = line
		line = SanitizeLine(line)
		lineInfos[i].Line = line
		lineInfos[i].IngredientsInString = GetIngredientsInString(line)
		if len(lineInfos[i].IngredientsInString) ==2 && len(lineInfos[i].IngredientsInString[1].Word) > len(lineInfos[i].IngredientsInString[0].Word) {
			lineInfos[i].IngredientsInString[0] = lineInfos[i].IngredientsInString[1]
		}
		lineInfos[i].AmountInString = GetNumbersInString(line)
		lineInfos[i].MeasureInString = GetMeasuresInString(line)

		score := 0.0
		// does it contain an ingredient?
		if len(lineInfos[i].IngredientsInString) > 0 {
			score++
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
		// is the line really long? (ingredient lines are short)
		if score > 0 && len(lineInfos[i].LineOriginal) > 100 {
			score--
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
		// log.Debugf("'%s' (%2.1f)", line, score)
		scores[i] = score
	}
	scores = scores[:i+1]
	lineInfos = lineInfos[:i+1]

	// //debugging purposes
	// lines = make([]string, len(lineInfos))
	// for i, li := range lineInfos {
	// 	lines[i] = li.Line
	// }
	// ioutil.WriteFile("out", []byte(strings.Join(lines, "\n")), 0644)

	// get the most likely location
	start, end := GetBestTopHatPositions(scores)

	log.Debugf("got %d line infos, %d - %d", len(lineInfos), start, end)
	if start < 3 || (end <= start) {
		rerr = fmt.Errorf("no recipe to parse")
		return
	}
	r.Lines = []LineInfo{}
	for _, lineInfo := range lineInfos[start-3 : end+3] {
		if len(strings.TrimSpace(lineInfo.Line)) < 3 {
			continue
		}

		lineInfo.Ingredient.Measure = Measure{}

		// get amount, continue if there is an error
		err := lineInfo.getTotalAmount()
		if err != nil {
			log.Debugf("[%s]: %s", lineInfo.LineOriginal, err.Error())
			continue
		}

		// get ingredient, continue if its not found
		err = lineInfo.getIngredient()
		if err != nil {
			log.Debugf("[%s]: %s", lineInfo.LineOriginal, err.Error())
			continue
		}

		// get measure
		err = lineInfo.getMeasure()
		if err != nil {
			log.Debugf("[%s]: %s", lineInfo.LineOriginal, err.Error())
		}

		// get comment
		if len(lineInfo.MeasureInString) > 0 && len(lineInfo.IngredientsInString) > 0 {
			lineInfo.Ingredient.Comment = getOtherInBetweenPositions(lineInfo.Line, lineInfo.MeasureInString[0], lineInfo.IngredientsInString[0])
		}

		// normalize into cups
		lineInfo.Ingredient.Measure.Cups, err = normalizeIngredient(
			lineInfo.Ingredient.Name,
			lineInfo.Ingredient.Measure.Name,
			lineInfo.Ingredient.Measure.Amount,
		)
		if err != nil {
			log.Debugf("[%s]: %s", lineInfo.LineOriginal, err.Error())
		} else {
			log.Debugf("[%s]: %+v", lineInfo.LineOriginal, lineInfo)
		}

		r.Lines = append(r.Lines, lineInfo)
	}
	rerr = r.ConvertIngredients()
	if rerr != nil {
		return
	}

	// consolidate ingredients
	ingredients := make(map[string]Ingredient)
	ingredientList := []string{}
	for _, line := range r.Lines {
		if _, ok := ingredients[line.Ingredient.Name]; ok {
			if ingredients[line.Ingredient.Name].Measure.Name == line.Ingredient.Measure.Name {
				ingredients[line.Ingredient.Name] = Ingredient{
					Name:    line.Ingredient.Name,
					Comment: ingredients[line.Ingredient.Name].Comment,
					Measure: Measure{
						Name:   ingredients[line.Ingredient.Name].Measure.Name,
						Amount: ingredients[line.Ingredient.Name].Measure.Amount + line.Ingredient.Measure.Amount,
						Cups:   ingredients[line.Ingredient.Name].Measure.Cups + line.Ingredient.Measure.Cups,
					},
				}
			} else {
				ingredients[line.Ingredient.Name] = Ingredient{
					Name:    line.Ingredient.Name,
					Comment: ingredients[line.Ingredient.Name].Comment,
					Measure: Measure{
						Name:   ingredients[line.Ingredient.Name].Measure.Name,
						Amount: ingredients[line.Ingredient.Name].Measure.Amount,
						Cups:   ingredients[line.Ingredient.Name].Measure.Cups + line.Ingredient.Measure.Cups,
					},
				}
				log.Debugf("different measure!")
			}
		} else {
			ingredientList = append(ingredientList, line.Ingredient.Name)
			ingredients[line.Ingredient.Name] = Ingredient{
				Name:    line.Ingredient.Name,
				Comment: line.Ingredient.Comment,
				Measure: Measure{
					Name:   line.Ingredient.Measure.Name,
					Amount: line.Ingredient.Measure.Amount,
					Cups:   line.Ingredient.Measure.Cups + line.Ingredient.Measure.Cups,
				},
			}
		}
	}
	r.Ingredients = make([]Ingredient, len(ingredients))
	for i, ing := range ingredientList {
		r.Ingredients[i] = ingredients[ing]
	}

	return
}

func (r *Recipe) ConvertIngredients() (err error) {

	return
}

func (r *Recipe) PrintIngredientList() string {
	s := ""
	for _, ing := range r.Ingredients {
		if ing.Frequency > 0 {
			s += fmt.Sprintf("%s %s %s (%2.1f%%)\n", AmountToString(ing.Measure.Amount), ing.Measure.Name, ing.Name, 100*ing.Frequency)
		} else {
			s += fmt.Sprintf("%s %s %s\n", AmountToString(ing.Measure.Amount), ing.Measure.Name, ing.Name)
		}
	}
	return s
}

func (r *Recipe) PrintDirections() string {
	s := ""
	for i, d := range r.Directions {
		s += fmt.Sprintf("%d) %s\n", i+1, d)
	}
	return s
}

// IngredientList will return a string containing the ingredient list
func (r *Recipe) IngredientList() (ingredientList IngredientList) {
	ingredientList = IngredientList{make([]Ingredient, len(r.Lines))}
	for i, li := range r.Lines {
		ingredientList.Ingredients[i] = li.Ingredient
	}
	return
}

func (lineInfo *LineInfo) getTotalAmount() (err error) {
	lastPosition := -1
	totalAmount := 0.0
	wps := lineInfo.AmountInString
	for i := range wps {
		wps[i].Word = strings.TrimSpace(wps[i].Word)
		if lastPosition == -1 {
			totalAmount = ConvertStringToNumber(wps[i].Word)
		} else if math.Abs(float64(wps[i].Position-lastPosition)) < 6 {
			totalAmount += ConvertStringToNumber(wps[i].Word)
		}
		lastPosition = wps[i].Position + len(wps[i].Word)
	}
	if totalAmount == 0 && strings.Contains(lineInfo.Line, "whole") {
		totalAmount = 1
	}
	if totalAmount == 0 {
		err = fmt.Errorf("no amount found")
	} else {
		lineInfo.Ingredient.Measure.Amount = totalAmount
	}
	return
}

func (lineInfo *LineInfo) getIngredient() (err error) {
	if len(lineInfo.IngredientsInString) == 0 {
		err = fmt.Errorf("no ingredient found")
		return
	}
	lineInfo.Ingredient.Name = lineInfo.IngredientsInString[0].Word
	return
}

func (lineInfo *LineInfo) getMeasure() (err error) {
	if len(lineInfo.MeasureInString) == 0 {
		lineInfo.Ingredient.Measure.Name = "whole"
		return
	}
	lineInfo.Ingredient.Measure.Name = lineInfo.MeasureInString[0].Word
	return
}
