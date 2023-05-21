package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var ints Ints
	for _, num := range i {
		if filter(num) {
			ints = append(ints, num)
		}
	}

	return ints
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var ints Ints
	for _, num := range i {
		if !filter(num) {
			ints = append(ints, num)
		}
	}

	return ints
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var lists Lists
	for _, list := range l {
		if filter(list) {
			lists = append(lists, list)
		}
	}

	return lists
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var strs Strings
	for _, str := range s {
		if filter(str) {
			strs = append(strs, str)
		}
	}

	return strs
}
