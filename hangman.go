package hangman

import (
	"fmt"
	"math/rand"
	"time"
)

var Hp int = 10
var Result []rune
var AsciiOrNot = false
var Save string
var ASave []rune

func Pendu(mot string) {
	Clear()
	var compt []string
	var temp bool
	var choix string
	var tempresult []rune
	fmt.Println("Good Luck, you have ", Hp, " attempts.")
	fmt.Print("\n")

	if len(Result) == 0 {
		//  remplace l'affichage des lettres par des "_"
		for i := 0; i < len(mot); i++ {
			Result = append(Result, '_')
			tempresult = append(tempresult, '_')
		}

		// Placement de la lettre random
		numerolettrerdm := rand.Intn(max(1, len(mot)/2-1))
		lettre := string(mot[numerolettrerdm])
		for i := 0; i < len(mot)/2; i++ {
			if lettre == string(mot[i]) {
				Result[i] = rune(lettre[0])
				tempresult[i] = rune(lettre[0])
			}
		}
	}
	dessinpendu(Hp)
	ASave = tempresult

	// boucle tant que le mot n'est pas trouvé ou que le joueur a toute sa vie
	for string(Result) != mot && Hp != 0 {
		// Affiche le mot a trouvé
		fmt.Print("\n")
		Print(Result, string(Result))

		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Println("Choose a letter : ")
		fmt.Scan(&choix)
		Save = Save + choix + " "
		for !(rune(choix[0]) >= 'A' && rune(choix[0]) <= 'Z') && !(rune(choix[0]) >= 'a' && rune(choix[0]) <= 'z') {
			Clear()
			fmt.Print("\n")
			dessinpendu(Hp)
			fmt.Print("\n")

			// Affiche le mot a trouvé
			fmt.Print("\n")
			Print(Result, string(Result))

			//
			fmt.Print("\n")
			fmt.Println("Choose a letter ! ")
			fmt.Scan(&choix)
			Save = Save + choix + " "
		}
		choix = ToUpper(choix)
		if len(choix) >= 2 && choix != "STOP" {
			advanced_feature(Hp, choix, mot)
			if Hp <= 0 {
				Loose(mot, Hp)
			}
		}
		for j := 0; j < len(compt); j++ {
			for choix == compt[j] {
				Clear()
				fmt.Print("\n")
				dessinpendu(Hp)
				//Affiche le mot a trouvé
				fmt.Print("\n")
				Print(Result, string(Result))

				//
				fmt.Print("\n")
				fmt.Print("\n")
				fmt.Println("Letter already chose !")
				fmt.Println("Choose an other letter ! You have ", Hp, " attempts.")
				fmt.Scan(&choix)
				Save = Save + choix + " "
				for !(rune(choix[0]) >= 'A' && rune(choix[0]) <= 'Z') && !(rune(choix[0]) >= 'a' && rune(choix[0]) <= 'z') {
					Clear()
					fmt.Print("\n")
					dessinpendu(Hp)
					//Affiche le mot a trouvé
					fmt.Print("\n")
					Print(Result, string(Result))

					//
					fmt.Print("\n")
					fmt.Println("Choose a letter ! ")
					fmt.Scan(&choix)
					Save = Save + choix + " "
				}
				choix = ToUpper(choix)
			}
		}
		if choix == "STOP" {
			Stop(Hp, Result, mot)
		}
		compt = append(compt, choix)
		temp = false
		for i := 0; i < len(mot); i++ {
			if choix == string(mot[i]) {
				Result[i] = rune(choix[0])
				temp = true
			}
		}
		if temp == false {
			fmt.Print("Not present in the word, ", Hp-1, " attempts remaining")
			Hp--
			time.Sleep(2000 * time.Millisecond)
		}
		Clear()
		fmt.Print("\n")
		dessinpendu(Hp)
	}
	if string(Result) == mot {
		Win(mot)
	}
	if Hp <= 0 {
		Loose(mot, Hp)
	}
}
