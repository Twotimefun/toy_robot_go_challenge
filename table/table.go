package table

type MAX_MOVE struct {
	MAX_X int
	MAX_Y int
}

// 0 to 4, so 5 by extension
const X, Y int = 4, 4

func Table() *MAX_MOVE {
	return &MAX_MOVE{X, Y}
}

// Checks borders of table
func InBounds(x int, y int) bool {
	table := Table()
	if x > table.MAX_X || x < 0 || y > table.MAX_Y || y < 0 {
		return false
	}
	return true
}

// Checks if the direction provided is within scope
func DirectionFacing(facing string) bool {
	switch facing {
	case
		"NORTH",
		"EAST",
		"SOUTH",
		"WEST":
		return true
	}
	return false
}
