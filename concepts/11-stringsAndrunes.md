# 🔤 Strings & Runes in Go

## What is a String in Go?
A **string** in Go is a **sequence of bytes** (not characters!). This is a key difference from many other languages.

**Key Points:**
- Strings are immutable (cannot be changed after creation)
- Strings are stored as bytes (UTF-8 encoded)
- Each byte is a `uint8` value
- Strings can be indexed like arrays
- Default value of a string is `""` (empty string)

---

## The Problem: Bytes vs Characters
 
In Go, when you index a string, you get a **byte**, not a character!
 
```go
package main
 
import "fmt"
 
func main() {
    str := "Hello"
    fmt.Println(str[0])          // 72 (not 'H'!)
    fmt.Println(string(str[0]))  // H
}
```
 
**Output:**
 
```
72
H
```

This works fine for **ASCII characters** (English letters, numbers) because each character equals 1 byte.

**But what about special characters?**

```go
package main
 
import "fmt"
 
func main() {
    str := "héllo"  // é is a special character
    fmt.Println(len(str))  // 6 (not 5!)
}
```

**Output:**
 
```
6
```

Why 6? Because `é` takes **2 bytes** in UTF-8 encoding, not 1!

---

## UTF-8 Encoding Basics
 
Go uses **UTF-8 encoding** for strings. In UTF-8, characters can take **1 to 4 bytes** depending on the character:
 
| Character Type | Bytes Used | Example |
|-----------------|------------|---------|
| ASCII (English) | 1 byte | `A`, `1`, `!` |
| Latin extended | 2 bytes | `é`, `ñ` |
| Most other scripts | 3 bytes | Chinese, Japanese |
| Emojis, rare symbols | 4 bytes | 😀, 🎉 |
 
```go
package main
 
import "fmt"
 
func main() {
    fmt.Println(len("A"))   // 1 byte
    fmt.Println(len("é"))   // 2 bytes
    fmt.Println(len("中")) // 3 bytes
    fmt.Println(len("😀")) // 4 bytes
}
```

**Output:**
 
```
1
2
3
4
```

This is why counting length correctly matters - `len()` gives **bytes**, not characters!
 
---

## What is a Rune?
A **rune** is Go's way of representing a **single Unicode character** (code point).
 
`rune` is just an alias for `int32`.

```go
package main
 
import "fmt"
 
func main() {
    var r rune = 'A'
    fmt.Println(r)          // 65 (Unicode code point)
    fmt.Println(string(r))  // A
}
```
 
**Output:**
 
```
65
A
```
 
---

## Bytes vs Runes: The Core Difference
 
| Concept | Type | Size | Represents |
|---------|------|------|------------|
| **Byte** | `uint8` | 1 byte | Raw byte value |
| **Rune** | `int32` | 4 bytes | Unicode code point (character) |

**Visual Example:**
 
```go
package main
 
import "fmt"
 
func main() {
    str := "héllo"
    
    fmt.Println("Length (bytes):", len(str))           // 6
    fmt.Println("Length (runes):", len([]rune(str)))   // 5
}
```

**Output:**
 
```
Length (bytes): 6
Length (runes): 5
```
 
This is the **most important concept** to understand for interviews!
 
---

## Iterating Over Strings: The Right Way
 
### Method 1: Wrong Way (Byte by Byte)
 
```go
package main
 
import "fmt"
 
func main() {
    str := "héllo"
    
    for i := 0; i < len(str); i++ {
        fmt.Println(i, str[i])  // Prints byte values, breaks on é
    }
}
```
 
This will print garbled output for `é` because it's split into 2 bytes.

### Method 2: Right Way (Using Range - Rune by Rune)
 
```go
package main
 
import "fmt"
 
func main() {
    str := "héllo"
    
    for index, runeValue := range str {
        fmt.Printf("%d: %c\n", index, runeValue)
    }
}
```

**Output:**
 
```
0: h
1: é
3: l
4: l
5: o
```

Notice the index jumps from 1 to 3 (skipping 2) because `é` takes 2 bytes! But `range` correctly gives you the **rune** (character), not the raw byte.
 
---

## Converting Between String, Byte Slice, and Rune Slice
 
