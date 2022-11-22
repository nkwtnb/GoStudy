package main

import (
	"fmt"
	"main/src"
	"time"
)

func main() {
    user := src.User{Name: "Naoki", Birth: time.Date(1990, 01, 01, 0, 0, 0, 0, time.Local)}
    user.AddAgeByRef()
    fmt.Println(user)
    // {Naoki 1991-01-01 00:00:00 +0900 JST}

    user2 := src.User{Name: "Naoki", Birth: time.Date(1990, 01, 01, 0, 0, 0, 0, time.Local)}
    user2.AddAgeByVal()
    fmt.Println(user2)
    // {Naoki 1990-01-01 00:00:00 +0900 JST}
}