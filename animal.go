package hpeg

import "strings"

var (
	AnimalNames             []string
	AnimalPhrases           []string
	AnimalComplementPhrases []string
	AnimalTitles            []string
)

func init() {
	resetAnimals()
}

func resetAnimals() {
	AnimalNames = []string{
		"Potam l'hippopotame",
		"Sophie la girafe",
		"Hector le castor",
		"Mireille l'abeille",
		"Jordy le colibri",
		"Momo le caribou",
		"Babar l'éléphant",
		"Sonic le hérisson",
		"Denver le dernier dinosaure",
		"Victor le tyrannosaure",
	}
	AnimalPhrases = []string{
		"Il y avait :NAME:",
		":NAME: aimait bien se promener",
		":NAME: avait beaucoup d'amis",
		":NAME: était un aventurier",
		":NAME: n'a peur de rien",
		":NAME: a peur de tout",
		"Pour :NAME:, une bonne journée se passe avec des amis",
	}
	AnimalComplementPhrases = []string{
		"sa passion est de manger des pommes",
		"son repas préféré est le gouter",
		"son plus grand rêve est d'aller voir un film au cinéma",
		"ses parents sont les plus gentils du monde",
		"son repas favoris est la pizza",
		"son matelas est moelleux",
	}
	AnimalTitles = []string{
		"Les fabuleuses aventures de :NAME:",
		":NAME: à la plage",
		":NAME: à la piscine",
		":NAME: à la ferme",
	}
}

type Animal struct {
	Name string
}

func (a Animal) GetPhrase() string {
	phrase := PickAndDelete(&AnimalPhrases)
	complement := PickAndDelete(&AnimalComplementPhrases)
	phrase += ", " + complement

	phrase = strings.Replace(phrase, ":NAME:", a.Name, -1)
	return phrase
}

func (a Animal) GetTitle() string {
	title := PickAndDelete(&AnimalTitles)
	title = strings.Replace(title, ":NAME:", a.Name, -1)
	return title
}

func NewAnimal() Animal {
	return Animal{
		Name: PickAndDelete(&AnimalNames),
	}
}
