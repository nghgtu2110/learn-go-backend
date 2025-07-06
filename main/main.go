package main

func main() {
	var p *int = nil // Making a nil pointer
	*p = 0           // việc hủy tham chiếu con trỏ nill là không được phép.
}

// panic: runtime error: invalid memory address or nil pointer dereference
