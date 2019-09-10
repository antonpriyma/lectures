package main

import (
	"bytes"
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestSum(t *testing.T) {
	s := "2 + 3"
	expected := "5"
	out := new(bytes.Buffer)
	err := calc(out, strings.NewReader(s))
	if !reflect.DeepEqual(expected, out.String()) || err != nil {
		t.Error("expected", expected, "result", out.String())
	}
}

func TestDiff(t *testing.T) {
	s := "3 - 2"
	expected := "1"
	out := new(bytes.Buffer)
	err := calc(out, strings.NewReader(s))
	if !reflect.DeepEqual(expected, out.String()) || err != nil {
		t.Error("expected", expected, "result", out.String())
	}
}

func TestMult(t *testing.T) {
	s := "4 * 5"
	expected := "20"
	out := new(bytes.Buffer)
	err := calc(out, strings.NewReader(s))
	if !reflect.DeepEqual(expected, out.String()) || err != nil {
		t.Error("expected", expected, "result", out.String())
	}
}

func TestDiv(t *testing.T) {
	s := "10 / 2"
	expected := "5"
	out := new(bytes.Buffer)
	err := calc(out, strings.NewReader(s))
	if !reflect.DeepEqual(expected, out.String()) || err != nil {
		t.Error("expected", expected, "result", out.String())
	}
}

func Test1(t *testing.T) {
	s := "((2*3)/2)*4"
	expected := "12"
	out := new(bytes.Buffer)
	err := calc(out, strings.NewReader(s))
	if !reflect.DeepEqual(expected, out.String()) || err != nil {
		t.Error("expected", expected, "result", out.String())
	}
}

func Test3(t *testing.T) {
	s := "1 * 2 * 3 * * * * = = = ="
	out := new(bytes.Buffer)
	err := calc(out, strings.NewReader(s))
	if err == nil {
		t.Error("expected error")
	}
}
func Test4(t *testing.T) {
	s := "1 2 3 / / / / = = = ="
	out := new(bytes.Buffer)
	err := calc(out, strings.NewReader(s))
	if err == nil {
		t.Error("expected error")
	}
}

func Test5(t *testing.T) {
	s := "10 /\n2"
	expected := "5"
	out := new(bytes.Buffer)
	err := calc(out, strings.NewReader(s))
	if !reflect.DeepEqual(expected, out.String()) || err != nil {
		t.Error("expected", expected, "result", out.String())
	}
}

func Test6(t *testing.T) {
	s := "10 / 2 ="
	expected := "5"
	out := new(bytes.Buffer)
	err := calc(out, strings.NewReader(s))
	if expected != out.String() || err != nil {
		t.Error("expected", expected, "result", out.String())
	}
}

func Test8(t *testing.T) {
	s := "10     /      2  ="
	expected := "5"
	out := new(bytes.Buffer)
	err := calc(out, strings.NewReader(s))
	if expected != out.String() || err != nil {
		t.Error("expected", expected, "result", out.String())
	}
}


func Test7(t *testing.T) {
	s := "*"
	expected := errors.New("Empty Stack")
	out := new(bytes.Buffer)
	err := calc(out, strings.NewReader(s))
	if err == nil {
		t.Error("expected", expected, "result", out.String())
	}
}

func TestEmpty1(t *testing.T) {
	s := "="
	expected := errors.New("Empty Stack")
	out := new(bytes.Buffer)
	err := calc(out, strings.NewReader(s))
	if err == nil {
		t.Error("expected", expected, "result", out.String())
	}
}

func TestEmpty2(t *testing.T) {
	s := "+"
	expected := errors.New("Empty Stack")
	out := new(bytes.Buffer)
	err := calc(out, strings.NewReader(s))
	if err == nil {
		t.Error("expected", expected, "result", out.String())
	}
}

func TestEmpty3(t *testing.T) {
	s := "5 +"
	expected := errors.New("Empty Stack")
	out := new(bytes.Buffer)
	err := calc(out, strings.NewReader(s))
	if err == nil {
		t.Error("expected", expected, "result", out.String())
	}
}
func TestEmpty4(t *testing.T) {
	s := "5 /"
	expected := errors.New("Empty Stack")
	out := new(bytes.Buffer)
	err := calc(out, strings.NewReader(s))
	if err == nil {
		t.Error("expected", expected, "result", out.String())
	}
}

func TestEmpty5(t *testing.T) {
	s := "/"
	expected := errors.New("Empty Stack")
	out := new(bytes.Buffer)
	err := calc(out, strings.NewReader(s))
	if err == nil {
		t.Error("expected", expected, "result", out.String())
	}
}

func TestEmpty6(t *testing.T) {
	s := "-"
	expected := errors.New("Empty Stack")
	out := new(bytes.Buffer)
	err := calc(out, strings.NewReader(s))
	if err == nil {
		t.Error("expected", expected, "result", out.String())
	}
}

func TestEmpty7(t *testing.T) {
	s := "5 - "
	expected := errors.New("Empty Stack")
	out := new(bytes.Buffer)
	err := calc(out, strings.NewReader(s))
	if err == nil {
		t.Error("expected", expected, "result", out.String())
	}
}

func TestEmpty8(t *testing.T) {
	s := "    "
	expected := errors.New("Empty Stack")
	out := new(bytes.Buffer)
	err := calc(out, strings.NewReader(s))
	if err == nil {
		t.Error("expected", expected, "result", out.String())
	}
}

func TestEmpty9(t *testing.T) {
	s := "    )"
	expected := errors.New("Empty Stack")
	out := new(bytes.Buffer)
	err := calc(out, strings.NewReader(s))
	if err == nil {
		t.Error("expected", expected, "result", out.String())
	}
}







