package main

import "fmt"

type Mover interface {
	Move()
}

type Locker interface {
	Lock()
	Unlock()
}

type MoveLocker interface {
	Mover
	Locker
}

type bike struct {}

func (bike)Move()  {
	fmt.Println("Locking the bike")
}

func (bike)Lock()  {
	fmt.Println("Locking the bike")
}

func (bike)Unlock()  {
	fmt.Println("Unlocking the bike")
}

func main() {
	var m1 MoveLocker
	var m Mover

	m1 = bike{}
	m = m1
	//m1 = m
	b := m.(bike)
	m1 = b
}
