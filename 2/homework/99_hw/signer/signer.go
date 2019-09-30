package main

import (
	"sort"
	"strconv"
	"strings"
	"sync"
)

var mu = &sync.Mutex{}

func SingleHash(in chan interface{}, out chan interface{}) {

	wg := &sync.WaitGroup{}
	for temp := range in {
		wg.Add(1)
		p, _ := temp.(int)
		mu.Lock()
		help := DataSignerMd5(strconv.Itoa(p))
		mu.Unlock()
		go func() {
			defer wg.Done()
			SingleHashHelp(p, help, out)
		}()
	}
	wg.Wait()
}

func SingleHashHelp(p int, s string, out chan interface{}) {
	var result string
	ch := make(chan string)
	ch1 := make(chan string)
	go func(num int, s string, ch chan string) {
		ch <- DataSignerCrc32(s)
		return
	}(0, strconv.Itoa(p), ch)
	go func(num int, s string, ch chan string) {
		ch <- DataSignerCrc32(s)
		return
	}(1, s, ch1)
	result += <-ch
	result += "~" + <-ch1
	out <- result
}

func MultiHashHelp(p string, out chan interface{}) {
	result := ""
	arr := make([]byte, 6, 6)
	for i := '0'; i < '6'; i++ {
		index, _ := strconv.Atoi(string(i))
		arr[index] = byte(i)
	}
	resultArray := make([]string, 6, 6)
	wg := &sync.WaitGroup{}
	for index, elem := range arr {
		wg.Add(1)
		go func(num int, s string, arr []string) {
			defer wg.Done()
			arr[num] = DataSignerCrc32(s)
		}(index, string(string(elem)+p), resultArray)
	}
	wg.Wait()
	for _, elem := range resultArray {
		result += elem
	}
	out <- result
}

func MultiHash(in chan interface{}, out chan interface{}) {
	wg := &sync.WaitGroup{}
	for temp := range in {
		p, _ := temp.(string)
		wg.Add(1)
		go func() {
			defer wg.Done()
			MultiHashHelp(p, out)
		}()
	}
	wg.Wait()
}

func CombineResults(in chan interface{}, out chan interface{}) {
	var result string
	var data []string
	for temp := range in {
		p, _ := temp.(string)
		data = append(data, p)
	}
	sort.Strings(data)
	result = strings.Join(data, "_")
	out <- result
}

func ExecutePipeline(arr ...job) {
	testCh := make([]chan interface{}, len(arr)+1, len(arr)+1)
	for i := 0; i < len(arr)+1; i++ {
		testCh[i] = make(chan interface{})
	}
	wg := &sync.WaitGroup{}
	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			defer close(testCh[i+1])
			arr[i](testCh[i], testCh[i+1])
		}(i)
	}
	wg.Wait()
}
