package main

import (
	"context"
	"io"
	"os"
	"sync"
	"time"
)

// greeter writes greeting messages every second
type greeter struct {
	writer io.StringWriter

	wg      sync.WaitGroup
	ctx     context.Context
	cancel  context.CancelFunc
	ticker  *time.Ticker
	tickerC <-chan time.Time
}

func newGreeter(writer io.StringWriter) *greeter {
	ticker := time.NewTicker(time.Second)
	ctx, cancel := context.WithCancel(context.Background())

	return &greeter{
		writer:  writer,
		ctx:     ctx,
		cancel:  cancel,
		ticker:  ticker,
		tickerC: ticker.C,
	}
}

// start allows to decouple the creation of greeter from the moment it starts working.
// This is useful for testing as we can safely mutate greeter fields.
func (g *greeter) start() {
	g.wg.Add(1)
	go func() {
		g.writeLoop()
		g.wg.Done()
	}()
}

func (g *greeter) writeLoop() {
	for {
		select {
		case <-g.ctx.Done():
			return
		case <-g.tickerC:
			_, _ = g.writer.WriteString("Hello world!\n")
		}
	}
}

func (g *greeter) stop() {
	g.ticker.Stop()
	g.cancel()
	g.wg.Wait()
}

func main() {
	greeter := newGreeter(os.Stdout)

	greeter.start()
	time.Sleep(5 * time.Second)
	greeter.stop()
}
