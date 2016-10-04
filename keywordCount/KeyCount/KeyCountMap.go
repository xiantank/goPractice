package KeyCount

import "fmt"

type KeyCountMap struct {
	keyMap map[string]int
}

func NewKeyCountMap() *KeyCountMap {
	keyCount := new(KeyCountMap)
	keyCount.keyMap = make(map[string]int)
	return keyCount
}

func (k *KeyCountMap) Add(pattern string) {
	count, ok := k.keyMap[pattern]
	if !ok {
		count = 0
	}
	count++
	k.keyMap[pattern] = count
}

func (k *KeyCountMap) Count(pattern string) int {
	return k.keyMap[pattern]
}

func (k *KeyCountMap) Print() {
	for key, count := range k.keyMap {
		fmt.Println(key, count)
	}
}
func (k *KeyCountMap) Total() int {
	total := 0
	for _, value := range k.keyMap {
		total += value
	}
	return total
}
func (k *KeyCountMap) PrintType() string {
	return "map"
}
