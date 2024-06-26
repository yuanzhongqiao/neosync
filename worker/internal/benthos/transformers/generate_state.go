package transformers

import (
	"math/rand"

	"github.com/benthosdev/benthos/v4/public/bloblang"
	transformers_dataset "github.com/nucleuscloud/neosync/worker/internal/benthos/transformers/data-sets"
)

func init() {
	spec := bloblang.NewPluginSpec()

	err := bloblang.RegisterFunctionV2("generate_state", spec, func(args *bloblang.ParsedParams) (bloblang.Function, error) {
		return func() (any, error) {
			return generateRandomState(), nil
		}, nil
	})
	if err != nil {
		panic(err)
	}
}

// Generates a randomly selected state that exists in the United States and returns the two-letter state code.
func generateRandomState() string {
	addresses := transformers_dataset.Addresses

	//nolint:gosec
	randomIndex := rand.Intn(len(addresses))
	return addresses[randomIndex].State
}
