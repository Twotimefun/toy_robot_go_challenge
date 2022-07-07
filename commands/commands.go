package commands

import (
	// "github.com/twotimefun/toy_robot_go_challenge/robot"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"toy_robot_go_challenge/robot"
)

func Start() {
	var robot_ *robot.Robot
	var err error
	fmt.Println("Welcome to the Toy Robot Simulator!", "\n", "Use a valid command", "\n", "For example: PLACE 0,0,NORTH,", "\n", "MOVE, LEFT, RIGHT, REPORT, or QUIT")
	inputReader := bufio.NewReader(os.Stdin)
	for {
		untrimmedCommand, _ := inputReader.ReadString('\n')
		trimmed_command := strings.TrimSpace(untrimmedCommand)
		robot_, err = commandInput(trimmed_command, robot_)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func Quit() {
	fmt.Println("The program is now exiting.")
	os.Exit(0)
}

func commandInput(command string, robot_ *robot.Robot) (*robot.Robot, error) {
	if regexp.MustCompile("QUIT").MatchString(command) {
		Quit()
	}

	if regexp.MustCompile("PLACE [0-9]+,[0-9]+,[A-Z]+").MatchString(command) {
		return placeReader(command, robot_)
	}

	if robot_ == nil {
		return nil, fmt.Errorf("you must first place the robot before you can run any other commands")
	}

	switch {
	case command == "MOVE":
		return robot.Move(robot_), nil
	case command == "LEFT":
		return robot.Left(robot_), nil
	case command == "RIGHT":
		return robot.Right(robot_), nil
	case command == "REPORT":
		fmt.Println(robot.Report(robot_))
		return robot_, nil
	case command == "QUIT":
		Quit()
	case command == "" || command == "\n":
		return robot_, nil
	}
	return robot_, fmt.Errorf("command: %s not recognized", command)
}

func placeReader(command string, robot_ *robot.Robot) (*robot.Robot, error) {
	plc_cmd := strings.Fields(command)[1]
	// Seperates by ',' to reader and translate input
	place_command := strings.Split(plc_cmd, ",")
	// Converts the first argument to an integer
	x, err := strconv.Atoi(place_command[0])

	if err != nil {
		panic(err)
	}

	// Converts the second argument to an integer
	y, err := strconv.Atoi(place_command[1])
	if err != nil {
		panic(err)
	}

	// Converts the third argument to a string
	facing := place_command[2]

	// Places the robot
	robot_placement, err := robot.Place(x, y, facing)
	if err != nil {
		return robot_, fmt.Errorf("invalid PLACE command")
	}
	return robot_placement, nil
}
