# DSA - Arrays

Practice problems solved in this folder, with the problem statement and a link to the Go solution.

## Index

1. [Reverse an Array](#1-reverse-an-array)
2. [Reverse Array in Groups of K](#2-reverse-array-in-groups-of-k)
3. [Rotate Array by D (Counter Clock Wise -> Left)](#3-rotate-array-by-d-left)
4. [Move Zeros to End](#4-move-zeros-to-end)
5. [Remove Duplicates from Sorted Array](#5-remove-duplicates-from-sorted-array)
6. [Maximum Consecutive Ones/Zeros](#6-maximum-consecutive-oneszeros)
7. [Second Largest Element](#7-second-largest-element)
8. [Third Largest Element](#8-third-largest-element)
9. [Maximum Product of a Triplet](#9-maximum-product-of-a-triplet)
10. [Maximum Product Subarray](#10-maximum-product-subarray)

---

### 1. Reverse an Array

Given an array `arr[]`, reverse its elements in place.

**Example**
```
Input:  [1, 2, 3, 4, 5]
Output: [5, 4, 3, 2, 1]
```

Solution: [`reverse.go`](reverse.go) — `ReverseArray`

---

### 2. Reverse Array in Groups of K

Given an array `arr[]` and an integer `k`, reverse every consecutive group of `k` elements in place. If the last group has fewer than `k` elements, reverse it as it is.

**Example**
```
Input:  arr = [1, 2, 3, 4, 5, 6, 7, 8], k = 3
Output: [3, 2, 1, 6, 5, 4, 8, 7]
```

Solution: [`reverse-group.go`](reverse-group.go) — `ReverseArrayGroupWise`

---

### 3. Rotate Array by D (Left)

Given an array of integers `arr[]` of size `n`, rotate the array elements to the left by `d` positions.

**Example**
```
Input:  arr = [1, 2, 3, 4, 5, 6], d = 2
Output: [3, 4, 5, 6, 1, 2]
```

Solution: [`rotate-array.go`](rotate-array.go) — `RotateArray`

---

### 4. Move Zeros to End

Given an array `nums[]`, move all `0`s to the end while maintaining the relative order of the non-zero elements. Must be done in place.

**Example**
```
Input:  [0, 1, 0, 3, 12]
Output: [1, 3, 12, 0, 0]
```

Solution: [`move-zeros-end.go`](move-zeros-end.go) — `MoveZerosToEnd`

---

### 5. Remove Duplicates from Sorted Array

Given a **sorted** array `nums[]`, remove duplicates in place so each element appears only once, and return the resulting slice.

**Example**
```
Input:  [1, 1, 2, 2, 3]
Output: [1, 2, 3]
```

Solution: [`remove-duplicates-array.go`](remove-duplicates-array.go) — `RemoveDuplicatesFromArray`

---

### 6. Maximum Consecutive Ones/Zeros

Given a binary array `arr[]` consisting of only `0`s and `1`s, find the length of the longest contiguous sequence of either `1`s or `0`s.

**Example**
```
Input:  [1, 1, 0, 0, 0, 1]
Output: 3
```

Solution: [`maximum-consecutive.go`](maximum-consecutive.go) — `MaxConsecBits`

---

### 7. Second Largest Element

Given an array `arr[]`, find the second largest distinct element.

**Example**
```
Input:  [12, 35, 1, 10, 34, 1]
Output: 34
```

Solution: [`second-largest.go`](second-largest.go) — `SecondLargest`

---

### 8. Third Largest Element

Given an array `arr[]`, find the third largest distinct element.

**Example**
```
Input:  [1, 14, 2, 16, 10, 20]
Output: 14
```

Solution: [`third-largest.go`](third-largest.go) — `ThirdLargest`

---

### 9. Maximum Product of a Triplet

Given an integer array, find the maximum product of any triplet (subsequence of size 3) in the array.

**Example**
```
Input:  [10, 3, 5, 6, 20]
Output: 1200   // 10 * 6 * 20

Input:  [-10, -3, -5, -6, -20]
Output: -90

Input:  [1, -4, 3, -6, 7, 0]
Output: 168
```

Solution: [`three-great-candidates.go`](three-great-candidates.go) — `maximumProduct`

---

### 10. Maximum Product Subarray

Given an integer array `nums`, find a **contiguous** subarray that has the largest product, and return the product.

**Example**
```
Input:  [2, 4, -3, 5]
Output: 8      // [2, 4]

Input:  [-3, 0, -2]
Output: 0
```

Solution: [`maximum-product-subarray.go`](maximum-product-subarray.go) — `MaximumProductSubarray`
