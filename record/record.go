package record

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type Record struct {
	UUID        uuid.UUID
	Name        string
	Description string
	Link        string
}

func (rec *Record) getId() uuid.UUID {
	return rec.UUID
}
func (rec *Record) getName() string {
	return rec.Name
}
func (rec *Record) getDescription() string {
	return rec.Description
}
func (rec *Record) getLink() string {
	return rec.Link
}

func (rec *Record) String() string {
	return fmt.Sprintf("ID: %v\tName: %s\tDescription: %s\t Link: %s", rec.getId(), rec.getName(), rec.getDescription(), rec.getLink())
}

type RecordConvert interface {
	ToMap()
}

type RecordSlice []Record

func (recSlice *RecordSlice) ToMap() MapRecord {
	mapRecords := make(MapRecord)
	for _, record := range *recSlice {
		mapRecords[record.UUID] = record
	}
	return mapRecords
}

// / Read data from Json to slice
func (reader *RecordSlice) FromJson(data []byte) error {
	if err := json.Unmarshal(data, &reader); err != nil {
		return err
	}
	return nil
}
