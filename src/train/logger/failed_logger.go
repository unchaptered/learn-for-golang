package main

import "log"

func failed_logger() {
	log.SetPrefix("[KEVIN]")
	log.Print("[failed_logger]", "Start")
	// log.Fatal("[failed_logger]", "Cracked before End")
	log.Print("[failed_logger]", "End")
}
