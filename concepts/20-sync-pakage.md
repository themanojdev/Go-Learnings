# 🔒 Sync Package: Mutex, RWMutex, Once & Atomic

## What is it
Go's `sync` and `sync/atomic` packages provide tools to safely share data between goroutines. They protect against **race conditions** — situations where multiple goroutines access the same data at the same time, causing incorrect or unpredictable results.

**Key Points:**
- These tools protect **shared memory** (variables/structs multiple goroutines touch)
- A different family of tools — channels and context — solves synchronization by **communication** instead (see channels concept)
- Go's philosophy: *"Don't communicate by sharing memory; share memory by communicating"* — but locks are still needed and widely used in real code.

---

## The Problem: Race Conditions
 
```go
type Counter struct {
    Count int
}
 
func (c *Counter) Increment() {
    c.Count++
}
 
func main() {
    c := Counter{}
    var wg sync.WaitGroup
 
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            c.Increment()
        }()
    }
 
    wg.Wait()
    fmt.Println("Final count:", c.Count)
}
```

You'd expect `1000`, but often get less — like `987`. This is because `c.Count++` secretly involves **3 steps**: read the value, add 1, write it back. If two goroutines both read the same value before either writes, one increment gets silently lost:

```
Goroutine 1: read 5
Goroutine 2: read 5          <- both read the SAME value before either writes
Goroutine 1: add 1 → write 6
Goroutine 2: add 1 → write 6  <- overwrites with the same 6, one increment LOST
```

Running one goroutine at a time avoids this — each increment fully completes (read, add, write) before the next starts. Concurrency introduces the overlap that causes the bug.
 
---

## Method 1: `sync.Mutex` — One Goroutine at a Time
A `Mutex` (mutual exclusion lock) makes sure only one goroutine can run a critical section of code at a time.
 
```go
type Counter struct {
    mu    sync.Mutex
    Count int
}
 
func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.Count++
}
```

**How it works:** `Lock()` means "I'm entering, nobody else can come in until I `Unlock()`." Any other goroutine calling `Lock()` while it's held simply waits.

**Output (now guaranteed correct):**
```
Final count: 1000
```

**Always use `defer` for `Unlock()`** — this guarantees the lock is released even if the function returns early or panics. Forgetting `Unlock()` entirely leaves the lock held forever: every other goroutine waiting on it blocks forever too. If something (like `main()`) is waiting on those stuck goroutines, Go detects it as a full deadlock (`fatal error: all goroutines are asleep - deadlock!`); if nothing is waiting on them, they just leak silently in the background.
 
---

## Method 2: `sync.RWMutex` — Many Readers, One Writer
Reading data doesn't change it, so multiple goroutines reading at the same time is always safe. `RWMutex` takes advantage of this: it allows many simultaneous readers, but still only one writer at a time — and while a writer is active, everyone (readers and writers) must wait.
 
```go
type Config struct {
    mu    sync.RWMutex
    value string
}
 
func (c *Config) Get() string {
    c.mu.RLock()         // read lock - many goroutines can hold this at once
    defer c.mu.RUnlock()
    return c.value
}
 
func (c *Config) Set(newValue string) {
    c.mu.Lock()          // write lock - only one at a time, blocks everyone
    defer c.mu.Unlock()
    c.value = newValue
}
```

**Simple picture:** like a notice board — many people can look at it at once, but while someone is changing it, nobody (not even other readers) should look, or they might catch a half-written, inconsistent message.

### `Mutex` vs `RWMutex`
 
| | `Mutex` | `RWMutex` |
|---|---|---|
| Use when | Roughly equal reads/writes, or unsure | Reads far outnumber writes |
| Readers block each other? | Yes | No |
| Writers block everyone? | Yes | Yes |
 
**Rule of thumb:** start with `Mutex`. Switch to `RWMutex` only when something is read constantly but written rarely — a config, a cache, a lookup table.
 
---

## Method 3: `sync.Once` — Run Exactly One Time
Some setup — connecting to a database, loading a config, initializing a logger — should run **exactly once**, even if many goroutines might trigger it concurrently.
 
```go
var once sync.Once
 
func setupDatabase() {
    fmt.Println("Connecting to database...")
}
 
func main() {
    var wg sync.WaitGroup
 
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            once.Do(setupDatabase)
        }()
    }
 
    wg.Wait()
}
```
 
**Output (only prints once, no matter how many goroutines call it):**
```
Connecting to database...
```

**Why not just use a plain `if !initialized` check?** That's unsafe — two goroutines could both pass the check before either sets `initialized = true`, and both would run the setup. This is the exact same kind of race condition as `Counter.Increment()`. `sync.Once` handles this correctly under the hood.
 
---

## Practical Code Example — Database Connection Pool (Real-World `sync.Once`)
 
```go
package main
 
import (
    "database/sql"
    "fmt"
    "sync"
 
    _ "github.com/lib/pq"
)
 
var (
    db   *sql.DB
    once sync.Once
)
 
func getDB() *sql.DB {
    once.Do(func() {
        fmt.Println("Opening database connection pool...")
        var err error
        db, err = sql.Open("postgres", "your-connection-string")
        if err != nil {
            panic(err)
        }
    })
    return db
}
 
func main() {
    var wg sync.WaitGroup
 
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            database := getDB()
            fmt.Println("Got DB connection:", database != nil)
        }()
    }
 
    wg.Wait()
}
```
 
