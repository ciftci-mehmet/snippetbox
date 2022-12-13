package main

import (
	"path/filepath"
	"text/template"

	"github.com/ciftci-mehmet/snippetbox/pkg/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	// init map to act as a cache
	cache := map[string]*template.Template{}

	// get all files ending with '.page.tmpl'
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	// loop pages
	for _, page := range pages {
		// extract full file name
		name := filepath.Base(page)

		// parse page in to a template set
		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// add layouts to the template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		// add the template set to the cache using name of the page as the key like 'home.page.tmpl'
		cache[name] = ts
	}

	return cache, nil
}
