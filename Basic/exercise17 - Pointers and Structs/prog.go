package main

import (
	"fmt"
)

type ElectionResult struct {
	Name  string
	Votes int
}

func main() {
	var initialVotes int
	initialVotes = 2

	var counter *int
	counter = NewVoteCounter(initialVotes)

	fmt.Println(*counter == initialVotes)

	var votes int
	votes = 3

	var voteCounter *int
	voteCounter = &votes

	fmt.Println(VoteCount(voteCounter))
	// Output: 3

	var nilVoteCounter *int
	fmt.Println(VoteCount(nilVoteCounter))

	IncrementVoteCount(voteCounter, 2)

	fmt.Println(votes == 5)        // true
	fmt.Println(*voteCounter == 5) // true

	var result *ElectionResult
	result = &ElectionResult{
		Name:  "John",
		Votes: 32,
	}

	fmt.Println(DisplayResult(result))

	var finalResults = map[string]int{
		"Mary": 10,
		"John": 51,
	}

	DecrementVotesOfCandidate(finalResults, "Mary")

	fmt.Println(finalResults["Mary"])

}

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(numOfVotes int) *int {
	return &numOfVotes
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}

	return *counter
}

// IncrementVoteCount increments the value in a vote counter
func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
}

// NewElectionResult creates a new election result
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	return &ElectionResult{
		Name:  candidateName,
		Votes: votes,
	}
}

// DisplayResult creates a message with the result to be displayed
func DisplayResult(result *ElectionResult) string {
	return fmt.Sprintf("%s (%d)", result.Name, result.Votes)
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	_, ok := results[candidate]
	if ok {
		results[candidate]--
	}
}
