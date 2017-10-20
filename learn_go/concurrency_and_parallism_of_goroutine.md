Concurrency, Goroutines and GOMAXPROCS

## go可以通过设置runtime.GOMAXPROCS(processors_num)来使用cpu多核实现真的并行!!默认为1.

Jan 29, 2014

Introduction
When new people join the Go-Miami group they always write that they want to learn more about Go’s concurrency model. Concurrency seems to be the big buzz word around the language. It was for me when I first started hearing about Go. It was Rob Pike’s Go Concurrency Patterns video that finally convinced me I needed to learn this language.

To understand how Go makes writing concurrent programs easier and less prone to errors, we first need to understand what a concurrent program is and the problems that result from such programs. I will not be talking about CSP (Communicating Sequential Processes) in this post, which is the basis for Go’s implementation of channels. This post will focus on what a concurrent program is, the role that goroutines play and how the GOMAXPROCS environment variable and runtime function affects the behavior of the Go runtime and the programs we write.

Processes and Threads
When we run an application, like the browser I am using to write this post, a process is created by the operating system for the application. The job of the process is to act like a container for all the resources the application uses and maintains as it runs. These resources include things like a memory address space, handles to files, devices and threads.

A thread is a path of execution that is scheduled by the operating system to execute the code we write in our functions against a processor. A process starts out with one thread, the main thread, and when that thread terminates the process terminates. This is because the main thread is the origin for the application. The main thread can then in turn launch more threads and those threads can launch even more threads.

The operating system schedules a thread to run on an available processor regardless of which process the thread belongs to. Each operating system has its own algorithms that make these decisions and it is best for us to write concurrent programs that are not specific to one algorithm or the other. Plus these algorithms change with every new release of an operating system, so it is dangerous game to play.

Goroutines and Parallelism
Any function or method in Go can be created as a goroutine. We can consider that the main function is executing as a goroutine, however the Go runtime does not start that goroutine. Goroutines are considered to be lightweight because they use little memory and resources plus their initial stack size is small. Prior to version 1.2 the stack size started at 4K and now as of version 1.4 it starts at 8K. The stack has the ability to grow as needed.

The operating system schedules threads to run against available processors and the Go runtime schedules goroutines to run within a logical processor that is bound to a single operating system thread. By default, the Go runtime allocates a single logical processor to execute all the goroutines that are created for our program. Even with this single logical processor and operating system thread, hundreds of thousands of goroutines can be scheduled to run concurrently with amazing efficiency and performance. It is not recommended to add more that one logical processor, but if you want to run goroutines in parallel, Go provides the ability to add more via the GOMAXPROCS environment variable or runtime function.

Concurrency is not Parallelism. Parallelism is when two or more threads are executing code simultaneously against different processors. If you configure the runtime to use more than one logical processor, the scheduler will distribute goroutines between these logical processors which will result in goroutines running on different operating system threads. However, to have true parallelism you need to run your program on a machine with multiple physical processors. If not, then the goroutines will be running concurrently against a single physical processor, even though the Go runtime is using multiple logical processors.

Concurrency Example
Let’s build a small program that shows Go running goroutines concurrently. In this example we are running the code with one logical processor:

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
)

