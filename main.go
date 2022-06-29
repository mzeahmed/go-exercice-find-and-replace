package main

import (
	"fmt"
	"strings"
)

// ProcessLine searches for old in line to replace it by new
// it returns found=true, if the pattern was found, res withthe resulting string
// and occ with the number of occurence of old
func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	oldLower := strings.ToLower(old)
	newLower := strings.ToLower(new)
	res = line

	if strings.Contains(line, old) || strings.Contains(line, oldLower) {
		found = true
		occ += strings.Count(line, old)
		occ += strings.Count(line, oldLower)
		res = strings.Replace(line, old, new, -1)
		res = strings.Replace(res, oldLower, newLower, -1)
	}

	return found, res, occ
}

func main() {
	found, res, occ := ProcessLine(
		"Go was conceived in 2007 to improve programming productivity at Google",
		"Go",
		"Pyton",
	)

	fmt.Println(found, res, occ)
}
