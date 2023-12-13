package hangman

func advanced_feature(hp int, choix string, mot string) {
	if choix != mot {
		hp = hp + 2
		Hp = hp / 2
	}
	if choix == mot {
		Win(mot)
	}
}
