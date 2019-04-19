package lib

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	// User ...
	User struct {
		ID   int
		Name string
	}

	// Users ...
	Users []*User

	// Group ...
	Group struct {
		ID      int
		Players []*Player
	}

	// Groups ...
	Groups []*Group

	// Player ...
	Player struct {
		ID   int
		User *User
	}
)

// Robin ...
func Robin() {
	usersCnt := 16
	byGroup := 4

	u1 := &User{ID: 1, Name: "aaa"}
	u2 := &User{ID: 2, Name: "bbb"}
	u3 := &User{ID: 3, Name: "ccc"}
	u4 := &User{ID: 4, Name: "ddd"}
	u5 := &User{ID: 5, Name: "eee"}
	u6 := &User{ID: 6, Name: "fff"}
	u7 := &User{ID: 7, Name: "ggg"}
	u8 := &User{ID: 8, Name: "hhh"}
	u9 := &User{ID: 9, Name: "iii"}
	u10 := &User{ID: 10, Name: "jjj"}
	u11 := &User{ID: 11, Name: "kkk"}
	u12 := &User{ID: 12, Name: "lll"}

	var users Users = []*User{u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12}

	fmt.Println(usersCnt, byGroup, users)

	users.print()
	users.shuffle()

	grpCnt := ((len(users) - 1) / byGroup) + 1
	var groups Groups = make([]*Group, grpCnt)
	for i := 0; i < grpCnt; i++ {
		groups[i] = &Group{
			ID:      i + 1,
			Players: make([]*Player, 0, byGroup),
		}
	}
	for i, u := range users {
		remain := i % grpCnt
		p := &Player{ID: i + 1, User: u}
		groups[remain].Players = append(groups[remain].Players, p)
	}
	groups.print()
	fmt.Println(groups, grpCnt)

}

func (us Users) shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := range us {
		j := rand.Intn(i + 1)
		us[i], us[j] = us[j], us[i]
	}
}

func (us Users) print() {
	for _, u := range us {
		fmt.Println(u.ID, u.Name)
	}
}

func (gs Groups) print() {
	for _, g := range gs {
		fmt.Println("GroupID: ", g.ID)
		fmt.Println("")
		for _, p := range g.Players {
			p.print()
			fmt.Println("")
		}
		fmt.Println("")
	}
}

func (p *Player) print() {
	fmt.Println("PlayerID: ", p.ID)
	fmt.Println("UserID: ", p.User.ID)
	fmt.Println("UserName: ", p.User.Name)
}
