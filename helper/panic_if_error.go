package helper

func PanicIfError(err error) {
	if err != nil {
		Logger().Error(err.Error())
		panic(err)
	}
}
