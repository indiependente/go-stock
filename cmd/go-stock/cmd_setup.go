package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/indiependente/go-stock/config"
)

func setup(apis []string) (string, error) {
	var (
		apiIdx   int
		filename string
		key      string
		err      error
	)
	maxTentatives := 3
	fmt.Printf("Which API do you want the tool to use?\n")
	for i, api := range apis {
		fmt.Printf("%d) %s\n", i+1, api)
	}
	reader := bufio.NewReader(os.Stdin)

	tentatives := 0
	for tentatives < maxTentatives {
		fmt.Print("Enter your choice: ")
		choice, err := reader.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf("could not read choice: %w", err)
		}
		choice = strings.TrimSuffix(choice, "\n")
		apiIdx, err = strconv.Atoi(choice)
		if err != nil {
			fmt.Printf("The choice must be a number.\n")
			tentatives++
			continue
		}
		if apiIdx < 1 || apiIdx > len(apis) {
			fmt.Printf("Please enter a valid choice.\n")
			tentatives++
			continue
		}
		break
	}
	if tentatives == maxTentatives {
		return "", errors.New("reached maximum number of tentatives")
	}
	apiIdx -= 1
	tentatives = 0
	for tentatives < maxTentatives {
		fmt.Printf("Enter your API key for %s: ", apis[apiIdx])
		key, _ = reader.ReadString('\n')
		if key == "" {
			fmt.Printf("Please enter a valid choice.\n")
			tentatives++
			continue
		}
		key = strings.TrimSuffix(key, "\n")
		break
	}
	if tentatives == maxTentatives {
		return "", errors.New("reached maximum number of tentatives")
	}
	tentatives = 0
	for tentatives < maxTentatives {
		fmt.Printf("Enter new config file name: ")
		filename, _ = reader.ReadString('\n')
		if filename == "" {
			fmt.Printf("Please enter a valid file name.\n")
			tentatives++
			continue
		}
		filename = strings.TrimSuffix(filename, "\n")
		break
	}
	if tentatives == maxTentatives {
		return "", errors.New("reached maximum number of tentatives")
	}
	c := config.Config{
		URL:    apis[apiIdx],
		APIKey: key,
	}

	err = config.Save(c, filename)
	if err != nil {
		return "", fmt.Errorf("could not save configuration to file %s: %w", filename, err)
	}
	return filename, nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
