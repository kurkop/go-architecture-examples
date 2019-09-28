package main

var counters map[string]int

func createCounter(ID string) {
	counters[ID] = 0
}

func incrementCounter(ID string) {
	counters[ID]++
}