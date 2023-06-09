package repositories

import (
	"encoding/csv"
	"os"

	"github.com/dafailyasa/event-micro/pkg/logger/domain/models"
	"github.com/dafailyasa/event-micro/pkg/logger/domain/ports"
)

type CsvFile struct {
	path string
}

var _ ports.LoggerRepository = (*CsvFile)(nil)

func NewCsvFile(path string) *CsvFile {
	return &CsvFile{
		path: path,
	}
}

func (c *CsvFile) Save(log *models.Log) error {
	file, err := os.Create(c.path)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	// write data to csv file
	err = csvWriter.Write([]string{log.ID, log.Level, log.Message, log.CreatedAt.String()})
	if err != nil {
		return err
	}

	return nil
}
