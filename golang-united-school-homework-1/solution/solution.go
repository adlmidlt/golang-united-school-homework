package solution

import "github.com/kyokomi/emoji"

func GetMessage() string {
	msg := emoji.Sprint("Hello 🗺️ !")

	return msg
}
