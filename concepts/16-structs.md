# 🧱 Structs in Go
 
## What is it
A **struct** is a way to group related data into one custom type. Go has no classes — structs are how you model real-world things like a `User`, `Car`, or `Order`.
 
**Key Points:**
- A struct is a collection of named fields
- Go has no `class` keyword — structs + methods replace it
- Struct fields get **zero values** automatically if not set
- Structs can contain other structs (nesting or embedding)

---

## Syntax
 
```go
type StructName struct {
    FieldName1 Type1
    FieldName2 Type2
}
```

**Example:**
```go
type User struct {
    Name string
    Age  int
    City string
}
```

---

## Method 1: Creating Struct Instances
 
```go
// With field names (recommended - clear and safe)
user1 := User{Name: "Alice", Age: 30, City: "Bangalore"}
 
// Positional (order matters, risky, avoid in real code)
user2 := User{"Alice", 30, "Bangalore"}
 
// Zero value first, then assign
var user3 User
user3.Name = "Alice"
user3.Age = 30
```

---

If a field isn't set, Go gives it a **zero value**: `""` for strings, `0` for numbers, `false` for bools, `nil` for pointers/slices/maps.
 
---

## Method 2: Structs with Pointers
Passing a struct normally **copies** it. Passing a pointer lets you **modify the original**.
 
```go
func birthday(u User) {
    u.Age++  // only changes the copy
}
 
func birthdayPtr(u *User) {
    u.Age++  // changes the original
}
 
birthdayPtr(&user1)
```

This is the same rule used for **pointer receiver methods** — pointers are needed anytime you want to change the original data.
 
---

## Method 3: Nested Structs
A struct can contain another struct as a named field.
 
```go
type Address struct {
    City  string
    State string
}
 
type User struct {
    Name    string
    Age     int
    Address Address
}
 
u := User{Name: "Alice", Address: Address{City: "Bangalore", State: "KA"}}
fmt.Println(u.Address.City)  // dot-chain to reach nested field
```
 
---

## Practical Code Example
 
```go
package main
 
import "fmt"
 
type Address struct {
    City  string
    State string
}
 
type User struct {
    Name    string
    Age     int
    Address Address
}
 
func main() {
    u := User{
        Name: "Alice",
        Age:  30,
        Address: Address{
            City:  "Bangalore",
            State: "KA",
        },
    }
 
    fmt.Println(u.Name)
    fmt.Println(u.Address.City)
}
```
 
**Output:**
```
Alice
Bangalore
```
 
---

## Struct Embedding (Composition, Not Inheritance)
Embedding means placing a struct inside another **without a field name** — just the type. The outer struct then gets direct access to the inner struct's fields and methods.
 
```go
type Animal struct {
    Name string
}
 
func (a Animal) Speak() string {
    return a.Name + " makes a sound"
}
 
type Dog struct {
    Animal        // embedded - no field name, just the type
    Breed string
}
 
func main() {
    d := Dog{Animal: Animal{Name: "Rex"}, Breed: "Labrador"}
 
    fmt.Println(d.Name)     // Rex        <- promoted field
    fmt.Println(d.Speak())  // Rex makes a sound  <- promoted method
}
```

**Output:**
```
Rex
Rex makes a sound
```

**Nested vs Embedded — the key difference:**
 
| | Nested Struct | Embedded Struct |
|---|---|---|
| Field name | Has one (`Address Address`) | None (just `Animal`) |
| Access | `user.Address.City` | `dog.Name` (direct) |
| Purpose | Group related data | Reuse fields/methods |
 
---

## Struct Tags (Used for JSON and Databases)
Struct tags are metadata written after a field, telling libraries like `encoding/json` how to map that field.

```go
type User struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email,omitempty"`
}
```

```go
package main
 
import (
    "encoding/json"
    "fmt"
)
 
type User struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email,omitempty"`
}
 
