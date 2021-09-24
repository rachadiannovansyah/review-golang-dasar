package stackoverflow

import (
	"log"
)

func main() {
	loop(1000000)
}

func loop(n int) {
	if n == 0 {
		return
	}

	log.Println("loop ke ", +n)
	loop(n - 1)
}
