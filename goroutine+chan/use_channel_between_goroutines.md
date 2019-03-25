The Nature Of Channels In Go

Feb 17, 2014

Introduction
In my last post called Concurrency, Goroutines and GOMAXPROCS, I set the stage for talking about channels. We discussed what concurrency was and how goroutines played a role. With that foundation in hand, we can now understand the nature of channels and how they can be used to synchronize goroutines to share resources in a safe, less error prone and fun way.

What Are Channels
Channels are type safe message queues that have the intelligence to control the behavior of any goroutine attempting to receive or send on it. A channel acts as a conduit between two goroutines and will synchronize the exchange of any resource that is passed through it. It is the channel’s ability to control the goroutines interaction that creates the synchronization mechanism. When a channel is created with no capacity, it is called an unbuffered channel. In turn, a channel created with capacity is called a buffered channel.

To understand what the synchronization behavior will be for any goroutine interacting with a channel, we need to know the type and state of the channel. The scenarios are a bit different depending on whether we are using an unbuffered or buffered channel, so let’s talk about each one independently.

Unbuffered Channels
Unbuffered channels have no capacity and therefore require both goroutines to be ready to make any exchange. When a goroutine attempts to send a resource to an unbuffered channel and there is no goroutine waiting to receive the resource, the channel will lock the sending goroutine and make it wait. When a goroutine attempts to receive from an unbuffered channel, and there is no goroutine waiting to send a resource, the channel will lock the receiving goroutine and make it wait.


In the diagram above, we see an example of two goroutines making an exchange using an unbuffered channel. In step 1, the two goroutines approach the channel and then in step 2, the goroutine on the left sticks his hand into the channel or performs a send. At this point, that goroutine is locked in the channel until the exchange is complete. Then in step 3, the goroutine on the right places his hand into the channel or performs a receive. That goroutine is also locked in the channel until the exchange is complete. In step 4 and 5 the exchange is made and finally in step 6, both goroutines are free to remove their hands and go on their way.

