package hangman

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type ListeDeMots struct {
	Text string `json:"mot"`
}

var mots []ListeDeMots

func DictionnaireAPI() string {
	WordList("words.txt")
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/add", AddMot)
	r.GET("/list", Liste)
	r.GET("/motrdm", MotRandom)

	go func() {
		r.Run(":8080")
	}()

	randomWord, err := RecupMot()
	if err != nil {
		os.Exit(0)
	}
	Clear()
	donneeJson := randomWord
	var data map[string]string
	err = json.Unmarshal([]byte(donneeJson), &data)
	if err != nil {
		os.Exit(0)
	}
	mot := data["mot"]
	return mot
}

func WordList(fichier string) {
	donnee, err := os.ReadFile(fichier)
	if err != nil {
		os.Exit(0)
	}

	Lignes := strings.Split(string(donnee), "\n")
	for _, ligne := range Lignes {
		ligne = strings.TrimSpace(ligne)
		if ligne != "" {
			mots = append(mots, ListeDeMots{Text: ligne})
		}
	}
}

func AddMot(c *gin.Context) {
	var word ListeDeMots
	if err := c.BindJSON(&word); err != nil {
		os.Exit(0)
	}

	mots = append(mots, word)
	c.JSON(http.StatusCreated, word)
}

func MotRandom(c *gin.Context) {
	if len(mots) == 0 {
		os.Exit(0)
	}

	rdmPos := rand.Intn(len(mots))
	randomWord := mots[rdmPos]

	c.JSON(http.StatusOK, randomWord)
}

func Liste(c *gin.Context) {
	c.JSON(http.StatusOK, mots)
}

func RecupMot() (string, error) {
	apiURL := "http://localhost:8080/motrdm"

	requete, err := http.Get(apiURL)
	if err != nil {
		os.Exit(0)
	}
	defer requete.Body.Close()

	if requete.StatusCode != http.StatusOK {
		os.Exit(0)
	}

	data, err := io.ReadAll(requete.Body)
	if err != nil {
		os.Exit(0)
	}

	return string(data), nil
}
