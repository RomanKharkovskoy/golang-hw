package main

func StudentScores() map[string]int {
	scores := make(map[string]int)
	scores["Анна"] = 3
	scores["Борис"] = 0
	scores["Вера"] = 5
	scores["Анна"] += 2
	delete(scores, "Вера")
	return scores
}

func main() {
}