Synchronization is inherent in the interaction between the send and the receive. One can not happen without the other. The nature of an unbuffered channel is guaranteed synchronization.
![p1](https://www.goinggo.net/images/goinggo/Screen+Shot+2014-02-16+at+10.10.54+AM.png)

Buffered Channels
Buffered channels have capacity and therefore can behave a bit differently. When a goroutine attempts to send a resource to a buffered channel and the channel is full, the channel will lock the goroutine and make it wait until a buffer becomes available. If there is room in the channel, the send can take place immediately and the goroutine can move on. When a goroutine attempts to receive from a buffered channel and the buffered channel is empty, the channel will lock the goroutine and make it wait until a resource has been sent.

![p1](https://www.goinggo.net/images/goinggo/Screen+Shot+2014-02-17+at+8.38.15+AM.png)
In the diagram above, we see an example of two goroutines adding and removing items from a buffered channel independently. In step 1, the goroutine on the right is removing a resource from the channel or performing a receive. In step 2, the goroutine on the right can remove the resource independent of the goroutine on the left adding a new resource to the channel. In step 3, both goroutines are adding and removing a resource from the channel at the same time and in step 4 both goroutines are done.

Synchronization still occurs within the interactions of receives and sends, however when the queue has buffer availability, the sends will not lock. Receives will not lock when there is something to receive from the channel. Consequently, if the buffer is full or if there is nothing to receive, a buffered channel will behave very much like an unbuffered channel.

Relay Race
If you have ever watched a track meet you may have seen a relay race. In a relay race there are four athletes who run around the track as fast as they can as a team. The key to the race is that only one runner per team can be running at a time. The runner with the baton is the only one allowed to run, and the exchange of the baton from runner to runner is critical to winning the race.

Let’s build a sample program that uses four goroutines and a channel to simulate a relay race. The goroutines will be the runners in the race and the channel will be used to exchanged the baton between each runner. This is a classic example of how resources can be passed between goroutines and how a channel controls the behavior of the goroutines that interact with it.
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Create an unbuffered channel
    baton := make(chan int)

    // First runner to his mark
    go Runner(baton)

    // Start the race
    baton <- 1

    // Give the runners time to race
    time.Sleep(500 * time.Millisecond)
}

func Runner(baton chan int) {
    var newRunner int

    // Wait to receive the baton
    runner := <-baton

    // Start running around the track
    fmt.Printf("Runner %d Running With Baton\n", runner)

    // New runner to the line
    if runner != 4 {
        newRunner = runner + 1
        fmt.Printf("Runner %d To The Line\n", newRunner)
        go Runner(baton)
    }

    // Running around the track
    time.Sleep(100 * time.Millisecond)

    // Is the race over
    if runner == 4 {
        fmt.Printf("Runner %d Finished, Race Over\n", runner)
        return
    }

    // Exchange the baton for the next runner
    fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)
    baton <- newRunner
}
```
When we run the sample program we get the following output:

Runner 1 Running With Baton          
Runner 2 To The Line      
Runner 1 Exchange With Runner 2      
Runner 2 Running With Baton      
Runner 3 To The Line      
Runner 2 Exchange With Runner 3      
Runner 3 Running With Baton      
Runner 4 To The Line      
Runner 3 Exchange With Runner 4      
Runner 4 Running With Baton      
Runner 4 Finished, Race Over      

The program starts out creating an unbuffered channel:

// Create an unbuffered channel
baton := make(chan int)

Using an unbuffered channel forces the goroutines to be ready at the same time to make the exchange of the baton. This need for both goroutines to be ready creates the guaranteed synchronization.

If we look at the rest of the main function, we see a goroutine created for the first runner in the race and then the baton is handed off to that runner. The baton in this example is an integer value that is being passed between each runner. The sample is using a sleep to let the race complete before main terminates and ends the program:

// Create an unbuffered channel
baton := make(chan int)

// First runner to his mark
go Runner(baton)

// Start the race
baton <- 1

// Give the runners time to race
time.Sleep(500 * time.Millisecond)

If we just focus on the core parts of the Runner function, we can see how the baton exchange takes place until the race is over. The Runner function is launched as a goroutine for each runner in the race. Every time a new goroutine is launched, the channel is passed into the goroutine. The channel is the conduit for the exchange, so the current runner and the one waiting to go next need to reference the channel:

func Runner(baton chan int)

The first thing each runner does is wait for the baton exchange. That is simulated with the receive on the channel. The receive immediately locks the goroutine until the baton is sent into the channel. Once the baton is send into the channel, the receive will release and the goroutine will simulate the next runner sprinting down the track. If the fourth runner is running, no new runner will enter the race. If we are still in the middle of the race, a new goroutine for the next runner is launched.

// Wait to receive the baton
runner := <-baton

// New runner to the line
if runner != 4 {
    newRunner = runner + 1
    go Runner(baton)
}

Then we sleep to simulate some time it takes for the runner to run around the track. If this is the fourth runner, the goroutine terminates after the sleep and the race is complete. If not, the baton exchange takes place with the send into the channel. There is a goroutine already locked and waiting for this exchange. As soon as the baton is sent into the channel, the exchange is made and the race continue:

// Running around the track
time.Sleep(100 * time.Millisecond)

// Is the race over
if runner == 4 {
    return
}

// Exchange the baton for the next runner
baton <- newRunner

Conclusion
The example showcases a real world event, a relay race between runners, being implemented in a way that mimics the actual events. This is one of the beautiful things about channels. The code flows in a way that simulates how these types of exchanges can happen in the real world.

Now that we have an understanding of the nature of unbuffered and buffered channels, we can look at different concurrency patterns we can implement using channels. Concurrency patterns allow us to implement more complex exchanges between goroutines that simulate real world computing problems like semaphores, generators and multiplexers.


