package main

import (
	"time"
)

type status int
type ByteSize float64

type task struct {
	title string
	done  bool
	due   *time.Time
}

type task2 struct {
	title  string
	status status
	due    *time.Time
}

const (
	UNKNOWN status = 0
	TODO    status = 1
	DONE    status = 2
)

const (
	BAD status = iota
	SOSO
	GOOD
)

const (
	minsu = iota
	chulsoo
	minhee
)

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	var myTask = task{"work", false, nil}
}
