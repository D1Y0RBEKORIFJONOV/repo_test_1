package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)

const (
	HEIGHT = 20
	WIDTH  = 40
	KYEL   = "\x1B[33m"

	STOP  int = 0
	RIGHT int = 1
	UP    int = 2
	DOWN  int = 3
	LEFT  int = 4
)

var (
	userX, userY, foodX, foodY int
	bigFood, score             int
	nTail                      int
	tailX, tailY               [100]int
	GameOver, isStop           bool
	highLevel                  bool
	dir                        int
)

func SetUp() {
	bigFood = 0
	var d int
	fmt.Println("Oyn Darajasini tanlang")
	fmt.Print("[0][1]: ")
	fmt.Scan(&d)
	if d == 1 {
		highLevel = true
	} else {
		highLevel = false
	}
	nTail = 0
	GameOver = false
	isStop = false
	dir = STOP
	score = 0
	userX = WIDTH / 2
	userY = HEIGHT / 2
	foodX = rand.Intn(WIDTH - 1)
	foodY = rand.Intn(HEIGHT - 1)
}

func Draw() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	fmt.Println("\033[0;32mDIYOR_DEV: SNAKE_GAME! 🐍")
	for i := 0; i < WIDTH+1; i++ {
		fmt.Print(KYEL + "☠️")
	}
	fmt.Println()

	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			if j == 0 || j == WIDTH-1 {
				fmt.Print(KYEL + "☠️")
			}
			if i == userY && j == userX-1 {
				fmt.Print("😺")
				j += 1
			} else if i == foodY && j == foodX {
				if bigFood == 4 {
					fmt.Print("🐭")
					j += 1
				} else {
					fmt.Print("🐭")
					j += 1
				}
			} else {
				isPrint := false
				if !isStop {
					for k := 0; k < nTail; k++ {
						if tailX[k] == j && tailY[k] == i {
							if k == nTail-1 {
								fmt.Print("\033[0;31mo")
							} else {
								fmt.Print("\033[0;34m=")
							}
							isPrint = true
						}
					}
				}
				if !isPrint {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}

	for i := 0; i < WIDTH+1; i++ {
		fmt.Print(KYEL + "☠️")
	}
	fmt.Println()
	fmt.Println("\x1B[0mSCORE:", score)
}

func Input() {
	var input string
	fmt.Scan(&input)
	switch input {
	case "a":
		dir = LEFT
	case "s":
		dir = DOWN
	case "d":
		dir = RIGHT
	case "w":
		dir = UP
	case "q":
		dir = STOP
	case "x":
		GameOver = true
	}
}

func Logic() {
	prevx := tailX[0]
	prevy := tailY[0]
	var prevx2, prevy2 int
	tailX[0] = userX
	tailY[0] = userY
	for i := 1; i < nTail; i++ {
		prevx2 = tailX[i]
		prevy2 = tailY[i]
		tailX[i] = prevx
		tailY[i] = prevy
		prevx = prevx2
		prevy = prevy2
	}
	switch dir {
	case LEFT:
		userX--
		isStop = false
	case RIGHT:
		userX++
		isStop = false
	case UP:
		userY--
		isStop = false
	case DOWN:
		userY++
		isStop = false
	case STOP:
		isStop = true
		userX = userX
		userY = userY
	}

	if (userX-1 == foodX && foodY == userY) || (userX == foodX && foodY == userY+1) {
		if bigFood == 6 {
			score += 50
			bigFood = 0
		} else {
			score += 10
			foodX = rand.Intn(WIDTH - 2)
			foodY = rand.Intn(HEIGHT - 2)
		}
		nTail++
		bigFood++
	}

	if !highLevel {
		if userX < 0 {
			userX = WIDTH - 1
		} else if userX >= WIDTH-1 {
			userX = 0
		}

		if userY < 0 {
			userY = HEIGHT
		} else if userY >= HEIGHT {
			userY = 0
		}
	} else if highLevel == true {
		if userX <= 0 || userX > WIDTH-2 || userY < 0 || userY+1 > HEIGHT {
			GameOver = true
		}
	}

	if !isStop {
		for i := 0; i < nTail; i++ {
			if tailX[i] == userX && tailY[i] == userY {
				GameOver = true
			}
		}
	}
}

func DelAy(a int) {
	var add int
	var i int
	var time int
	time = a * 10000000
	for i = 0; i < time; i++ {
		add *= i
		add++
		add++
	}
}

func main() {
start:
	SetUp()

	for !GameOver {
		Draw()
		Input()
		Logic()
		DelAy(1)
	}

	fmt.Println("Your score:", score)
	var c string
	fmt.Print("ENTER[1] TO REGAME: ")
	fmt.Scan(&c)
	if c == "\n" {
		goto start
	}
}
