# waiter
a simple golang package that simply blocks while waiting for a signal

## Usage
```
import (
  "fmt"
  
  "github.com/haukened/waiter"
)

func main() {
  ...
  waiter.WaitForSignal(os.Interrupt)
  fmt.Println("Done!")
}
```
