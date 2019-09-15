package main

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"testing"
)


var  sortSimpleKeys  = keys{
	caseIgnore:    false,
	showOnlyFirst: false,
	sortReverse:   false,
	sortNumbers:   false,
	sortByColumn:  false,
	columnSort:    0,
	outputFile:    "",
}
func TestSum(t *testing.T) {
	expected := "1\n2"
	out := new(bytes.Buffer)
	err := proceedFile(out,"tests/test_1.txt", sortSimpleKeys)
	if !reflect.DeepEqual(expected, out.String()) || err != nil {
		t.Error("expected", expected, "result", out.String())
	}
}



func TestSumReverse(t *testing.T) {
	var  sortKeys  = keys{
		caseIgnore:    false,
		showOnlyFirst: false,
		sortReverse:   true,
		sortNumbers:   false,
		sortByColumn:  false,
		columnSort:    0,
		outputFile:    "",
	}
	expected := "2\n1"
	out := new(bytes.Buffer)
	err := proceedFile(out,"tests/test_1.txt", sortKeys)
	if !reflect.DeepEqual(expected, out.String()) || err != nil {
		t.Error("expected", expected, "result", out.String())
	}
}



func Test3(t *testing.T) {
	var  sortKeys  = keys{
		caseIgnore:    false,
		showOnlyFirst: false,
		sortReverse:   false,
		sortNumbers:   false,
		sortByColumn:  true,
		columnSort:    1,
		outputFile:    "",
	}
	expected := "2 1\n1 2"
	out := new(bytes.Buffer)
	err := proceedFile(out,"tests/test_2.txt", sortKeys)
	if !reflect.DeepEqual(expected, out.String()) || err != nil {
		t.Error("expected", expected, "result", out.String())
	}
}

func Test4(t *testing.T) {
	var  sortKeys  = keys{
		caseIgnore:    true,
		showOnlyFirst: false,
		sortReverse:   false,
		sortNumbers:   false,
		sortByColumn:  true,
		columnSort:    1,
		outputFile:    "",
	}
	expected := "1 Mem\n2 MEM"
	out := new(bytes.Buffer)
	err := proceedFile(out,"tests/test_3.txt", sortKeys)
	if !reflect.DeepEqual(expected, out.String()) || err != nil {
		t.Error("expected", expected, "result", out.String())
	}
}

func Test5(t *testing.T) {
	var  sortKeys  = keys{
		caseIgnore:    true,
		showOnlyFirst: false,
		sortReverse:   false,
		sortNumbers:   false,
		sortByColumn:  false,
		columnSort:    0,
		outputFile:    "",
	}
	expected := "Mem\nMEM"
	out := new(bytes.Buffer)
	err := proceedFile(out,"tests/test_4.txt", sortKeys)
	if !reflect.DeepEqual(expected, out.String()) || err != nil {
		t.Error("expected", expected, "result", out.String())
	}
}

func Test6(t *testing.T) {
	var  sortKeys  = keys{
		caseIgnore:    false,
		showOnlyFirst: true,
		sortReverse:   false,
		sortNumbers:   false,
		sortByColumn:  false,
		columnSort:    0,
		outputFile:    "",
	}
	expected := "MEM"
	out := new(bytes.Buffer)
	err := proceedFile(out,"tests/test_5.txt", sortKeys)
	if !reflect.DeepEqual(expected, out.String()) || err != nil {
		t.Error("expected", expected, "result", out.String())
	}
}

func Test7(t *testing.T) {
	var  sortKeys  = keys{
		caseIgnore:    false,
		showOnlyFirst: true,
		sortReverse:   false,
		sortNumbers:   false,
		sortByColumn:  false,
		columnSort:    0,
		outputFile:    "",
	}

	out := new(bytes.Buffer)
	err := proceedFile(out,"FILE.TXT", sortKeys)
	if err == nil {
		t.Error("expected", "error", "result", out.String())
	}
}

func Test8(t *testing.T) {
	var  sortKeys  = keys{
		caseIgnore:    false,
		showOnlyFirst: false,
		sortReverse:   false,
		sortNumbers:   false,
		sortByColumn:  false,
		columnSort:    0,
		outputFile:    "output_test.txt",
	}
	expected:="1\n2"
	out := new(bytes.Buffer)
	err := proceedFile(out,"tests/test_1.txt", sortKeys)
	result,_:=ioutil.ReadFile("output_test.txt")

	if !reflect.DeepEqual(expected, string(result)) || err != nil {
		t.Error("expected", expected, "result", out.String())
	}
}

func Test69(t *testing.T) {
	var  sortKeys  = keys{
		caseIgnore:    true,
		showOnlyFirst: true,
		sortReverse:   false,
		sortNumbers:   false,
		sortByColumn:  false,
		columnSort:    0,
		outputFile:    "",
	}
	expected := "Mem"
	out := new(bytes.Buffer)
	err := proceedFile(out,"tests/test_4.txt", sortKeys)
	if !reflect.DeepEqual(expected, out.String()) || err != nil {
		t.Error("expected", expected, "result", out.String())
	}
}