func main() {
    runtime.GOMAXPROCS(1)

    var wg sync.WaitGroup
    wg.Add(2)

    fmt.Println("Starting Go Routines")
    go func() {
        defer wg.Done()

        for char := ‘a’; char < ‘a’+26; char++ {
            fmt.Printf("%c ", char)
        }
    }()

    go func() {
        defer wg.Done()

        for number := 1; number < 27; number++ {
            fmt.Printf("%d ", number)
        }
    }()

    fmt.Println("Waiting To Finish")
    wg.Wait()

    fmt.Println("\nTerminating Program")
}
```

This program launches two goroutines by using the keyword go and declaring two anonymous functions. The first goroutine displays the english alphabet using lowercase letters and the second goroutine displays numbers 1 through 26. When we run this program we get the following output:

Starting Go Routines
Waiting To Finish   
a b c d e f g h i j k l m n o p q r s t u v w x y z 1 2 3 4 5 6 7 8 9 10 11   
12 13 14 15 16 17 18 19 20 21 22 23 24 25 26   
Terminating Program   

When we look at the output we can see that the code was run concurrently. Once the two goroutines are launched, the main goroutine waits for the goroutines to complete. We need to do this because once the main goroutine terminates, the program terminates. Using a WaitGroup is a great way for goroutines to communicate when they are done.

We can see that the first goroutine completes displaying all 26 letters and then the second goroutine gets a turn to display all 26 numbers. Because it takes less than a microsecond for the first goroutine to complete its work, we don’t see the scheduler interrupt the first goroutine before it finishes its work. We can give a reason to the scheduler to swap the goroutines by putting a sleep into the first goroutine:

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func main() {
    runtime.GOMAXPROCS(1)

    var wg sync.WaitGroup
    wg.Add(2)

    fmt.Println("Starting Go Routines")
    go func() {
        defer wg.Done()

        time.Sleep(1 * time.Microsecond)
        for char := ‘a’; char < ‘a’+26; char++ {
            fmt.Printf("%c ", char)
        }
    }()

    go func() {
        defer wg.Done()

        for number := 1; number < 27; number++ {
            fmt.Printf("%d ", number)
        }
    }()

    fmt.Println("Waiting To Finish")
    wg.Wait()

    fmt.Println("\nTerminating Program")
}
```

This time we add a sleep in the first goroutine as soon as it starts. Calling sleep causes the scheduler to swap the two goroutines:

Starting Go Routines   
Waiting To Finish   
1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 a   
b c d e f g h i j k l m n o p q r s t u v w x y z   
Terminating Program   

This time the numbers display first and then the letters. The sleep causes the scheduler to stop running the first goroutine and let the second goroutine do its thing.

Parallel Example
In our past two examples the goroutines were running concurrently, but not in parallel. Let’s make a change to the code to allow the goroutines to run in parallel. All we need to do is add a second logical processor to the scheduler to use two threads:
```go
package main

import (
    "fmt"
    "runtime"
    "sync"
)

func main() {
    runtime.GOMAXPROCS(2)

    var wg sync.WaitGroup
    wg.Add(2)

    fmt.Println("Starting Go Routines")
    go func() {
        defer wg.Done()

        for char := ‘a’; char < ‘a’+26; char++ {
            fmt.Printf("%c ", char)
        }
    }()

    go func() {
        defer wg.Done()

        for number := 1; number < 27; number++ {
            fmt.Printf("%d ", number)
        }
    }()

    fmt.Println("Waiting To Finish")
    wg.Wait()

    fmt.Println("\nTerminating Program")
}
```
Here is the output for the program:

Starting Go Routines    
Waiting To Finish    
a b 1 2 3 4 c d e f 5 g h 6 i 7 j 8 k 9 10 11 12 l m n o p q 13 r s 14    
t 15 u v 16 w 17 x y 18 z 19 20 21 22 23 24 25 26    
Terminating Program    

Every time we run the program we are going to get different results. The scheduler does not behave exactly the same for each and every run. We can see that the goroutines are truly running in parallel. Both goroutines start running immediately and you can see them both competing for standard out to display their results.

Conclusion
Just because we can add multiple logical processors for the scheduler to use doesn’t mean we should. There is a reason the Go team has set the defaults to the runtime the way they did. Especially the default for only using a single logical processor. Just know that arbitrarily adding logical processors and running goroutines in parallel will not necessarily provide better performance for your programs. Always profile and benchmark your programs and make sure the Go runtime configuration is only changed if absolutely required.

The problem with building concurrency into our applications is eventually our goroutines are going to attempt to access the same resources, possibly at the same time. Read and write operations against a shared resource must always be atomic. In other words reads and writes must happen by one goroutine at a time or else we create race conditions in our programs. To learn more about race conditions read my post.

Channels are the way in Go we write safe and elegant concurrent programs that eliminate race conditions and make writing concurrent programs fun again. Now that we know how goroutines work, are scheduled and can be made to run in parallel, channels are the next thing we need to learn.