```go
package main
 
import "fmt"
 
func main() {
    str := "hello"
    
    // String to byte slice
    byteSlice := []byte(str)
    fmt.Println(byteSlice)  // [104 101 108 108 111]
    
    // String to rune slice
    runeSlice := []rune(str)
    fmt.Println(runeSlice)  // [104 101 108 108 111]
    
    // Byte slice back to string
    backToString := string(byteSlice)
    fmt.Println(backToString)  // hello
}
```
 
**Output:**
 
```
[104 101 108 108 111]
[104 101 108 108 111]
hello
```
 
---

## Why Strings are Immutable
 
You **cannot** change a character in a string directly:
 
```go
package main
 
func main() {
    str := "Hello"
    str[0] = 'h'  // ERROR! Cannot assign to str[0]
}
```

### Why Go Made Strings Immutable (Interview Favorite!)
 
1. **Memory safety** - Multiple variables can safely share the same underlying data without unexpected changes
2. **Thread safety** - Since strings can't change, multiple goroutines can read the same string safely without locks
3. **Hashing efficiency** - Strings are used as map keys; immutability guarantees the hash stays consistent
4. **Performance** - The compiler can optimize and safely reuse string literals
**To "modify" a string, convert it first:**

```go
package main
 
import "fmt"
 
func main() {
    str := "Hello"
    
    byteSlice := []byte(str)
    byteSlice[0] = 'h'
    
    newStr := string(byteSlice)
    fmt.Println(newStr)  // hello
}
```
 
**Output:**
 
```
hello
```
 
---

## String Concatenation Methods
 
There are multiple ways to join strings together. Choosing the right one matters for **performance**!
 
### Method 1: Using `+` Operator
Simple but **inefficient** for many concatenations (creates a new string each time).

```go
package main
 
import "fmt"
 
func main() {
    str1 := "Hello"
    str2 := "World"
    
    result := str1 + ", " + str2 + "!"
    fmt.Println(result)  // Hello, World!
}
```

**Output:**
 
```
Hello, World!
```

### Method 2: Using `fmt.Sprintf`
 
Good for formatting with variables, but slower than `strings.Builder` for loops.
 
```go
package main
 
import "fmt"
 
func main() {
    name := "Manoj"
    age := 25
    
    result := fmt.Sprintf("Name: %s, Age: %d", name, age)
    fmt.Println(result)  // Name: Manoj, Age: 25
}
```
 
**Output:**
 
```
Name: Manoj, Age: 25
```
### Method 3: Using `strings.Builder` (Most Efficient)
 
**Best choice** when concatenating strings in a loop - avoids creating multiple intermediate strings.
 
```go
package main
 
import (
    "fmt"
    "strings"
)
 
func main() {
    var builder strings.Builder
    
    builder.WriteString("Hello")
    builder.WriteString(", ")
    builder.WriteString("World")
    builder.WriteString("!")
    
    result := builder.String()
    fmt.Println(result)  // Hello, World!
}
```
 
**Output:**
 
```
Hello, World!
```

### Why strings.Builder is Faster (Interview Point!)
 
Using `+` repeatedly in a loop creates a **new string in memory every time** because strings are immutable. `strings.Builder` uses an internal byte buffer that **grows dynamically**, avoiding repeated allocations.
 
```go
package main
 
import (
    "fmt"
    "strings"
)
 
func main() {
    // Inefficient way (avoid in loops!)
    result := ""
    for i := 0; i < 5; i++ {
        result += "a"  // Creates new string each iteration
    }
    fmt.Println(result)  // aaaaa
    
    // Efficient way
    var builder strings.Builder
    for i := 0; i < 5; i++ {
        builder.WriteString("a")  // Appends to buffer
    }
    fmt.Println(builder.String())  // aaaaa
}
```
 
**Output:**
 
```
aaaaa
aaaaa
```
 
---

## String Comparison
 
### Method 1: Using `==` Operator
 
Compares strings **byte by byte**. Returns `true` if identical.
 
```go
package main
 
import "fmt"
 
func main() {
    str1 := "hello"
    str2 := "hello"
    str3 := "Hello"
    
    fmt.Println(str1 == str2)  // true
    fmt.Println(str1 == str3)  // false (case-sensitive)
}
```
 
**Output:**
 
```
true
false
```

## String Comparison
 
### Method 1: Using `==` Operator
 
Compares strings **byte by byte**. Returns `true` if identical.
 
