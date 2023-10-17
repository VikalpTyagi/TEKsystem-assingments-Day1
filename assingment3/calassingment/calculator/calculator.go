package calculator

func Square(num int) int {
	return num * num
}
func Sum(num1, num2 int) int {
	return num1 + num2
}
func Differentiator(num1, num2 int) int {
	if num1 > num2 {
		return num1 - num2
	}
	return num2 - num1
}
func Multiplication(num1, num2 int) int {
	return num1 * num2
}
func Divider(num1, num2 int) (int, int) {
	quotient := num1 % num2
	remainder := num1 / num2
	return quotient, remainder
}
