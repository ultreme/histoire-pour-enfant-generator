package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/codegangsta/cli"
	"github.com/ultreme/histoire-pour-enfant-generator"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {

	app := cli.NewApp()
	app.Name = "histoire-pour-enfant-generator"
	app.Usage = "Generates cool histoires pour enfants"

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "animals",
			Value: 4,
			Usage: "Amount of animals",
		},
	}

	app.Action = Action

	app.Run(os.Args)
}

func Action(c *cli.Context) {
	story := hpeg.NewStory()

	for i := 0; i < c.Int("animals"); i++ {
		story.AddElement(hpeg.NewAnimal())
	}

	fmt.Println(strings.Join(story.Tell(), "\n"))
}
