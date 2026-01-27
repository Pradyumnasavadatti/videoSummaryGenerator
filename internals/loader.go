package internals

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func StartWithContext(ctx context.Context, wgp *sync.WaitGroup,message string) {
	spinner := []rune{'|', '/', '-', '\\'}
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	defer wgp.Done()
	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\r")
			fmt.Println();
			return
		case <-ticker.C:
			fmt.Printf("\r%s %c", message, spinner[i%4])
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("\r%s %s", message," ")
			i++
		}
	}
}
