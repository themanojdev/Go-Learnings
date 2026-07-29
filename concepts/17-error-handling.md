# ⚠️ Error Handling in Go

## What is it
In Go, an error is simply a **normal return value** that tells you something went wrong. There are no exceptions, no `try/catch`, no jumping around in your code — a function that can fail just returns an `error` alongside its result.

**Key Points:**
- Errors are values, not special interruptions
- A function that can fail returns `(result, error)`
- You check `if err != nil` right after calling it
- `error` is just a built-in interface: `type error interface { Error() string }`
- Any type with an `Error() string` method satisfies this interface

---

## Syntax
 
```go
func DoSomething() (ResultType, error) {
    if somethingWentWrong {
        return zeroValue, errors.New("what went wrong")
    }
    return result, nil
}
```

**The standard check:**
```go
result, err := DoSomething()
if err != nil {
    // handle the problem
    return
}
// safe to use result here
```
 
---

## Method 1: Creating Errors — `errors.New()`
Use this for a plain, fixed message that never changes.
 
```go
err := errors.New("cannot divide by zero")
```
 
---

## Method 2: Creating Errors — `fmt.Errorf()`
Use this when you want to insert details (a name, number, ID) into the message. It works like `fmt.Printf()`, but builds an error instead of printing text.
 
```go
err := fmt.Errorf("cannot divide %d by zero", a)
```

**Simple rule:**
> Use `errors.New()` when the message never changes. Use `fmt.Errorf()` when you need to plug in specific details.
 
---

