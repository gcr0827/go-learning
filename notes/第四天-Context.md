# Context

## 1. `ctx.Done()` 什么时候返回

`ctx.Done()` 返回的 channel 在以下情况会被关闭（从而 `<-ctx.Done()` 会返回）：

- **手动取消**：调用了 `cancel()`（WithCancel / WithTimeout / WithDeadline 返回的 cancel 函数）
- **超时**：WithTimeout 设置的时长到了
- **截止时间到达**：WithDeadline 设置的 deadline 到了
- **父 context 被取消**：父 context 取消会级联取消所有子 context
- **进程结束**：`context.Background()` 和 `context.TODO()` 的 Done() 返回 nil，永远不会关闭

所以更精确的说法是：**当 context 被取消（无论是手动、超时、截止时间到、还是父级取消）时，Done 通道被关闭。**

------

## 2. WithTimeout 和 WithCancel 的区别

| 特性     | WithCancel                                   | WithTimeout                                                  |
| :------- | :------------------------------------------- | :----------------------------------------------------------- |
| 取消方式 | 只能手动调用 `cancel()`                      | 到达指定时长后**自动取消**，也可以手动调用 `cancel()` 提前取消 |
| 底层实现 | 直接创建可取消的 context                     | 本质是 `WithDeadline(parent, time.Now().Add(timeout))`       |
| 适用场景 | 需要手动控制取消时机（如收到信号、请求结束） | 限制一个操作的最长执行时间                                   |
| 返回值   | `(ctx, cancel)`                              | `(ctx, cancel)`，**cancel 仍然可用**                         |

**关键点补充**：WithTimeout 内部也是通过 WithDeadline 实现的，超时后会自动触发取消，但手动调用 `cancel()` 可以提前释放资源，两者都返回 cancel 函数用于提前取消。

------

## 3. 为什么 context 是协作式取消，调用方不检查 Done() 就不会被中断

**核心原因**：context 的取消机制只是**关闭一个 channel（信号通知）**，它本身不具备强制中断 goroutine 的能力。Go 语言没有提供从外部强制终止一个 goroutine 的机制（这是 Go 的设计哲学——goroutine 之间应该是协作的，而不是抢占式的）。