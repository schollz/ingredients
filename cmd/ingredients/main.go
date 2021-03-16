package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/schollz/ingredients"
	log "github.com/schollz/logger"
)

func main() {
	log.SetLevel("error")
	if len(os.Args) < 2 {
		log.Error("usage: ingredients [file/url]")
		os.Exit(1)
	}

	h := md5.New()
	io.WriteString(h, os.Args[1])
	fname := fmt.Sprintf("%x.json", h.Sum(nil))
	if _, err := os.Stat(fname); err == nil {
		return
	}
	var r *ingredients.Recipe
	var err error
	type Result struct {
		Ingredients []ingredients.Ingredient `json:"ingredients"`
		Origin      string                   `json:"origin"`
	}
	r, err = ingredients.NewFromFile(os.Args[1])
	if err != nil {
		r, err = ingredients.NewFromURL(os.Args[1])
		if err != nil {
			log.Error("usage: ingredients [file/url]")
			os.Exit(1)
		}
	}
	ing := r.IngredientList()
	var re Result
	re.Ingredients = ing.Ingredients
	re.Origin = os.Args[1]
	if len(ing.Ingredients) > 0 {
		b, _ := json.MarshalIndent(re, "", "    ")
		ioutil.WriteFile(fname, b, 0644)
		fmt.Printf("wrote '%s'\n", fname)
	} else {
		fmt.Printf("no ingredients for '%s'\n", os.Args[1])
	}
}
