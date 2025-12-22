package wizard

import "github.com/manifoldco/promptui"

func ask(label, defaultValue string) (string, error) {
	p := promptui.Prompt{
		Label:   label,
		Default: defaultValue,
	}
	return p.Run()
}
