func main() {
	// longest sub string in a string
	fmt.Println(rpnCalc("     1      3 * ksd 2 -"))
}
func rpnCalc(s string) string {
	operator := "-+/%*"
	rpn := strings.Fields(s)
	stack := []int{}
	for i := 0; i < len(rpn); i++ {
		op := ""
		nu, err := strconv.Atoi(rpn[i])
		if err != nil {
			if strings.ContainsAny(operator, rpn[i]) && stack != nil {
				op = rpn[i]
				err := opSolve(op, &stack)
				if err != nil {
					return err.Error()
				}
			} else {
				return "ERROR"
			}
		} else {
			stack = append(stack, nu)
		}
		fmt.Println(stack)
	}
	if len(stack) != 1 {
		return "ERROR"
	}
	result := strconv.Itoa(stack[0])
	return result
}

func opSolve(op string, st *[]int) (err error) {
	var num int
	defer func() {
		if re := recover(); re != nil {
			err = fmt.Errorf("ERROR")
		}
	}() // this is to catch index out of range panic error

	last, bLast := (*st)[len((*st))-1], (*st)[len((*st))-2]
	*st = slices.Delete(*st, len(*st)-2, len(*st))
	switch op {
	case "+":
		num = bLast + last
	case "*":
		num = bLast * last
	case "-":
		num = bLast - last
	case "/":
		num = bLast / last
	case "%":
		num = bLast % last
	}
	*st = append(*st, num)
	return nil
}