**Output:**
```
Opening database connection pool...
Got DB connection: true
Got DB connection: true
Got DB connection: true
Got DB connection: true
Got DB connection: true
```

The connection pool is created only once, regardless of how many goroutines call `getDB()` concurrently.
 
---

## Method 4: `sync/atomic` — A Lighter Alternative for Simple Values
For simple values (integers, booleans, pointers), `atomic` operations are a faster alternative to `Mutex`. An atomic operation happens as one indivisible action — no other goroutine can ever catch it half-done.
 
```go
import "sync/atomic"
 
type Counter struct {
    Count atomic.Int64
}
 
func (c *Counter) Increment() {
    c.Count.Add(1)  // atomic add - no lock needed
}
 
func main() {
    c := Counter{}
    var wg sync.WaitGroup
 
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            c.Increment()
        }()
    }
 
    wg.Wait()
    fmt.Println("Final count:", c.Count.Load())  // .Load() reads safely too
}
```
 
**Output (guaranteed correct, every time):**
```
Final count: 1000
```

### The Trap: Atomic Doesn't Compose Across Multiple Fields
If two related fields need to update **together consistently** — e.g. `Count` and `LastUpdated` — using two separate atomic operations still leaves a gap between them where a reader could catch a half-updated state:
 
```go
c.Count.Add(1)           // atomic - safe alone
c.LastUpdated.Store(now) // atomic - safe alone, but a reader could catch the GAP between these two lines
```
 
**The fix:** use `Mutex` instead, so both fields update together as one locked block:
 
```go
type Counter struct {
    mu          sync.Mutex
    Count       int
    LastUpdated time.Time
}
 
func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.Count++
    c.LastUpdated = time.Now()
}
```

**Rule:** use `atomic` only for a single, standalone value. The moment two or more related fields must change together consistently, use `Mutex` instead.

---

## `Mutex` vs `atomic` — When to Use Which
 
| | `Mutex` | `atomic` |
|---|---|---|
| Protects | Any code, any complexity (multiple lines/fields) | A single simple value only |
| Performance | Slightly slower (lock/unlock overhead) | Faster — no locking involved |
| Use case | Structs with multiple related fields | A single counter, flag, or ID |
 
---

## Full Comparison — The Sync Toolkit
 
| Tool | Problem It Solves |
|---|---|
| `sync.WaitGroup` | Waiting for goroutines to finish |
| `sync.Mutex` | Only one goroutine touches shared data at a time |
| `sync.RWMutex` | Same as Mutex, but allows many simultaneous readers |
| `sync.Once` | Run something exactly once, no matter how many goroutines try |
| `sync/atomic` | Fast, lock-free updates for a single simple value |

---

## When to Use
 
- **`Mutex`:** protecting a critical section involving one or more fields, especially with more complex logic
- **`RWMutex`:** data read constantly but written rarely (config, cache, lookup table)
- **`Once`:** one-time setup that might be triggered by concurrent goroutines (DB connection pool, config load, logger init)
- **`atomic`:** a single, standalone value (counter, flag) needing fast, safe updates with no related fields

---

## Interview Questions This Covers
 
- What is a race condition, and how does `c.Count++` cause one under concurrency?
- What does `sync.Mutex` do, and why must `Unlock()` always be deferred?
- What happens if you forget to call `Unlock()`?
- What is the difference between `Mutex` and `RWMutex`, and when would you choose each?
- Why is it safe for multiple goroutines to read concurrently, but not to write concurrently?
- What problem does `sync.Once` solve, and why can't a simple `if !initialized` check replace it?
- What is an atomic operation, and how is it different from using a `Mutex`?
- Why doesn't `atomic` work safely across multiple related fields?

## 💡 Memory Points
 
1. A race condition happens when multiple goroutines access/modify shared data at the same time without protection
2. `c.Count++` is not one step — it's read, add, write — and overlapping these steps across goroutines loses updates
3. `sync.Mutex` allows only one goroutine into a locked section at a time; always `defer Unlock()`
4. Forgetting `Unlock()` leaves the lock held forever — leads to deadlock (if something waits on it) or a silent goroutine leak (if nothing does)
5. Reading data concurrently is always safe; writing concurrently is not, because writing changes data mid-read for others
6. `RWMutex` allows many simultaneous readers (`RLock`/`RUnlock`) but only one writer at a time (`Lock`/`Unlock`), which blocks everyone
7. Use `RWMutex` when reads vastly outnumber writes; otherwise plain `Mutex` is simpler and usually sufficient
8. `sync.Once` guarantees a function runs exactly once, no matter how many goroutines call it concurrently
9. A manual `if !initialized` check is unsafe — multiple goroutines can pass the check simultaneously before either sets the flag
10. `sync/atomic` provides lock-free, indivisible operations for single simple values (integers, booleans, pointers)
11. Use `.Add()` to atomically update and `.Load()` to atomically read a value
12. Atomic operations do NOT compose safely across multiple related fields — a gap between two atomic calls can expose inconsistent state
13. When multiple related fields must update together consistently, use `Mutex`, not multiple separate atomic operations
14. `go run -race` is Go's built-in race detector, used during development/testing to catch these bugs before production