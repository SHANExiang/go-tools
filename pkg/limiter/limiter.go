package main

import (
    "golang.org/x/time/rate"
    "log"
    "sync"
    "time"
)


type Limiters struct {
    limiters   *sync.Map
}

type Limiter struct {
    limiter      *rate.Limiter
    lastGet      time.Time
    key          string
}

var (
	globalLimiters = Limiters{limiters: &sync.Map{}}
	once           = sync.Once{}
)

func NewLimiter(rate rate.Limit, b int, key string) *Limiter {
    once.Do(func() {
        go globalLimiters.clearLimiter()
    })
    keyLimiter := globalLimiters.getLimiter(rate, b, key)
    return keyLimiter
}

func (l *Limiter) Allow() bool {
    l.lastGet = time.Now()
    return l.limiter.Allow()
}

func (l *Limiters) getLimiter(r rate.Limit, b int, key string) *Limiter {
    limiter, ok := l.limiters.Load(key)
    if ok {
        return limiter.(*Limiter)
    }
    lim := &Limiter{
        limiter: rate.NewLimiter(r, b),
        lastGet: time.Now(),
        key: key,
    }
    l.limiters.Store(key, lim)
    return lim
}

func (l *Limiters) clearLimiter() {
    for {
        time.Sleep(1 * time.Minute)
        l.limiters.Range(func(key, value any) bool {
            if time.Now().Unix() - value.(*Limiter).lastGet.Unix() > 60 {
                l.limiters.Delete(key)
            }
            return true
        })
    }
}

func useLimiter() {
    limiter := NewLimiter(rate.Every(10 * time.Millisecond), 100, "")
    if !limiter.Allow() {
        log.Println("Failed to write, your visits are too frequent, please try again later")
    } else {
        log.Println("Write success")
    }
}

func main() {
    l := NewLimiter(rate.Every(time.Millisecond * 31), 1, "152****86")
    for i := 0; i < 10; i++ {
        if l.Allow() {
            log.Println("success")
        } else {
            log.Println("您的访问过于频繁")
        }
        time.Sleep(time.Millisecond * 20)
    }
}
