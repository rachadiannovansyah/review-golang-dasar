package main

import "fmt"

func main() {

	fmt.Println("------Kenalan Pertama--------")
	fullName := "rachadian novansyah"
	age := 24
	healthy := true
	mmr := 7500.78

	fmt.Printf("hello, my name is %s! \n", fullName)
	fmt.Printf("Umur saya %d! \n", age)
	fmt.Printf("Saya sehat? %t \n", healthy)
	fmt.Printf("MMR Dota saya adalah %.3f \n", mmr)

	fmt.Println("--------Kondisional----------")
	// kondisional
	if mmr == 10000 {
		fmt.Println("selamat anda lulus dengan nilai sempurna")
	} else if mmr > 5000 {
		fmt.Println("lulus gan")
	} else if mmr > 3000 && mmr <= 5000 {
		fmt.Println("hampir lulus gan")
	} else {
		fmt.Println("waduh anda tidak lulus")
	}

	if percent := mmr / 100; percent >= 100 {
		fmt.Printf("%.1f %s perfect! \n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%1.f %s bagus gan \n", percent, "%")
	} else {
		fmt.Printf("%1.f %s not bad! \n", percent, "%")
	}

	switch {
	case mmr == 3000:
		fmt.Println("standar")
	case (mmr > 3000) && (mmr < 5000):
		fmt.Println("mulai imba")
	case (mmr > 5000) && (mmr < 8500):
		fmt.Println("IMBA Gan!")
		fallthrough
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better bro!")
		}
	}

	fmt.Println("--------Perulangan----------")
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	for i := 1; i < 6; i++ {
		for j := i; j < 6; j++ {
			fmt.Print(j, "_")
		}

		fmt.Println()
	}

outerloop:
	for i := 1; i < 6; i++ {
		for j := 1; j < 6; j++ {
			if i == 3 {
				break outerloop
			}

			fmt.Print("Martiks [", i, "][", j, "]", "\n")
		}
	}

	fmt.Println("--------Array----------")

	var agility = [4]string{"rikimaru", "nevermore", "juggernaut", "lanaya"}

	fmt.Println("Jumlah heroes \t \t", len(agility))
	fmt.Println("Mereka adalah \t", agility)

	fmt.Println("--------Slice----------")
	var strength = []string{"huskar", "axe", "tuskar", "butcher"}
	fmt.Println("Jumlah heroes \t \t", len(strength))
	fmt.Println("Mereka adalah \t", strength)

	var strength1 = strength[0:3]
	fmt.Println("Jumlah heroes \t \t \t", len(strength1))
	fmt.Println("Jumlah heroes cap \t \t", cap(strength1))
	fmt.Println("Mereka adalah \t", strength1)

}
