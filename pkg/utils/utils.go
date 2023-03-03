package utils

func ProcessError(err error) *string {
	if err != nil {
		return GetStringPointer(err.Error())
	}
	return nil
}

func GetStringPointer(s string) *string {
	return &s
}
