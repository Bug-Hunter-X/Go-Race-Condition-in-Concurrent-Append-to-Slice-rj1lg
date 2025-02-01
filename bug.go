```Go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var data []int
        const numRoutines = 1000

        for i := 0; i < numRoutines; i++ {
                wg.Add(1)
                go func(index int) {
                        defer wg.Done()
                        data = append(data, index)
                }(i)
        }

        wg.Wait()
        fmt.Println("Length of data:", len(data))
}
```