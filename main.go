package main

import "fmt"

type merger interface{ merge() }

type user struct {
	names []string
	name  string
}

func (u *user) merge() {
	u.name = u.names[0]
}

func InterfaceTest(mergeItem merger, mergeItem2 merger) {
	mergeItem.merge()
	mergeItem2.merge()
}

func main() {
	var names []string
	names = append(names, "yuya", "okumura")
	user := &user{names: names}
	people := &people{peopleNames: names}
	fmt.Printf("User: %s", user)
	fmt.Printf("People:%s", people)
	InterfaceTest(user, people)
}

type people struct {
	peopleNames []string
	peopleName string
}

func(p *people) merge() {
	p.peopleName = "ゴンザレス"
}