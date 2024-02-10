package main

import "log"

func panic_logger() {
	log.Print("[panic_logger]", "Start")
	log.Panic("[panic_logger]", "panic before End")
	log.Print("[panic_logger]", "End")
}
