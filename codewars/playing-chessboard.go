package codewars

import "fmt"
import "math"

type Fraction struct {
	num int
	den int
}

func displayFraction(frac Fraction) []int {
	num := frac.num
	den := frac.den
	if den == 1 {
		return []int{num}
	} else {
		return []int{num, den}
	}
}

func simpleFraction(frac Fraction) Fraction {
	num := frac.num
	den := frac.den
	res_gcd := gcd(num, den)
	new_num := num / res_gcd
	new_den := den / res_gcd
	return Fraction{num: new_num, den: new_den}
}

func addFraction(a Fraction, b Fraction) Fraction {
	a_num := a.num
	b_num := b.num
	a_den := a.den
	b_den := b.den

	if a_num == 0 {
		return Fraction{num: b_num, den: b_den}
	}
	if b_num == 0 {
		return Fraction{num: a_num, den: a_den}
	}

	lcm_den := lcm(a_den, b_den)
	a_mul := lcm_den / a_den
	b_mul := lcm_den / b_den

	new_a_num := a_mul * a_num
	new_b_num := b_mul * b_num

	result_frac := Fraction{num: new_a_num + new_b_num, den: lcm_den}
	simple_frac := simpleFraction(result_frac)
	return simple_frac
}

func gcd(a int, b int) int {
	for b != 0 {
		temp_a := a
		a = b
		b = temp_a % b
	}
	if a == 0 {
		return 1
	} else {
		return a
	}
}

func lcm(a int, b int) int {
	return a * b / gcd(a, b)
}

func Game(n int) []int {
	if n == 0 {
		return []int{0}
	} else {
		result := Fraction{num: int(math.Pow(float64(n), 2)), den: 2}
		result = simpleFraction(result)
		disp_result := displayFraction(result)
		fmt.Println(disp_result)
		return disp_result
	}
}
