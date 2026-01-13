Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
=======
> *Concurrency means that more then one task is working, but not nessesery at the same time. Parallelism means at the same time.*
>>>>>>> ef6e22e96c1cb7975e7fc54edb7b201872d255c2

What is the difference between a *race condition* and a *data race*? 
> *Data race are when more then one thread acsess the same memory at the same time. Race condition is a logic problem, like delay form the logick gates leads to en unsertenty in bit flips. This ilustration is very nice: https://en.wikipedia.org/wiki/Race_condition#/media/File:Race_condition.svg* 
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
> *A scheduler desides witch thread who gets to go on the CPU at what time.* 


### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
> *Threads are used ehrn the program needs to do more then one thing at the same time. Threads makes it posibale to held the system "active" still when other things is happening. Use the CPU-kjerner more effektivt.*

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> *Fibers and coroutines are easier to use, threads are styrt form the OS. They are often used beacuse they have shorter overhead then threads, they are cheaper and easier to controll*

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> *Both, it make it easer to reflect real life problems and better respons and ytelse. In the same time it makes it more diffecoult becuse it comes with complex things like race conditions, locks and syncronicing. This makes it less generall*

What do you think is best - *shared variables* or *message passing*?
> *Message pasing is more robust and easier to read. Shared can be faster and is often used ehen u have many threads, and ehrn u work close to the embedded system.*



