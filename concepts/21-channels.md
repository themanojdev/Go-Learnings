# 📡 Channels in Go

## What is it
A **channel** is how goroutines **talk to each other** safely. Think back to the goroutine problem: once you launch `go sayHello()`, that goroutine runs independently — you have no direct way to get data back from it or know when it's really done, except by guessing (`time.Sleep`) or counting (`WaitGroup`).

A channel solves this differently: it's like a **pipe** between goroutines. One goroutine can **send** a value into the pipe, and another can **receive** it out the other end.
 
**Key Points:**
- A channel is a typed pipe: `chan int`, `chan string`, `chan bool`, etc.
- One goroutine sends, another receives — this is how goroutines pass data safely
- Sending and receiving can **block** (pause) the goroutine until the other side is ready
- Channels are how Go avoids messy shared-memory bugs between goroutines

---

## Syntax
 
```go
ch := make(chan int)   // create a channel that carries int values
 
ch <- 5      // send 5 into the channel
value := <-ch // receive a value from the channel
```
 
The arrow `<-` always points in the direction data is flowing.

---

## Method 1: Unbuffered Channel — The Default
An unbuffered channel has **no storage space**. A send only completes once someone is ready to receive it — like handing something directly to another person's hand, not dropping it in a box.
 
```go
package main
 
import "fmt"
 
func main() {
    ch := make(chan string)
 
    go func() {
        ch <- "Hello from goroutine"
    }()
 
    msg := <-ch
    fmt.Println(msg)
}
```
 
**Output:**
```
Hello from goroutine
```

Here, the goroutine's `ch <- "..."` **pauses** until `main()` reaches `<-ch` to receive it. This waiting is exactly what replaces `time.Sleep()` — the channel naturally synchronizes the two goroutines.
 
---

## Method 2: Buffered Channel — Has Storage Space
A buffered channel **can hold a limited number of values** before blocking. Think of it like a small mailbox with a fixed number of slots — sending doesn't block until the mailbox is full.
 
```go
ch := make(chan int, 2)  // buffer size 2
 
ch <- 1  // doesn't block, 1 slot filled
ch <- 2  // doesn't block, 2 slots filled (buffer full now)
// ch <- 3  // this WOULD block - buffer is full, no room
```
 
```go
package main
 
import "fmt"
 
func main() {
    ch := make(chan int, 2)
 
    ch <- 1
    ch <- 2
 
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
```
 
**Output:**
```
1
2
```

---

## Method 3: Closing a Channel with `close()`
`close()` tells everyone: **"no more values are coming through this channel."** This is important when a receiver needs to know when to stop waiting.
 
```go
ch := make(chan int)
 
go func() {
    for i := 1; i <= 3; i++ {
        ch <- i
    }
    close(ch)  // signal: done sending
}()
 
for value := range ch {
    fmt.Println(value)
}
```
 
**Output:**
```
1
2
3
```

`range` over a channel automatically keeps receiving values **until the channel is closed** — then the loop ends cleanly on its own.
 
---

## The `value, ok` Pattern with Channels
Just like map lookups, receiving from a channel can tell you whether the channel is still open:
 
```go
value, ok := <-ch
if !ok {
    fmt.Println("Channel is closed, no more values")
}
```

`ok` is `false` only when the channel is **closed AND empty**. If it's closed but still has buffered values left, you'll keep getting those values with `ok = true` until they run out.
 
---

## Practical Code Example
 
```go
package main
 
import "fmt"
 
func worker(ch chan<- int) {
    for i := 1; i <= 3; i++ {
        ch <- i * 10
    }
    close(ch)
}
 
func main() {
    ch := make(chan int)
 
    go worker(ch)
 
    for value := range ch {
        fmt.Println("Received:", value)
    }
}
```
 
**Output:**
```
Received: 10
Received: 20
Received: 30
```
 
---

## Directional Channels (`chan<-` and `<-chan`)
You can restrict a channel to **only sending** or **only receiving** when passing it into a function — this makes your code's intent clear and prevents mistakes.
 
```go
func send(ch chan<- int) {  // this function can only SEND into ch
    ch <- 5
}
 
func receive(ch <-chan int) {  // this function can only RECEIVE from ch
    value := <-ch
    fmt.Println(value)
}
```
 
