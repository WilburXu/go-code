package gotest


var testData = GenerateData()

// generate billion slice data
func GenerateData() []int {
	data := make([]int, 1000000000)
	for key, _ := range data {
		data[key] = key % 128
	}

	return data
}

// get len
func GetDataLen() int {
	return len(testData)
}

// case one
func CaseSumOne(result *int) {
	data := GenerateData()
	for i := 0; i < GetDataLen(); i++ {
		*result += data[i]
	}
}

// case two
func CaseSumTwo(result *int) {
	data := GenerateData()
	dataLen := GetDataLen()
	for i := 0; i < dataLen; i++ {
		*result += data[i]
	}
}

// case three
func CaseSumThree(result *int) {
	data := GenerateData()
	dataLen := GetDataLen()
	tmp := *result
	for i:= 0; i < dataLen; i++ {
		tmp += data[i]
	}
	*result = tmp
}
