# List down the main features of Golang ?
- Strong typing and garbage collection
- Concurrency support through goroutines and channels
- Fast compilation and execution
- Static linking of dependencies
- Built-in support for creating efficient network applications
- Cross-platform compatibility

# What is goroutine in Golang?
- Goroutine is a lightweight, independently executing thread in Go.
- Goroutines allow you to execute functions concurrently without the overhead of traditional threads.
- They are managed by the **Go runtime** and are more efficient than OS threads.
- Goroutines are defined using the **go** keyword **followed by a function call**.

# Difference between function and method in Golang?
- Functions are **standalone blocks of code** that can be called with arguments.
- Methods are functions associated with a specific type and are invoked on instances of that type.

# Difference between threads and goroutine?
- **Threads** are managed by the **operating system**, while **goroutines** are managed by the **Go runtime**.
- Goroutines are **lightweight** compared to OS threads, allowing for thousands to be created without significant performance impact.
- 
- **Switching between goroutines is managed by the Go runtime scheduler**, 
- while thread context switching involves **more overhead**.

# How many ways we can communicate between multiple goroutine?
- **Channels:**  - **Synchronized** communication using **chan data types**.
- **Mutexes:** **Synchronization** using mutual exclusion lock**s (sync.Mutex)**.
- **Wait Groups:** **Coordination** and **synchronization** using **sync.WaitGroup**.
- **Atomic Operations:** Safe atomic operations using **sync/atomic** package.
- **Select Statement:** Used for waiting on multiple communication operations-


# Vardic Functions and usage?
- Variadic functions in Go accept a variable number of arguments of a specific type.
- They are defined using an ellipsis (...)

# Working of Interface in Go?
- An interface in Go defines a **set of method signatures**. 
- A type implicitly satisfies an interface if it implements all the methods declared by the interface.
- This **allows for polymorphism and abstract behavior** without explicitly stating a type relationship.

# How we manage dependencies in GO project?
- Go uses a tool **go mod** for managing dependencies. 
- You can create a **go.mod file** ==>  lists the required packages along with their versions.
- The **go tool will automatically download and manage** the specified dependencies.
  ```
    go mod init myproject
    go get github.com/gin-gonic/gin@v1.7.2
  ```
- This will create a go.mod file and download the gin framework v1.7.2 as a dependency.

# Difference between pointer receiver method and value receiver method?
- **Pointer Receiver:** 
  - Method operates on the actual value pointed to by the receiver.
  - Changes made are reflected outside the method.
  
- **Value Receiver:** 
  - Method operates on a copy of the receiver value.
  - Changes within the method do not affect the original value.

# What is the use of select statement and switch statements in go?
- **select statement:** 
  - Used to **wait on multiple channel operations**. 
  - It blocks until one of the cases can proceed.
  
- **switch statement:** 
  - Used for conditionally executing different code blocks based on the value of an expression.
  - In Go, switch cases don't fall through by default.