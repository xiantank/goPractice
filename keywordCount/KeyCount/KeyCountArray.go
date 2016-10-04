package KeyCount

import "fmt"

type KeyCountArray struct {
	keyMap []KeyNode
}

func NewKeyCountArray() *KeyCountArray {
	keyCount := new(KeyCountArray)
	keyCount.keyMap = make([]KeyNode, 0, 1024)
	return keyCount
}

func (k *KeyCountArray) Add(pattern string) {
	index := k.indexOf(pattern)
	if index == -1 {
		index = len(k.keyMap)
		k.keyMap = append(k.keyMap, KeyNode{pattern, 0})
	}
	k.keyMap[index].Count++
}

func (k *KeyCountArray) indexOf(pattern string) int {
	for index, keyNode := range k.keyMap {
		if keyNode.Key == pattern {
			return index
		}
	}
	return -1
}

func (k *KeyCountArray) Count(pattern string) int {
	index := k.indexOf(pattern)
	if index == -1 {
		return 0
	}
	return k.keyMap[index].Count
}

func (k *KeyCountArray) Print() {
	for _, keyNode := range k.keyMap {
		fmt.Println(keyNode.Key, keyNode.Count)
	}
}
func (k *KeyCountArray) Total() int {
	total := 0
	for _, keyNode := range k.keyMap {
		fmt.Println(keyNode.Key, keyNode.Count)
		total += keyNode.Count
	}
	return total
}
func (k *KeyCountArray) PrintType() string {
	return "array"
}
