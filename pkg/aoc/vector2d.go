package aoc

// Vector2D represents a 3d vector
type Vector2D struct {
	x, y int
}

// MakeVector2D creates a Vector2D for the specified x, y
func MakeVector2D(x, y int) Vector2D {
	return Vector2D{x, y}
}

// X retuns the x value for the vector
func (v Vector2D) X() int {
	return v.x
}

// Y retuns the x value for the vector
func (v Vector2D) Y() int {
	return v.y
}

// Add adds another vector to the vector
func (v Vector2D) Add(other Vector2D) Vector2D {
	return Vector2D{v.x + other.x, v.y + other.y}
}

// ManhattanLength returns the Manhattan distance from {0,0}
func (v Vector2D) ManhattanLength() int {
	return AbsInt(v.x) + AbsInt(v.y)
}
