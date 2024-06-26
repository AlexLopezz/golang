package variable

import "fmt"

func SampleVariable() {
	var myStringVariable string = "My string variable :D"
	var myInt8Variable int8 = 123                   // 8 bits == short
	var myInt16Variable int16 = 12343               // 16bits
	var myInt32Variable int32 = 1234345678          //32 bits
	var myInt64Variable int64 = 1234234533657897532 // 64bits == long
	var myIntVariable int = 123434224234324242      // Default int64 its depends the variable size

	var myFloatVariable float32 = 32.1234567890764343534543                                // float
	var myFloat64Variable float64 = 328.12313126365325235237521736521352132153215169999888 // double

	fmt.Println(myStringVariable)

	fmt.Println(myInt8Variable)
	fmt.Println(myInt16Variable)
	fmt.Println(myInt32Variable)
	fmt.Println(myInt64Variable)
	fmt.Println(myIntVariable)

	fmt.Println(myFloatVariable)
	fmt.Println(myFloat64Variable)
}

func ShortForm() {
	// var + variableName + data-type = value <-- Long form
	var myLongVariableString string = "Long form for declare and initialize variable."

	// variableName := value <-- Short form
	myShortVariableString := "Short form for declare and initialize variable."

	fmt.Println(myLongVariableString)
	fmt.Println(myShortVariableString)
}
