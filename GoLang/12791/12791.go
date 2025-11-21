package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	albumYear := []int{
		1967,
		1969,
		1970,
		1971,
		1972,
		1973,
		1973,
		1974,
		1975,
		1976,
		1977,
		1977,
		1979,
		1980,
		1983,
		1984,
		1987,
		1993,
		1995,
		1997,
		1999,
		2002,
		2003,
		2013,
		2016,
	}
	albumName := []string{
		"DavidBowie",
		"SpaceOddity",
		"TheManWhoSoldTheWorld",
		"HunkyDory",
		"TheRiseAndFallOfZiggyStardustAndTheSpidersFromMars",
		"AladdinSane",
		"PinUps",
		"DiamondDogs",
		"YoungAmericans",
		"StationToStation",
		"Low",
		"Heroes",
		"Lodger",
		"ScaryMonstersAndSuperCreeps",
		"LetsDance",
		"Tonight",
		"NeverLetMeDown",
		"BlackTieWhiteNoise",
		"1.Outside",
		"Earthling",
		"Hours",
		"Heathen",
		"Reality",
		"TheNextDay",
		"BlackStar",
	}

	var q int
	fmt.Fscanln(reader, &q)

	var ans int
	var s, e int
	for i := 0; i < q; i++ {
		fmt.Fscanln(reader, &s, &e)

		ans = 0
		for j := 0; j < 25; j++ {
			if albumYear[j] >= s && albumYear[j] <= e {
				ans++
			}
		}
		fmt.Fprintln(writer, ans)
		for j := 0; j < 25; j++ {
			if albumYear[j] >= s && albumYear[j] <= e {
				fmt.Fprintln(writer, albumYear[j], albumName[j])
			}
		}
	}
	writer.Flush()
}
