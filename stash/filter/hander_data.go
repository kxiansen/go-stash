package filter

import (
	"fmt"

	"github.com/kevwan/go-stash/stash/config"
)

func HanderData(p config.Cluster) {
	fmt.Println(p)
}
