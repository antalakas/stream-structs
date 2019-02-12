package stream_structs

import (
	"github.com/go-redis/redis"
	"github.com/quipo/statsd"
)

// Counter struct
type Counter struct {
	dbM         *redis.Client
	stats       *statsd.StatsdBuffer
	counterChan chan<- *CounterMessage
}

// NewCounter function
func NewCounter(dbM *redis.Client, s *statsd.StatsdBuffer,
	counterChan chan<- *CounterMessage) *Counter {
	var counter Counter
	counter.dbM = dbM
	counter.stats = s
	counter.counterChan = counterChan

	return &counter
}

// ThreatIQCounter struct
type ThreatIQCounter struct {
	basicRedisKey     	string
	discard5m			bool
	basic5mRedisKey     string
	stats             	*statsd.StatsdBuffer
	theEvent 	 	  	CounterCollationEvent
	counterChan       	chan<- *ThreatIQCounterMessage
	uniqueCounterChan 	chan<- *ThreatIQCounterMessage
	counter5mChan       chan<- *ThreatIQCounterMessage
	uniqueCounter5mChan chan<- *ThreatIQCounterMessage
}

// NewThreatIQCounter function
func NewThreatIQCounter(
	s *statsd.StatsdBuffer,
	counterChan chan<- *ThreatIQCounterMessage,
	uniqueCounterChan chan<- *ThreatIQCounterMessage,
	counter5mChan chan<- *ThreatIQCounterMessage,
	uniqueCounter5mChan chan<- *ThreatIQCounterMessage) *ThreatIQCounter {

	var counter ThreatIQCounter
	counter.stats = s
	counter.counterChan = counterChan
	counter.uniqueCounterChan = uniqueCounterChan
	counter.counter5mChan = counter5mChan
	counter.uniqueCounter5mChan = uniqueCounter5mChan
	return &counter
}
