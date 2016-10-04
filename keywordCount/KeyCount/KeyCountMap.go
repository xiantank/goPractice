package KeyCount

import "fmt"

type KeyNode struct {
	Key   string
	Count int
}

type KeyCount struct {
	keyMap map[string]int
}

func NewKeyCount() *KeyCount {
	keyCount := new(KeyCount)
	keyCount.keyMap = make(map[string]int)
	return keyCount
}

func (k *KeyCount) Add(pattern string) {
	count, ok := k.keyMap[pattern]
	if !ok {
		count = 0
	}
	count++
	k.keyMap[pattern] = count
}

func (k *KeyCount) Count(pattern string) int {
	return k.keyMap[pattern]
}

func (k *KeyCount) Print() {
	for key, count := range k.keyMap {
		fmt.Println(key, count)
	}
}
func (k *KeyCount) Total() int {
	total := 0
	for _, value := range k.keyMap {
		total += value
	}
	return total
}
