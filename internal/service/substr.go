package service

type SubstrService struct {
}

func NewSubstrService() *SubstrService {
	return &SubstrService{}
}

func (s *SubstrService) LongestSubstrFind(str string) string {
	var n = len(str)
	var i = 0
	var st int
	var curlen int
	var maxlen int
	var start int
	pos := make(map[byte]int)
	pos[str[i]] = 0
	for i = 1; i < n; i++ {
		if _, ok := pos[str[i]]; !ok {
			pos[str[i]] = i
		} else {
			if pos[str[i]] >= st {
				curlen = i - st
				if maxlen < curlen {
					maxlen = curlen
					start = st
				}
				st = pos[str[i]] + 1
			}
			pos[str[i]] = i
		}
	}
	if maxlen < i-st {
		maxlen = i - st
		start = st
	}
	return str[start : start+maxlen]
}
