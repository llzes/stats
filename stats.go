package stats

import "sync"

type ValueType float64
type StatsType map[string]ValueType

type StatsCollector struct {
	lock  sync.RWMutex
	stats StatsType
}

func New() *StatsCollector {
	return new(StatsCollector).Reset()
}

// Add adds a float value to a key in a collector.
func (stat *StatsCollector) Add(key string, change ValueType) (val ValueType) {
	stat.lock.Lock()
	val = stat.stats[key]
	val += change
	stat.stats[key] = val
	stat.lock.Unlock()
	return
}

// Del deletes a key.
func (stat *StatsCollector) Del(key string) {
	stat.lock.Lock()
	delete(stat.stats, key)
	stat.lock.Unlock()
}

// Get returns a value for a given key.
func (stat *StatsCollector) Get(key string) (val ValueType) {
	stat.lock.RLock()
	val = stat.stats[key]
	stat.lock.RUnlock()
	return
}

// Set sets a value for a given key.
func (stat *StatsCollector) Set(key string, val ValueType) {
	stat.lock.Lock()
	stat.stats[key] = val
	stat.lock.Unlock()
}

// Data returns a copy of StatsCollector.
func (stat *StatsCollector) Data() StatsType {
	copy := make(StatsType)
	stat.lock.RLock()
	for key, value := range stat.stats {
		copy[key] = value
	}
	stat.lock.RUnlock()
	return copy
}

// Reset removes all key/value pairs from StatsCollector.
func (stat *StatsCollector) Reset() {
	stat.lock.Lock()
	stat.stats = make(StatsType)
	stat.lock.Unlock()
}
