package utils

import (
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
)

const (
	channel         = 1000
	quantityRequest = 10000
)

func SyncGenUUID() {

	// uuid: thoi gian gen khoang 1.000.000 ID la khoang
	wg := &sync.WaitGroup{}
	wg.Add(channel)
	total := 0
	startTime := time.Now()
	for i := 0; i < channel; i++ {
		count := make(chan int)
		go genUUID(count, wg, &total)
	}
	wg.Wait()
	log.Println("total: ", total)
	log.Println("thoi gian gen 1.000.000 ID voi uuid la: ", time.Now(), startTime)

}

func genUUID(ch chan int, wg *sync.WaitGroup, total *int) {
	m := make(map[string]int)
	for i := 0; i < quantityRequest; i++ {
		id := uuid.New()
		m[id.String()] = 1
	}
	*total += len(m)
	wg.Done()
	ch <- len(m)
}

func SyncGenSnowflake() {
	// snowflake: thoi gian gen khoang 10.000.000 ID la khoang
	wg := &sync.WaitGroup{}
	wg.Add(channel)
	total := 0
	startTime := time.Now()
	for i := 0; i < channel; i++ {
		count := make(chan int)
		go genSnowflake(count, wg, &total)
	}
	wg.Wait()
	log.Println("total: ", total)
	log.Println("thoi gian gen 10.000.000 ID voi snowflake la: ", time.Now(), startTime)
}

func genSnowflake(ch chan int, wg *sync.WaitGroup, total *int) {
	m := make(map[string]int)
	for i := 0; i < quantityRequest; i++ {
		id := uuid.New()
		m[id.String()] = 1
	}
	*total += len(m)
	wg.Done()
	ch <- len(m)
}

// trong TH nhieu channel thi thoi gian gen id cua uuid xap xi snowflake
