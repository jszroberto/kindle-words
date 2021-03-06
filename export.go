package main

import (
	"bufio"
	// "github.com/gocarina/gocsv"
	"github.com/atotto/encoding/csv"
	"github.com/jszroberto/kindle-words/kindledb"
	"os"
	"strings"
)

func exportVocabularyCom(path string, words []kindledb.Word) error {
	outputs := map[string]string{}

	for _, word := range words {
		if word.IsEnglish() {
			if outputs[word.Book] != "" {
				outputs[word.Book] += ","
			}
			outputs[word.Book] += strings.ToLower(word.Value)
		}

	}

	if err := os.MkdirAll(path+"/", 0755); err != nil {
		panic("Unable to create directory for tagfile! - " + err.Error())
	}

	for key, value := range outputs {
		file, err := os.Create(path + "/" + key)
		if err != nil {
			return err
		}
		defer file.Close()

		writer := bufio.NewWriter(file)
		_, err = writer.WriteString(value)
		if err != nil {
			return err
		}
		writer.Flush()
	}

	return nil
}

func exportHtml(path string, words []kindledb.Word) error {
	return exportToFolders(path, "html", "templates/book.html", words, true)
}

func exportEvernote(path string, words []kindledb.Word) error {
	return exportToFolders(path, "html", "templates/evernote.html", words, false)
}

func exportToCSV(path string, words []kindledb.Word) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	w.WriteStructAll(words) // calls Flush internally

	if err := w.Error(); err != nil {
		return err
	}

	// err = gocsv.MarshalFile(&words, file)
	//
	// if err != nil {
	// 	return err
	// }

	return nil
}
