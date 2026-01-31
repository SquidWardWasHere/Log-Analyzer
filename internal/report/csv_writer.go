package report

import (
	"encoding/csv"
	"log-analyzer/internal/analyzer"
	"os"
)

func ExportToCSV(filename string, data []analyzer.LogResult) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Zaman Damgasi", "Seviye", "Kaynak", "Mesaj"})

	for _, row := range data {
		writer.Write([]string{row.Timestamp, row.Level, row.Source, row.Message})
	}

	return nil
}
