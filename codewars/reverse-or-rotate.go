package codewars

import "strconv"
import "math"

func is_reverse(s string) bool {
	sum_value := 0
	for _, value := range s {
		int_value, _ := strconv.Atoi(string(value))
		sum_value = sum_value + int(math.Pow(float64(int_value), 3))
	}
	return sum_value%2 == 0
}

func reverse(s string) string {
	var result string
	for _, value := range s {
		result = string(value) + result
	}
	return result
}

func rotate(s string) string {
	var result string
	n_s := len(s)
	if n_s > 1 {
		for _, value := range s[1:] {
			result = result + string(value)
		}
		result = result + string(s[0])
	}
	return result
}

func chunkify(s string, n int) []string {
	n_s := len(s)
	chunk_s := []string{}
	for start_idx := 0; start_idx <= n_s-n; start_idx = start_idx + n {
		chunk_s = append(chunk_s, s[start_idx:start_idx+n])
	}
	return chunk_s
}

func process(chunk_s []string) []string {
	var proc_val string
	proc_s := []string{}

	for _, val := range chunk_s {
		if is_reverse(val) {
			proc_val = reverse(val)
		} else {
			proc_val = rotate(val)
		}
		proc_s = append(proc_s, proc_val)
	}
	return proc_s
}

func Revrot(s string, n int) string {
	result := ""
	n_s := len(s)
	if n == 0 {
		return result
	} else if n > n_s {
		return result
	} else {
		chunk_s := chunkify(s, n)
		proc_s := process(chunk_s)
		for _, text := range proc_s {
			result = result + text
		}
	}
	return result
}