**How to remember the arrow direction:**
- `chan<- int` → arrow points **into** `chan` → send-only
- `<-chan int` → arrow points **out of** `chan` → receive-only

---

## Blocking Rules (Memorize This Table)
 
| Action | Unbuffered Channel | Buffered Channel (not full) | Buffered Channel (full) |
|---|---|---|---|
| Send (`ch <-`) | Blocks until received | Doesn't block | Blocks until space frees up |
| Receive (`<-ch`) | Blocks until sent | Doesn't block if data exists | N/A |
| Receive from closed channel | Returns zero value, `ok = false` | Returns remaining values, then zero value + `ok = false` | Same |
 
---

## Deadlock — A Classic Interview Trap
 
A **deadlock** happens when goroutines are stuck waiting on each other forever, and Go detects this and crashes the program.
 
```go
func main() {
    ch := make(chan int)
    ch <- 5  // blocks forever - nobody is receiving, and no goroutine to help
    fmt.Println(<-ch)
}
```
 
**Output:**
```
fatal error: all goroutines are asleep - deadlock!
```

**Why:** `ch <- 5` blocks, waiting for a receiver. But `main()` is stuck on that same line — it never reaches the `<-ch` receive line to unblock itself. Nothing else is running to receive it, so it waits forever.
 
**The fix:** always make sure sends and receives happen in **different goroutines**, or use a buffered channel with enough room.
 
---

## Nil Channel Behavior (Another Interview Trap)
 
A channel that's declared but never initialized with `make()` is `nil`. Sending or receiving on a `nil` channel **blocks forever** — it doesn't panic, it just hangs.
 
```go
var ch chan int  // nil channel, never made with make()
ch <- 5           // blocks forever, this goroutine will never proceed
```
 
This is a real bug source — forgetting `make()` silently creates a channel that looks valid but never works.
 
---
 
## When to Use
 
- **Unbuffered channel:** when you want strict handoff — sender and receiver must both be ready at the same moment (good for synchronization)
- **Buffered channel:** when the sender should be able to move on without waiting immediately (a small queue of pending work)
- **`close()`:** when a sender knows it's done producing values, and receivers need a clean way to know when to stop
- **Directional channels:** when passing a channel into a function that should only send or only receive, to prevent misuse
---
 
## Real-World / Project Usage
 
1. **Worker pools** — multiple goroutines pull jobs from a shared channel and send results back through another channel
2. **Pipelines** — one goroutine's output channel becomes the next goroutine's input channel, forming a processing chain
3. **Fan-in / fan-out** — multiple goroutines send into one channel (fan-in), or one channel's work is spread across multiple goroutines (fan-out)
4. **Signaling completion** — `chan struct{}` (the empty struct you already learned) is often used purely as a signal, with no actual data needed
5. **Graceful shutdown** — servers use a channel to signal "stop accepting new requests" across many goroutines at once
---
 
## Interview Questions This Covers
 
- What is a channel, and how does it help goroutines communicate safely?
- What's the difference between a buffered and unbuffered channel?
- What happens when you send/receive on a channel — when does it block?
- What does `close()` do, and what happens if you receive from a closed channel?
- What is a deadlock, and can you give an example?
- What happens when you send or receive on a `nil` channel?
- What are directional channels, and why use them?
---
 
## 💡 Memory Points
 
1. A channel is a typed pipe for goroutines to send and receive data safely
2. Create with `make(chan Type)`; buffered version: `make(chan Type, size)`
3. `ch <- value` sends; `value := <-ch` receives
4. Unbuffered channel: send blocks until someone receives (direct handoff)
5. Buffered channel: send only blocks once the buffer is full
6. `close(ch)` signals "no more values coming" — required for `range` loops to end cleanly
7. `value, ok := <-ch` — `ok` is `false` only when the channel is closed AND empty
8. `range` over a channel automatically stops when the channel is closed
9. `chan<- Type` = send-only channel; `<-chan Type` = receive-only channel
10. Deadlock = goroutines stuck waiting on each other forever; Go detects it and crashes with a fatal error
11. A `nil` channel (declared but never `make()`'d) blocks forever on send or receive — no panic, just a silent hang
12. Only the sender should `close()` a channel — never the receiver, and never close a channel twice (causes a panic)
13. Channels are Go's preferred way to share data between goroutines instead of using shared memory + locks directly