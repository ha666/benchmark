package main

import (
	"context"
	"github.com/ha666/benchmark/go-micro/proto/arith"
	"github.com/ha666/logs"
	"github.com/montanaflynn/stats"
	"sync"
	"sync/atomic"
	"time"
)

var (
	concurrency      int
	requestPerClient = 10000
	clients          map[int]arith.ArithService
)

func main() {

	n := concurrency
	m := requestPerClient

	logs.Info("concurrency: %d，requests per client: %d", n, m)

	args := &arith.Args{
		A: 10,
		B: 20,
	}

	var wg sync.WaitGroup
	wg.Add(n * m)

	var trans uint64
	var transOK uint64

	d := make([][]int64, n, n)

	//it contains warmup time but we can ignore it
	totalT := time.Now().UnixNano()
	for i := 0; i < n; i++ {
		dt := make([]int64, 0, m)
		d = append(d, dt)
		go func(i int) {
			for j := 0; j < m; j++ {
				t := time.Now().UnixNano()
				reply, err := clients[i].Mul(context.Background(), args)
				if err != nil {
					logs.Emergency("err:%s", err.Error())
				}
				t = time.Now().UnixNano() - t

				d[i] = append(d[i], t)

				if err == nil && reply.C > 0 {
					atomic.AddUint64(&transOK, 1)
				}
				atomic.AddUint64(&trans, 1)
				wg.Done()
			}
		}(i)
	}

	wg.Wait()
	totalT = time.Now().UnixNano() - totalT
	totalT = totalT / 1000000
	logs.Info("took %d ms for %d requests", totalT, n*m)

	totalD := make([]int64, 0, n*m)
	for _, k := range d {
		totalD = append(totalD, k...)
	}
	totalD2 := make([]float64, 0, n*m)
	for _, k := range totalD {
		totalD2 = append(totalD2, float64(k))
	}

	mean, _ := stats.Mean(totalD2)
	median, _ := stats.Median(totalD2)
	max, _ := stats.Max(totalD2)
	min, _ := stats.Min(totalD2)
	p99, _ := stats.Percentile(totalD2, 99.9)

	logs.Info("sent     requests    : %d", n*m)
	logs.Info("received requests    : %d", atomic.LoadUint64(&trans))
	logs.Info("received requests_OK : %d", atomic.LoadUint64(&transOK))
	logs.Info("throughput  (TPS)    : %d", int64(n*m)*1000/totalT)
	logs.Info("mean: %.f ns, median: %.f ns, max: %.f ns, min: %.f ns, p99: %.f ns", mean, median, max, min, p99)
	logs.Info("mean: %d ms, median: %d ms, max: %d ms, min: %d ms, p99: %d ms", int64(mean/1000000), int64(median/1000000), int64(max/1000000), int64(min/1000000), int64(p99/1000000))
}
