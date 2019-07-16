package main

import (
	"fmt"
	"os"

	log "github.com/schollz/logger"
	"github.com/schollz/recipeingredients"
)

func main() {
	log.SetLevel("debug")
	if len(os.Args) < 2 {
		log.Error("need to have argument")
		os.Exit(1)
	}

	r, err := recipeingredients.NewFromFile(os.Args[1])
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	err = r.Parse()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	fmt.Println(r.PrintIngredientList())
}