```go
package main
 
import "fmt"
 
func main() {
    str1 := "hello"
    str2 := "hello"
    str3 := "Hello"
    
    fmt.Println(str1 == str2)  // true
    fmt.Println(str1 == str3)  // false (case-sensitive)
}
```
 
**Output:**
 
```
true
false
```

### Method 2: Using `strings.Compare()`
 
Returns `0` if equal, `-1` if first is smaller, `1` if first is larger (lexicographic order).
 
```go
package main
 
import (
    "fmt"
    "strings"
)
 
func main() {
    fmt.Println(strings.Compare("apple", "apple"))   // 0
    fmt.Println(strings.Compare("apple", "banana"))  // -1
    fmt.Println(strings.Compare("banana", "apple"))  // 1
}
```

**Output:**
 
```
0
-1
1
```

### Method 3: Case-Insensitive Comparison
 
```go
package main
 
import (
    "fmt"
    "strings"
)
 
func main() {
    str1 := "Hello"
    str2 := "hello"
    
    fmt.Println(strings.EqualFold(str1, str2))  // true
}
```
 
**Output:**
 
```
true
```
 
---

## Common String Manipulation Functions
Go's `strings` package provides many useful functions used constantly in real projects.

### Case Conversion
 
```go
package main
 
import (
    "fmt"
    "strings"
)
 
func main() {
    str := "Hello, World!"
    
    fmt.Println(strings.ToUpper(str))  // HELLO, WORLD!
    fmt.Println(strings.ToLower(str))  // hello, world!
}
```
 
**Output:**
 
```
HELLO, WORLD!
hello, world!
```

### Trimming Whitespace and Characters
 
```go
package main
 
import (
    "fmt"
    "strings"
)
 
func main() {
    str := "  Hello, World!  "
    
    fmt.Println(strings.TrimSpace(str))       // "Hello, World!"
    fmt.Println(strings.Trim(str, " !"))      // "Hello, World"
    fmt.Println(strings.TrimLeft(str, " "))   // "Hello, World!  "
    fmt.Println(strings.TrimRight(str, " "))  // "  Hello, World!"
}
```

**Output:**
 
```
Hello, World!
Hello, World
Hello, World!  
  Hello, World!
```

### Splitting Strings
 
```go
package main
 
import (
    "fmt"
    "strings"
)
 
func main() {
    str := "apple,banana,orange"
    
    parts := strings.Split(str, ",")
    fmt.Println(parts)  // [apple banana orange]
    
    sentence := "  Hello   World  Go  "
    fields := strings.Fields(sentence)  // Splits by whitespace, ignores extra spaces
    fmt.Println(fields)  // [Hello World Go]
}
```
 **Output:**
 
```
[apple banana orange]
[Hello World Go]
```

### Joining Strings
 
```go
package main
 
import (
    "fmt"
    "strings"
)
 
func main() {
    words := []string{"Hello", "World", "Go"}
    
    result := strings.Join(words, " ")
    fmt.Println(result)  // Hello World Go
    
    csv := strings.Join(words, ",")
    fmt.Println(csv)  // Hello,World,Go
}
```
**Output:**
 
```
Hello World Go
Hello,World,Go
```

### Searching Within Strings
 
```go
package main
 
import (
    "fmt"
    "strings"
)
 
func main() {
    str := "Hello, World!"
    
    fmt.Println(strings.Contains(str, "World"))     // true
    fmt.Println(strings.HasPrefix(str, "Hello"))     // true
    fmt.Println(strings.HasSuffix(str, "!"))         // true
    fmt.Println(strings.Index(str, "World"))         // 7 (position found)
    fmt.Println(strings.Count(str, "l"))             // 3 (occurrences)
}
```
 
**Output:**
 
```
true
true
true
7
3
```

### Replacing Substrings
 
```go
package main
 
import (
    "fmt"
    "strings"
)
 
func main() {
    str := "Hello, World! World is great."
    
    // Replace first n occurrences (use -1 for all)
    fmt.Println(strings.Replace(str, "World", "Go", 1))   // Hello, Go! World is great.
    fmt.Println(strings.Replace(str, "World", "Go", -1))  // Hello, Go! Go is great.
    fmt.Println(strings.ReplaceAll(str, "World", "Go"))   // Hello, Go! Go is great.
}
```
 
