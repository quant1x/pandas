package stat

import "sort"

const (
	kMaximumPercentage = float64(0.9999) // 最大百分比
	kMinimumPercentage = float64(0.0001) // 最小百分比
	kMaxScale          = 10000           // 最大尺寸
)

func FindPercent(zScore float64) (percent float64) {
	index := sort.SearchFloat64s(__percentToZscore, zScore)
	percent = float64(index) / float64(kMaxScale)
	return percent
}

func FindZScore(percent float64) (zScore float64) {
	// 第一步约束 percentage在0~9999范围内
	index := int(percent*(kMaxScale)) % kMaxScale
	return __percentToZscore[index]
}

// ConfidenceIntervalToZscore 通过置信区间百分比
func ConfidenceIntervalToZscore(confidenceInterval float64) (zScore float64) {
	// 约束 percentage在0~9999范围内
	index := int(confidenceInterval*(kMaxScale)) % kMaxScale
	return __z_table[index]
}

func ZscoreToConfidenceInterval(zScore float64) (confidenceInterval float64) {
	index := __SearchFloat64s(__z_table, zScore)
	confidenceInterval = float64(index) / float64(kMaxScale)
	return confidenceInterval
}

func __SearchFloat64s_v1(a []float64, x float64) int {
	n := sort.Search(len(a), func(i int) bool { return a[i] >= x })
	return n
}

func __SearchFloat64s(a []float64, x float64) int {
	n, found := sort.Find(len(a), func(i int) int {
		m := x - a[i]
		return int(m * kMaxScale)
	})
	if !found {
		n = n - 1
	}
	return n
}
