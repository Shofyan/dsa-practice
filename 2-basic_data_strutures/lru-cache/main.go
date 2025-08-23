package main

import (
	"container/list"
	"fmt"
)

// Node merepresentasikan elemen dalam cache, yang akan disimpan di linked list.
type Node struct {
	key   string
	value int
}

// LRUCache adalah struktur utama dari cache kita.
type LRUCache struct {
	capacity int
	cache    map[string]*list.Element // Hash map untuk akses cepat O(1)
	list     *list.List               // Doubly linked list untuk melacak urutan penggunaan
}

// NewLRUCache adalah constructor untuk membuat instance LRUCache baru.
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

// Get mengambil nilai dari cache berdasarkan key.
// Jika key ditemukan, elemen tersebut akan dipindahkan ke depan list.
func (lru *LRUCache) Get(key string) (int, bool) {
	if elem, ok := lru.cache[key]; ok {
		// Pindahkan elemen yang diakses ke depan list (paling baru digunakan)
		lru.list.MoveToFront(elem)
		// Kembalikan nilainya
		return elem.Value.(*Node).value, true
	}
	// Jika tidak ditemukan, kembalikan nilai default dan false
	return -1, false
}

// Put menambahkan atau memperbarui key-value pair ke dalam cache.
func (lru *LRUCache) Put(key string, value int) {
	// Jika key sudah ada, update nilainya dan pindahkan ke depan
	if elem, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(elem)
		elem.Value.(*Node).value = value
		return
	}

	// Jika cache sudah penuh, buang elemen yang paling lama tidak digunakan (di akhir list)
	if lru.list.Len() == lru.capacity {
		// Dapatkan elemen terakhir
		lastElem := lru.list.Back()
		if lastElem != nil {
			// Hapus dari map dan list
			delete(lru.cache, lastElem.Value.(*Node).key)
			lru.list.Remove(lastElem)
		}
	}

	// Buat node baru dan tambahkan ke depan list serta ke map
	newNode := &Node{key: key, value: value}
	elem := lru.list.PushFront(newNode)
	lru.cache[key] = elem
}

// Display untuk menampilkan isi cache (dari yang terbaru ke terlama).
func (lru *LRUCache) Display() {
	fmt.Print("Cache content (most-recent -> least-recent): ")
	for elem := lru.list.Front(); elem != nil; elem = elem.Next() {
		node := elem.Value.(*Node)
		fmt.Printf("[%s: %d] ", node.key, node.value)
	}
	fmt.Println()
}

func main() {
	// Buat cache dengan kapasitas 3
	lruCache := NewLRUCache(3)

	fmt.Println("Menambahkan 'a':1, 'b':2, 'c':3")
	lruCache.Put("a", 1)
	lruCache.Put("b", 2)
	lruCache.Put("c", 3)
	lruCache.Display() // Output: [c: 3] [b: 2] [a: 1]

	fmt.Println("\nMengakses 'a'...")
	lruCache.Get("a")
	lruCache.Display() // Output: [a: 1] [c: 3] [b: 2] -> 'a' pindah ke depan

	fmt.Println("\nMenambahkan 'd':4 (cache penuh, 'b' akan terbuang)")
	lruCache.Put("d", 4)
	lruCache.Display() // Output: [d: 4] [a: 1] [c: 3] -> 'b' (paling lama tidak digunakan) hilang

	fmt.Println("\nMengupdate 'c' menjadi 30")
	lruCache.Put("c", 30)
	lruCache.Display() // Output: [c: 30] [d: 4] [a: 1] -> 'c' diupdate dan pindah ke depan
}