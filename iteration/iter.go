package iteration

func main() {
	Repeat("a", 5)
}

func Repeat(s string, iter int) string {
	var repeated string
	for i := 0; i < iter; i++ {
		repeated += s
	}
	return repeated
}
