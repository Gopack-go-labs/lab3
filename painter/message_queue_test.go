package painter

import (
	"golang.org/x/exp/shiny/screen"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestMessageQueueLen(t *testing.T) {
	mq := initMessageQueue()
	if mq.len() != 0 {
		t.Error("Unexpected length of the queue")
	}
	mq.push(OperationFunc(func(screen.Texture) {}))
	if mq.len() != 1 {
		t.Error("Unexpected length of the queue")
	}
	mq.push(OperationFunc(func(screen.Texture) {}))
	if mq.len() != 2 {
		t.Error("Unexpected length of the queue")
	}
	mq.pull()
	if mq.len() != 1 {
		t.Error("Unexpected length of the queue")
	}
	mq.pull()
	if mq.len() != 0 {
		t.Error("Unexpected length of the queue")
	}
}

func TestWaitOnPullEmpty(t *testing.T) {
	mq := initMessageQueue()

	go func() {
		time.Sleep(100 * time.Millisecond)
		mq.push(OperationFunc(func(screen.Texture) {}))
	}()

	t1 := time.Now()
	mq.pull()
	t2 := time.Now()

	if t2.Sub(t1) < 100*time.Millisecond {
		t.Error("Unexpected pull time")
	}
}

func TestPushPull(t *testing.T) {
	mq := initMessageQueue()
	var ops []int
	var (
		op1 = OperationFunc(func(screen.Texture) { ops = append(ops, 1) })
		op2 = OperationFunc(func(screen.Texture) { ops = append(ops, 2) })
		op3 = OperationFunc(func(screen.Texture) { ops = append(ops, 3) })
	)
	mq.push(op1)
	mq.push(op2)
	mq.push(op3)

	mq.pull().Do(nil)
	mq.pull().Do(nil)
	mq.pull().Do(nil)

	if !reflect.DeepEqual(ops, []int{1, 2, 3}) {
		t.Error("Bad order:", ops)
	}
}

func TestConcurrency(t *testing.T) {
	mq := initMessageQueue()
	const N = 1000
	wg := sync.WaitGroup{}
	wg.Add(N)

	go func() {
		for i := 0; i < N; i++ {
			mq.push(OperationFunc(func(screen.Texture) {}))
		}
	}()

	go func() {
		for i := 0; i < N; i++ {
			mq.pull()
			wg.Done()
		}
	}()

	wg.Wait()
	if mq.len() != 0 {
		t.Error("Unexpected length of the queue")
	}
}
