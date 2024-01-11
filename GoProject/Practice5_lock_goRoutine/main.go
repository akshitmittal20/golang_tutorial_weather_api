package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var m = sync.Mutex{}
var wg = sync.WaitGroup{} //first we hve made a program for DB delay simulation. now we will try to add this function to have a parallel execution of task without concurrency
var dbdata = []string{"1d1", "id2", "id3", "id4", "id5"}
var results = []string{}

func dbcall(i int) { //simulate db call delay
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is :", dbdata[i])
	m.Lock()
	results = append(results, dbdata[i])
	m.Unlock()
	wg.Done()
}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbdata); i++ {
		wg.Add(1)
		go dbcall(i)
	}
	wg.Wait()
	fmt.Printf(("\n Total executon time is : %v"), time.Since((t0)))
	fmt.Printf("\n The results are %v", results)
}
