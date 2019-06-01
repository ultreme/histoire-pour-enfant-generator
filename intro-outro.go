package hpeg

var (
	Intros []string
	Outros []string
)

func init() {
	Intros = []string{
		"Il était une fois... dans un monde fabuleux...",
		"En 50 avant jésus-christ...",
		"L'autre jour...",
		"Il était une fois...",
	}
	Outros = []string{
		"Et ils vécurent heureux et eurent beaucoup d'enfants",
		"Elle devint une princesse",
		"Et c'était la fête",
		"Et c'était plutôt cool",
		"Et c'est ainsi que l'homme marcha sur la lune",
		"Et il eut son poids en pomme",
		"Et elle devint astronaute",
		"Et il mangea ses céréales au gouter",
		"Et il devint magicien",
		"Et il devint un lapin",
		"Fin",
	}
}

func GetIntro() string {
	return PickAndDelete(&Intros)
}

func GetOutro() string {
	return PickAndDelete(&Outros)
}
