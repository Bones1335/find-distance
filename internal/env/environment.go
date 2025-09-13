package env

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SetEnv(filename string) error {
	if filename != ".env" {
		return fmt.Errorf("can't load non .env file")
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.SplitN(line, "=", 2)
		if len(splitLine) != 2 {
			return fmt.Errorf("current line variable in .env causing errors: %v", line)
		}

		err = os.Setenv(splitLine[0], strings.ReplaceAll(splitLine[1], "\"", ""))
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
