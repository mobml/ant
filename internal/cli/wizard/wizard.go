package wizard

func Run(fields []Field) error {
	for _, field := range fields {
		current := field.Value()

		// If optional and has current value, skip
		if current != "" && field.Optional {
			continue
		}

		input, err := ask(field.Label, current)
		if err != nil {
			return err
		}

		// ENTER to keep current value
		if input == "" && current != "" {
			continue
		}

		if err := field.SetValue(input); err != nil {
			return err
		}
	}

	return nil
}
