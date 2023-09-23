package main

import "testing"

func TestGame(t *testing.T) {
	player := NewPlayer()
	go startUIloop(player)
	startGameLoop(player)
}
