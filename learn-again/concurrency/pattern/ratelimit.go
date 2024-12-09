package pattern

import (
	"context"
	"golang.org/x/time/rate"
	"log"
	"os"
	"sort"
	"sync"
	"time"
)

type RateLimiter interface {
	Wait(ctx context.Context) error
	Limit() rate.Limit
}

type MultiLimiter struct {
	limiters []RateLimiter
}

func NewMultiLimiter(limiters ...RateLimiter) *MultiLimiter {
	byLimit := func(i, j int) bool {
		return limiters[i].Limit() < limiters[j].Limit()
	}
	sort.Slice(limiters, byLimit)
	//for _, l := range limiters {
	//	fmt.Printf("%v ", l.Limit())
	//}
	//fmt.Println()
	return &MultiLimiter{limiters: limiters}
}

func (ml *MultiLimiter) Wait(ctx context.Context) error {
	for _, l := range ml.limiters {
		if err := l.Wait(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (ml *MultiLimiter) Limit() rate.Limit {
	return ml.limiters[0].Limit()
}

func Per(eventCount int, duration time.Duration) rate.Limit {
	return rate.Every(duration / time.Duration(eventCount))
}

func open() *APIConnection {
	//apiLimit := NewMultiLimiter(
	//	rate.NewLimiter(rate.Every(time.Second/100), 2),
	//	rate.NewLimiter(rate.Every(time.Minute/10), 10),
	//)
	//networkLimit := NewMultiLimiter(rate.NewLimiter(rate.Every(time.Second/3), 3))
	//discLimit := NewMultiLimiter(rate.NewLimiter(rate.Every(time.Second), 1))

	//secondLimit := rate.NewLimiter(rate.Limit(rate.Every(time.Second/2)), 1)

	apiLimit := NewMultiLimiter(
		rate.NewLimiter(Per(4, time.Second), 4),
		rate.NewLimiter(Per(10, time.Minute), 10),
	)
	diskLimit := NewMultiLimiter(
		rate.NewLimiter(rate.Limit(1), 1),
	)
	networkLimit := NewMultiLimiter(
		rate.NewLimiter(Per(3, time.Second), 3),
	)
	return &APIConnection{

		fileLimit:    NewMultiLimiter(apiLimit, diskLimit),
		addressLimit: NewMultiLimiter(apiLimit, networkLimit),
	}
}

type APIConnection struct {
	//diskLimit    RateLimiter
	//apiLimit     RateLimiter
	//networkLimit RateLimiter

	addressLimit,
	fileLimit RateLimiter
}

func (c *APIConnection) ResolveFile(ctx context.Context) error {
	if err := c.fileLimit.Wait(ctx); err != nil {
		return err
	}

	return nil
}

func (c *APIConnection) ResolveAddress(ctx context.Context) error {
	if err := c.addressLimit.Wait(ctx); err != nil {
		return err
	}

	return nil
}

func rateLimit() {
	defer log.Println("Done")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	var wg sync.WaitGroup
	apiConnection := open()
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			if err := apiConnection.ResolveFile(context.Background()); err != nil {
				log.Println("can't read file")
			}
			log.Println("read file")
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			if err := apiConnection.ResolveAddress(context.Background()); err != nil {
				log.Println("can't reach network")
			}
			log.Println("resolve address")
		}()
	}

	wg.Wait()
}
