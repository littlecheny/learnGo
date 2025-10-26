package main

import (

)

type ListNode struct {
	val int;
	next *ListNode;
}

type LRUCache struct {
    m map[int]int
    head *ListNode
    capacity int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
        m: make(map[int]int, capacity),
        head: &ListNode{},
        capacity: capacity,
    }
}