package pkg

import "fmt"

func LogPrint(msg string) string {
	return fmt.Sprintf("[%v was decorated]", msg)
}
