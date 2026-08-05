# 🧵 Goroutines in Go
 
## What is it
A goroutine is a function that runs **independently**, without making the rest of the program wait for it to finish. It's how Go achieves concurrency — doing multiple things without one fully blocking the other.

**Key Points:**
- Created with the `go` keyword in front of a function call
- Goroutines run concurrently — Go's scheduler decides when each one actually executes
- When `main()` finishes, the entire program exits immediately, even if goroutines are still running
- Concurrency ≠ parallelism (see below) — Go gives you concurrency always, and parallelism depending on the machine

---

## Concurrency vs Parallelism
- **Concurrency:** managing multiple tasks, making progress on each by switching between them. Works fine with a single CPU core.
- **Parallelism:** doing multiple tasks at the exact same instant. Requires multiple CPU cores.
| | Concurrency | Parallelism |
|---|---|---|
| What it means | Managing multiple tasks, making progress on each | Doing multiple tasks at the exact same instant |
| Needs multiple cores? | No | Yes |
| Analogy | One cook switching between rice and veggies | Two cooks, one on each |
 
Go's goroutines always give you concurrency. Whether that also becomes parallelism depends on how many CPU cores the machine has and how many Go is allowed to use (`GOMAXPROCS`).
 
---

## Syntax
 
```go
go functionName(arguments)
```

Just add `go` in front of any function call, and it runs as a goroutine instead of blocking the current code.
 
---

## Method 1: Basic Goroutine
 
```go
func sayHello() {
    fmt.Println("Hello")
}
 
func main() {
    go sayHello() // runs independently, doesn't block main
    fmt.Println("Main function")
}
```

**Output (usually):**
```
Main function
```

`"Hello"` may not print at all — because `main()` reaches its end and the program exits before the goroutine gets a chance to run.
 
---

## Method 2: Waiting for Goroutines with `sync.WaitGroup`
`time.Sleep()` can "fix" the above by pausing `main()` for a guessed amount of time, but this is unreliable — too short and the goroutine gets cut off, too long and time is wasted. The real fix is `sync.WaitGroup`, which waits for an exact signal instead of a guess.
 
**The three pieces:**
- `wg.Add(n)` — "I'm about to start `n` goroutines" (increases the counter)
- `wg.Done()` — "I finished" (decreases the counter by 1), called inside the goroutine
- `wg.Wait()` — pauses here until the counter hits `0`
```go
func sayHello(wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println("Hello")
}
 
func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go sayHello(&wg)
    wg.Wait()
    fmt.Println("Main function")
}
```

## Method 3: Anonymous Goroutines
Instead of calling a named function, you can write the function inline with no name, and call it immediately.
 
```go
go func() {
    fmt.Println("Hello")
}()
```
 
With a `WaitGroup`, in a loop:
 
```go
func main() {
    var wg sync.WaitGroup
 
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            fmt.Println("Hello")
        }()
    }
 
    wg.Wait()
    fmt.Println("Main function")
}
```
 
**Output:**
```
Hello
Hello
Hello
Main function
```
 
---

## Practical Code Example
 
```go
package main
 
import (
    "fmt"
    "sync"
)
 
func sayHello(wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println("Say Hello to world!")
}
 
func main() {
    var wg sync.WaitGroup
 
    wg.Add(3)
    for i := 0; i < 3; i++ {
        go sayHello(&wg)
    }
    wg.Wait()
 
    fmt.Println("Main Function")
}
```

**Output:**
```
Say Hello to world!
Say Hello to world!
Say Hello to world!
Main Function
```

All 3 lines are guaranteed to print before "Main Function" — but the order among the 3 goroutines themselves is never guaranteed.
 
---

## `Add`/`Done` Mismatch — Two Ways It Breaks
`Add(n)` must exactly equal the number of times `Done()` will actually be called.
 
| Mismatch Type | What Happens |
|---|---|
| `Add` count HIGHER than actual `Done()` calls | Counter never reaches 0 → `Wait()` blocks forever → **deadlock**: `fatal error: all goroutines are asleep - deadlock!` |
| `Add` count LOWER than actual `Done()` calls | Counter hits 0 early, then a further `Done()` tries to go below 0 → **panic**: `panic: sync: negative WaitGroup counter` |

**Example of the "too few Done calls" deadlock:**
```go
wg.Add(3)
go sayHello(&wg)  // only 1 goroutine launched, calls Done() once
wg.Wait()          // counter stuck at 2, never reaches 0 - main hangs forever
```

**Example of missing the `go` keyword entirely:**
```go
sayHello(&wg)  // no "go" - runs like a normal function call, not a goroutine
```
This just runs synchronously — no concurrency happens at all, and any `Wait()` afterward may hang if the counter was set expecting more goroutines than actually ran.
 
---

## The Loop Variable Capture Trap (Classic Interview Question)
 
**Pre-Go 1.22 behavior:** a `for` loop reuses the same variable across all iterations. If a goroutine reads that variable directly, it may see the wrong (final) value by the time it actually runs, since the loop has already finished changing it.
 
```go
for i := 0; i < 3; i++ {
    go func() {
        fmt.Println("Iteration:", i)
    }()
}
```
 
**Old, buggy output (pre-1.22):**
```
Iteration: 3
Iteration: 3
Iteration: 3
```

**Go 1.22+ behavior:** each loop iteration automatically gets its own private copy of the loop variable. This bug no longer happens by default — the code above correctly prints `0, 1, 2` (in random order) on Go 1.22 and later.
 
### The Old Fixes (still asked in interviews, needed for pre-1.22 code)
 
