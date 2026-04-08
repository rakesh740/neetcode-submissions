type TimeMap struct {
m map[string][]data
}

type data struct {
	v     string
	t int
}


func Constructor() TimeMap {
	return TimeMap{
		m: make(map[string][]data),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], data{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	data := this.m[key]
	l, r := 0, len(data)-1
	for l <= r {
		m := (l+r) / 2

		if data[m].t <= timestamp {
			if m == len(data) - 1 || data[m+1].t > timestamp {
				return data[m].v
			}
			l = m + 1
		} else if data[m].t > timestamp {
			r = m - 1
		} 
	}

 return ""
}
