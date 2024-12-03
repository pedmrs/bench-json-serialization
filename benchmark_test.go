package main

import (
	"encoding/json"
	"testing"

	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
)

type SampleStruct struct {
	FirstString   string
	SecondString  string
	FirstInteger  int
	SecondInteger int
}

var sampleStruct = SampleStruct{
	FirstString:   "12345",
	SecondString:  "67890",
	FirstInteger:  12345,
	SecondInteger: 67890,
}

func BenchmarkBasicData(b *testing.B) {
	b.Run("DefaultLib", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data, err := json.Marshal(sampleStruct)
			if err != nil {
				b.Fatal(err)
			}
			var ss SampleStruct
			err = json.Unmarshal(data, &ss)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("JSONIterator", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data, err := jsoniter.Marshal(sampleStruct)
			if err != nil {
				b.Fatal(err)
			}
			var ss SampleStruct
			err = jsoniter.Unmarshal(data, &ss)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("GoJSON", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data, err := gojson.Marshal(sampleStruct)
			if err != nil {
				b.Fatal(err)
			}
			var ss SampleStruct
			err = gojson.Unmarshal(data, &ss)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkLargeData(b *testing.B) {
	largeData := make([]SampleStruct, 1000)
	for i := range largeData {
		largeData[i] = sampleStruct
	}

	b.Run("DefaultLibLarge", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data, err := json.Marshal(largeData)
			if err != nil {
				b.Fatal(err)
			}
			var ss []SampleStruct
			err = json.Unmarshal(data, &ss)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("JSONIteratorLarge", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data, err := jsoniter.Marshal(largeData)
			if err != nil {
				b.Fatal(err)
			}
			var ss []SampleStruct
			err = jsoniter.Unmarshal(data, &ss)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("GoJSONLarge", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			data, err := gojson.Marshal(largeData)
			if err != nil {
				b.Fatal(err)
			}
			var ss []SampleStruct
			err = gojson.Unmarshal(data, &ss)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkParallel(b *testing.B) {
	b.Run("DefaultLibParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data, err := json.Marshal(sampleStruct)
				if err != nil {
					b.Fatal(err)
				}
				var ss SampleStruct
				err = json.Unmarshal(data, &ss)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	})

	b.Run("JSONIteratorParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data, err := jsoniter.Marshal(sampleStruct)
				if err != nil {
					b.Fatal(err)
				}
				var ss SampleStruct
				err = jsoniter.Unmarshal(data, &ss)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	})

	b.Run("GoJSONParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data, err := gojson.Marshal(sampleStruct)
				if err != nil {
					b.Fatal(err)
				}
				var ss SampleStruct
				err = gojson.Unmarshal(data, &ss)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	})
}

func BenchmarkErrorHandling(b *testing.B) {
	invalidJSON := []byte(`{"FirstString": -1, "SecondString": 0, "FirstInteger": "not_int", "SecondInteger": "also_not_int"}`)

	b.Run("DefaultLibInvalidInput", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var ss SampleStruct
			err := json.Unmarshal(invalidJSON, &ss)
			if err == nil {
				b.Fatal("expected error but got nil")
			}
		}
	})

	b.Run("JSONIteratorInvalidInput", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var ss SampleStruct
			err := jsoniter.Unmarshal(invalidJSON, &ss)
			if err == nil {
				b.Fatal("expected error but got nil")
			}
		}
	})

	b.Run("GoJSONInvalidInput", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var ss SampleStruct
			err := gojson.Unmarshal(invalidJSON, &ss)
			if err == nil {
				b.Fatal("expected error but got nil")
			}
		}
	})
}

func main() {}
