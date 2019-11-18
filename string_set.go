package main

type StringSet map[string]struct{}

func SetFrom(array []string) StringSet {
	set := make(StringSet)
	for _, item := range array {
		set[item] = struct{}{}
	}
	return set
}

func (set StringSet) ToArray() []string {
	keys := make([]string, 0, len(set))
	for key := range set {
		keys = append(keys, key)
	}
	return keys
}

func Intersection(set1 StringSet, set2 StringSet) StringSet {
	intersection := make(StringSet)
	for item := range set1 {
		if _, ok := set2[item]; ok {
			intersection[item] = struct{}{}
		}
	}
	return intersection
}
