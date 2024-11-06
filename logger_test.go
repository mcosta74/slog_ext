// Package slogext exposes some utilities to create slog.Logger.
package slogext

import (
	"bytes"
	"cmp"
	"encoding/json"
	"log/slog"
	"maps"
	"reflect"
	"strings"
	"testing"
	"time"
)

func compareSliceCount[T cmp.Ordered](a, b []T) bool {

	countFunc := func(a []T) map[T]int {
		result := make(map[T]int)
		for _, s := range a {
			result[s]++
		}
		return result
	}

	counterA := countFunc(a)
	counterB := countFunc(b)

	return reflect.DeepEqual(counterA, counterB)
}

func TestNew(t *testing.T) {
	keyValuesFromText := func(text string) map[string]string {
		keyValues := make(map[string]string)
		pairs := strings.Split(text, " ")
		for _, pair := range pairs {
			keyValue := strings.Split(pair, "=")
			keyValues[keyValue[0]] = keyValue[1]
		}
		return keyValues
	}

	checkKeys := func(t *testing.T, keyValues map[string]string, want []string) {
		t.Helper()

		got := make([]string, 0)
		for key := range maps.Keys(keyValues) {
			got = append(got, key)
		}

		if !compareSliceCount(got, want) {
			t.Errorf("Keys are not the same, got = %v, want = %v", got, want)
		}
	}

	t.Run("default", func(t *testing.T) {
		writer := &bytes.Buffer{}
		logger := New(writer)
		logger.Info("hello")

		keyValues := keyValuesFromText(writer.String())
		checkKeys(t, keyValues, []string{"time", "level", "msg"})

		timeStr := keyValues["time"]
		logTime, err := time.Parse(time.RFC3339, timeStr)
		if err != nil {
			t.Fatalf("invalid timeStr = %q", timeStr)
		}

		if location := logTime.Location().String(); location == "UTC" {
			t.Errorf("Expected Local location, got = %q", location)
		}

	})

	t.Run("with source", func(t *testing.T) {
		writer := &bytes.Buffer{}
		logger := New(writer, WithSource(true))
		logger.Info("hello")

		keyValues := keyValuesFromText(writer.String())
		checkKeys(t, keyValues, []string{"time", "level", "msg", "source"})
	})

	t.Run("with level", func(t *testing.T) {
		writer := &bytes.Buffer{}
		logger := New(writer, WithLevel(slog.LevelError))
		logger.Info("hello")

		if output := writer.String(); output != "" {
			t.Errorf("Expected no output, got = %q", output)
		}
	})

	t.Run("with UTC", func(t *testing.T) {

		writer := &bytes.Buffer{}
		logger := New(writer, WithUseUTC(true))
		logger.Info("hello")

		keyValues := keyValuesFromText(writer.String())
		timeStr := keyValues["time"]
		logTime, err := time.Parse(time.RFC3339, timeStr)
		if err != nil {
			t.Fatalf("invalid timeStr = %q", timeStr)
		}

		if location := logTime.Location().String(); location != "UTC" {
			t.Errorf("Expected UTC location, got = %q", location)
		}
	})

	t.Run("with JSON", func(t *testing.T) {
		writer := &bytes.Buffer{}
		logger := New(writer, WithJSON(true))
		logger.Info("hello")

		var jsonData map[string]string
		if err := json.Unmarshal(writer.Bytes(), &jsonData); err != nil {
			t.Errorf("Expected JSON output, err = %v", err)
		}
		checkKeys(t, jsonData, []string{"time", "level", "msg"})
	})

}
