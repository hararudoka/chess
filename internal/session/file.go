package session

// File is a line with letter name
type File int

// Rank is a line with number name
type Rank int

func ByteToFile(x byte) File {
	m := map[byte]File{
		'A': 0,
		'B': 1,
		'C': 2,
		'D': 3,
		'E': 4,
		'F': 5,
		'G': 6,
		'H': 7,

		'a': 0,
		'b': 1,
		'c': 2,
		'd': 3,
		'e': 4,
		'f': 5,
		'g': 6,
		'h': 7,
	}
	if _, ok := m[x]; !ok {
		return -1
	}
	return m[x]
}

func ByteToRank(y byte) Rank {
	m := map[byte]Rank{
		'1': 7,
		'2': 6,
		'3': 5,
		'4': 4,
		'5': 3,
		'6': 2,
		'7': 1,
		'8': 0,
	}
	if _, ok := m[y]; !ok {
		return -1
	}
	return m[y]
}
