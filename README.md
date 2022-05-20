# clock

提供时钟接口，并提供设置全局时钟和获取全局时钟的当前时间。

- 提供基于系统时间的时钟（SystemClock）

- 提供基于NTP服务的时钟（NTPClock）

- 提供只能前进滚动的时钟（ForwardOnlyClock），即该时钟只允许往大调（递增），不能往小调（递减）

- 可设置全局时钟（只能设置一次，防止运行期间被修改）

## 用法

- 基于系统时钟（SystemClock）的只能前进滚动的时钟（ForwardOnlyClock）

```go
package main

import (
    "fmt"
    "time"

    "github.com/berkaroad/clock"
)

func main() {
    var startTs int64 = 1655431380 // 2022-06-17 10:03:00
    // 将系统时钟作为参照物
    var refClock clock.Clock = clock.NewSystemClock()
    // 设置 ForwardOnlyClock 为全局时钟
    clock.SetGlobal(clock.NewForwardOnlyClock(startTs, refClock))

    for i := 0; i < 10; i++ {
        <-time.After(time.Second)
        // 打印系统时钟、全局时钟
        fmt.Printf("system clock: %v, global clock: %v\n", refClock.Now(), clock.Now())
    }
}
```

- 基于NTP时钟（NTPClock）的只能前进滚动的时钟（ForwardOnlyClock）

```go
package main

import (
    "fmt"
    "time"

    "github.com/berkaroad/clock"
)

func main() {
    var startTs int64 = 1655431380 // 2022-06-17 10:03:00
    // 将NTP时钟作为参照物
    var refClock clock.Clock = clock.NewNTPClock([]string{"pool.ntp.org"}, ntp.QueryOptions{Timeout: time.Second})
    // 设置 ForwardOnlyClock 为全局时钟
    clock.SetGlobal(clock.NewForwardOnlyClock(startTs, refClock))

    for i := 0; i < 10; i++ {
        <-time.After(time.Second)
        // 打印ntp时钟、全局时钟
        fmt.Printf("ntp clock: %v, global clock: %v\n", refClock.Now(), clock.Now())
    }
}
```
