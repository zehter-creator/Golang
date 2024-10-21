package chapter_1

import (
	"log"
)

func main() {
	pound := money.New(100, money.GBP)
	twoPounds, err := pound.Add(pound)

	if err != nil {
		log.Fatal(err)
	}

	parties, err := twoPounds.Split(3)

	if err != nil {
		log.Fatal(err)
	}

	parties[0].Display() // £0.67
	parties[1].Display() // £0.67
	parties[2].Display() // £0.66
}
