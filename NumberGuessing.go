package main

func main() {
	for {
		setupGame()
		playGame()

		if !askToPlayAgain() {
			break
		}

	}

}
