package main

import (
	"flag"
	"github.com/mgutz/logxi/v1"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

const none = -1

type keys struct {
	caseIgnore    bool
	showOnlyFirst bool
	sortReverse   bool
	sortNumbers   bool
	sortByColumn  bool
	columnSort    int
	outputFile    string
}

func sortSlice(words []string, keys keys) []string { //Сортировка массива строк
	if keys.columnSort != none {
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

func proceedFile(words []string, keys keys) []string { //Обработка и вывод
	words = sortSlice(words, keys)
	if keys.showOnlyFirst {
		words = removeDuplicates(keys.caseIgnore, words)
	}
	if keys.sortReverse {
		words = reverse(words)
	}
	return words
}

func output(out io.Writer, words []string, keys keys) (err error) {
	if keys.outputFile != "" {
		err = writeToFile(keys, words)
		if err != nil {
			return err
		}
	} else {
		if out == nil {
			out = os.Stdout
		}
		err := printToConsole(out, words)
		if err != nil {
			return err
		}
	}
	return nil
}

func readFile(path string) (words []string, err error) {
	words, err = parseFile(path)
	if err != nil {
		return
	}
	return
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
		err = errors.Wrap(err, "Error while creating file")
		return err
	}
	_, err = f.WriteString(output)
	if err != nil {
		err = errors.Wrap(err, "Error while writing to file")
		return err
	}
	err = f.Close()
	if err != nil {
		err = errors.Wrap(err, "Error while closing file")
		return err
	}
	return nil
}

func parseFile(path string) ([]string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		err = errors.Wrap(err, "Error while reading file")
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
	list = make([]string, 0)
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

func sortFile(out io.Writer, path string, keys keys) error {
	words, err := readFile(path)
	if err != nil {
		return err
	}
	words = proceedFile(words, keys)
	err = output(out, words, keys)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("Usage: go run sort.go <filename>")
	}
	var keys keys
	log.Debug("Args", args)

	path := flag.String("p", "test.txt", "input file")
	flag.BoolVar(&keys.caseIgnore, "f", false, "case ignore")
	flag.BoolVar(&keys.showOnlyFirst, "u", false, "show only first element")
	flag.BoolVar(&keys.sortReverse, "r", false, "sort desc")
	flag.StringVar(&keys.outputFile, "o", "", "output file")
	flag.IntVar(&keys.columnSort, "k", -1, "sort by column")
	flag.Parse()

	err := sortFile(nil, *path, keys)

	if err != nil {
		log.Fatal(err.Error())
	}

}
