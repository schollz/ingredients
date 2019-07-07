package extractrecipe

import (
	"fmt"
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

func TestParseREcipes(t *testing.T) {
	files := []string{
		"testing/sites/1017060-doughnuts",
	}
	for _, f := range files {
		log.Infof("working on %s", f)
		r, err := NewFromFile(f)
		assert.Nil(t, err)
		err = r.Parse()
		assert.Nil(t, err)
		fmt.Println(r.IngredientList())
	}
}
