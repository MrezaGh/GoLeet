package pattern

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func ctx() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if e := printGreeting(ctx); e != nil {
			fmt.Printf("failed to print greet: %v\n", e)
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if e := printFarewel(ctx); e != nil {
			fmt.Printf("failed to print farewell: %v\n", e)
		}
	}()

	wg.Wait()

}

func printGreeting(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	//ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()
	switch locale, err := locale(ctx); {
	case err != nil:
		return err
	case locale == "EN/US":
		fmt.Printf("hello world!\n")
		return nil
	}

	return fmt.Errorf("unsupported locale")

}

func printFarewel(ctx context.Context) error {
	switch locale, err := locale(ctx); {
	case err != nil:
		return err
	case locale == "EN/US":
		fmt.Printf("goodbye world!\n")
		return nil
	}
	return fmt.Errorf("unsupported locale")
}

func locale(ctx context.Context) (string, error) {
	if deadline, ok := ctx.Deadline(); ok && deadline.UnixMilli() < time.Now().Add(1*time.Minute).UnixMilli() {
		return "", context.DeadlineExceeded
	}
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(1 * time.Minute):
		//case <-time.After(1 * time.Second):
	}
	return "EN/US", nil
}
