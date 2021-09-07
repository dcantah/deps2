package deps2

import (
	"fmt"

	api "k8s.io/api/core/v1"
)

func Deps() {
	fmt.Println(api.StreamTypeError)
}