## Method 3: Wrapping Errors with `%w`
`%w` works like `%s` (it inserts the error's message) but also keeps the original error attached inside, so it can be found again later.

```go
originalErr := errors.New("user not found")
wrappedErr := fmt.Errorf("failed to load profile: %w", originalErr)
 
fmt.Println(wrappedErr)              // failed to load profile: user not found
fmt.Println(errors.Unwrap(wrappedErr)) // user not found
```

 This is how errors travel through layered code (repository → service → handler) — each layer adds context while keeping the original error intact underneath.
 
---

## Practical Code Example
 
```go
package main
 
import (
    "errors"
    "fmt"
)
 
var ErrUserNotFound = errors.New("user not found")
 
func getUserFromDB(id int) error {
    return ErrUserNotFound
}
 
func getUserService(id int) error {
    err := getUserFromDB(id)
    if err != nil {
        return fmt.Errorf("getUserService: %w", err)
    }
    return nil
}
 
func main() {
    err := getUserService(5)
    fmt.Println(err)
}
```

**Output:**
```
getUserService: user not found
```
 
---

## Sentinel Errors
A **sentinel error** is a fixed, named error value created once and reused everywhere, so your code can check "did THIS exact problem happen?"

```go
var ErrUserNotFound = errors.New("user not found")
```

| | Regular error | Sentinel error |
|---|---|---|
| Created | Anywhere, one-off | Once, at the top of the file |
| Reused? | No | Yes, same variable everywhere |
| Purpose | Just describe a problem | Let other code check for this exact problem |
 
Real examples from the standard library: `io.EOF`, `sql.ErrNoRows` — both created once and checked everywhere with `errors.Is()`.
 
---

## `errors.Is()` — Checking for a Specific Error
Checks if a specific known error is hiding anywhere inside a wrapped error, even after multiple layers of wrapping.
 
```go
package main
 
import (
    "errors"
    "fmt"
)
 
var ErrUserNotFound = errors.New("user not found")
 
func getUserFromDB(id int) error {
    return ErrUserNotFound
}
 
func getUserService(id int) error {
    err := getUserFromDB(id)
    if err != nil {
        return fmt.Errorf("getUserService: %w", err)
    }
    return nil
}
 
func main() {
    err := getUserService(5)
 
    if errors.Is(err, ErrUserNotFound) {
        fmt.Println("Handled: user does not exist")
    } else {
        fmt.Println("Some other error occurred")
    }
}
```
 
**Output:**
```
Handled: user does not exist
```

Direct `==` comparison would fail here because `err` is a new wrapped error, not literally the same value as `ErrUserNotFound`. `errors.Is()` unwraps it layer by layer to find the match.
 
---

## `errors.As()` — Checking for a Specific Error Type
Used when the error carries extra data (like a field name) inside a custom error type. `errors.As()` digs through the wrapped error, finds a matching type, and fills in your variable so you can read its fields.
 
```go
package main
 
import (
    "errors"
    "fmt"
)
 
type ValidationError struct {
    Field string
}
 
func (e *ValidationError) Error() string {
    return "invalid field: " + e.Field
}
 
func validateUser(email string) error {
    if email == "" {
        return &ValidationError{Field: "email"}
    }
    return nil
}
 
func processUser(email string) error {
    err := validateUser(email)
    if err != nil {
        return fmt.Errorf("processUser failed: %w", err)
    }
    return nil
}
 
func main() {
    err := processUser("")
 
    var valErr *ValidationError
    if errors.As(err, &valErr) {
        fmt.Println("Validation problem on field:", valErr.Field)
    } else {
        fmt.Println("Some other error:", err)
    }
}
```
 
**Output:**
```
Validation problem on field: email
```
 
---

## `errors.Is()` vs `errors.As()`
 
| | `errors.Is()` | `errors.As()` |
|---|---|---|
| Question it answers | "Is this THAT exact error?" | "Is this THAT KIND of error, and can I read its data?" |
| Works with | Sentinel errors | Custom error types (structs with `Error()`) |
| Result | true/false | true/false + fills in the struct so you can read its fields |
 
---

## Error vs Panic
Go gives you two ways to handle something going wrong — but they are meant for very different situations.
 
- **`error`** → for problems you **expect** could happen (bad input, file not found, network timeout). Handle it calmly with `if err != nil`.
- **`panic`** → for problems that should **never** happen in correct code (nil pointer dereference, array index out of bounds, programmer bugs). It stops the program immediately unless recovered.

```go
// error - expected, recoverable
func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
 
// panic - unexpected, should crash loudly
func MustPositive(n int) int {
    if n < 0 {
        panic("negative number not allowed")
    }
    return n
}
```
 
**Recovering from a panic** (usually done at the top level, like in a web server, so one bad request doesn't crash the whole app):

```go
func safeCall() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    panic("something broke")
}
```
 
**Output:**
```
Recovered from panic: something broke
```

### Error vs Panic Comparison
 
| | `error` | `panic` |
|---|---|---|
| Use case | Expected, recoverable failures | Unexpected, unrecoverable bugs |
| Handling | `if err != nil` | `recover()` inside `defer` |
| Convention | Normal Go code path | Rare — avoid using for regular error handling |
| Example | Invalid user input, missing file | Nil pointer, index out of range |
 
**Simple rule to remember:**
> If it can realistically happen during normal use, return an `error`. If it means the program reached a state that should be impossible, `panic`.
 
---

## When to Use
 
- **`errors.New()`** — a fixed, simple error message
- **`fmt.Errorf()`** — an error message that needs specific details plugged in
- **`%w` wrapping** — adding context to an error while a layer passes it upward
- **Sentinel errors** — a known, specific failure other code needs to check for (not found, permission denied)
- **Custom error types** — an error that needs to carry extra data (which field, which status code)
- **`errors.Is()`** — checking if a specific known error occurred
- **`errors.As()`** — checking for a specific error type and reading its fields
- **`error` return** — anything that can realistically happen (bad input, network failure)
- **`panic`** — only for truly impossible states (programmer bugs, corrupted memory) — rare in normal Go code
---
 
## Real-World / Project Usage
 
1. **API handlers** — check the error with `errors.Is()`/`errors.As()` and convert it to the right HTTP status code (404, 403, 500)
2. **Database layers** — wrap SQL errors with context using `%w`, e.g. `fmt.Errorf("query failed: %w", err)`
3. **Validation** — custom error types (like `ValidationError`) carry which field failed and why
4. **Standard library sentinels** — `sql.ErrNoRows` and `io.EOF` are checked constantly with `errors.Is()` in real projects
5. **Layered architecture** — errors travel from repository → service → handler, gaining context at each layer without losing the original cause
---
 
## Interview Questions This Covers
 
- Why doesn't Go use exceptions like other languages?
- What is the `error` interface?
- Difference between `errors.New()` and `fmt.Errorf()`?
- What does `%w` do, and why is it useful?
- What is a sentinel error?
- Difference between `errors.Is()` and `errors.As()`?
- How do you create a custom error type?
- When would you use `panic` instead of returning an `error`?
- How does `recover()` work, and where should it be used?
---
 
## 💡 Memory Points
 
1. An error is just a normal return value — Go has no exceptions
2. A function that can fail returns `(result, error)`; check with `if err != nil`
3. `error` is a built-in interface requiring only `Error() string`
4. `errors.New()` → fixed message that never changes
5. `fmt.Errorf()` → message with details plugged in, works like `Printf`
6. `%w` wraps an error, keeping the original accessible underneath
7. `errors.Unwrap()` pulls the original error back out of a wrapped one
8. A sentinel error is a fixed, named error value created once and reused everywhere
9. `errors.Is()` checks if a specific known error is hiding inside a wrapped error
10. Direct `==` comparison breaks once an error is wrapped — use `errors.Is()` instead
11. `errors.As()` checks for a specific error type and extracts its fields
12. Custom error types are just structs with an `Error() string` method
13. `io.EOF` and `sql.ErrNoRows` are real sentinel errors from the standard library
14. Each layer in a project (repository → service → handler) should add context with `%w`, not swallow the original error
15. Use `error` for expected, recoverable problems; use `panic` only for impossible states that indicate a bug
16. `recover()` only works inside a `deferred` function, and is usually placed at the top level to stop one failure from crashing the whole program

