package main

import (
	"time"
)

func spawnScanner() {
	go func() {
		for {
			if err := scanAllDiscussions(); err != nil {
				Println("[COMMENT-CACHE] failed to scan for discussions:", err)
			}
			if err := scanAllManga(); err != nil {
				Println("[COMMENT-CACHE] failed to scan for comments:", err)
			}
			time.Sleep(time.Duration(*interval * int(time.Minute.Nanoseconds())))
		}
	}()
}
