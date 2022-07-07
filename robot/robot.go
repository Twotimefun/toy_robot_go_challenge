package robot

import (
	// "github.com/twotimefun/toy_robot_go_challenge/commands"
	"fmt"
	"toy_robot_go_challenge/table"
)

// Creates and sustains values for Robot
type Robot struct {
	X, Y   int
	Facing string
}

// To place, accepts x, y cords, and direction facing.
// Returns Robot with new values, and error if invalid.
func Place(x int, y int, facing string) (*Robot, error) {
	if !table.InBounds(x, y) {
		return nil, fmt.Errorf("The robot cannot be out of bounds")
	}
	if !table.DirectionFacing(facing) {
		return nil, fmt.Errorf("Robot cannot face an invalid direction")
	}

	return &Robot{x, y, facing}, nil
}

// Requires current direction and location to comprehend movement
func Move(robot *Robot) *Robot {
	x, y, facing := robot.X, robot.Y, robot.Facing

	switch facing {
	case "NORTH":
		y += 1
	case "SOUTH":
		y -= 1
	case "EAST":
		x += 1
	case "WEST":
		x -= 1
	}
	if !table.InBounds(x, y) {
		fmt.Println("The robot can not leave the table")
		return robot
	}
	return &Robot{x, y, facing}
}

func Report(robot *Robot) string {
	// Returns current location and direction, as a string
	return fmt.Sprintf("%d,%d,%s", robot.X, robot.Y, robot.Facing)
}

func Left(robot *Robot) *Robot {
	var new_facing string

	switch robot.Facing {
	case "NORTH":
		new_facing = "WEST"
	case "SOUTH":
		new_facing = "EAST"
	case "EAST":
		new_facing = "NORTH"
	case "WEST":
		new_facing = "SOUTH"
	}
	return &Robot{robot.X, robot.Y, new_facing}
}

func Right(robot *Robot) *Robot {
	var new_facing string

	switch robot.Facing {
	case "NORTH":
		new_facing = "EAST"
	case "SOUTH":
		new_facing = "WEST"
	case "EAST":
		new_facing = "SOUTH"
	case "WEST":
		new_facing = "NORTH"
	}
	return &Robot{robot.X, robot.Y, new_facing}
}
