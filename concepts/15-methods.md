# 🔧 Methods in Go

## What is it
A **method** is a function that is attached to a specific type using a **receiver**. It's how Go achieves object-oriented behavior without using classes.

A **function** is standalone — it does not belong to any type.

**Key Points:**
- A method = a function + a receiver
- The receiver sits between `func` and the method name: `func (r Type) MethodName()`
- Methods can be attached to structs, and even to custom named types (like `type Celsius float64`)
- Go does not have classes or inheritance — methods + structs + interfaces replace that entirely
- There are two kinds of receivers: **value receiver** and **pointer receiver**

---

## Syntax
 
```go
// Function - no receiver, standalone
func FunctionName(params) returnType {
    // body
}
 
// Method - has a receiver, attached to a type
func (r ReceiverType) MethodName(params) returnType {
    // body
}
```
 
---

## Method 1: Value Receiver
The method receives a **copy** of the original value. Changes made inside do NOT affect the original.
 
```go
type Dog struct {
    Name string
    Age  int
}
 
func (d Dog) Bark() string {
    return d.Name + " says Woof!"
}
```

**Use when:**
- You are only reading data, not modifying it
- The struct is small (cheap to copy)

---

## Method 2: Pointer Receiver
The method receives a **pointer** to the original value. Changes made inside DO affect the original.
 
```go
func (d *Dog) HaveBirthday() {
    d.Age++
}
```
 
**Use when:**
- You need to modify the receiver's fields
- The struct is large (avoids copying overhead)
- You want consistent behavior across all methods of that type

---

## Method 3: Methods on Non-Struct Types
Go allows methods on any named type, not just structs. This is heavily used in real projects for validation and formatting logic.
 
```go
type Celsius float64
 
func (c Celsius) ToFahrenheit() float64 {
    return float64(c)*9/5 + 32
}
```
 
---

## Method 3: Methods on Non-Struct Types
Go allows methods on any named type, not just structs. This is heavily used in real projects for validation and formatting logic.
 
```go
type Celsius float64
 
func (c Celsius) ToFahrenheit() float64 {
    return float64(c)*9/5 + 32
}
```
 
---

## Practical Code Example
 
```go
package main
 
import "fmt"
 
type Account struct {
    Owner   string
    Balance float64
}
 
// Value receiver - only reads data
func (a Account) GetBalance() float64 {
    return a.Balance
}
 
// Pointer receiver - modifies data
func (a *Account) Deposit(amount float64) {
    a.Balance += amount
}
 
func (a *Account) Withdraw(amount float64) bool {
    if a.Balance >= amount {
        a.Balance -= amount
        return true
    }
    return false
}
 
func main() {
    acc := Account{Owner: "Alice", Balance: 1000}
 
    fmt.Println("Balance:", acc.GetBalance())
 
    acc.Deposit(500)
    fmt.Println("After deposit:", acc.GetBalance())
 
    success := acc.Withdraw(300)
    fmt.Println("Withdrew 300:", success)
    fmt.Println("Final balance:", acc.GetBalance())
}
```
 
**Output:**
```
Balance: 1000
After deposit: 1500
Withdrew 300: true
Final balance: 1200
```
---

## Function vs Method Comparison
 
| Aspect | Function | Method |
|---|---|---|
| Declaration | `func Name()` | `func (r Type) Name()` |
| Attached to | Nothing | A specific type |
| Called as | `Name()` | `variable.Name()` |
| Modifies original | Only via pointer arguments | Via pointer receiver |
| Use case | General-purpose logic | Behavior tied to a type |
 
---

## Value Receiver vs Pointer Receiver Rules
 
- If a method needs to **modify** the receiver → use pointer receiver
- If the struct is **large** → use pointer receiver (avoids copying)
- If you're only **reading** and the struct is small → value receiver is fine
- A value can call both value and pointer receiver methods automatically (Go handles referencing)
- A pointer can call both value and pointer receiver methods automatically (Go handles dereferencing)
- **Best practice:** Don't mix receiver types on the same struct's methods — pick one style and stay consistent

---

## When to Use
 
- **Value Receiver:** Read-only operations, small structs, no need to track changes (e.g., `GetBalance()`, `String()` formatting)
- **Pointer Receiver:** State-changing operations, large structs, working with slices/maps inside the struct (e.g., `Deposit()`, `UpdateStatus()`, `SetName()`)

---

## Real-World / Project Usage
 
This is where methods matter beyond just interviews — they show up constantly in real Go codebases:
 
1. **API Handlers** — HTTP handler structs use methods to attach behavior: `func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request)`
2. **Database Models** — Structs representing DB rows use methods like `func (u *User) Save() error` or `func (u User) Validate() bool`
3. **Interfaces** — Methods are what make a type satisfy an interface. Without methods, interfaces in Go would not work at all
4. **Custom Error Types** — `func (e *MyError) Error() string` is the standard way to create custom errors
5. **Configuration Structs** — Methods like `func (c *Config) Load() error` are used to build clean, chainable setup logic
6. **JSON Marshaling** — Custom `MarshalJSON()` / `UnmarshalJSON()` methods control how structs convert to/from JSON

---

## Interview Questions This Covers
 
- What is the difference between a method and a function in Go?
- What is the difference between a value receiver and a pointer receiver?
- Why does Go not have classes, and how do methods replace them?
- Can you define a method on a non-struct type?
- How do methods relate to interfaces in Go?
- When would you choose a pointer receiver over a value receiver?

---

## 💡 Memory Points
 
1. A method = a function + a receiver attached to a type
2. Functions are standalone; methods belong to a type
3. Receiver syntax: `func (r Type) MethodName()`
4. Value receiver → works on a copy, original stays unchanged
5. Pointer receiver → works on the original, changes persist
6. Use pointer receiver to modify data or avoid copying large structs
7. Use value receiver for small, read-only operations
8. Methods can be attached to any named type, not just structs
9. Go achieves OOP-like behavior through methods + structs + interfaces (no classes/inheritance)
10. Don't mix value and pointer receivers on the same type without good reason
11. Methods are essential for satisfying interfaces — no method, no interface match
12. Real-world usage: HTTP handlers, DB models, custom errors, JSON marshaling all rely heavily on methods
13. Go auto-handles referencing/dereferencing when calling methods, but method sets still differ between value and pointer types (important for interface satisfaction)
14. Called using dot notation: `variable.MethodName()`
 