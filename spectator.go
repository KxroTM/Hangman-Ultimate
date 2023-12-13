package hangman

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func PenduSpect(file string) {
	var mot, result, input = Recupdata(file)
	Clear()
	var compt []string
	var temp bool
	var choix string

	fmt.Println("Good Luck, you have ", Hp, " attempts.")
	fmt.Print("\n")

	dessinpendu(Hp)

	// boucle tant que le mot n'est pas trouvé ou que le joueur a toute sa vie
	for string(result) != mot && Hp != 0 {
		for i := 0; i < len(input); i++ {
			choix = input[i]
			// Affiche le mot a trouvé
			fmt.Print("\n")
			Print(result, string(result))

			fmt.Print("\n")
			fmt.Print("\n")
			fmt.Println("Choose a letter : ")
			fmt.Print("\n")
			fmt.Print("User's input : " + choix)
			fmt.Print("\n")
			time.Sleep(1700 * time.Millisecond)
			for !(rune(choix[0]) >= 'A' && rune(choix[0]) <= 'Z') && !(rune(choix[0]) >= 'a' && rune(choix[0]) <= 'z') {
				Clear()
				fmt.Print("\n")
				dessinpendu(Hp)
				fmt.Print("\n")

				// Affiche le mot a trouvé
				fmt.Print("\n")
				Print(result, string(result))

				//
				fmt.Print("\n")
				fmt.Println("Choose a letter ! ")
				i++
				choix = input[i]
				fmt.Print("\n")
				fmt.Print("User's input : " + choix)
				fmt.Print("\n")
				time.Sleep(1700 * time.Millisecond)
			}
			choix = ToUpper(choix)
			if len(choix) >= 2 && choix != "STOP" {
				advanced_featureSpec(Hp, choix, mot)
				if Hp <= 0 {
					LooseSpect(mot, Hp)
				}
			}
			for j := 0; j < len(compt); j++ {
				for choix == compt[j] {
					Clear()
					fmt.Print("\n")
					dessinpendu(Hp)
					//Affiche le mot a trouvé
					fmt.Print("\n")
					Print(result, string(result))

					//
					fmt.Print("\n")
					fmt.Print("\n")
					fmt.Println("Letter already chose !")
					fmt.Println("Choose an other letter ! You have ", Hp, " attempts.")
					i++
					choix = input[i]
					fmt.Print("\n")
					fmt.Print("User's input : " + choix)
					fmt.Print("\n")
					time.Sleep(1700 * time.Millisecond)
					for !(rune(choix[0]) >= 'A' && rune(choix[0]) <= 'Z') && !(rune(choix[0]) >= 'a' && rune(choix[0]) <= 'z') {
						Clear()
						fmt.Print("\n")
						dessinpendu(Hp)
						//Affiche le mot a trouvé
						fmt.Print("\n")
						Print(result, string(result))

						//
						fmt.Print("\n")
						fmt.Println("Choose a letter ! ")
						i++
						choix = input[i]
						fmt.Print("\n")
						fmt.Print("User's input : " + choix)
						fmt.Print("\n")
						time.Sleep(1700 * time.Millisecond)
					}
					choix = ToUpper(choix)
				}
			}
			if choix == "STOP" {
				os.Exit(0)
			}
			compt = append(compt, choix)
			temp = false
			for i := 0; i < len(mot); i++ {
				if choix == string(mot[i]) {
					result[i] = rune(choix[0])
					temp = true
				}
			}
			if temp == false {
				fmt.Print("Not present in the word, ", Hp-1, " attempts remaining")
				Hp--
				time.Sleep(1500 * time.Millisecond)
			}
			Clear()
			fmt.Print("\n")
			dessinpendu(Hp)
		}
		if string(result) == mot {
			WinSpect(mot)
		}
		if Hp <= 0 {
			LooseSpect(mot, Hp)
		}
	}
}

func Recupdata(nomFichier string) (motspec string, result []rune, inputs []string) {
	fichier, err := os.Open(nomFichier)
	if err != nil {
		os.Exit(0)
	}
	defer fichier.Close()

	scanner := bufio.NewScanner(fichier)

	premiereLigne := true
	deuxiemeLigne := true

	for scanner.Scan() {
		if premiereLigne {
			motspec = scanner.Text()
			premiereLigne = false
		} else if deuxiemeLigne {
			result = []rune(scanner.Text())
			deuxiemeLigne = false
		} else {
			inputs = append(inputs, scanner.Text())
		}
	}

	return motspec, result, inputs
}

func WinSpect(mot string) {
	fmt.Print("\n")
	fmt.Println("Congrats !")
	fmt.Print("The word was ", mot)
	os.Exit(0)
}

func LooseSpect(mot string, Hp int) {
	fmt.Print("\n")
	dessinpendu(Hp)
	fmt.Println("Nice try !")
	fmt.Print("The word was ", mot)
	os.Exit(0)
}

func advanced_featureSpec(hp int, choix string, mot string) {
	if choix != mot {
		hp = hp + 2
		Hp = hp / 2
	}
	if choix == mot {
		WinSpect(mot)
	}
}
