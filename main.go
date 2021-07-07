package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type playerData struct {
	id          int
	dice_amount int
	point       int
}

func check_player(data []playerData) bool {
	playing := 0
	for i := 0; i < len(data); i++ {
		if data[i].dice_amount > 0 {
			playing += 1
		}
	}
	if playing > 1 {
		return true
	} else {
		return false
	}
}

func dice_roll(dice int) []int {
	var result []int
	var roll int
	for i := 0; i < dice; i++ {
		roll = rand.Intn(6-1) + 1
		result = append(result, roll)
	}
	return result

}

func dice_roll_result(data []playerData) map[int][]int {
	res := make(map[int][]int, len(data))
	var rollResult []int
	for i := 0; i < len(data); i++ {
		if data[i].dice_amount == 0 {
			res[i] = []int{}
		} else {
			rollResult = dice_roll(data[i].dice_amount)
			res[i] = rollResult
		}
	}
	return res
}

func main() {
	var player, dice int
	_, err := fmt.Scan(&player, &dice)
	if err != nil {
		fmt.Println(err)
	}

	var dataPlayers []playerData

	for i := 0; i < player; i++ {
		temp := playerData{i, dice, 0}
		fmt.Println(temp)
		dataPlayers = append(dataPlayers, temp)
	}
	a := 1

	for true {
		contPlay := check_player(dataPlayers)
		if contPlay == false {
			break
		}
		fmt.Println("turn " + strconv.Itoa(a))
		roll_result := dice_roll_result(dataPlayers)
		fmt.Println("roll_result")
		fmt.Println(roll_result)
		a += 1
		if a > 2 {
			break
		}
		for i := 0; i < len(roll_result); i++ {
			fmt.Println("player " + strconv.Itoa(i))
			fmt.Println(roll_result[i])

			if len(roll_result[i]) > 0 {
				for j, _ := range roll_result[i] {
					if j == 6 {
						dataPlayers[i].point += 1
					} else if j == 1 {
						dataPlayers[i].dice_amount -= 1
						if i < len(dataPlayers)-1 {
							dataPlayers[i+1].dice_amount += 1
						} else {
							dataPlayers[0].dice_amount += 1
						}
					}
				}
			} else {
				continue
			}
		}
		fmt.Println("data after roll")
		fmt.Println(dataPlayers)
	}

	fmt.Println(dataPlayers)
}
