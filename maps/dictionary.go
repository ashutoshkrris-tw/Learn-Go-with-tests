package maps

type Dictionary map[string]int

func (d Dictionary) Search(letter string) int {
	return d[letter]
}
