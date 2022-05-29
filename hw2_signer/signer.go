package main

import (
	"sync"
)

// сюда писать код
// https://github.com/Terminator637/hw2_signer
func ExecutePipeline(jobs ...job) {
	in := make(chan any)
	wg := &sync.WaitGroup{}

	wg.Add(len(jobs))
	for _, j := range jobs {
		out := make(chan any)
		//
		go func(j job, in, out chan any, wg *sync.WaitGroup) {
			defer wg.Done()
			defer close(out)

			j(in, out)
		}(j, in, out, wg)
		//
		in = out
	}
	wg.Wait()
}

func SingleHash(in, out chan any) {
	for v := range in {
		//v := v.(string)
		//out <- DataSignerCrc32(v) + "~" + DataSignerCrc32(DataSignerMd5(v))
		out <- v
	}
}
func MultiHash(in, out chan any) {
	for v := range in {
		//v := v.(string)

		//for i := 0; i < 6; i++ {
		//	i := strconv.Itoa(i)
		//	out <- DataSignerCrc32(i + v)
		//}
		out <- v
	}
}

func CombineResults(in, out chan any) {
	for v := range in {
		//v := v.(string)
		out <- v
	}
}
