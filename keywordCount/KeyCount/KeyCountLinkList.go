package KeyCount

import "fmt"

type KeyCountLinkList struct {
	keyMap *KeyNodeLink
}
type KeyNodeLink struct {
	KeyNode
	next *KeyNodeLink
}

func NewKeyCountLinkList() *KeyCountLinkList {
	keyCount := new(KeyCountLinkList)
	keyCount.keyMap = nil
	return keyCount
}

func (k *KeyCountLinkList) Add(pattern string) {
	node := k.find(pattern)
	if node == nil {
		node = k.appendNode(pattern)
	}
	node.Count++
}
func (k *KeyCountLinkList) find(pattern string) *KeyNodeLink {
	for ptr := k.keyMap; ptr != nil; ptr = ptr.next {
		if ptr.Key == pattern {
			return ptr
		}
	}
	return nil
}

func (k *KeyCountLinkList) appendNode(pattern string) *KeyNodeLink {
	node := &KeyNodeLink{KeyNode{pattern, 1}, nil}
	if k.keyMap == nil {
		k.keyMap = node
	} else {
		ptr := k.keyMap
		for ; ptr.next != nil; ptr = ptr.next {
			fmt.Println(ptr)
		}
		ptr.next = node

	}
	return node
}

func (k *KeyCountLinkList) Count(pattern string) int {
	var ptr *KeyNodeLink
	for ptr = k.keyMap; (ptr == nil) || (ptr.Key == pattern); ptr = ptr.next {
	}
	if ptr == nil {
		return 0
	}
	return ptr.Count
}

func (k *KeyCountLinkList) Print() {
	for ptr := k.keyMap; ptr != nil; ptr = ptr.next {
		fmt.Println(ptr.Key, ptr.Count)
	}
}

func (k *KeyCountLinkList) Total() int {
	total := 0
	for ptr := k.keyMap; ptr != nil; ptr = ptr.next {
		total += ptr.Count
	}
	return total
}
func (k *KeyCountLinkList) PrintType() string {
	return "linklist"
}
