package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/danicc097/yaegi-script/internal"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func main() {
	content, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	content = removeShebang(content)

	i := interp.New(interp.Options{
		Unrestricted: true,
	})

	if err := i.Use(stdlib.Symbols); err != nil {
		log.Fatalf("Failed to load stdlib: %v", err)
	}

	if err := i.Use(internal.Symbols); err != nil {
		log.Fatalf("Failed to load internal symbols: %v", err)
	}

	_, err = i.Eval(string(content))
	if err != nil {
		log.Fatalf("Evaluation error: %v", err)
	}
}

func readInput() ([]byte, error) {
	if len(os.Args) > 1 {
		return os.ReadFile(os.Args[1])
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return io.ReadAll(os.Stdin)
	}

	return nil, fmt.Errorf("please specify a file or pipe input")
}

func removeShebang(content []byte) []byte {
	scanner := bufio.NewScanner(bytes.NewReader(content))
	if scanner.Scan() {
		firstLine := scanner.Text()
		if strings.HasPrefix(firstLine, "#!") {
			rest := []byte{}
			for scanner.Scan() {
				rest = append(rest, append(scanner.Bytes(), '\n')...)
			}
			return rest
		}
	}
	return content
}
