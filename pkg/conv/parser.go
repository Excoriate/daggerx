package conv

import (
	"dagger.io/dagger"
	"fmt"
)

func ToFnContainer[T any](input interface{}) *T {
	ctr, ok := input.(*T)
	if !ok {
		return nil
	}

	return ctr
}

func ToFnContainerE[T any](input interface{}) (*T, error) {
	ctr, ok := input.(*T)
	if !ok {
		return nil, fmt.Errorf("failed to assert type, the passed type is not a *Container")
	}

	return ctr, nil
}

func ToDaggerContainerE(input interface{}) (*dagger.Container, error) {
	ctr, ok := input.(*dagger.Container)
	if !ok {
		return nil, fmt.Errorf("failed to assert type, the passed type is not a *dagger.Container")
	}

	return ctr, nil
}

func ToDaggerContainer(input interface{}) *dagger.Container {
	ctr, ok := input.(*dagger.Container)
	if !ok {
		return nil
	}

	return ctr
}
