package main

import (
	"time"
	"encoding/json"
	"os"
	"log"
	"strconv"
	"encoding/csv"
	"reflect"
)

var (
	TenantID int
	metric []float64
)

type Data struct {
	A float64
	B float64
	C float64
	D float64
	E float64
	F float64
	G float64
	H float64
	I float64
	J float64
}

func main() {
	a := time.Now().Format("15:04")

	TenantID = 1234
	var d Data

	bs := []byte(`{"A":53.2,"B":45.2,"C":58.2,"D":78,"E":78.25,"F":45893.25,"G":784.2564,"H":748.23651,"I":9865.25696,"J":56.2546}`)
	json.Unmarshal(bs, &d)

	// Create the csv file
	file, err := os.Create("test.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	v := reflect.ValueOf(d)
	csvWriter := csv.NewWriter(file)

	for i := 0; i < 10; i++ {
		value := v.Field(i).Interface()
		flt, _ := value.(float64)
		metric := strconv.FormatFloat(flt, 'f',2,64)
		data := []string{a, strconv.Itoa(TenantID), string(i + 65), metric}
		strWrite := [][]string{data}
		csvWriter.WriteAll(strWrite)
	}
	csvWriter.Flush()
}