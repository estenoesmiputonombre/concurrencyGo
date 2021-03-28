package main

import (
	"sync"
	"testing"
	"time"
)

func TestDataRace(t *testing.T) {
	max := 100
	var (
		a, b int
		wg   sync.WaitGroup
	)
	for ; max > 0; max-- {
		wg.Add(1)
		if group(&wg, dataRace) == 0 {
			a++
		} else {
			b++
		}
	}
	wg.Wait()
	t.Logf("Number of times that the function returns 0: %d\nNumber of times that the function returns 1: %d\n", a, b)
}

func TestDataRaceTimeout(t *testing.T) {
	max := 100
	var (
		a, b int
		wg   sync.WaitGroup
	)
	for ; max > 0; max-- {
		wg.Add(1)
		if dataRaceTimeout(&wg, time.Millisecond) == 0 {
			a++
		} else {
			b++
		}
	}
	wg.Wait()
	t.Logf("Number of times that the function returns 0: %d\nNumber of times that the function returns 1: %d\n", a, b)
}

func TestDataRaceLock(t *testing.T) {
	max := 100
	var (
		a, b int
		wg   sync.WaitGroup
	)
	for ; max > 0; max-- {
		wg.Add(1)
		if dataRaceLock(&wg) == 0 {
			a++
		} else {
			b++
		}
	}
	wg.Wait()
	t.Logf("Number of times that the function returns 0: %d\nNumber of time that the function returns 1: %d\n", a, b)
}
