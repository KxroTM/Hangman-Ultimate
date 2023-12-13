package hangman

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func ToUpper(s string) string {
	var txt []rune
	tab := []rune(s)
	for i := 0; i <= len(tab)-1; i++ {
		if tab[i] >= 97 && tab[i] <= 122 {
			txt = append(txt, tab[i]-32)
		} else {
			txt = append(txt, tab[i])
		}
	}
	return string(txt)
}

func ReadFile(fichier string, start int, LignesSouhaite int) {
	f, err := os.Open(fichier)
	if err != nil {
		fmt.Print("Erreur : Vous n'avez pas tous les fichiers requis !")
		os.Exit(0)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	compt := 0
	ligneFinale := LignesSouhaite

	for scanner.Scan() {
		compt++
		if compt >= start && compt <= ligneFinale {
			fmt.Println(scanner.Text())
		}
	}
}

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Print(Result []rune, lettre string) {

	if AsciiOrNot == true {
		PrintAscii(lettre)
	} else {
		for i := 0; i < len(Result); i++ {
			fmt.Print(string(Result[i]))
			fmt.Print(" ")
		}
	}
}

func Win(mot string) {
	Clear()
	fmt.Print("\n")
	fmt.Println("Congrats !")
	fmt.Print("The word was ", mot)
	SpectRq(Save, mot)
	os.Exit(0)
}

func Loose(mot string, Hp int) {
	Clear()
	fmt.Print("\n")
	dessinpendu(Hp)
	fmt.Println("Nice try !")
	fmt.Print("The word was ", mot)
	SpectRq(Save, mot)
	os.Exit(0)
}

func SpectRq(save string, mot string) {
	var choix string
	var name = "!"
	for choix != "y" && choix != "n" {
		fmt.Print("\n")
		fmt.Println("Do you want to save the game ? (y for yes n for no ) ")
		fmt.Scan(&choix)
	}
	if choix == "y" {
		for CaractereIncorrect(name) {
			fmt.Println("What's the save's name : (only letters and digits) ")
			fmt.Scan(&name)
		}
		SaveSpect(name, Save, mot, ASave)

	} else if choix == "n" {
		os.Exit(0)
	}
}

func CaractereIncorrect(s string) bool {
	for _, char := range s {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') && (char < '0' || char > '9') {
			return true
		}
	}
	return false
}

func SaveSpect(fichier string, contenu string, mot string, lettre []rune) {
	file, err := os.Create("savespect/" + fichier + ".txt")
	if err != nil {
		os.Exit(0)
	}
	defer file.Close()

	_, err = file.WriteString(mot)
	if err != nil {
		os.Exit(0)
	}
	_, err = file.WriteString("\n")
	if err != nil {
		os.Exit(0)
	}
	_, err = file.WriteString(string(lettre))
	if err != nil {
		os.Exit(0)
	}
	_, err = file.WriteString("\n")
	if err != nil {
		os.Exit(0)
	}
	for i := 0; i < len(contenu); i++ {
		if string(contenu[i]) == " " {
			_, err = file.WriteString("\n")
			if err != nil {
				os.Exit(0)
			}
		} else {
			_, err = file.WriteString(string(contenu[i]))
			if err != nil {
				os.Exit(0)
			}
		}
	}
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("Save complete.")
}

func ListeFichiers(chemin string) []string {
	dir, err := os.Open(chemin)
	if err != nil {
		os.Exit(0)
	}
	defer dir.Close()

	fichiers, err := dir.Readdir(-1)
	if err != nil {
		os.Exit(0)
	}

	var fichier []string
	for _, status := range fichiers {
		if !status.IsDir() {
			fichier = append(fichier, status.Name())
		}
	}

	return fichier
}
