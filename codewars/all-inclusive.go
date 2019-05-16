package codewars

func getRots(strng string) []string {
	var rots_array []string
	n_string := len(strng)
	for start_idx := 0; start_idx < n_string; start_idx++ {
		rots_string := ""
		for add_idx := 0; add_idx < n_string; add_idx++ {
			idx := (start_idx + add_idx) % n_string
			rots_string = rots_string + string(strng[idx])
		}
		rots_array = append(rots_array, rots_string)
	}

	return rots_array
}

func inList(arr []string, strng string) bool {
	result := false
	for _, val := range arr {
		if strng == val {
			result = true
		}
	}
	return result
}

func ContainAllRots(strng string, arr []string) bool {
	rots_array := getRots(strng)
	result := true
	for _, rots_strng := range rots_array {
		if !inList(arr, rots_strng) {
			result = false
		}
	}
	return result
}
