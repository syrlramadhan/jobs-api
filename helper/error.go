package helper

func SendPanicError(err error){
	if err != nil {
		panic(err)
	}
}