**Output:**
 
```
Hello, Go! World is great.
Hello, Go! Go is great.
Hello, Go! Go is great.
```
 
---

## String Conversion with strconv Package
 
The `strconv` package handles conversion between strings and other types - **used constantly in real projects** (parsing input, config files, APIs).
 
### String to Int
 
```go
package main
 
import (
    "fmt"
    "strconv"
)
 
func main() {
    str := "123"
    
    num, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(num)  // 123
}
```

**Output:**
 
```
123
```

### Int to String
 
```go
package main
 
import (
    "fmt"
    "strconv"
)
 
func main() {
    num := 456
    
    str := strconv.Itoa(num)
    fmt.Println(str)  // 456
}
```
 
**Output:**
 
```
456
```

### String to Float
 
```go
package main
 
import (
    "fmt"
    "strconv"
)
 
func main() {
    str := "3.14"
    
    num, err := strconv.ParseFloat(str, 64)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(num)  // 3.14
}
```
 
**Output:**
 
```
3.14
```

### String to Bool
 
```go
package main
 
import (
    "fmt"
    "strconv"
)
 
func main() {
    str := "true"
    
    val, err := strconv.ParseBool(str)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(val)  // true
}
```
 **Output:**
 
```
true
```



 ### Handling Invalid Conversion
 
Always check the error when converting - invalid input will return an error, not panic.
 
```go
package main
 
import (
    "fmt"
    "strconv"
)
 
func main() {
    str := "abc"
    
    num, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println("Conversion failed:", err)  // Conversion failed: ...
        return
    }
    fmt.Println(num)
}
```
 
**Output:**
 
```
Conversion failed: strconv.Atoi: parsing "abc": invalid syntax
```
 
---

## Practice Exercise: Password Validator

**Problem Statement:**
A messaging platform is enforcing a new password policy. Write a validator that checks a password string against the rules below.

**Tasks:**

1. Write a function that takes a password `string` and returns a `bool` indicating whether it is valid.
2. A password is valid only if **all** of these hold:
   - At least 5 characters long, but no more than 12
   - Contains at least one uppercase letter
   - Contains at least one digit
3. Iterate over the password's characters using `range` and classify each one with the `unicode` package (`unicode.IsUpper`, `unicode.IsDigit`).

Hint: think about whether the length check should count **bytes** (`len(password)`) or **runes** — they differ the moment a non-ASCII character shows up.

Solution: [`11-strings-runes/password-validator.go`](example-code/11-strings-runes/password-validator.go)

---

## Key Differences: String vs Byte vs Rune
 
| Type | Represents | Size | Example |
|------|-----------|------|---------|
| **string** | Sequence of bytes (UTF-8) | Variable | `"Hello"` |
| **byte** | Single raw byte | 1 byte (`uint8`) | `str[0]` gives byte |
| **rune** | Single Unicode character | 4 bytes (`int32`) | `'A'`, `'é'`, `'😀'` |
 
---
 
## When to Use What
 
Use **len(str)** when:
- You need byte count (for memory/storage purposes)
Use **len([]rune(str))** when:
- You need actual character count (for display purposes)
- Working with non-English text
Use **range** when:
- Iterating over a string to get actual characters
- Working with Unicode text safely
Use **strings.Builder** when:
- Concatenating strings in a loop
- Building large strings (better performance)
Use **strconv** when:
- Converting user input (usually strings) to numbers
- Converting numbers to strings for display or storage

---

## 💡 Memory Points
 
1. **String** = Sequence of bytes (UTF-8 encoded), immutable
2. **Byte** = `uint8`, 1 byte, raw value from indexing a string
3. **Rune** = `int32`, represents one Unicode character (1-4 bytes)
4. **len(str)** = Gives byte count, NOT character count
5. **len([]rune(str))** = Gives actual character count
6. **range over string** = Correctly iterates rune by rune (handles multi-byte chars)
7. **Immutability reasons** = Memory safety, thread safety, hashing, performance
8. **strings.Builder** = Most efficient way to concatenate in loops
9. **strconv.Atoi/Itoa** = Convert String ↔ Int
10. **Always check errors** = strconv functions return errors for invalid input
11. **strings.Split/Join** = Break apart and combine strings with delimiters
12. **strings.Contains/Index** = Search within strings
 