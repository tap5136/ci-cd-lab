package main

import "testing"

func TestBaseball(t *testing.T) {
    team := BestBaseballTeam()
    if team != "Boston Red Sox" {
        t.Errorf("Nope, not the %s, the best team in baseball is clearly the Boston Red Sox", team)
    }
}

func TestSum(t *testing.T) {
    x := 2
    y := 3
    sum := Sum(x, y)
    if sum != x + y {
        t.Errorf("Sum(%d, %d) should return %d, not return %d", x, y, x + y, sum)
    }
}