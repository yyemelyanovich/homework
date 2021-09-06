package main

import (
	"sync"
	"time"
)

const (
	maxScore         = 15
	player1, player2 = "Player 1", "Player 2"
)

func play(players [2]chan struct{}) (winner string) {
	var points [2]int
	var pointsMutex sync.Mutex
	var group sync.WaitGroup

	run := func(player1Ch chan struct{}, player2Ch chan struct{}, pointsPlayer2 *int) {
		defer group.Done()
		for {
			pointsMutex.Lock()
			if gameEnded(points) {
				pointsMutex.Unlock()
				return
			}
			pointsMutex.Unlock()
			<-player1Ch
			if ballMissed() {
				pointsMutex.Lock()
				(*pointsPlayer2)++
				pointsMutex.Unlock()
			}
			player2Ch <- struct{}{}
		}
	}

	group.Add(2)

	players[1] <- struct{}{}
	go run(players[0], players[1], &points[1])
	go run(players[1], players[0], &points[0])

	group.Wait()
	if points[0] > points[1] {
		return player1
	}
	return player2
}

func ballMissed() bool {
	now := time.Now()
	return now.UnixNano()%9 == 0
}

func gameEnded(points [2]int) bool {
	for _, p := range points {
		if p > maxScore {
			return true
		}
	}

	return false
}

func PlayGame() string {
	players := [2]chan struct{}{
		make(chan struct{}, 1),
		make(chan struct{}, 1),
	}

	defer func() {
		for _, p := range players {
			close(p)
		}
	}()

	return play(players)
}
