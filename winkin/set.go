package winkin

var (
	sharedValue = -1
)

// SetValue updates the shared value. (This should use a lock but does not.)
func SetValue(v int) {
	sharedValue = v
}

// GetValue returns the shared value. (This should use a lock but does not.)
func GetValue() int {
	return sharedValue
}
