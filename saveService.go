package main

import (
	"os"
)

func saveGameResult(rankingsJson []byte) {
	file, err := os.OpenFile(
		"results.json",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write(rankingsJson)
	file.Write([]byte("\n"))

}
