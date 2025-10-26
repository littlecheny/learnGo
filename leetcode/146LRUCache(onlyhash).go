package main

import (
	"math"
)

type LRUCache struct {
	pid      map[int][]int
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		pid:      make(map[int][]int, capacity),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.pid[key]; ok {
		for k := range this.pid {
			if k != key {
				this.pid[k][1]++
			}
		}
		this.pid[key][1] = 0
		return val[0]
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.pid[key]; ok {
		this.pid[key][0] = value
		for k := range this.pid {
			if k != key {
				this.pid[k][1]++
			}
		}
		this.pid[key][1] = 0
	} else if len(this.pid) < this.capacity {
		this.pid[key] = []int{value, 0}
		for k := range this.pid {
			if k != key {
				this.pid[k][1]++
			}
		}
	} else {
		index := 0
		max := math.MinInt
		for i := range this.pid {
			if this.pid[i][1] > max {
				max = this.pid[i][1]
				index = i
			}
		}
		delete(this.pid, index)
		for k := range this.pid {
			if k != key {
				this.pid[k][1]++
			}
		}
		this.pid[key] = []int{value, 0}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
/*
func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
}*/
