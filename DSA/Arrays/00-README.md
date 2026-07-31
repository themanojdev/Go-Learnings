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
11. [Sort Array in Wave Form](#11-sort-array-in-wave-form)
12. [Add One to a Number (Digit Array)](#12-add-one-to-a-number-digit-array)

---

### 1. Reverse an Array

Given an array `arr[]`, reverse its elements in place.

**Example**
```
Input:  [1, 2, 3, 4, 5]
Output: [5, 4, 3, 2, 1]
```

Solution: [`01-reverse.go`](01-reverse.go) — `ReverseArray`

---

### 2. Reverse Array in Groups of K

Given an array `arr[]` and an integer `k`, reverse every consecutive group of `k` elements in place. If the last group has fewer than `k` elements, reverse it as it is.

**Example**
```
Input:  arr = [1, 2, 3, 4, 5, 6, 7, 8], k = 3
Output: [3, 2, 1, 6, 5, 4, 8, 7]
```

Solution: [`02-reverse-group.go`](02-reverse-group.go) — `ReverseArrayGroupWise`

---

### 3. Rotate Array by D (Left)

Given an array of integers `arr[]` of size `n`, rotate the array elements to the left by `d` positions.

**Example**
```
Input:  arr = [1, 2, 3, 4, 5, 6], d = 2
Output: [3, 4, 5, 6, 1, 2]
```

Solution: [`03-rotate-array.go`](03-rotate-array.go) — `RotateArray`

---

### 4. Move Zeros to End

Given an array `nums[]`, move all `0`s to the end while maintaining the relative order of the non-zero elements. Must be done in place.

**Example**
```
Input:  [0, 1, 0, 3, 12]
Output: [1, 3, 12, 0, 0]
```

Solution: [`04-move-zeros-end.go`](04-move-zeros-end.go) — `MoveZerosToEnd`

---

### 5. Remove Duplicates from Sorted Array

Given a **sorted** array `nums[]`, remove duplicates in place so each element appears only once, and return the resulting slice.

**Example**
```
Input:  [1, 1, 2, 2, 3]
Output: [1, 2, 3]
```

Solution: [`05-remove-duplicates-array.go`](05-remove-duplicates-array.go) — `RemoveDuplicatesFromArray`

---

### 6. Maximum Consecutive Ones/Zeros

Given a binary array `arr[]` consisting of only `0`s and `1`s, find the length of the longest contiguous sequence of either `1`s or `0`s.

**Example**
```
Input:  [1, 1, 0, 0, 0, 1]
Output: 3
```

Solution: [`06-maximum-consecutive.go`](06-maximum-consecutive.go) — `MaxConsecBits`

---

### 7. Second Largest Element

Given an array `arr[]`, find the second largest distinct element.

**Example**
```
Input:  [12, 35, 1, 10, 34, 1]
Output: 34
```

Solution: [`07-second-largest.go`](07-second-largest.go) — `SecondLargest`

---

### 8. Third Largest Element

Given an array `arr[]`, find the third largest distinct element.

**Example**
```
Input:  [1, 14, 2, 16, 10, 20]
Output: 14
```

Solution: [`08-third-largest.go`](08-third-largest.go) — `ThirdLargest`

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

Solution: [`09-three-great-candidates.go`](09-three-great-candidates.go) — `maximumProduct`

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

Solution: [`10-maximum-product-subarray.go`](10-maximum-product-subarray.go) — `MaximumProductSubarray`

---

### 11. Sort Array in Wave Form

Given an array `nums[]`, rearrange it into a wave form by swapping every adjacent pair (indices `0↔1`, `2↔3`, `4↔5`, ...).

**Example**
```
Input:  [1, 2, 3, 4, 5, 6]
Output: [2, 1, 4, 3, 6, 5]
```

Note: this swaps adjacent pairs unconditionally — it produces a valid ascending/descending wave (`nums[0] >= nums[1] <= nums[2] >= ...`) only when the input is already sorted in ascending order first.

Solution: [`11-wave-form.go`](11-wave-form.go) — `SortArrayInWave`

---

### 12. Add One to a Number (Digit Array)

Given a large integer represented as an array of digits (most significant digit first, no leading zeros), increment the number by one and return the resulting digit array.

**Example**
```
Input:  [1, 2, 3]
Output: [1, 2, 4]
```

⚠️ **Note:** this solution converts the digit array into a single `int` (`result`), adds one, then splits it back into digits. That only works for numbers small enough to fit in an `int`. On LeetCode, `digits` can have up to ~100 elements — far more digits than an `int`/`int64` can hold — so `result` silently overflows and this solution **will fail** on the platform's larger test cases. The correct approach avoids converting to a number at all: walk the array from the last digit, increment with carry propagation in place, and only prepend a new leading `1` if every digit was a `9` (e.g. `[9,9,9] → [1,0,0,0]`).

Solution: [`12-add-digitTo-number.go`](12-add-digitTo-number.go) — `AddingOneDigitToGivenNumber`
