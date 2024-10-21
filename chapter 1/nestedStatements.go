package chapter_1

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
)

var onlyOnce sync.Once
var dice = []int{1, 2, 3, 4, 5, 6}
var r *rand.Rand

func rollDice(SEED int64) int {
	onlyOnce.Do(func() {
		r = rand.New(rand.NewSource(SEED))
		fmt.Println(r.Uint64())
		fmt.Println(r.Uint64())
	})
	return dice[r.Intn(len(dice))]
}

func _() {
	var elma, karpuz, nar, dolunay, moonStone = 12, 15, 24, 10, 5
	fmt.Println("GEMIDEN BATAN MALLAR BUNLAR VATANDAS"+"ELMAM ", elma, karpuz, nar, dolunay, moonStone)
	fmt.Println("Bak sen şu işe! Çok fazla metin taşım var! Bir tane almak ister misin? Ama şunu bil ki aldığın zaman gökten bir hazine düşecek. kaç tane istiyorsun?")

	var iGot int
	_, err := fmt.Scan(&iGot)
	if err != nil {
		log.Println("Girdi hatası:", err)
		return
	}

	fmt.Println("Girdiğiniz değer:", iGot)

	switch {
	case iGot <= 12:
		fmt.Println("Taşları alıdın, pazardan yavaşça uzaklaşıyorsun.")
		fmt.Println(" Kaderde seni ne beklediğini öğrenmek için zar attık. Zar sensin yol sensin 1 den 6'ya kadar bir sayı seç")

		var zar = rollDice(0)
		var playerNumber int
		_, err := fmt.Scan(&playerNumber)
		if err != nil {
			log.Fatal(err)
		}
		res := zar - playerNumber
		fmt.Println("kaderindeki sayı:", res)
	default:
		fmt.Println("Seçimin uygun değil.")
	}
}
