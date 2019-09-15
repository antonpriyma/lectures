package main

import (
	"github.com/mgutz/logxi/v1"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type keys struct {
	caseIgnore    bool
	showOnlyFirst bool
	sortReverse   bool
	sortNumbers   bool
	sortByColumn  bool
	columnSort    int
	outputFile    string
}

func sortSlice(words []string, keys keys) []string {//Сортировка массива строк
	if keys.sortByColumn {
		if keys.caseIgnore {
			sort.Slice(words, func(i, j int) bool {
				log.Debug("Words[i]: ", strings.Fields(words[i]))
				return strings.ToLower(strings.Fields(words[i])[keys.columnSort]) < strings.ToLower(strings.Fields(words[j])[keys.columnSort])
			})
		} else {
			sort.Slice(words, func(i, j int) bool {
				log.Debug("Words[i]: ", strings.Fields(words[i]))
				return strings.Fields(words[i])[keys.columnSort] < strings.Fields(words[j])[keys.columnSort]
			})
		}
	} else {
		if keys.caseIgnore {
			sort.Slice(words, func(i, j int) bool {
				return strings.ToLower(words[i]) < strings.ToLower(words[j])
			})
		} else {
			sort.Strings(words)
		}
	}

	log.Info("Sorted array: ", words)
	return words
}

func proceedFile(out io.Writer, path string, keys keys) error {//Обработка и вывод
	words, err := parseFile(path)
	if err != nil {
		return err
	}
	words = sortSlice(words, keys)
	if keys.showOnlyFirst {
		words = removeDuplicates(keys.caseIgnore, words)
	}
	if keys.sortReverse {
		words = reverse(words)
	}

	if keys.outputFile != "" {
		err = writeToFile(keys, words)
		if err != nil {
			return err
		}
	} else {
		err := printToConsole(out, words)
		if err != nil {
			return err
		}
	}

	return nil
}

func printToConsole(out io.Writer, words []string) error {
	output := strings.Join(words, "\n")
	_, err := out.Write([]byte(output))
	return err
}

func writeToFile(keys keys, words []string) error {
	output := strings.Join(words, "\n")
	f, err := os.Create(keys.outputFile)
	if err != nil {
		return err
	}
	_, err = f.WriteString(output)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}

func parseFile(path string) ([]string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	log.Info("Read from file: ", b)
	f := func(c rune) bool {
		return c == '\n'
	}
	words := strings.FieldsFunc(string(b), f)
	log.Info("Split array of bytes: ", words)
	return words, nil
}

func reverse(words []string) []string {
	for i := 0; i < len(words)/2; i++ {
		j := len(words) - i - 1
		words[i], words[j] = words[j], words[i]
	}
	return words
}

func removeDuplicates(ignoreCase bool, words []string) (list []string) {
	keys := make(map[string]bool)
	list := make([]string, 0)
	for i, word := range words {
		if ignoreCase {
			word = strings.ToLower(word)
		}
		if _, value := keys[word]; !value {
			keys[word] = true
			list = append(list, words[i])
		}
	}
	return
}

func main() {
	args := os.Args
	if len(args)<2{
		panic("Usage: go run sort.go <filename>")
	}
	var keys keys
	log.Debug("Args", args)
	for i, arg := range args {
		switch arg {
		case "-f":
			keys.caseIgnore = true
		case "-u":
			keys.showOnlyFirst = true
		case "-r":
			keys.sortReverse = true
		case "-n":
			keys.sortNumbers = true
		case "-o":
			keys.outputFile = args[i+1]
		case "-k":
			keys.sortByColumn = true
			keys.columnSort, _ = strconv.Atoi(args[i+1])

		default:

		}
	}
	path := args[1]
	err := proceedFile(os.Stdout, path, keys)
	if err != nil {
		log.Fatal(err.Error())
	}

}
