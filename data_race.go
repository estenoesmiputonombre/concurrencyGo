package main

import (
	"sync"
	"time"
)

func dataRace() int {
	var i int
	go func() {
		i++
	}()
	return i
}

func dataRaceTimeout(wg *sync.WaitGroup, t time.Duration) int {
	defer wg.Done()
	var i int
	go func() {
		i++
	}()
	time.Sleep(t)
	return i
}

// we dont resolve the data race condition, but we can ensure data syncronization.
func dataRaceLock(wg *sync.WaitGroup) int {
	defer wg.Done()
	var (
		i            int
		memoryAccess sync.Mutex
	)
	go func() {
		memoryAccess.Lock()
		i++
		memoryAccess.Unlock()
	}()
	return i
}

func group(wg *sync.WaitGroup, f func() int) int {
	defer wg.Done()
	return f()
}
