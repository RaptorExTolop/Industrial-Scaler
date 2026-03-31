package main

type State int32
const (
	StateIdle State = iota
	StateRunning 
	StateBrokenLvl1
	StateBrokenLvl2
	StateBrokenLvl3
	StateBrokenFinal
)

type Factory struct {
	Input int32 // the item the factory accepts
	Output int32 // the factories output
	Efficiency float32 // how much output per input
	Speed float32 // how many times the factory runs per cycle
	State State
}
