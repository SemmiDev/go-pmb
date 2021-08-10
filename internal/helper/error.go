package helper

func PanicIfNeeded(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrorOrNil(err error) error {
	if err != nil {
		return err
	}
	return nil
}
