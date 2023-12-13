package hangman

func dessinpendu(compt int) {
	if compt == 1 {
		pendujpg10()
	}
	if compt == 2 {
		pendujpg9()
	}
	if compt == 3 {
		pendujpg8()
	}
	if compt == 4 {
		pendujpg7()
	}
	if compt == 5 {
		pendujpg6()
	}
	if compt == 6 {
		pendujpg5()
	}
	if compt == 7 {
		pendujpg4()
	}
	if compt == 8 {
		pendujpg3()
	}
	if compt == 9 {
		pendujpg2()
	}
	if compt == 10 {
		pendujpg1()
	}
}

func pendujpg1() {
	ReadFile("hangman.txt", 1, 7)
}

func pendujpg2() {
	ReadFile("hangman.txt", 9, 15)
}

func pendujpg3() {
	ReadFile("hangman.txt", 17, 23)
}

func pendujpg4() {
	ReadFile("hangman.txt", 25, 31)
}

func pendujpg5() {
	ReadFile("hangman.txt", 33, 39)
}

func pendujpg6() {
	ReadFile("hangman.txt", 41, 47)
}

func pendujpg7() {
	ReadFile("hangman.txt", 49, 55)
}

func pendujpg8() {
	ReadFile("hangman.txt", 57, 63)
}

func pendujpg9() {
	ReadFile("hangman.txt", 65, 71)
}

func pendujpg10() {
	ReadFile("hangman.txt", 73, 79)
}
