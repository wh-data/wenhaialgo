package back_track

var Res []string
var str string
var AlphaMap = make(map[string][]string, 0)
var InputArr = make([][]string, 0)

func PhoneSpell(input string) {
	initAlphaMap(AlphaMap)
	parseInput(input)
	max := len(InputArr)
	pickPath(1, max)
}

func initAlphaMap(alphaMap map[string][]string) {
	alphaMap["2"] = []string{"a", "b", "c"}
	alphaMap["3"] = []string{"d", "e", "f"}
	alphaMap["4"] = []string{"g", "h", "i"}
	alphaMap["5"] = []string{"j", "k", "l"}
	alphaMap["6"] = []string{"m", "n", "o"}
	alphaMap["7"] = []string{"p", "q", "r", "s"}
	alphaMap["8"] = []string{"t", "u", "v"}
	alphaMap["9"] = []string{"w", "x", "y", "z"}
}

func parseInput(input string) {
	arr := []rune(input)
	for _, v := range arr {
		InputArr = append(InputArr, AlphaMap[string(v)])
	}
}

func pickPath(level, max int) {
	if level > max {
		Res = append(Res, str)
		return
	}
	for i := 0; i < len(InputArr[level-1]); i++ {
		str = str + InputArr[level-1][i]
		pickPath(level+1, max)
		str = str[0 : len(str)-1] //recover
	}
}
