package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ProcessLine recherche l'ancienne ligne pour la remplacer par le nouveau
// il renvoie 'found=true', si le motif a été trouvé, 'res' avec la chaîne résultante
// et 'occ' avec le nombre d'occurrences de l'ancien
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

func FindReplaceFile(src string, dst string, old string, new string) (occ int, lines []int, err error) {
	// On ouvre le fichier
	srcFile, err := os.Open(src)
	if err != nil {
		return occ, lines, err
	}
	defer srcFile.Close()

	// On créé un nouveau fichier pour l'insertion du nouveau texte
	dstFile, err := os.Create(dst)
	if err != nil {
		return occ, lines, err
	}
	defer dstFile.Close()

	// On gere les cas go serait present dans un mot, exemple google
	old = old + " "
	new = new + " "

	lineIdx := 1
	scanner := bufio.NewScanner(srcFile)
	writer := bufio.NewWriter(dstFile)
	defer writer.Flush()
	for scanner.Scan() {
		found, res, o := ProcessLine(scanner.Text(), old, new)

		// Si on trouve le mot
		if found {
			// on augmente le compteur d'occurence global par rappor tà celui qui nous ai renvoyé
			occ += o
			lines = append(lines, lineIdx)
		}

		fmt.Fprintf(writer, res)
		lineIdx++
	}

	return occ, lines, nil
}

func main() {
	old := "Go"
	new := "Python"

	occ, lines, err := FindReplaceFile("wikigo.txt", "wikipyton.txt", old, new)
	if err != nil {
		fmt.Printf("Error while executing find replace: %v \n", err)
	}

	fmt.Println("== Summary ==")
	defer fmt.Println(" == End of summarry ==")
	fmt.Printf("Number of occurences of %v: %v \n", old, occ)
	fmt.Printf("Number of lines: %d \n", len(lines))

	// On affiche les lignes où les occurrences sont presente
	fmt.Print("Lines: [ ")
	len := len(lines)
	for i, l := range lines {
		fmt.Printf("%v", l)
		if i < len-1 {
			fmt.Printf(" - ")
		}
	}
	fmt.Println(" ]")
}
