package chapter_1

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	marketName := "Malatya"
	fmt.Println(marketName + "pazarı")
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	var fruits, elma int = 200, random.Intn(100)
	karpuzVar := fruits - elma
	fmt.Println(karpuzVar, "alma")
}
