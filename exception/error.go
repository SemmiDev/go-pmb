package exception

func ErrorOrNil(err error) error {
	if err != nil {
		return err
	}
	return nil
}
