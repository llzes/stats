package stats

var standardCollector = New()

// Add adds a value to a key to the standard collector.
func Add(key string, change ValueType) ValueType {
	return standardCollector.Add(key, change)
}

// Del deletes a key to the standard collector.
func Del(key string) {
	standardCollector.Del(key)
}

// Get returns a value for a given key in the standard collector.
func Get(key string) ValueType {
	return standardCollector.Get(key)
}

// Set sets a value for a given key in the standard collector.
func Set(key string, val ValueType) {
	standardCollector.Set(key, val)
}

// Data returns a copy of the StatsCollector in the standard collector.
func Data() StatsType {
	return standardCollector.Data()
}

// Reset removes all key/value pairs from StatsCollector in the standard collector.
func Reset() {
	standardCollector.Reset()
}
