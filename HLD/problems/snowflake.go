package main 

/* 
64 bit snowflake ID generator

1 - Sign bit
5 - Datacentre bits 
5 - worker/instances
41- timestamp
12 - sequence bits


*/

import (
	"sync"
	"time"
	"errors"
	"strings"
	"fmt"
	"log"
)


const (
	workerIDBits uint8 = 5 
	datacenterIDBits uint8 = 5
	sequenceBits uint8 = 12
	
	maxWorkerID     int64 = -1 ^ (-1 << workerIDBits)
	maxDatacenterID int64 = -1 ^ (-1 << datacenterIDBits)
	sequenceMask    int64 = -1 ^ (-1 << sequenceBits)

	workerIDShift     uint8 = sequenceBits
	datacenterIDShift uint8 = sequenceBits + workerIDBits
	timestampShift    uint8 = sequenceBits + workerIDBits + datacenterIDBits

	// Custom Epoch (Jan 1, 2025, in milliseconds)
	// You can get this with:
	// time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano() / 1e6
	customEpoch int64 = 1735689600000
)

// SnowflakeIDGenerator holds the state for generating IDs
type SnowflakeIDGenerator struct {
	mu            sync.Mutex
	lastTimestamp int64
	workerID      int64
	datacenterID  int64
	sequence      int64
}

// NewSnowflakeIDGenerator creates a new ID generator
func NewSnowflakeIDGenerator(datacenterID, workerID int64) (*SnowflakeIDGenerator, error) {
	if workerID < 0 || workerID > maxWorkerID {
		return nil, fmt.Errorf("worker ID must be between 0 and %d", maxWorkerID)
	}
	if datacenterID < 0 || datacenterID > maxDatacenterID {
		return nil, fmt.Errorf("datacenter ID must be between 0 and %d", maxDatacenterID)
	}

	return &SnowflakeIDGenerator{
		workerID:      workerID,
		datacenterID:  datacenterID,
		lastTimestamp: -1,
		sequence:      0,
	}, nil
}

// _currentMilliseconds returns the current time in milliseconds
func (s *SnowflakeIDGenerator) _currentMilliseconds() int64 {
	return time.Now().UnixNano() / 1e6 // Convert nanoseconds to milliseconds
}

// _waitForNextMS waits until the next millisecond
func (s *SnowflakeIDGenerator) _waitForNextMS(lastTimestamp int64) int64 {
	timestamp := s._currentMilliseconds()
	for timestamp <= lastTimestamp {
		timestamp = s._currentMilliseconds()
	}
	return timestamp
}

func (s *SnowflakeIDGenerator) GenerateID() (int64,error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	timestamp := s._currentMilliseconds()

	// --- Clock Drift Handling ---
	if timestamp < s.lastTimestamp {
		return 0, errors.New("clock moved backwards, refusing to generate ID")
	}

	// --- Same Millisecond Handling ---
	if timestamp == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & sequenceMask
		// Sequence overflow
		if s.sequence == 0 {
			timestamp = s._waitForNextMS(s.lastTimestamp)
		}
	} else {
		// New millisecond, reset sequence
		s.sequence = 0
	}

	relativeTimestamp := timestamp - customEpoch

	newID := (relativeTimestamp << timestampShift) |
		(s.datacenterID << datacenterIDShift) |
		(s.workerID << workerIDShift) |
		s.sequence

	return newID, nil

}

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Base62Encode encodes a positive integer into a Base62 string
func Base62Encode(number int64) string {
	if number == 0 {
		return "0"
	}

	var sb strings.Builder
	
	for number > 0 {
		remainder := number % 62
		sb.WriteByte(charset[remainder])
		number = number / 62
	}

	// The string is built in reverse order, so we must reverse it
	return reverseString(sb.String())
}

// Helper function to reverse a string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	// --- Example Usage ---

	// Initialize the generator (e.g., in Datacenter 1, Worker 5)
	generator, err := NewSnowflakeIDGenerator(1, 5)
	if err != nil {
		log.Fatalf("Failed to create Snowflake generator: %v", err)
	}

	fmt.Println("Generating 3 Snowflake IDs...")
	
	// Generate some IDs
	id1, err := generator.GenerateID()
	if err != nil {
		log.Fatalf("Error generating ID: %v", err)
	}

	id2, err := generator.GenerateID()
	if err != nil {
		log.Fatalf("Error generating ID: %v", err)
	}
	
	id3, err := generator.GenerateID()
	if err != nil {
		log.Fatalf("Error generating ID: %v", err)
	}

	fmt.Printf("ID 1 (raw): %d\n", id1)
	fmt.Printf("ID 2 (raw): %d\n", id2)
	fmt.Printf("ID 3 (raw): %d\n", id3)

	fmt.Println("\n--- Encoding for TinyURL ---")

	// Encode them to get the short keys
	shortKey1 := Base62Encode(id1)
	shortKey2 := Base62Encode(id2)

	fmt.Printf("Snowflake ID %d becomes short key: %s\n", id1, shortKey1)
	fmt.Printf("Snowflake ID %d becomes short key: %s\n", id2, shortKey2)
}