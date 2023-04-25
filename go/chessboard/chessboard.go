package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	column, ok := cb[file]
	if !ok {
		return 0
	}

	var count int
	for _, square := range column {
		if square {
			count += 1
		}
	}

	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}

	var count int
	for _, column := range cb {
		if column[rank-1] {
			count += 1
		}
	}

	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var count int
	for _, column := range cb {
		count += len(column)
	}

	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var count int
	for fileName := range cb {
		count += CountInFile(cb, fileName)
	}

	return count
}
