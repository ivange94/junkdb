package storage

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/ivange94/junkdb/internal/config"
)

type Engine struct {
	mu   sync.RWMutex
	path string
}

var ErrKeyNotFound = errors.New("key not found")

const recordSeparator = ","

func NewEngine() (*Engine, error) {
	dataDir, err := config.Dir()
	if err != nil {
		return nil, fmt.Errorf("resolve data directory: %w", err)
	}
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		return nil, fmt.Errorf("create data directory: %w", err)
	}
	dataFilePath, err := config.DataFilePath()
	if err != nil {
		return nil, fmt.Errorf("resolve data file path: %w", err)
	}

	return &Engine{path: dataFilePath}, nil
}

func (e *Engine) Put(key, value string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	file, err := os.OpenFile(e.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // TODO: handle concurrent access
	if err != nil {
		return fmt.Errorf("open data file: %w", err)
	}
	defer func() {
		_ = file.Close()
	}()

	entry := formatRecord(key, value)
	if _, err := file.WriteString(entry); err != nil {
		return fmt.Errorf("write entry: %w", err)
	}
	return nil
}

func (e *Engine) Get(key string) (string, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	content, err := os.ReadFile(e.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("%w: %s", ErrKeyNotFound, key)
		}
		return "", fmt.Errorf("read data file: %w", err)
	}

	var value string
	var found bool

	scanner := bufio.NewScanner(bytes.NewReader(content))
	for scanner.Scan() {
		recordKey, recordValue, err := parseRecord(scanner.Text())
		if err != nil {
			return "", err
		}
		if recordKey == key {
			value = recordValue
			found = true
		}
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("scan data file: %w", err)
	}
	if found {
		return value, nil
	}
	return "", fmt.Errorf("%w: %s", ErrKeyNotFound, key)
}

func formatRecord(key, value string) string {
	return fmt.Sprintf("%s%s%s\n", key, recordSeparator, value)
}

func parseRecord(line string) (string, string, error) {
	if strings.TrimSpace(line) == "" {
		return "", "", nil
	}

	parts := strings.SplitN(line, recordSeparator, 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("corrupted data entry: %q", line)
	}
	return parts[0], parts[1], nil
}
