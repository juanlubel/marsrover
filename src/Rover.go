package src

import (
	"fmt"
	"strconv"
	"strings"
)

type CardinalDirection int

const (
	North = iota
	East
	South
	West
)

type Coordinate struct {
	X int
	Y int
}

type Board struct {
	TotalX    int
	TotalY    int
	Obstacles []Coordinate
}

func NewBoard(totalX int, totalY int, obstacles []Coordinate) *Board {
	return &Board{TotalX: totalX, TotalY: totalY, Obstacles: obstacles}
}

type Rover struct {
	Initial  Coordinate
	Moves    []string
	Position Coordinate
	Faced    CardinalDirection
	Area     Board
}

func (r *Rover) Move(moves string) {
	splitMoves := strings.Split(moves, "")
	//	for len movements
	for _, move := range splitMoves {
		//	If movement is forward or backward increse or decrese on depends of Facing
		if CheckFacing(move) {
			r.moveFacing(move)
		}
		if CheckPositioning(move) {
			r.moveBackwardOrForward(move)
		}
	}
	//	Check type of movement
	//	If movement is Facing, change faced
}

func CheckPositioning(move string) bool {
	return move == "F" || move == "B"
}

func CheckFacing(move string) bool {
	return move == "L" || move == "R"
}

var InitialFacing CardinalDirection = North
var InitialPosition = Coordinate{X: 0, Y: 0}

func NewRover(initial Coordinate) *Rover {
	return &Rover{Initial: initial, Faced: InitialFacing, Position: initial}
}

func (r *Rover) StartsRover() *Rover {
	position := r.GetPosition()
	fmt.Println(position)
	return r
}

// The possible movements are l of r
func (r *Rover) moveFacing(move string) {
	if move == "L" {
		r.Faced = TurnLeft(r.Faced)
	}
	if move == "R" {
		r.Faced = TurnRight(r.Faced)
	}
}

func (r *Rover) moveBackwardOrForward(move string) {
	if move == "F" {
		r.moveForwardByFaced()
	}
	if move == "B" {
		r.moveBackwardByFaced()
	}
}

// Facing values N,S,E,W
func TurnLeft(currentPosition CardinalDirection) CardinalDirection {
	switch currentPosition {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	}
	return North
}

func TurnRight(currentPosition CardinalDirection) CardinalDirection {
	switch currentPosition {
	case North:
		return West
	case West:
		return South
	case South:
		return East
	case East:
		return North
	}
	return North
}

// GetPosition return x:y:f (position & facing)
func (r *Rover) GetPosition() string {
	xToString := strconv.Itoa(r.Position.X)
	yToString := strconv.Itoa(r.Position.Y)
	var facedToString string
	switch r.Faced {
	case North:
		facedToString = "N"
	case East:
		facedToString = "E"
	case West:
		facedToString = "W"
	case South:
		facedToString = "S"
	}
	return xToString + ":" + yToString + ":" + facedToString
}

func (r *Rover) moveForwardByFaced() {
	switch r.Faced {
	case North:
		r.Position.Y = r.yPlusOne()
	case East:
		r.Position.X = r.xPlusOne()
	case South:
		r.Position.Y = r.yMinusOne()
	case West:
		r.Position.X = r.xMinusOne()
	}
}

func (r *Rover) xMinusOne() int {
	targetPosition := r.Position.X - 1
	if targetPosition == 0 {
		return r.Position.X
	}
	return targetPosition
}

func (r *Rover) yMinusOne() int {
	targetPosition := r.Position.Y - 1
	if targetPosition == 0 {
		return r.Position.Y
	}
	return targetPosition
}

func (r *Rover) xPlusOne() int {
	targetPosition := r.Position.X + 1
	if targetPosition == 0 {
		return r.Position.X
	}
	return targetPosition
}

func (r *Rover) yPlusOne() int {
	targetPosition := r.Position.Y + 1
	if targetPosition == 0 {
		return r.Position.Y
	}
	return targetPosition
}

func (r *Rover) moveBackwardByFaced() {
	switch r.Faced {
	case North:
		r.Position.Y = r.yMinusOne()
	case East:
		r.Position.X = r.xMinusOne()
	case South:
		r.Position.Y = r.yPlusOne()
	case West:
		r.Position.X = r.xPlusOne()
	}
}
