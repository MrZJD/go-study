package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 设计一个工作池
type Job struct { // 任务结构体
	id      int
	randNum int
}

type Result struct { // 任务结果结构体
	job     Job
	randRes int
}

var jobs = make(chan Job, 10)

// var results = make(chan Result, 10)

func calc(i int) int {
	sum := 0
	num := i
	for num != 0 {
		digit := num % 10
		sum += digit
		num /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		res := calc(job.randNum)
		_ = Result{
			job:     job,
			randRes: res,
		}
		// results <- output
		fmt.Println("[Resolve Mission]: id:", job.id, "input:", job.randNum, "output:", res, "unresolve mission length: ", len(jobs))
	}
	fmt.Println("[Worker Done]")
	wg.Done()
}

func createWorkerPool(numOfWorks int) { // 创建工作池
	var wg sync.WaitGroup
	for i := 0; i < numOfWorks; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	fmt.Println("[All Worker Done]")
	// close(results)
}

func allcMission(numOfJobs int) { // 分配任务
	for i := 0; i < numOfJobs; i++ {
		r := rand.Intn(999)
		jobs <- Job{
			i,
			r,
		}
		fmt.Println("[Allction Mission]: id:", i, "input:", r)
	}
	close(jobs)
}

// func getResult(done chan bool) {
// 	for result := range results {
// 		fmt.Printf("Job id: %d, input: %d, output: %d\n", result.job.id, result.job.randNum, result.randRes)
// 	}
// 	done <- true
// }

func main() {
	startTime := time.Now()
	numsOfJobs := 50
	go allcMission(numsOfJobs) // 分配一百个任务

	// done := make(chan bool)
	// go getResult(done)

	numsOfWorkers := 10 // 分配工作协程
	createWorkerPool(numsOfWorkers)

	// <-done
	endTime := time.Now()
	fmt.Println("Total Using Time: ", endTime.Sub(startTime).Seconds(), "seconds")
}
