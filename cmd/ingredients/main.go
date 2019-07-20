package main

import (
	"encoding/json"
	"fmt"
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

	var r *ingredients.Recipe
	var err error

	r, err = ingredients.NewFromFile(os.Args[1])
	if err != nil {
		r, err = ingredients.NewFromURL(os.Args[1])
		if err != nil {
			log.Error("usage: ingredients [file/url]")
			os.Exit(1)
		}
	}
	ing := r.IngredientList()
	b, _ := json.MarshalIndent(ing, "", "    ")
	fmt.Println(string(b))
}
