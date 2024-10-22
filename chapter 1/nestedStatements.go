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

// Zar atma fonksiyonu
func rollDice(SEED int64) int {
	onlyOnce.Do(func() {
		r = rand.New(rand.NewSource(SEED))
		fmt.Println("Random seed:", r.Uint64())
		fmt.Println("Random seed:", r.Uint64())
	})
	return dice[r.Intn(len(dice))]
}

func main() {
	var elma, karpuz, nar, dolunay, moonStone = 12, 15, 24, 10, 5
	fmt.Println("GEMIDEN BATAN MALLAR BUNLAR VATANDAS"+" ELMAM ", elma, karpuz, nar, dolunay, moonStone)
	fmt.Println("Bak sen şu işe! Çok fazla metin taşım var! Bir tane almak ister misin? Ama şunu bil ki aldığın zaman gökten bir hazine düşecek. Kaç tane istiyorsun?")

	// Girdi alma
	var iGot int
	_, err := fmt.Scan(&iGot) // Girdiyi belleğe almak için &iGot kullanılır
	if err != nil {
		log.Println("Girdi hatası:", err)
		return
	}

	fmt.Println("Girdiğiniz değer:", iGot)

	switch {
	case iGot <= 12:
		fmt.Println("Taşları aldın, pazardan yavaşça uzaklaşıyorsun.")
		fmt.Println(" Kaderde seni ne beklediğini öğrenmek için zar attık. Zar sensin, yol sensin. 1 den 6'ya kadar bir sayı seç")

		// Zar atma
		var zar = rollDice(0)
		var playerNumber int
		_, err := fmt.Scan(&playerNumber)
		if err != nil {
			log.Fatal("Girdi hatası:", err)
		}

		// Zar sonucunu hesaplama
		res := zar - playerNumber
		fmt.Println("Kaderindeki sayı:", res)
	default:
		fmt.Println("Seçimin uygun değil.")
	}
}
