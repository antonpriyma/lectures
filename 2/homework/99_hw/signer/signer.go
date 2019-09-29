package main

import (
	"sort"
	"strconv"
	"sync"
)

func SingleHash(in chan interface{}, out chan interface{}) {
	wg := &sync.WaitGroup{}
	for temp := range in {
		wg.Add(1)
		p, _ := temp.(int)
		help := DataSignerMd5(strconv.Itoa(p))
		go SingleHashHelp(p, help, out, wg)
	}
	wg.Wait()
}

func SingleHashHelp(p int, s string, out chan interface{}, group *sync.WaitGroup) {
	defer group.Done()
	var result string
	ch := make(chan string)
	ch1 := make(chan string)
	wg1 := &sync.WaitGroup{}
	wg1.Add(1)
	go func(num int, s string, ch chan string, group *sync.WaitGroup) {
		defer group.Done()
		ch <- DataSignerCrc32(s)
		return
	}(0, strconv.Itoa(p), ch, wg1)
	wg1.Add(1)
	go func(num int, s string, ch chan string, group *sync.WaitGroup) {
		defer group.Done()
		ch <- DataSignerCrc32(s)
		return
	}(1, s, ch1, wg1)
	result += <-ch
	result += "~" + <-ch1
	out <- result
}

func MultiHashHelp(p string, out chan interface{}, group *sync.WaitGroup) {
	defer group.Done()
	result := ""
	arr := make([]byte, 6, 6)
	for i := '0'; i < '6'; i++ {
		index, _ := strconv.Atoi(string(i))
		arr[index] = byte(i)
	}
	resarr := make([]string, 6, 6)
	wg := &sync.WaitGroup{}
	for index, elem := range arr {
		wg.Add(1)
		go func(num int, s string, arr []string, group *sync.WaitGroup) {
			defer group.Done()
			arr[num] = DataSignerCrc32(s)
		}(index, string(string(elem)+p), resarr, wg) // HelpHash(index, string(string(elem)+p), resarr, wg)
	}
	wg.Wait()
	for _, elem := range resarr {
		result += elem
	}
	out <- result
}

func MultiHash(in chan interface{}, out chan interface{}) {
	wg := &sync.WaitGroup{}
	for temp := range in {
		p, _ := temp.(string)
		wg.Add(1)
		go MultiHashHelp(p, out, wg)
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
	lenght := len(data) - 1
	for i := range data {
		result += data[i]
		if i != lenght {
			result += "_"
		}
	}
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
