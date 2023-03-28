package main

import (
	"log"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func SequentialGenID() {
	log.Println("GEN TUAN TU 10.000.000 ID !!!")

	// nanoid gen ID tuan tu: thoi gian  gen 10.000.000 ID khoang 5.16s
	log.Printf("time start nanoid: %s", time.Now())
	for i := 0; i < 10000000; i++ {
		id, _ := gonanoid.New()
		if i == 9999999 {
			log.Println(id)
		}
	}
	log.Printf("end start nanoid: %s", time.Now())
	log.Println("                             ")

	// uuid gen ID tuan tu: thoi gian  gen 10.000.000 ID khoang 3.30s
	log.Printf("time start uuid: %s", time.Now())
	for i := 0; i < 10000000; i++ {
		id := uuid.New()
		if i == 9999999 {
			log.Println(id)
		}
	}
	log.Printf("end start uuid: %s", time.Now())
	log.Println("                             ")

	// snowflake gen ID tuan tu: thoi gian gen 10.000.000 ID khoang 2,44s
	n, _ := snowflake.NewNode(1)
	log.Printf("time start snowflake : %s", time.Now())
	for i := 0; i < 10000000; i++ {
		// tạo ID
		id := n.Generate()
		if i == 9999999 {
			log.Println(id)
		}
	}
	log.Printf("time start snowflake : %s", time.Now())
}
