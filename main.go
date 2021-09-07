package deps2

import (
	"fmt"

	runtime "k8s.io/cri-api/pkg/apis/runtime/v1"
)

func Deps() {
	fmt.Println(&runtime.AuthConfig{})
}
