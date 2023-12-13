package main

import (
	"bufio"
	"flag"
	"hangman"
	"math/rand"
	"os"
	"strings"
)

func main() {
	fichierDictionnaire := "./words.txt"
	mot := Dictionnaire(fichierDictionnaire)

	var letterFile string
	var startWith string
	var dictionnaireAPI string
	var spectator string

	flag.StringVar(&letterFile, "letterFile", "", "standard.txt")
	flag.StringVar(&startWith, "startWith", "", "save.txt")
	flag.StringVar(&dictionnaireAPI, "dictionnaireAPI", "", "api")
	flag.StringVar(&spectator, "spectator", "", "")
	flag.Parse()

	files := hangman.ListeFichiers("./savespect")
	for i := 0; i < len(files); i++ {
		if spectator == files[i] {
			filecible := "savespect/" + files[i]
			hangman.PenduSpect(filecible)
		}
	}
	if letterFile == "standard.txt" {
		hangman.AsciiOrNot = true
		hangman.Starting()
		hangman.Pendu(hangman.ToUpper(mot))
	} else if startWith == "save.txt" {
		hangman.Hp, hangman.Result, mot, _ = hangman.Start("save.txt")
		word := strings.Trim(mot, "\"")
		hangman.Starting()
		hangman.Pendu(hangman.ToUpper(word))
	} else if dictionnaireAPI == "api" {
		hangman.Starting()
		hangman.Pendu(hangman.ToUpper((hangman.DictionnaireAPI())))
	} else {
		hangman.Starting()
		hangman.Pendu(hangman.ToUpper(mot))
	}
}

func Dictionnaire(fichierDictionnaire string) string {
	fichier, err := os.Open(fichierDictionnaire)
	if err != nil {
		os.Exit(0)
	}
	defer fichier.Close()
	var mots []string
	scanner := bufio.NewScanner(fichier)
	for scanner.Scan() {
		mots = append(mots, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		os.Exit(0)
	}
	index := rand.Intn(len(mots))
	return mots[index]
}
