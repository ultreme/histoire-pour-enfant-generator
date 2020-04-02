package hpeg

/*
 [x] avoir un début et une fin
 [x] ça parle d'animaux
 [ ] ça parle d'aventure
 [ ] il y a des trucs répétitifs
 [ ] ça parle des parties du corps
 [ ] ça parle de la maitresse
 [ ] ça parle de la maman et du papa
 [ ] ça parle des jours de la semaine
 [ ] ça parle de gateaux au chocolat
 [ ] ça parle de jardin, de foret, de jungle
*/

type Elem interface {
	GetPhrase() string
	GetTitle() string
}

type Story struct {
	Elems []Elem
}

func NewStory() Story {
	resetIntros()
	resetAnimals()
	return Story{
		Elems: make([]Elem, 0),
	}
}

func (s *Story) AddElement(elem Elem) {
	s.Elems = append(s.Elems, elem)
}

func (s *Story) Tell() []string {
	output := []string{}

	output = append(output, s.Elems[0].GetTitle(), "")

	output = append(output, GetIntro())

	for _, elem := range s.Elems {
		output = append(output, elem.GetPhrase()+".")
	}

	output = append(output, GetOutro()+".")

	return output
}
