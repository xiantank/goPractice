package KeyCount

import "fmt"

type KeyNode struct {
	Key   string
	Count int
}

type KeyCount struct {
	keyMap []KeyNode
}

func NewKeyCount() *KeyCount {
	keyCount := new(KeyCount)
	keyCount.keyMap = make([]KeyNode, 0, 1024)
	return keyCount
}

func (k *KeyCount) Add(pattern string) {
	index := k.indexOf(pattern)
	if index == -1 {
		index = len(k.keyMap)
		append(k.keyMap, &KeyNode{pattern, 0})
	}
	k.keyMap[index].Count++
}

func (k *KeyCount) indexOf(pattern string) int {
	for index, keyNode := range k {
		if keyNode.Key == pattern {
			return index
		}
	}
	return -1
}

func (k *KeyCount) Count(pattern string) int {
	index := k.indexOf(pattern)
	if index == -1 {
		return 0
	}
	return k.keyMap[index]
}

func (k *KeyCount) Print() {
	for keyNode := range k.keyMap {
		fmt.Println(keyNode.Key, keyNode.Count)
	}
}
func (k *KeyCount) Total() int {
	total := 0
	for keyNode := range k.keyMap {
		fmt.Println(keyNode.Key, keyNode.Count)
		total += keyNode.Count
	}
	return total
}
