package main

import "fmt"

func main() {
	fmt.Println("Don't run...test")
}

func partitionLabels(s string) []int {
	m := make(map[string]int)
	partitions := []string{}

	// create array of lower case letters
	j := 1
	for i := 97; i < 123; i++ {
		char := rune(i)
		m[string(char)] = j
		j++
	}
	a := []int{}

	j = 1
	partition := ""
	nextLetter := ""
	// iterate characters in string
	for i := 0; i < len(s)-1; i++ {
		currLetter := s[i : i+1]
		currLetterNum := m[s[i:i+1]]
		nextLetter = s[i+1 : i+2]
		nextLetterNum := m[s[i+1:i+2]]
		diffBetweenLetters := 0

		// measure distance between chars
		if currLetterNum > nextLetterNum {
			diffBetweenLetters = currLetterNum - nextLetterNum
		} else {
			diffBetweenLetters = nextLetterNum - currLetterNum
		}

		// build next partition
		partition += currLetter

		// if char distance > 2 and it's been enough chars since last partition,
		// make a new one
		if diffBetweenLetters > 2 && j > 7 {
			a = append(a, len(partition))
			partitions = append(partitions, partition)
			partition = ""
			j = 1
		}
		j++
	}
	partition += nextLetter
	if len(partition) < 7 {
		a[len(a)-1] += len(partition)
		partitions[len(a)-1] += partition

	} else {
		a = append(a, len(partition))
		partitions = append(partitions, partition)
	}

	// clean up partitions based on char frequency
	/*charFreq := make(map[string]int)
	for _, partition := range partitions {
		for i := 0; i < len(partition); i++ {
			currLetter := s[i : i+1]
			_, r := charFreq[currLetter]
			if r {
				charFreq[currLetter]++
			} else {
				charFreq[currLetter] = 1
			}
		}
		fmt.Printf("\t%s -> %v\n", partition, charFreq)

	}
	*/

	// if letter is next to letter that has freq>1 - partition

	return a
}