**Fix 1: Shadow the variable inside the loop**
```go
for i := 0; i < 3; i++ {
    i := i  // new local variable, a private copy for this iteration
    go func() {
        fmt.Println("Iteration:", i)
    }()
}
```

**Fix 2: Pass it as a function argument**
```go
for i := 0; i < 3; i++ {
    go func(i int) {
        fmt.Println("Iteration:", i)
    }(i)  // copied at the moment the call is made
}
```

**Why both work:** function arguments and `:=` both create a **brand-new variable, copied at that exact moment** — a private copy the goroutine reads, instead of the shared loop variable that keeps changing.
 
---

## Goroutine vs OS Thread
 
This is one of the most commonly asked Go interview questions.
 
| | OS Thread | Goroutine |
|---|---|---|
| Managed by | Operating system | Go runtime |
| Initial stack size | ~1-2 MB (fixed) | ~2 KB (grows as needed) |
| Creation cost | Expensive | Very cheap |
| Typical scale | Thousands | Hundreds of thousands+ |
| Switching cost | Expensive (kernel involved) | Cheap (handled by Go scheduler) |
 
**Why goroutines are "lightweight":** an OS thread is created and managed by the operating system, with a large fixed stack. A goroutine is created and managed by the Go runtime, starting with a tiny stack that grows or shrinks as needed — making it dramatically cheaper to create and switch between.

### The GMP Model (High-Level)
 
Go doesn't give each goroutine its own OS thread. Instead, it uses a scheduler with three pieces:
- **G** = Goroutine (the actual task)
- **M** = Machine (an OS thread)
- **P** = Processor (a context that manages scheduling)
Many **G**s get multiplexed onto a much smaller number of **M**s, coordinated by **P**s — an **M:N model** (many goroutines, few OS threads), instead of a **1:1 model** (one thread per task).

### Comparison with Java Threads
 
Traditional Java (`new Thread()`) uses a **1:1 model** — every Java thread maps directly to one real OS thread, inheriting all its costs (large stack, expensive creation).
 
| | Java (traditional `Thread`) | Go (`goroutine`) |
|---|---|---|
| Mapping to OS threads | 1:1 | M:N (many share few) |
| Managed by | Operating system | Go runtime |
| Stack size | ~1 MB fixed | ~2 KB, grows as needed |
 
**Note:** Java 21 (2023) introduced **virtual threads** (Project Loom), which behave more like goroutines — lightweight and JVM-managed, not directly 1:1 with OS threads. Go was designed this way from the start; Java added it later as a new feature.
 
---

## When to Use
 
- **Goroutines:** any time you want a task to run without blocking the rest of the program (background work, handling multiple requests, running independent steps concurrently)
- **`sync.WaitGroup`:** whenever `main()` (or any function) needs to wait for a known number of goroutines to finish before continuing
- **Anonymous goroutines:** quick, one-off concurrent tasks that don't need a separately named function
- **Loop variable fixes:** any time goroutines are launched inside a loop and need the correct per-iteration value (always double-check the Go version in use)

---

## Real-World / Project Usage
 
1. **HTTP servers** — Go's `net/http` automatically runs each incoming request in its own goroutine, handling many clients concurrently
2. **Background jobs** — sending an email or writing a log without making the main request wait for it to finish
3. **Fan-out processing** — launching a goroutine per item (e.g., per file, per API call) to process a batch faster
4. **Graceful shutdown** — using `WaitGroup` to make sure all in-flight goroutines finish cleanly before the program exits
5. **Worker pools** — a fixed number of goroutines pulling work from a shared queue (builds on this same `Add`/`Done`/`Wait` pattern)

---

## Interview Questions This Covers
 
- What is a goroutine, and how is it different from an OS thread?
- What is the difference between concurrency and parallelism?
- Why does a program exit before a goroutine finishes if `main()` returns first?
- What does `sync.WaitGroup` do, and what are `Add`, `Done`, and `Wait` each responsible for?
- What happens if `Add()` and `Done()` counts don't match? (deadlock vs panic)
- What is the loop variable capture bug, and how was it fixed differently before and after Go 1.22?
- How do you pass the current loop variable safely into a goroutine?

---

## 💡 Memory Points
 
1. A goroutine is a function that runs independently, launched with the `go` keyword
2. Concurrency = managing multiple tasks by switching between them; parallelism = doing them at the same instant (needs multiple cores)
3. Go goroutines always give concurrency; parallelism depends on the machine's CPU cores
4. When `main()` finishes, the whole program exits immediately — running goroutines get cut off, not waited for
5. `time.Sleep()` is an unreliable way to "wait" for a goroutine — it guesses a duration instead of waiting for a real signal
6. `sync.WaitGroup` tracks a counter: `Add(n)` increases it, `Done()` decreases it, `Wait()` blocks until it hits zero
7. `Add(n)` must exactly match the number of `Done()` calls that will actually happen
8. Too many `Add()`s relative to `Done()`s → deadlock (`Wait()` hangs forever)
9. Too many `Done()`s relative to `Add()`s → panic (negative WaitGroup counter)
10. Forgetting the `go` keyword just runs the function normally — no concurrency happens at all
11. An anonymous goroutine is an inline, unnamed function launched immediately: `go func() { ... }()`
12. Pre-Go 1.22, loop variables were shared across iterations, so goroutines could all read the same (final) value
13. Go 1.22+ gives each loop iteration its own private copy of the loop variable automatically, removing this bug by default
14. Old-style fixes still worth knowing: shadow with `i := i`, or pass `i` as a function argument to the goroutine

 
 


 