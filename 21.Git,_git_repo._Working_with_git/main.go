package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Name string
	Age  int
}

type SafeMap struct {
	mu   sync.Mutex
	data map[string]string
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]string),
	}
}

func (s *SafeMap) Set(key string, value string, wg *sync.WaitGroup) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
	wg.Done()

}
func (s *SafeMap) Get(key string, ch chan string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	ch <- s.data[key]

}
func (s *SafeMap) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}

func main() {
	safeMap := NewSafeMap()
	wg := sync.WaitGroup{}
	data := make(chan string, 7)
	key := []string{"a", "b", "c", "d", "e", "f", "g"}
	value := []string{"1", "2", "3", "4", "5", "6", "7"}

	for i := 0; i < len(key); i++ {
		wg.Add(1)
		go safeMap.Set(key[i], value[i], &wg)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeMap.Get(key[i], data)
		}(i)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeMap.Delete(key[i])
		}(i)
	}
	wg.Wait()
	for i := 0; i < len(data); i++ {
		fmt.Print(<-data, " ")
	}

}
