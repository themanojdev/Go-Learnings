# 🧭 If Statement with Initial Statement

## What is it

Go's `if` statement can include an **initial statement** before the condition — a small setup step that runs right before the check happens. Any variable created there is only visible **inside** that `if` block (and its `else if` / `else` blocks) — nowhere else.

**Key Points:**
- Syntax: `if INITIAL_STATEMENT; CONDITION { }`
- The variable created in the initial statement is scoped **only** to the if/else if/else chain
- It's shorter than declaring the variable on its own line beforehand
- It prevents the variable from leaking into the rest of the function, where it isn't needed

---

## Syntax

```go
if INITIAL_STATEMENT; CONDITION {
    // variable is available here
} else if CONDITION2 {
    // still available here
} else {
    // still available here
}
// NOT available here - out of scope
```

---

## Method 1: Without Initial Statement (the long way)

```go
length := getLength(email)
if length < 10 {
    fmt.Printf("Email must be at least 10 characters, is %d\n", length)
}
```

Here, `length` stays alive for the **rest of the function**, even though it's only needed for this one check.

---

## Method 2: With Initial Statement (the short way)

```go
if length := getLength(email); length < 10 {
    fmt.Printf("Email must be at least 10 characters, is %d\n", length)
}
```

Same result, but `length` only exists inside this `if` block. Once the block ends, `length` is gone — you can't accidentally use it later in the function.

---

## Practical Code Example

```go
package main

import "fmt"

func getLength(s string) int {
    return len(s)
}

func main() {
    email := "abc@x.com"

    if length := getLength(email); length < 10 {
        fmt.Printf("Email must be at least 10 characters, is %d\n", length)
    } else {
        fmt.Printf("Email length is fine: %d\n", length)
    }

    // length is NOT accessible here - it's out of scope
}
```

**Output:**
```
Email must be at least 10 characters, is 9
```

---

## Why Use This

1. **Shorter code** — combines declaration and check into one line
2. **Limits scope** — the variable only exists where it's actually needed, so it can't be misused or accidentally referenced later in the function

---

## When to Use

- When a variable is only needed to make **one decision** and nothing else afterward
- Very common with functions that return `(value, error)` or `(value, ok)`, e.g.:

```go
if val, ok := myMap["key"]; ok {
    fmt.Println("Found:", val)
}

if err := doSomething(); err != nil {
    fmt.Println("Error:", err)
}
```

---

## Real-World / Project Usage

1. **Error checking** — the most common pattern in Go: `if err := doSomething(); err != nil { ... }` keeps `err` scoped to just that check
2. **Map lookups** — `if val, ok := myMap[key]; ok { ... }` avoids leaking `val`/`ok` into the rest of the function
3. **Cleaner handler code** — in HTTP handlers and service functions, this pattern keeps temporary variables from cluttering the function scope

---

## Interview Questions This Covers

- What is the initial statement in an `if` block, and what problem does it solve?
- What is the scope of a variable declared in an `if` statement's initial statement?
- Why is this pattern so common in Go's error handling style?

---

## 💡 Memory Points

1. Syntax: `if INITIAL_STATEMENT; CONDITION { }`
2. The initial statement runs once, right before the condition is checked
3. Variables created there are scoped to the `if`, `else if`, and `else` blocks only
4. Outside that chain, the variable no longer exists
5. This keeps variables from leaking into the rest of the function when they're only needed for one check
6. Extremely common with error handling: `if err := doSomething(); err != nil { }`
7. Also common with map lookups: `if val, ok := myMap[key]; ok { }`
8. It's a scoping convenience, not a different kind of `if` — the logic behaves exactly like a normal `if`