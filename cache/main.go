package main

import (
	"fmt"
	"log"
	"time"
)

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

type Memory struct {
	function Function
	cache    map[int]FunctionResult
}

type Function func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

func NewCache(function Function) *Memory {
	return &Memory{
		function: function,
		cache:    make(map[int]FunctionResult),
	}
}

func (memory *Memory) Get(key int) (interface{}, error) {
	result, exists := memory.cache[key]
	if !exists {
		result.value, result.err = memory.function(key)
		memory.cache[key] = result
	}
	return result.value, result.err
}

func GetFibonacci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

func main() {
	cache := NewCache(GetFibonacci)
	fibonacci := []int{42, 40, 41, 42, 38}

	for _, n := range fibonacci {
		start := time.Now()
		value, err := cache.Get(n)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%d, %s, %d\n", n, time.Since(start), value)
	}
}
