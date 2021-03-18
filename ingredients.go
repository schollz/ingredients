package ingredients

//go:generate go run corpus/main.go
//go:generate gofmt -w corpus.go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"

	"github.com/jinzhu/inflection"
	log "github.com/schollz/logger"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Recipe contains the info for the file and the lines
type Recipe struct {
	FileName    string       `json:"filename"`
	FileContent string       `json:"file_content"`
	Lines       []LineInfo   `json:"lines"`
	Ingredients []Ingredient `json:"ingredients"`
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
	Name    string  `json:"name,omitempty"`
	Comment string  `json:"comment,omitempty"`
	Measure Measure `json:"measure,omitempty"`
	Line    string  `json:"line,omitempty"`
}

// Measure includes the amount, name and the cups for conversions
type Measure struct {
	Amount float64 `json:"amount"`
	Name   string  `json:"name"`
	Cups   float64 `json:"cups"`
	Weight float64 `json:"weight,omitempty"`
}

// IngredientList is a list of ingredients
type IngredientList struct {
	Ingredients []Ingredient `json:"ingredients"`
}

func (il IngredientList) String() string {
	s := ""
	for _, ing := range il.Ingredients {
		name := ing.Name
		if ing.Measure.Amount > 1 && ing.Measure.Name == "whole" {
			name = inflection.Plural(name)
		}
		s += fmt.Sprintf("%s %s %s", AmountToString(ing.Measure.Amount), ing.Measure.Name, name)
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

// ParseTextIngredients parses a list of ingredients and
// returns an ingredient list back
func ParseTextIngredients(text string) (ingredientList IngredientList, err error) {
	r := &Recipe{FileName: "lines"}
	r.FileContent = text
	lines := strings.Split(text, "\n")
	i := 0
	goodLines := make([]string, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		goodLines[i] = line
		i++
	}
	_, r.Lines = scoreLines(goodLines)
	err = r.parseRecipe()
	if err != nil {
		return
	}

	ingredientList = r.IngredientList()
	return
}

// NewFromFile generates a new parser from a HTML file
func NewFromFile(fname string) (r *Recipe, err error) {
	r = &Recipe{FileName: fname}
	b, err := ioutil.ReadFile(fname)
	r.FileContent = string(b)
	err = r.parseHTML()
	return
}

// NewFromString generates a new parser from a HTML string
func NewFromString(htmlString string) (r *Recipe, err error) {
	r = &Recipe{FileName: "string"}
	r.FileContent = htmlString
	err = r.parseHTML()
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

	return NewFromHTML(url, string(html))
}

// NewFromHTML generates a new parser from a HTML text
func NewFromHTML(name, htmlstring string) (r *Recipe, err error) {
	r = &Recipe{FileName: name}
	r.FileContent = htmlstring
	err = r.parseHTML()
	return
}

func IngredientsFromURL(url string) (ingredients []Ingredient, err error) {
	r, err := NewFromURL(url)
	if err != nil {
		return
	}
	ingredients = r.Ingredients
	return
}

// Parse is the main parser for a given recipe.
func (r *Recipe) parseHTML() (rerr error) {
	if r == nil {
		r = &Recipe{}
	}
	if r.FileContent == "" || r.FileName == "" {
		rerr = fmt.Errorf("no file loaded")
		return
	}

	r.Lines, rerr = getIngredientLinesInHTML(r.FileContent)
	return r.parseRecipe()

}

func (r *Recipe) parseRecipe() (rerr error) {
	goodLines := make([]LineInfo, len(r.Lines))
	j := 0
	for _, lineInfo := range r.Lines {
		if len(strings.TrimSpace(lineInfo.Line)) < 3 || len(strings.TrimSpace(lineInfo.Line)) > 150 {
			continue
		}

		// singularlize
		lineInfo.Ingredient.Measure = Measure{}

		// get amount, continue if there is an error
		err := lineInfo.getTotalAmount()
		if err != nil {
			log.Tracef("[%s]: %s (%+v)", lineInfo.Line, err.Error(), lineInfo.AmountInString)
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

func getIngredientLinesInHTML(htmlS string) (lineInfos []LineInfo, err error) {
	doc, err := html.Parse(bytes.NewReader([]byte(htmlS)))
	if err != nil {
		return
	}
	var f func(n *html.Node, lineInfos *[]LineInfo) (s string, done bool)
	f = func(n *html.Node, lineInfos *[]LineInfo) (s string, done bool) {
		childrenLineInfo := []LineInfo{}
		// log.Tracef("%+v", n)
		score := 0
		isScript := n.DataAtom == atom.Script
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if isScript {
				// try to capture JSON and if successful, do a hard exit
				lis, errJSON := extractLinesFromJavascript(c.Data)
				if errJSON == nil && len(lis) > 2 {
					log.Tracef("got ingredients from JSON")
					*lineInfos = lis
					done = true
					return
				}
			}
			var childText string
			childText, done = f(c, lineInfos)
			if done {
				return
			}
			if childText != "" {
				scoreOfLine, lineInfo := scoreLine(childText)
				childrenLineInfo = append(childrenLineInfo, lineInfo)
				score += scoreOfLine
			}
		}
		if score > 2 && len(childrenLineInfo) < 25 && len(childrenLineInfo) > 2 {
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
		return
	}
	f(doc, &lineInfos)
	return
}

func extractLinesFromJavascript(jsString string) (lineInfo []LineInfo, err error) {

	var arrayMap = []map[string]interface{}{}
	var regMap = make(map[string]interface{})
	err = json.Unmarshal([]byte(jsString), &regMap)
	if err != nil {
		err = json.Unmarshal([]byte(jsString), &arrayMap)
		if err != nil {
			return
		}
		if len(arrayMap) == 0 {
			err = fmt.Errorf("nothing to parse")
			return
		}
		parseMap(arrayMap[0], &lineInfo)
		err = nil
	} else {
		parseMap(regMap, &lineInfo)
		err = nil
	}

	return
}

func parseMap(aMap map[string]interface{}, lineInfo *[]LineInfo) {
	for _, val := range aMap {
		switch val.(type) {
		case map[string]interface{}:
			parseMap(val.(map[string]interface{}), lineInfo)
		case []interface{}:
			parseArray(val.([]interface{}), lineInfo)
		default:
			// fmt.Println(key, ":", concreteVal)
		}
	}
}

func parseArray(anArray []interface{}, lineInfo *[]LineInfo) {
	concreteLines := []string{}
	for _, val := range anArray {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			parseMap(val.(map[string]interface{}), lineInfo)
		case []interface{}:
			parseArray(val.([]interface{}), lineInfo)
		default:
			switch v := concreteVal.(type) {
			case string:
				concreteLines = append(concreteLines, v)
			}
		}
	}

	score, li := scoreLines(concreteLines)
	log.Trace(score, li)
	if score > 20 {
		*lineInfo = li
	}

	return
}

func scoreLines(lines []string) (score int, lineInfo []LineInfo) {
	if len(lines) < 2 {
		return
	}
	lineInfo = make([]LineInfo, len(lines))
	for i, line := range lines {
		var scored int
		scored, lineInfo[i] = scoreLine(line)
		score += scored
	}
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
	if len(lineInfo.Line) > 30 {
		score = score - (len(lineInfo.Line) - 30)
	}
	if len(lineInfo.Line) > 250 {
		score = 0
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

// IngredientList will return a string containing the ingredient list
func (r *Recipe) IngredientList() (ingredientList IngredientList) {
	ingredientList = IngredientList{make([]Ingredient, len(r.Lines))}
	for i, li := range r.Lines {
		ingredientList.Ingredients[i] = li.Ingredient
		ingredientList.Ingredients[i].Line = li.LineOriginal
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
	lineInfo.Ingredient.Name = inflection.Singular(lineInfo.IngredientsInString[0].Word)
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
