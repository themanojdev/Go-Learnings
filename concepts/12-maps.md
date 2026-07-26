# 🗺️ Maps in Go

## What is a Map?
A **map** is a collection of **key-value pairs**, where each key is unique and maps to a specific value.

**Key Points:**
- Stores data as key-value pairs
- Keys must be unique
- Keys must be a comparable type (string, int, bool, etc.)
- Values can be any type
- Unordered - no guaranteed order when iterating
- Reference type (like slices)

---

## Syntax
 
```go
var mapName map[KeyType]ValueType
```

**Example:**
 
```go
var ages map[string]int  // Key: string, Value: int
```

---

## Declaring and Initializing Maps
### Method 1: Declare with var (Nil Map)

```go
package main
 
import "fmt"
 
func main() {
    var ages map[string]int
    
    fmt.Println(ages)        // map[]
    fmt.Println(ages == nil) // true
}
```
 
**Output:**
 
```
map[]
true
```

**Warning:** A nil map cannot have values added to it!
 
```go
var ages map[string]int
ages["Manoj"] = 25  // PANIC! assignment to entry in nil map
```

### Method 2: Using make()
 
```go
package main
 
import "fmt"
 
func main() {
    ages := make(map[string]int)
    
    ages["Manoj"] = 25
    ages["Priya"] = 30
    
    fmt.Println(ages)  // map[Manoj:25 Priya:30]
}
```

**Output:**
 
```
map[Manoj:25 Priya:30]
```

### Method 3: Map Literal (Declare with Values)
 
```go
package main
 
import "fmt"
 
func main() {
    ages := map[string]int{
        "Manoj": 25,
        "Priya": 30,
        "Raj":   28,
    }
    
    fmt.Println(ages)
}
```
 
**Output:**
 
```
map[Manoj:25 Priya:30 Raj:28]
```
 
---

## Why "var" Creates a Nil Map (Important!)
When you declare a map with `var` and no value, Go doesn't create an actual map structure in memory. It just creates a variable holding the **zero value** for a map type, which is `nil`.

### Why This Matters
A map in Go is internally a **pointer to a hash table structure** that Go manages. That hash table needs to be **allocated** (given actual memory) before you can insert anything into it.

- `var ages map[string]int` → creates a map variable that points to **nothing** (nil)
- `make(map[string]int)` → actually **allocates** the hash table and gives you a valid pointer to it
Think of it like this:
- `var ages map[string]int` = You have a **label** "ages" but no actual filing cabinet exists yet
- `ages["Manoj"] = 25` = You're trying to put a file into a cabinet that was never built → **panic!**
- `make(map[string]int)` = Now the filing cabinet is actually built, so you can add files

### Why Reading Works but Writing Doesn't
 
Reading from a nil map is completely safe:
 
```go
package main
 
import "fmt"
 
func main() {
    var ages map[string]int
    
    value := ages["Manoj"]  // Works! Returns 0 (zero value)
    fmt.Println(value)      // 0
}
```
 
**Output:**
 
```
0
```

Go allows reads on nil maps because a lookup on an empty structure logically just means "not found," so it safely returns the zero value. But a **write** requires actual memory to store the key-value pair, and since no memory was ever allocated, Go panics instead of silently failing.

### The Fix
 
```go
package main
 
import "fmt"
 
func main() {
    ages := make(map[string]int)  // Allocates actual map memory
    
    ages["Manoj"] = 25  // Works now!
    fmt.Println(ages)   // map[Manoj:25]
}
```
 
**Output:**
 
```
map[Manoj:25]
```

Or using a map literal, which also allocates memory:
 
```go
ages := map[string]int{}  // Also allocates - safe to write to
```
 
### Quick Reference
 
| Declaration | Allocated? | Can Read? | Can Write? |
|-------------|-----------|-----------|------------|
| `var ages map[string]int` | No (nil) | Yes (returns zero value) | **No - panics** |
| `make(map[string]int)` | Yes | Yes | Yes |
| `map[string]int{}` | Yes | Yes | Yes |
 
---

## Accessing Values
 
```go
package main
 
import "fmt"
 
func main() {
    ages := map[string]int{
        "Manoj": 25,
        "Priya": 30,
    }
    
    fmt.Println(ages["Manoj"])   // 25
    fmt.Println(ages["Unknown"]) // 0 (zero value, no error!)
}
```

**Output:**
 
```
25
0
```
 
Notice: accessing a **non-existent key** doesn't cause an error - it returns the **zero value** of the value type!

---

## The "Comma OK" Idiom (Check if Key Exists)
 
Since missing keys return zero values, how do you know if a key **actually exists** vs has a zero value?
 
```go
package main
 
import "fmt"
 
func main() {
    ages := map[string]int{
        "Manoj": 25,
        "Priya": 0,  // Actual value is 0
    }
    
    value, exists := ages["Priya"]
    fmt.Println(value, exists)  // 0 true (key exists!)
    
    value, exists = ages["Unknown"]
    fmt.Println(value, exists)  // 0 false (key doesn't exist!)
}
```

