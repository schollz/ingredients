package recipeingredients

//go:generate go run corpus/main.go
//go:generate gofmt -s -w corpus.go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"

	log "github.com/schollz/logger"
	"golang.org/x/net/html"
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
	b, err := ioutil.ReadFile(fname)
	r.FileContent = string(b)
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
	r.FileContent = string(html)
	return
}

// Parse is the main parser for a given recipe.
func (r *Recipe) Parse() (rerr error) {
	if r == nil {
		r = &Recipe{}
	}
	if r.FileContent == "" || r.FileName == "" {
		rerr = fmt.Errorf("no file loaded")
		return
	}

	r.Lines, rerr = GetIngredientLinesInHTML(r.FileContent)

	goodLines := make([]LineInfo,len(r.Lines))
	j := 0
	for _,lineInfo := range r.Lines {
		if len(strings.TrimSpace(lineInfo.Line)) < 3 {
			continue
		}

		lineInfo.Ingredient.Measure = Measure{}

		// get amount, continue if there is an error
		err := lineInfo.getTotalAmount()
		if err != nil {
			log.Tracef("[%s]: %s (%+v)", lineInfo.Line, err.Error(),lineInfo.AmountInString)
			continue
		}

		// get ingredient, continue if its not found
		err = lineInfo.getIngredient()
		if err != nil {
			log.Tracef("[%s]: %s", lineInfo.Line, err.Error())
			continue
		}

		// get measure
		err = lineInfo.getMeasure()
		if err != nil {
			log.Tracef("[%s]: %s", lineInfo.Line, err.Error())
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
			log.Tracef("[%s]: %s", lineInfo.LineOriginal, err.Error())
		} else {
			log.Tracef("[%s]: %+v", lineInfo.LineOriginal, lineInfo)
		}

	goodLines[j] = lineInfo
	j++
	}
	r.Lines = goodLines[:j]

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

func GetIngredientLinesInHTML(htmlS string) (lineInfos []LineInfo, err error) {
	doc, err := html.Parse(bytes.NewReader([]byte(htmlS)))
	if err != nil {
		return
	}
	var f func(n *html.Node, lineInfos *[]LineInfo) (s string)
	f = func(n *html.Node, lineInfos *[]LineInfo) (s string) {
		childrenLineInfo := []LineInfo{}
		// fmt.Printf("%+v\n", n)
		score := 0
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			var childText string
			childText = f(c, lineInfos)
			if childText != "" {
				scoreOfLine, lineInfo := scoreLine(childText)
				childrenLineInfo = append(childrenLineInfo, lineInfo)
				score += scoreOfLine
			}
		}
		if score > 5 && len(childrenLineInfo) < 15 && len(childrenLineInfo) > 1 {
			*lineInfos = append(*lineInfos, childrenLineInfo...)
			for _, child := range childrenLineInfo {
				log.Tracef("[%s]", child.LineOriginal)
			}
		}
		if len(childrenLineInfo) > 0 {
			// fmt.Println(childrenLineInfo)
			childrenText := make([]string, len(childrenLineInfo))
			for i := range childrenLineInfo {
				childrenText[i] = childrenLineInfo[i].LineOriginal
			}
			s = strings.Join(childrenText, " ")
		} else if n.DataAtom == 0 && strings.TrimSpace(n.Data) != "" {
			s = strings.TrimSpace(n.Data)
		}
		return s
	}
	f(doc, &lineInfos)
	return
}

func scoreLine(line string) (score int, lineInfo LineInfo) {
	lineInfo = LineInfo{}
	lineInfo.LineOriginal = line
	lineInfo.Line = SanitizeLine(line)

	lineInfo.IngredientsInString = GetIngredientsInString(lineInfo.Line)
	lineInfo.AmountInString = GetNumbersInString(lineInfo.Line)
	lineInfo.MeasureInString = GetMeasuresInString(lineInfo.Line)
	if len(lineInfo.IngredientsInString) == 2 && len(lineInfo.IngredientsInString[1].Word) > len(lineInfo.IngredientsInString[0].Word) {
		lineInfo.IngredientsInString[0] = lineInfo.IngredientsInString[1]
	}

	if len(lineInfo.LineOriginal) > 50 {
		return
	}

	// does it contain an ingredient?
	if len(lineInfo.IngredientsInString) > 0 {
		score++
	}

	// disfavor containing multiple ingredients
	if len(lineInfo.IngredientsInString) > 1 {
		score = score - len(lineInfo.IngredientsInString) + 1
	}

	// does it contain an amount?
	if len(lineInfo.AmountInString) > 0 {
		score++
	}
	// does it contain a measure (cups, tsps)?
	if len(lineInfo.MeasureInString) > 0 {
		score++
	}
	// does the ingredient come after the measure?
	if len(lineInfo.IngredientsInString) > 0 && len(lineInfo.MeasureInString) > 0 && lineInfo.IngredientsInString[0].Position > lineInfo.MeasureInString[0].Position {
		score++
	}
	// does the ingredient come after the amount?
	if len(lineInfo.IngredientsInString) > 0 && len(lineInfo.AmountInString) > 0 && lineInfo.IngredientsInString[0].Position > lineInfo.AmountInString[0].Position {
		score++
	}
	// does the measure come after the amount?
	if len(lineInfo.MeasureInString) > 0 && len(lineInfo.AmountInString) > 0 && lineInfo.MeasureInString[0].Position > lineInfo.AmountInString[0].Position {
		score++
	}

	// disfavor lots of puncuation
	puncuation := []string{".", ",", "!", "?"}
	for _, punc := range puncuation {
		if strings.Count(lineInfo.LineOriginal, punc) > 1 {
			score--
		}
	}

	// disfavor long lines
	if len(lineInfo.Line) > 50 {
		score = score - (len(lineInfo.Line) - 50)
	}

	// does it start with a list indicator (* or -)?
	fields := strings.Fields(lineInfo.Line)
	if len(fields) > 0 && (fields[0] == "*" || fields[0] == "-") {
		score++
	}
	// if only one thing is right, its wrong
	if score == 1 {
		score = 0.0
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
