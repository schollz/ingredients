package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/schollz/ingredients"
	log "github.com/schollz/logger"
	"github.com/schollz/progressbar/v3"
)

func main() {
	log.SetLevel("error")
	fnames := []string{}
	err := filepath.Walk(os.Args[1],
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				fnames = append(fnames, path)
			}
			return nil
		})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var numJobs = len(fnames)

	type job struct {
		pathToFile string
	}

	type result struct {
		err error
	}

	jobs := make(chan job, numJobs)
	results := make(chan result, numJobs)
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		go func(jobs <-chan job, results chan<- result) {
			for j := range jobs {
				// step 3: specify the work for the worker
				var r result
				processFile(os.Args[1], j.pathToFile)
				results <- r
			}
		}(jobs, results)
	}

	// step 4: send out jobs
	for i := 0; i < numJobs; i++ {
		jobs <- job{fnames[i]}
	}
	close(jobs)

	// step 5: do something with results
	b := progressbar.Default(int64(numJobs))
	for i := 0; i < numJobs; i++ {
		b.Add(1)
		r := <-results
		if r.err != nil {
			// do something with error
		}
	}
}

func processFile(folderName, pathToFile string) {
	h := md5.New()
	io.WriteString(h, pathToFile)
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
	r, err = ingredients.NewFromFile(pathToFile)
	if err != nil {
		log.Errorf("%s: %s", pathToFile, err)
		return
	}
	ing := r.IngredientList()
	var re Result
	re.Ingredients = ing.Ingredients
	re.Origin = strings.TrimPrefix(pathToFile, folderName)
	re.Origin = strings.TrimPrefix(re.Origin, "/")
	if len(ing.Ingredients) > 0 {
		b, _ := json.MarshalIndent(re, "", "    ")
		ioutil.WriteFile(path.Join(folderName, fname), b, 0644)
		log.Debugf("wrote '%s'\n", pathToFile)
	} else {
		log.Debugf("no ingredients for '%s'\n", pathToFile)
	}
}
