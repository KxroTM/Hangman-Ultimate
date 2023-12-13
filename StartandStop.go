package hangman

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type GameData struct {
	HP     int
	Result []rune
	Word   string
}

func Stop(hp int, result []rune, mot string) {
	Clear()
	file, err := os.Create("save.txt")
	if err != nil {
		fmt.Println("error")
	}
	donnee1, err := json.Marshal(hp)
	donnee2, err := json.Marshal(result)
	donnee3, err := json.Marshal(mot)

	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = file.Write(donnee1)
	_, err = file.Write([]byte("\n"))
	_, err = file.Write(donnee2)
	_, err = file.Write([]byte("\n"))
	_, err = file.Write(donnee3)
	if err != nil {
		fmt.Println("Wrinting error : ", err)
		return
	}
	fmt.Println("Game Saved in save.txt.")
	os.Exit(0)
}

func Start(filename string) (int, []rune, string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return 0, nil, "", err
	}

	var hp int
	var result []rune
	var word string

	lines := strings.Split(string(file), "\n")
	if len(lines) >= 1 {
		if err := json.Unmarshal([]byte(lines[0]), &hp); err != nil {
			fmt.Print("Opening error")
			os.Exit(0)
		}
	}

	if len(lines) >= 2 {
		if err := json.Unmarshal([]byte(lines[1]), &result); err != nil {
			fmt.Print("Opening error")
			os.Exit(0)
		}
	}

	if len(lines) >= 3 {
		word = lines[2]
	}

	return hp, result, word, nil
}
