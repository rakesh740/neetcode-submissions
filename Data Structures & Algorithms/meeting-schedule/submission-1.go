/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

 type meets []Interval

 func(m meets) Len()  int {
 return len(m)
 }
  func(m meets) Less(i, j int)  bool {
 return m[i].start < m[j].start
 }
  func(m meets) Swap(i, j int)  {
	m[i], m[j] = m[j], m[i]
 }

func canAttendMeetings(intervals []Interval) bool {

	// sort intervals by start time
	sort.Sort(meets(intervals))

	for i := 1;i < len(intervals) ; i++{
		if intervals[i-1].end > intervals[i].start {
			return false
		}
	}

	return true
}
