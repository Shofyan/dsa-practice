
# Interview Problems

This document provides a detailed explanation of common software engineering interview problems, their solutions, and the reasoning behind the chosen approaches.

---

## 1. Two Sum

**Problem Statement:**

Given an array of integers `nums` and an integer `target`, return the indices of the two numbers such that they add up to `target`. You may assume that each input would have exactly one solution, and you may not use the same element twice.

**Solution Approach:**

We can use a hash map (or a dictionary in Python, or a map in Go) to solve this problem in a single pass. The hash map will store the numbers we have seen so far as keys, and their indices as values.

**Detailed Explanation:**

1.  Create an empty hash map.
2.  Iterate through the `nums` array. For each element `num` at index `i`:
    a.  Calculate the `complement` which is `target - num`.
    b.  Check if the `complement` exists in the hash map.
        i.  If it exists, we have found our pair. Return the index of the `complement` from the hash map and the current index `i`.
        ii. If it does not exist, add the current `num` and its index `i` to the hash map.
3.  If we finish the loop without finding a pair, it means there is no solution (though the problem statement guarantees one).

**Why this approach?**

*   **Time Complexity:** O(n). We iterate through the array once. Hash map lookups and insertions take, on average, O(1) time.
*   **Space Complexity:** O(n). In the worst case, we might store all the elements of the array in the hash map.

This is a significant improvement over the brute-force approach of checking every pair of numbers, which would have a time complexity of O(n^2).

---

## 2. Reverse a String

**Problem Statement:**

Write a function that reverses a string. The input string is given as an array of characters `s`.

**Solution Approach:**

We can use the **two-pointers** technique to solve this problem in-place.

**Detailed Explanation:**

1.  Initialize two pointers: `left` at the beginning of the array (index 0) and `right` at the end of the array (index `len(s) - 1`).
2.  While `left` is less than `right`:
    a.  Swap the characters at the `left` and `right` pointers.
    b.  Increment `left` and decrement `right`.
3.  The loop terminates when the pointers meet or cross, and the string is reversed.

**Why this approach?**

*   **Time Complexity:** O(n). We iterate through about half of the string.
*   **Space Complexity:** O(1). The reversal is done in-place, so we don't use any extra space.

---

## 3. Remove Duplicates from Sorted Array

**Problem Statement:**

Given a sorted array `nums`, remove the duplicates **in-place** such that each element appears only once and return the new length of the array.

**Solution Approach:**

We can use a two-pointer approach. One pointer will iterate through the array, and the other will keep track of the position for the next unique element.

**Detailed Explanation:**

1.  Initialize a pointer `insert_index` to 1. This is where the next unique element will be placed.
2.  Iterate through the array with a second pointer `i` starting from the second element (index 1).
3.  For each element `nums[i]`, compare it with the previous element `nums[i-1]`.
4.  If `nums[i]` is different from `nums[i-1]`, it's a unique element. Copy it to `nums[insert_index]` and increment `insert_index`.
5.  After the loop, `insert_index` will be the new length of the array with unique elements.

**Why this approach?**

*   **Time Complexity:** O(n). We iterate through the array once.
*   **Space Complexity:** O(1). The operation is done in-place.

---

## 4. Reverse a Linked List

**Problem Statement:**

Given the `head` of a singly linked list, reverse the list, and return the new head.

**Solution Approach:**

We can reverse the list iteratively. We'll need to keep track of the `previous`, `current`, and `next` nodes as we traverse the list.

**Detailed Explanation:**

1.  Initialize three pointers:
    *   `prev` to `nil`
    *   `current` to `head`
    *   `next_node` to `nil`
2.  Iterate while `current` is not `nil`:
    a.  Store the `current.next` in `next_node`.
    b.  Set `current.next` to `prev`. This is the actual reversal step.
    c.  Move `prev` to `current`.
    d.  Move `current` to `next_node`.
3.  When the loop finishes, `prev` will be pointing to the new head of the reversed list. Return `prev`.

**Why this approach?**

*   **Time Complexity:** O(n). We traverse the list once.
*   **Space Complexity:** O(1). We only use a few pointers, regardless of the list's size.

---

## 5. Implement Stack using Queues

**Problem Statement:**

Implement a Last-In-First-Out (LIFO) stack using only two queues. The implemented stack should support all the functions of a normal stack (`push`, `top`, `pop`, and `empty`).

**Solution Approach:**

We can use one queue for the main storage and a second queue as a helper to maintain the LIFO order.

**Detailed Explanation:**

*   **Push:**
    1.  Enqueue the new element to `queue2`.
    2.  Dequeue all elements from `queue1` and enqueue them to `queue2`.
    3.  Swap the names of `queue1` and `queue2`.

*   **Pop:**
    1.  Dequeue from `queue1`.

*   **Top:**
    1.  Peek at the front of `queue1`.

*   **Empty:**
    1.  Check if `queue1` is empty.

**Why this approach?**

This approach ensures that the most recently added element is always at the front of `queue1`, mimicking the behavior of a stack. The `push` operation is expensive (O(n)), while `pop` and `top` are cheap (O(1)). An alternative is to make `pop` expensive and `push` cheap. This problem tests the understanding of how to manipulate data structures to achieve a desired behavior.

---

## 6. Implement LRU Cache

**Problem Statement:**

Design a data structure that follows the constraints of a **Least Recently Used (LRU) cache**. Implement a `LRUCache` class that supports the following operations: `get` and `put`.

*   `get(key)`: Get the value of the key if the key exists in the cache, otherwise return -1.
*   `put(key, value)`: Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the `capacity` from this operation, evict the least recently used key.

**Solution Approach:**

The key to an efficient LRU cache is to have fast lookups and fast updates to the order of items. We can achieve this by combining a **hash map** and a **doubly linked list**.

*   **Hash Map:** Stores the keys and a reference (pointer) to the corresponding node in the doubly linked list. This gives us O(1) average time complexity for lookups.
*   **Doubly Linked List:** Stores the values in order of usage. The most recently used item will be at the head of the list, and the least recently used item will be at the tail.

**Detailed Explanation:**

1.  **Initialization:**
    *   Create a hash map.
    *   Create a doubly linked list with a `head` and `tail` sentinel node.
    *   Store the `capacity`.

2.  **`get(key)`:**
    *   Check if the `key` is in the hash map.
        *   If not, return -1.
        *   If it is, get the node from the hash map. Move this node to the front of the linked list (to mark it as most recently used) and return its value.

3.  **`put(key, value)`:**
    *   Check if the `key` is in the hash map.
        *   If it is, update the value of the corresponding node and move it to the front of the list.
        *   If not:
            *   Create a new node with the `key` and `value`.
            *   Add the new node to the front of the list.
            *   Add the `key` and a reference to the new node to the hash map.
            *   If the size of the cache now exceeds the `capacity`:
                *   Get the node at the tail of the list (the least recently used).
                *   Remove it from the linked list.
                *   Remove its key from the hash map.

**Why this approach?**

*   **Time Complexity:** O(1) for both `get` and `put` on average.
*   **Space Complexity:** O(capacity), as we store up to `capacity` key-value pairs.
