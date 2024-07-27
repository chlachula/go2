package main
import (
"fmt"
"time"
a "github.com/chlachula/go2/indexPngSlides"
)
func main() {
	defer func(start time.Time) {
		fmt.Printf("Elapsed time %s\n", time.Since(start))
	}(time.Now())
        
        fmt.Println(a.Hello)
}
