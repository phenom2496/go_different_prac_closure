package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(records)
	rec, err := json.Marshal(records)
	if err != nil {
		fmt.Println("err", err)
	}
	os.Stdout.Write(rec)
}