**Output:**
 
```
0 true
0 false
```

This is a **very important pattern** used constantly in Go!
 
---

## Adding and Updating Values
 
```go
package main
 
import "fmt"
 
func main() {
    ages := make(map[string]int)
    
    ages["Manoj"] = 25  // Add
    fmt.Println(ages)   // map[Manoj:25]
    
    ages["Manoj"] = 26  // Update (same key)
    fmt.Println(ages)   // map[Manoj:26]
}
```
 
**Output:**
 
```
map[Manoj:25]
map[Manoj:26]
```
 
---

## Deleting Keys
 
Use the built-in `delete()` function.
 
```go
package main
 
import "fmt"
 
func main() {
    ages := map[string]int{
        "Manoj": 25,
        "Priya": 30,
    }
    
    fmt.Println("Before:", ages)
    
    delete(ages, "Manoj")
    
    fmt.Println("After:", ages)
}
```
 
**Output:**
 
```
Before: map[Manoj:25 Priya:30]
After: map[Priya:30]
```

**Note:** Deleting a non-existent key is safe - no error occurs.
 
---

## Length of a Map
 
```go
package main
 
import "fmt"
 
func main() {
    ages := map[string]int{
        "Manoj": 25,
        "Priya": 30,
    }
    
    fmt.Println(len(ages))  // 2
}
```
 
**Output:**
 
```
2
```
 
---

```go
package main
 
import "fmt"
 
func main() {
    ages := map[string]int{
        "Manoj": 25,
        "Priya": 30,
        "Raj":   28,
    }
    
    for key, value := range ages {
        fmt.Println(key, value)
    }
}
```
 
**Output (order may vary each time!):**
 
```
Manoj 25
Priya 30
Raj 28
```

**Important:** Map iteration order is **not guaranteed** in Go! Every time you run this, the order might be different.
 
---

## Maps are Reference Types
 
Like slices, maps are reference types - copying a map variable doesn't create a new map, both point to the same underlying data.
 
```go
package main
 
import "fmt"
 
func main() {
    map1 := map[string]int{"a": 1}
    map2 := map1  // Both point to same map
    
    map2["a"] = 100
    
    fmt.Println(map1)  // map[a:100] - Also changed!
    fmt.Println(map2)  // map[a:100]
}
```
 
**Output:**
 
```
map[a:100]
map[a:100]
```
 
---

## Maps with Struct Values (Preview)
 
Maps can hold any type as values, including structs (covered in detail later, but here's a preview):
 
```go
package main
 
import "fmt"
 
type Person struct {
    Name string
    Age  int
}
 
func main() {
    people := map[string]Person{
        "p1": {Name: "Manoj", Age: 25},
        "p2": {Name: "Priya", Age: 30},
    }
    
    fmt.Println(people["p1"].Name)  // Manoj
}
```
 
**Output:**
 
```
Manoj
```
 
---

## Nested Maps
 
Maps can hold other maps as values.
 
```go
package main
 
import "fmt"
 
func main() {
    students := map[string]map[string]int{
        "Manoj": {"Math": 90, "Science": 85},
        "Priya": {"Math": 95, "Science": 88},
    }
    
    fmt.Println(students["Manoj"]["Math"])  // 90
}
```
 
**Output:**
 
```
90
```
 
---

## Common Mistakes
 
### Mistake 1: Using Nil Map Without make()
 
```go
var ages map[string]int
ages["Manoj"] = 25  // PANIC! nil map
```
 
### Mistake 2: Assuming Map Order
 
```go
// Don't rely on map order - it's random!
for key, value := range someMap {
    // Order changes between runs
}
```
 
### Mistake 3: Not Checking Key Existence
 
```go
value := ages["Unknown"]  // Returns 0, might think it's a real value
// Better: use comma-ok idiom
value, exists := ages["Unknown"]
```
 
---

## When to Use Maps
 
Use **maps** when:
- You need fast lookups by a unique key
- You're counting occurrences (word frequency, etc.)
- You need to check existence of items quickly
- Data doesn't need to maintain order
Use **slices** instead when:
- Order matters
- You need indexed access (0, 1, 2...)
- You're iterating in sequence
---

## 💡 Memory Points
 
1. **Map** = Collection of unique key-value pairs
2. **Syntax** = `map[KeyType]ValueType`
3. **var declaration** = Creates a **nil** map (cannot write to it)
4. **make()** = Allocates memory, safe to read and write
5. **Map literal `{}`** = Also allocates memory, safe to use
6. **Nil map** = Reading works (returns zero value), writing panics
7. **Missing key** = Returns zero value, no error
8. **Comma-ok idiom** = `value, exists := myMap[key]` to check existence
9. **delete()** = Removes a key, safe even if key doesn't exist
10. **Iteration order** = Never guaranteed, changes between runs
11. **Reference type** = Copying a map shares the same underlying data
12. **len()** = Returns number of key-value pairs
 