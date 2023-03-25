// Modified GPT-4 version

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const storeFilename = "store.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: dv <get|set|del> [args]")
		return
	}

	command := os.Args[1]

	switch command {
	case "set":
		if len(os.Args) != 4 {
			fmt.Println("Usage: dv set <key> <value>")
			return
		}

		key := os.Args[2]
		value := os.Args[3]

		err := set(key, value)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("ok")
		}
	case "get":
		if len(os.Args) != 3 {
			fmt.Println("Usage: dv get <key>")
			return
		}

		key := os.Args[2]

		value, err := get(key)

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("%s\n", value)
		}
	case "del":
		if len(os.Args) != 3 {
			fmt.Println("Usage: dv del <key>")
			return
		}

		key := os.Args[2]

		err := del(key)

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("ok")
		}
	default:
		fmt.Println("Unknown command:", command)
	}
}

func loadStore() (map[string]string, error) {
	data := make(map[string]string)

	file, err := os.Open(storeFilename)

	if errors.Is(err, os.ErrNotExist) {
		return data, nil
	} else if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func saveStore(data map[string]string) error {
	file, err := os.Create(storeFilename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return err
	}

	return nil
}

func set(key, value string) error {
	data, err := loadStore()

	if err != nil {
		return err
	}

	data[key] = value

	err = saveStore(data)
	if err != nil {
		return err
	}

	return nil
}

func get(key string) (string, error) {
	data, err := loadStore()
	if err != nil {
		return "", err
	}

	value, exists := data[key]
	if !exists {
		return "", fmt.Errorf("key not found: %s", key)
	}

	return value, nil
}

func del(key string) error {
	data, err := loadStore()

	if err != nil {
		return err
	}

	_, exists := data[key]

	if !exists {
		return fmt.Errorf("key not found: %s", key)
	}

	delete(data, key)

	err = saveStore(data)
	if err != nil {
		return err
	}

	return nil
}
