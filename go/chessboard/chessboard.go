package chessboard

// Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Chessboard contains a map of eight Ranks, accessed with values from 1 to 8
type Chessboard map[int]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank int) (ret int) {
	count := 0
	for _, hasOccupied := range cb[rank] {
		if hasOccupied {
			count++
		}
	}
	return count
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) (ret int) {
	count := 0
	for _, rank := range cb {
		if (file <= len(rank)) && rank[file-1] {
			count++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) (ret int) {
	count := 0
	for _, rank := range cb {
		count += len(rank)
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) (ret int) {
	count := 0
	for index := range cb {
		count += CountInRank(cb, index)
	}
	return count
}
