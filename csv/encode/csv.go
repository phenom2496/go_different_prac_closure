package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("test.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	new1 := [][]string{
		{`5`, `vikram`, `30`, `vjhak@gmail.com`},
		{`6`, `mukul`, `32`, `mgar@gmail.com`},
	}
	w := csv.NewWriter(f)
	w.WriteAll(new1)
	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}

}
