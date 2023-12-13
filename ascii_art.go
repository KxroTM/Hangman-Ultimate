package hangman

import (
	"bufio"
	"fmt"
	"os"
)

func PrintAscii(caractere string) {
	var debut int
	var fin int

	var lignes [][]string

	for i := 0; i < len(caractere); i++ {
		lettre := int(caractere[i])
		if lettre == 95 {
			debut = 234
			fin = 239
		} else {
			debut = (lettre - 65) * 9
			fin = debut + 6
		}

		fichier, err := os.Open("standard.txt")
		if err != nil {
			os.Exit(0)
		}
		defer fichier.Close()

		scanner := bufio.NewScanner(fichier)
		defaut := 0
		carctereLignes := []string{}
		for scanner.Scan() {
			if defaut >= debut && defaut <= fin {
				carctereLignes = append(carctereLignes, scanner.Text())
			}
			defaut++
		}
		lignes = append(lignes, carctereLignes)
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < len(caractere); j++ {
			if i < len(lignes[j]) {
				fmt.Print(lignes[j][i])
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
