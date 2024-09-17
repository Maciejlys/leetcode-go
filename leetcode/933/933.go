package leetcode

type RecentCounter struct {
	pings []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		pings: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.pings = append(this.pings, t)
	this.removeOld(t)

	return len(this.pings)
}

func (this *RecentCounter) removeOld(t int) {
	if len(this.pings) == 0 || this.pings[0] >= t-3000 {
		return
	}
	this.pings = this.pings[1:]
	this.removeOld(t)
}
