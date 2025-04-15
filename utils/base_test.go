package utils

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRandString(t *testing.T) {
	// Test with length of 10
	result := RandString(10)
	assert.Len(t, result, 10, "Expected string length to be 10")

	// Test to check that the string contains only valid characters
	for _, r := range result {
		assert.Contains(t, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", string(r), "Invalid character found in the string")
	}
}

func TestHashMd5(t *testing.T) {
	// Test with different inputs
	result1 := HashMd5("test")
	result2 := HashMd5("test")
	result3 := HashMd5("different")

	// Same input should produce the same result
	assert.Equal(t, result1, result2, "MD5 hash should be the same for identical inputs")

	// Different input should produce different result
	assert.NotEqual(t, result1, result3, "MD5 hash should differ for different inputs")
}

func TestToBson(t *testing.T) {
	// Test with a struct
	doc := struct {
		Name  string `bson:"name"`
		Value int    `bson:"value"`
	}{
		Name:  "test",
		Value: 123,
	}
	result := ToBson(doc)

	// Test that the conversion results in expected BSON
	assert.Equal(t, "test", result["name"], "Expected BSON 'name' to be 'test'")
	assert.EqualValues(t, 123, result["value"], "Expected BSON 'value' to be 123")
}

func TestTimeString(t *testing.T) {
	// Test with a valid time
	validTime := time.Date(2024, 11, 28, 0, 0, 0, 0, time.UTC)
	result := TimeString(validTime)
	assert.NotEmpty(t, result, "Expected time string to be non-empty")
	parsedTime, err := time.Parse(time.RFC3339Nano, result)
	assert.NoError(t, err, "Expected time string to be in RFC3339Nano format")
	assert.Equal(t, validTime.UTC(), parsedTime.UTC(), "Expected time string to match the input time")

	// Test with zero time (empty string)
	var zeroTime time.Time
	resultZero := TimeString(zeroTime)
	assert.Empty(t, resultZero, "Expected empty string for zero time")
}

func TestNewUuidV7String(t *testing.T) {
	// Test UUID v7 generation
	result := NewUuidV7String()
	_, err := uuid.Parse(result)
	assert.NoError(t, err, "Expected valid UUID to be generated")
}
