package record

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type MapRecord map[uuid.UUID]Record

func (mapRec *MapRecord) AddMapRecord(r Record) error {
	if r == (Record{}) {
		return fmt.Errorf("record is nil")
	}
	if v, ok := (*mapRec)[r.UUID]; ok {
		return fmt.Errorf("key %v with value %v is exist", r.UUID, v)
	} else {
		(*mapRec)[r.UUID] = r
		return nil
	}
}

func (mapRec *MapRecord) FindByID(Id uuid.UUID) (Record, bool) {
	if v, ok := (*mapRec)[Id]; ok {
		return v, true
	} else {
		return Record{}, false
	}
}

func (mapRec *MapRecord) FindByName(searchName string) ([]Record, bool) {
	findRecords := []Record{}
	for _, record := range *mapRec {
		if strings.Contains(strings.ToLower(record.Name), strings.ToLower(searchName)) {
			findRecords = append(findRecords, record)

		}
	}
	return findRecords, len(findRecords) > 0
}
