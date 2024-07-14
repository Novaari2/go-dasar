package main

import "fmt"

type validationError struct{
	Message string
}

type notFountError struct{
	Message string
}

func (v *validationError) Error() string{
	return v.Message
}

func (n *notFountError) Error() string{
	return n.Message
}

func SaveData(id string, data any) error {
	if id == ""{
		return &validationError{"Validation Error"}
	}

	if id != "nova"{
		return &notFountError{"data not found"}
	}

	// ok

	return nil
}

func main(){
	err := SaveData("nova", nil)
	if err != nil{
		if validationError, ok := err.(*validationError); ok{
			fmt.Println("validation error", validationError.Error())
		}else if notFountError, ok := err.(*notFountError); ok{
			fmt.Println("not found error:", notFountError.Error())
		}else{
			fmt.Println("unknown error:", err.Error())
		}
	}else{
		fmt.Println("Sukses")
	}
}