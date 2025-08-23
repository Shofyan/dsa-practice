package main

import "fmt"

// Node merepresentasikan satu simpul dalam pohon biner.
// Ini menyimpan kunci integer dan pointer ke anak kiri dan kanannya.
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Tree merepresentasikan pohon biner itu sendiri.
// Ini hanya berisi pointer ke simpul root.
type Tree struct {
	Root *Node
}

// Insert menambahkan simpul baru dengan kunci yang diberikan ke pohon.
// Ini mempertahankan properti pohon pencarian biner:
// - Kunci yang lebih kecil dari induk pergi ke kiri.
// - Kunci yang lebih besar atau sama dengan induk pergi ke kanan.
func (t *Tree) Insert(key int) {
	// Jika pohon kosong, simpul baru menjadi root.
	if t.Root == nil {
		t.Root = &Node{Key: key}
	} else {
		// Jika tidak, kita memulai proses penyisipan dari root.
		t.Root.insert(key)
	}
}

// insert adalah metode pembantu pribadi yang secara rekursif menemukan posisi yang benar
// untuk kunci baru dan menyisipkan simpul baru di sana.
func (n *Node) insert(key int) {
	// Jika kunci baru lebih kecil dari kunci simpul saat ini, pergi ke kiri.
	if key < n.Key {
		if n.Left == nil {
			// Jika tidak ada anak kiri, buat satu.
			n.Left = &Node{Key: key}
		} else {
			// Jika tidak, lanjutkan secara rekursif ke bawah subpohon kiri.
			n.Left.insert(key)
		}
		// Jika kunci baru lebih besar atau sama dengan kunci simpul saat ini, pergi ke kanan.
	} else {
		if n.Right == nil {
			// Jika tidak ada anak kanan, buat satu.
			n.Right = &Node{Key: key}
		} else {
			// Jika tidak, lanjutkan secara rekursif ke bawah subpohon kanan.
			n.Right.insert(key)
		}
	}
}

// InOrder melakukan traversal in-order dari pohon (Kiri, Root, Kanan).
// Ini akan mencetak kunci simpul dalam urutan menaik yang diurutkan.
func (t *Tree) InOrder() {
	fmt.Println("In-order traversal:")
	inOrder(t.Root)
	fmt.Println() // untuk baris baru setelah output
}

// inOrder adalah pembantu rekursif untuk traversal in-order.
func inOrder(node *Node) {
	if node != nil {
		inOrder(node.Left)
		fmt.Printf("%d ", node.Key)
		inOrder(node.Right)
	}
}

// PreOrder melakukan traversal pre-order dari pohon (Root, Kiri, Kanan).
func (t *Tree) PreOrder() {
	fmt.Println("Pre-order traversal:")
	preOrder(t.Root)
	fmt.Println()
}

// preOrder adalah pembantu rekursif untuk traversal pre-order.
func preOrder(node *Node) {
	if node != nil {
		fmt.Printf("%d ", node.Key)
		preOrder(node.Left)
		preOrder(node.Right)
	}
}

// PostOrder melakukan traversal post-order dari pohon (Kiri, Kanan, Root).
func (t *Tree) PostOrder() {
	fmt.Println("Post-order traversal:")
	postOrder(t.Root)
	fmt.Println()
}

// postOrder adalah pembantu rekursif untuk traversal post-order.
func postOrder(node *Node) {
	if node != nil {
		postOrder(node.Left)
		postOrder(node.Right)
		fmt.Printf("%d ", node.Key)
	}
}

// Display menampilkan struktur pohon secara visual.
func (t *Tree) Display() {
	fmt.Println("Struktur Pohon:")
	display(t.Root, 0, "  ")
}

// display adalah fungsi pembantu rekursif untuk mencetak pohon.
func display(node *Node, level int, prefix string) {
	if node == nil {
		return
	}
	display(node.Right, level+1, "┌──")
	for i := 0; i < level; i++ {
		fmt.Print("│   ")
	}
	fmt.Println(prefix, node.Key)
	display(node.Left, level+1, "└──")
}

func main() {
	// Buat instance pohon baru.
	tree := &Tree{}

	// Sisipkan serangkaian nilai.
	// Struktur pohon akan dibentuk berdasarkan urutan penyisipan.
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(18)

	// Tampilkan struktur pohon
	tree.Display()

	// Cetak pohon menggunakan tiga metode traversal yang berbeda.
	tree.InOrder()   // Output yang diharapkan: 3 5 7 10 12 15 18
	tree.PreOrder()  // Output yang diharapkan: 10 5 3 7 15 12 18
	tree.PostOrder() // Output yang diharapkan: 3 7 5 12 18 15 10
}
