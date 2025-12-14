package common

func ValidateID(id string) error {
	if id == "" {
		return ErrIDRequired
	}
	return nil
}
