package LamportClock

type LamportClock struct {
	internalClock int32
}

func (clock *LamportClock) Initialize() {
	clock.internalClock = 0
}

func (clock *LamportClock) Increment() {
	clock.internalClock += 1
}

func (receiver *LamportClock) SyncTime(senderTime int32) {
	receiver.internalClock = max(senderTime, receiver.internalClock) + 1
}

func (clock *LamportClock) GetTime() int32 {
	return clock.internalClock
}

func max(i int32, t int32) int32 {
	if i > t {
		return i
	} else {
		return t
	}
}
