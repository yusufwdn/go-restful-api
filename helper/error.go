package helper

// fungsi untuk mengecek apakah ada error atau tidak pada parameter yang diberikan
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
