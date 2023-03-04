package stat

import (
	"fmt"
	"testing"
)

func TestFindZScore(t *testing.T) {
	fmt.Println(FindZScore(0.975))
	fmt.Println(FindZScore(0.90))
	fmt.Println(FindZScore(0.95))
	fmt.Println(FindZScore(0.8621))
	fmt.Println(FindZScore(0.5))
	fmt.Println("------------------------------------------------------------")
	fmt.Println(FindPercent(1.6448551507250881))
	fmt.Println(FindPercent(1.6448))
	fmt.Println(FindPercent(1.644))
	fmt.Println(FindPercent(1.64))
	fmt.Println("------------------------------------------------------------")
	fmt.Println(FindPercent(1.959966356818851))
	fmt.Println(FindPercent(1.9599))
	fmt.Println(FindPercent(1.959))
	fmt.Println(FindPercent(1.95))
	fmt.Println(FindPercent(1.00))
	fmt.Println("------------------------------------------------------------")
}

func TestPercentToZscore(t *testing.T) {
	fmt.Println(ConfidenceIntervalToZscore(0.975))
	fmt.Println(0.9000, ConfidenceIntervalToZscore(0.9000))
	fmt.Println(ConfidenceIntervalToZscore(0.95))
	fmt.Println(ConfidenceIntervalToZscore(0.85))
	fmt.Println(ConfidenceIntervalToZscore(0.5000))
	fmt.Println("------------------------------------------------------------")
	fmt.Println(ZscoreToConfidenceInterval(1.6448551507250881))
	fmt.Println(ZscoreToConfidenceInterval(1.6448))
	fmt.Println(ZscoreToConfidenceInterval(1.644))
	fmt.Println(ZscoreToConfidenceInterval(1.64))
	fmt.Println("------------------------------------------------------------")
	fmt.Println(ZscoreToConfidenceInterval(1.959966356818851))
	fmt.Println(ZscoreToConfidenceInterval(1.9599))
	fmt.Println(ZscoreToConfidenceInterval(1.959))
	fmt.Println(ZscoreToConfidenceInterval(1.95))
	fmt.Println(1.2816 == 1.2815524)
	fmt.Println(1.2816, ZscoreToConfidenceInterval(1.2816))
	fmt.Println(ZscoreToConfidenceInterval(1.09))
	fmt.Println(ZscoreToConfidenceInterval(1.00))
	fmt.Println(ZscoreToConfidenceInterval(0))
	fmt.Println("------------------------------------------------------------")
}
