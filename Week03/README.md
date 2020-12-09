学习笔记
本周主要学习了goroutine、memory model、sync包、chan、context包等内容。

- goroutine：
协程之间应该通过通信来共享内存，而不应该通过加锁共享内存来通信。在并发时，尽量减少锁的使用，能提高程序的并发量。

- memory model：
cpu的缓存设计会引发可见性问题。golang的内存模型就是为了处理相关的问题，内存屏幕，还有指令重排等相关细节需要深入去研究。

- sync包：
copy on write的思想可以处理非常多的问题，如果存在一个读多写少的场景，可以使用atomic包，load加载读去数据。atomic基于cow思想实现。

- chan：
channel是golang并发编程的重要组件之一，channel的使用是必不可少的，完美的协程间通信组件。

- context:
context的显式传递虽然让传参看起来没那么优雅，但是也减少了编程过程中漏传、错用的情况。
context非常适合用于协程的上下文、元数据存储。使用context可以做一些协程的级联超时控制、级联退出等等。