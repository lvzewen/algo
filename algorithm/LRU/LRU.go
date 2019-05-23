package main

import "fmt"

func main()  {
	lrc := LRUCache{
		Capacity: 2,
		Ordered: []int{},
		Value: make(map[int]int),
	}
	lrc.put(1, 3)
	lrc.put(2, 2)
	fmt.Printf("1lrc %+v, lrc value %+v, lrc ordered %+v", lrc, lrc.Value, lrc.Ordered)
	lrc.put(3, 3)
	fmt.Printf("2lrc %+v, lrc value %+v, lrc ordered %+v", lrc, lrc.Value, lrc.Ordered)

	lrc.get(2)
	fmt.Printf("3lrc %+v, lrc value %+v, lrc ordered %+v", lrc, lrc.Value, lrc.Ordered)
}

// LRUCache ...
type LRUCache struct {
	Ordered  []int
	Value    map[int]int
	Capacity   int
}

func (lc *LRUCache) get(key int) int {
	lc.moveToEnd(key)
	return lc.Value[key]
}

func (lc *LRUCache) put(key, value int) {
	lc.moveToEnd(key)
	
	lc.Value[key] = value
}

func (lc *LRUCache) moveToEnd(key int) {
	index := intContains(lc.Ordered, key)
	if index > -1 {
		tmp := lc.Ordered[index:index+1]
		lc.Ordered = append(lc.Ordered[:index], lc.Ordered[index+1:]...)
		lc.Ordered = append(tmp, lc.Ordered...)
	} else if len(lc.Ordered) >= lc.Capacity {
		lc.Ordered = lc.Ordered[0:lc.Capacity-1]
		lc.Ordered = append([]int{key}, lc.Ordered...)
	} else {
		lc.Ordered = append([]int{key}, lc.Ordered...)
	}
}

func intContains (arr []int, val int) (index int) {
	index = -1
	for i := 0; i < len(arr); i++ {
		if val == arr[i] {
			index = i
		}
	}
	return index
}

