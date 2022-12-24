package day2

// calculateGameScore() takes a challenge and response (single letter) and returns:
// score: the honest score from challenge/response
// alternateScore: the rigged score where XYZ decides outcome
func calculateGameScore(challenge string, response string) (int, int) {
	// point for shape: 1 for Rock (X/A), 2 for Paper (Y/B), and 3 for Scissors (Z/C)
	// point for outcome: 0 if you lost, 3 if the round was a draw, and 6 if you won
	// rigged outcome: X = you must lose, Y = must draw, Z = you must win

	var score int
	var alternateScore int

	switch response {
	case "X": // you have rock OR you must lose
		score += 1
		switch challenge {
		case "A": // opponent picks rock
			score += 3          // draw
			alternateScore += 3 // pick scissors to lose
		case "B": // opponent picks paper
			// rock loses to paper
			alternateScore += 1 // pick rock to lose
		case "C": // opponent picks scissors
			score += 6          // rock beats scissors
			alternateScore += 2 // pick paper to lose
		}
	case "Y": // you have paper OR must draw
		score += 2
		switch challenge {
		case "A":
			score += 6                // paper beats rock
			alternateScore += (1 + 3) // mirror rock to draw
		case "B":
			score += 3                // draw
			alternateScore += (2 + 3) // mirror paper to draw
		case "C":
			// paper loses to scissors
			alternateScore += (3 + 3) // mirror scissors to draw
		}
	case "Z": // you have scissors OR you must win
		score += 3
		switch challenge {
		case "A":
			// scissors lose to rock
			alternateScore += (2 + 6) // pick paper to beat rock
		case "B":
			score += 6                // scissors beats paper
			alternateScore += (3 + 6) // pick scissors to beat paper
		case "C":
			score += 3                // draw
			alternateScore += (1 + 6) // pick rock to beat scissors
		}
	}

	return score, alternateScore
}