func main() {
    u := User{Name: "Alice", Age: 30}
    data, _ := json.Marshal(u)
    fmt.Println(string(data))
}
```

**Output:**
```
{"name":"Alice","age":30}
```

`Email` is missing from the output because it was empty and `omitempty` skips empty fields.
 
---

## Constructor Pattern (`NewXxx` functions)
Go has no constructors, so the convention is a function starting with `New` that builds and returns the struct.
 
```go
type User struct {
    Name string
    Age  int
}
 
func NewUser(name string, age int) (*User, error) {
    if age < 0 {
        return nil, fmt.Errorf("age cannot be negative")
    }
    return &User{Name: name, Age: age}, nil
}
```

This gives you a clean, safe place to add validation or default values before the struct is created.
 
---

## Empty Struct `struct{}`
An empty struct has no fields and takes **zero bytes** of memory. It's used when you only care that a key exists, not its value.
 
```go
set := make(map[string]struct{})
set["apple"] = struct{}{}
set["banana"] = struct{}{}
 
_, exists := set["apple"]
fmt.Println(exists)
```
 
**Output:**
```
true
```

Common uses: building a set (Go has no built-in set type), and signaling on channels (`chan struct{}`) when no actual data needs to be sent.
 
---

## Comparing Structs
Structs can be compared with `==` if every field inside is comparable.
 
```go
user1 := User{Name: "Alice", Age: 30}
user2 := User{Name: "Alice", Age: 30}
fmt.Println(user1 == user2)  // true
```

If a struct contains a slice or map, it can no longer be compared with `==` — Go will throw a compile error.
 
---

## Field Visibility (Exported vs Unexported)
Same capitalization rule as packages: uppercase = visible outside the package, lowercase = only visible inside it.
 
```go
type user struct {
    Name string  // exported - visible outside package
    age  int     // unexported - only visible inside package
}
```

Used in real projects to hide internal state and only expose what's needed through methods.
 
---

## When to Use
- **Plain struct:** grouping related fields together (a user, a product, a config)
- **Pointer to struct:** whenever the struct needs to be modified or is large
- **Nested struct:** when one entity clearly "has" another (User has an Address)
- **Embedded struct:** when one entity should reuse another's behavior (Dog is-a-kind-of Animal)
- **Struct tags:** whenever the struct is converted to/from JSON or stored in a database
- **Constructor function:** whenever creation needs validation or default values
- **Empty struct:** building sets or signaling without needing to send data

---

## Real-World / Project Usage
1. **API request/response models** — JSON payloads map directly onto structs, using struct tags
2. **Database models** — each table row is represented as a struct
3. **Configuration** — app settings loaded into a `Config` struct at startup
4. **Layered architecture** — `handler → service → repository`, where each layer often passes structs between them
5. **Custom errors** — error types are structs with an `Error()` method attached
6. **Sets and signals** — `map[string]struct{}` and `chan struct{}` for memory-efficient patterns

---

## Interview Questions This Covers
 
- What is a struct, and how is it different from a class?
- How does Go achieve code reuse without inheritance? (embedding)
- What's the difference between passing a struct by value vs by pointer?
- What is the difference between a nested struct and an embedded struct?
- When are two structs equal in Go?
- What are struct tags used for?
- Why does Go use `NewXxx()` functions instead of constructors?
- What is an empty struct, and why is it useful?

---

## 💡 Memory Points
 
1. A struct groups related fields into one custom type
2. Go has no classes — structs + methods do the same job
3. Unset fields automatically get zero values
4. Passing a struct copies it; passing a pointer lets you modify the original
5. Nested struct = named field holding another struct; access with `outer.inner.field`
6. Embedded struct = no field name, just the type; fields/methods are promoted directly to the outer struct
7. Embedding is Go's replacement for inheritance — it's composition, not "is-a" inheritance
8. Struct tags (`` `json:"name"` ``) control how fields map to JSON or database columns
9. `omitempty` skips a field in JSON output if it's empty
10. `NewXxx()` functions are Go's constructor pattern, used for validation and defaults
11. Empty struct `struct{}` takes zero memory — used for sets and channel signals
12. Structs are comparable with `==` only if every field inside is also comparable
13. Uppercase field = exported (public), lowercase field = unexported (private to package)
14. Structs are the backbone of real Go projects: API models, DB models, configs, and error types all use them


 