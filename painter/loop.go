package painter

import (
	"golang.org/x/exp/shiny/screen"
	"image"
	"sync"
)

// Receiver отримує текстуру, яка була підготовлена в результаті виконання команд у циклі подій.
type Receiver interface {
	Update(t screen.Texture)
}

// Loop реалізує цикл подій для формування текстури отриманої через виконання операцій отриманих з внутрішньої черги.
type Loop struct {
	Receiver Receiver

	next screen.Texture // текстура, яка зараз формується
	prev screen.Texture // текстура, яка була відправлення останнього разу у Receiver

	mq messageQueue

	stop    chan struct{}
	stopReq bool
}

var size = image.Pt(800, 800)

// Start запускає цикл подій. Цей метод потрібно запустити до того, як викликати на ньому будь-які інші методи.
func (l *Loop) Start(s screen.Screen) {
	l.next, _ = s.NewTexture(size)
	l.prev, _ = s.NewTexture(size)
	l.stop = make(chan struct{})
	l.mq = initMessageQueue()

	go func() {
		for !(l.stopReq && l.mq.empty()) {
			op := l.mq.pull()

			if ready := op.Do(l.next); ready {
				l.Receiver.Update(l.next)
				l.next, l.prev = l.prev, l.next
			}
		}

		close(l.stop)
	}()
}

// Post додає нову операцію у внутрішню чергу.
func (l *Loop) Post(op Operation) {
	l.mq.push(op)
}

// StopAndWait сигналізує про необхідність завершити цикл та блокується до моменту його повної зупинки.
func (l *Loop) StopAndWait() {
	l.mq.push(OperationFunc(func(t screen.Texture) {
		l.stopReq = true
	}))

	<-l.stop
}

type messageQueue struct {
	operations []Operation
	mutex      *sync.Mutex
	condition  *sync.Cond
}

func initMessageQueue() messageQueue {
	mutex := &sync.Mutex{}
	return messageQueue{
		mutex:      mutex,
		condition:  sync.NewCond(mutex),
		operations: make([]Operation, 0),
	}
}

func (mq *messageQueue) push(op Operation) {
	mq.mutex.Lock()
	defer mq.mutex.Unlock()

	mq.operations = append(mq.operations, op)
	mq.condition.Signal()
}

func (mq *messageQueue) pull() Operation {
	mq.mutex.Lock()
	defer mq.mutex.Unlock()

	for mq.empty() {
		mq.condition.Wait()
	}

	op := mq.operations[0]
	mq.operations = mq.operations[1:]
	return op
}

func (mq *messageQueue) len() int {
	return len(mq.operations)
}

func (mq *messageQueue) empty() bool {
	return mq.len() == 0
}
