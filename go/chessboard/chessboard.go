package chessboard

// Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Chessboard contains a map of eight Ranks, accessed with values from 1 to 8
type Chessboard map[int]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank int) (ret int) {
	for _, hasOccupied := range cb[rank] {
		if hasOccupied {
			ret++
		}
	}
	return
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) (ret int) {
	for _, rank := range cb {
		if (file <= len(rank)) && rank[file-1] {
			ret++
		}
	}
	return
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) (ret int) {
	for _, rank := range cb {
		ret += len(rank)
	}
	return
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) (ret int) {
	for index := range cb {
		ret += CountInRank(cb, index)
	}
	return
}
