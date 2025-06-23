package problems

import (
	"embed"
	"encoding/json"
	"io/fs"
)

//go:embed lists/*
var files embed.FS

func GetAllProblems() ([]Problem, error) {
	var problems []Problem
	entries, err := fs.ReadDir(files, "lists")
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			data, err := fs.ReadFile(files, "lists/"+entry.Name())
			if err != nil {
				return nil, err
			}
			var p Problem
			if err := json.Unmarshal(data, &p); err != nil {
				return nil, err
			}
			problems = append(problems, p)
		}
	}
	return problems, nil
}

func GetProblemByID(id string) (*Problem, error) {
	problems, err := GetAllProblems()
	if err != nil {
		return nil, err
	}
	for _, p := range problems {
		if p.ID == id {
			return &p, nil
		}
	}
	return nil, nil
}
