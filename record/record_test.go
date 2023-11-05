package record

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddMapRecord(t *testing.T) {
	m := make(MapRecord)

	strExitUUID := "ef2b2dd9-52e7-4f63-a803-9b5c8ca56d9d"
	existId, _ := uuid.Parse(strExitUUID)

	goodRecord := Record{UUID: uuid.New(), Name: "Example 1", Description: "An example record", Link: "https://example.com"}
	preExistRecord := Record{UUID: existId, Name: "Example 3", Description: "An example record", Link: "https://example.com"}
	existRecord := Record{UUID: existId, Name: "UUID Exist", Description: "An example record", Link: "https://example.com"}
	emptyRecord := Record{}

	assert.Nil(t, m.AddMapRecord(goodRecord))
	assert.Nil(t, m.AddMapRecord(preExistRecord))
	assert.Equal(t, fmt.Errorf("key %v with value %v is exist", existId, m[existId]), m.AddMapRecord(existRecord))
	assert.Equal(t, fmt.Errorf("record is nil"), m.AddMapRecord(emptyRecord))
	assert.Equal(t, 2, len(m))
}

func TestMapRecord_findByID(t *testing.T) {
	// Create a new MapRecord with some test data
	id := uuid.New()
	mapRec := MapRecord{
		id: Record{
			UUID:        id,
			Name:        "Test Record",
			Description: "An example record",
			Link:        "https://example.com",
		},
	}

	// Test that findByID returns the correct record
	expected := mapRec[id]
	actual, ok := mapRec.FindByID(id)
	assert.True(t, ok)
	assert.Equal(t, expected, actual)

	// Test that findByID returns an error for an invalid ID
	invalidID := uuid.New()
	_, ok = mapRec.FindByID(invalidID)
	assert.False(t, ok)
}

func TestMapRecord_findByName(t *testing.T) {
	// Create a MapRecord with some test data
	testData := MapRecord{
		uuid.New(): Record{
			UUID:        uuid.New(),
			Name:        "Test Record 1",
			Description: "This is the first test record",
			Link:        "https://example.com/test1",
		},
		uuid.New(): Record{
			UUID:        uuid.New(),
			Name:        "Test Record 2",
			Description: "This is the second test record",
			Link:        "https://example.com/test2",
		},
		uuid.New(): Record{
			UUID:        uuid.New(),
			Name:        "Another Record",
			Description: "This is another test record",
			Link:        "https://example.com/another",
		},
	}

	// Test searching for a record that exists
	findRecords, found := testData.FindByName("test")
	assert.True(t, found, "Expected to find records")
	assert.Len(t, findRecords, 2, "Expected to find 2 records")
	for _, record := range findRecords {
		assert.True(t, strings.Contains(strings.ToLower(record.Name), "test"), "Expected record name to contain 'test'")
	}

	// Test searching for a record that doesn't exist
	findRecords, found = testData.FindByName("foobar")
	assert.False(t, found, "Expected not to find records")
	assert.Len(t, findRecords, 0, "Expected to find 0 records")
}
