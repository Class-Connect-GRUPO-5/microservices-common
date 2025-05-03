package utils

import (
	"encoding/json"
	"fmt"

	"github.com/Class-Connect-GRUPO-5/microservices-common/models"
)

func ExtractSuccessData[T any](resp models.APIResponse) (T, error) {
	var zero T

	success, ok := resp.(models.SuccessDetails)
	if !ok {
		return zero, fmt.Errorf("expected SuccessDetails, got %T", resp)
	}

	var slice []T
	if err := json.Unmarshal([]byte(success.Data), &slice); err != nil {
		return zero, fmt.Errorf("failed to unmarshal Data: %w", err)
	}

	if len(slice) == 0 {
		return zero, fmt.Errorf("no data found in SuccessDetails")
	}

	return slice[0], nil
}
