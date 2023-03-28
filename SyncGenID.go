package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
)

func SyncGenUUID() {

	// uuid: thoi gian gen khoang 1.000.000 ID la 0.84s
	wg := &sync.WaitGroup{}
	wg.Add(4)
	total := 0
	startTime := time.Now()
	for i := 0; i < 4; i++ {
		count := make(chan int)
		go StandardUUID(count, wg, &total)
	}
	wg.Wait()
	log.Println("total: ", total)
	log.Println("thoi gian gen 1.000.000 ID voi uuid la: ", time.Now(), startTime)

}

func StandardUUID(ch chan int, wg *sync.WaitGroup, total *int) {
	m := make(map[string]int)
	for i := 0; i < 250000; i++ {
		id := uuid.New()
		m[id.String()] = 1
	}
	*total += len(m)
	wg.Done()
	ch <- len(m)
}

func SyncGenSnowflake() {
	// snowflake: thoi gian gen khoang 1.000.000 ID la 0.061s
	wg := &sync.WaitGroup{}
	wg.Add(4)
	startTime := time.Now()
	for i := 0; i < 4; i++ {
		fmt.Println("* create chanel ", i)
		go genSnowflake(i, wg)

	}
	wg.Wait()
	log.Println("thoi gian gen 1.000.000 ID voi snowflake la: ", time.Now(), startTime)
}

func genSnowflake(node int, wg *sync.WaitGroup) {
	n, err := snowflake.NewNode(int64(node))
	if err != nil {
		println(err)
		os.Exit(1)
	}
	for i := 0; i < 250000; i++ {
		// tạo ID
		id := n.Generate()
		if i == 249999 {
			log.Println(id)
		}

	}
	wg.Done()
}
