package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("failed to open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("failed to open file")
	}

	file.Close()
	return lines, nil
}

func WriteJSON(path string, data any) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("failed to create file")
	}

	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		file.Close()
		return errors.New("failed to encode data")
	}

	file.Close()
	return nil
}
