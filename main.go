package main

import(
	"fmt"
	"time"
	"math/rand"
)

type User struct {
	Id int
	Likes int
	Gender string
}

// About 75% of Tinder users identify as male

func genMaleUsers() []User {
	var list []User
	for i := 1; i < 751; i++ {
		var us User
		us.Id = i 
		us.Gender = "M"
		us.Likes = 0
		list = append(list, us)
	}
	return list
}

func foo(a int) {

}

func genFemUsers() []User {
	var list []User
	for i := 1; i < 251; i++ {
		var us User
		us.Id = i 
		us.Gender = "F"
		us.Likes = 0
		list = append(list, us)
	}
	return list
}

func main() {
	start := time.Now()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	ListMaleUsers := genMaleUsers()
	ListFemaleUsers := genFemUsers()
	
	for n := 1; n <= 101; n++{
		for i := range ListMaleUsers {
			foo(i)
			randomNumber := r.Float64()
			if randomNumber < 0.4 {
				randomFemaleIndex := r.Intn(len(ListFemaleUsers))
				ListFemaleUsers[randomFemaleIndex].Likes++
			}
		}
	}

	for _, user := range ListFemaleUsers {
		fmt.Printf("User %d (F) - Likes Recebidos: %d\n", user.Id, user.Likes)
	}

	for n := 1; n <= 101; n++ {
		for i := range ListFemaleUsers {
			foo(i)
			randomNumber := r.Float64()
			if randomNumber < 0.07 {
				randomMaleIndex := r.Intn(len(ListMaleUsers))
				ListMaleUsers[randomMaleIndex].Likes++
			}
		}
	}

	for _, user := range ListMaleUsers {
		fmt.Printf("User %d (M) - Likes Recebidos: %d\n", user.Id, user.Likes)
	}

	fmt.Println(time.Since(start))
}