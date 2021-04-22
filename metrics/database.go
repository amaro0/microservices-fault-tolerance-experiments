package metrics

import (
	"encoding/csv"
	"log"
	"os"
)

var w *csv.Writer

func initDb() chan Model {
	file, err := os.Create("result.csv")
	if err != nil {
		log.Panicln("Cannot open result file" + err.Error())
	}

	// File can not be closed via defer
	//defer file.Close()

	w = csv.NewWriter(file)
	if err := w.Write(getCSVHeader()); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}
	w.Flush()
	metricsChan := make(chan Model, 1000)

	go handleNewMetric(w, metricsChan)

	return metricsChan
}

func handleNewMetric(w *csv.Writer, c chan Model) {
	for {
		select {
		case metric := <-c:
			csvRow := metric.prepareForCSV()

			if err := w.Write(csvRow); err != nil {
				log.Fatalln("error writing record to csv:", err)
			}

			w.Flush()
		}
	}
}
