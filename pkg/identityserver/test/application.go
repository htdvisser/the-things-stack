// Copyright © 2017 The Things Network Foundation, distributed under the MIT license (see LICENSE file)

package test

import (
	"fmt"
	"time"

	"github.com/TheThingsNetwork/ttn/pkg/identityserver/types"
	"github.com/smartystreets/assertions"
)

func defaultApplication(in interface{}) (*types.DefaultApplication, error) {
	if u, ok := in.(types.Application); ok {
		return u.GetApplication(), nil
	}

	if d, ok := in.(types.DefaultApplication); ok {
		return &d, nil
	}

	if ptr, ok := in.(*types.DefaultApplication); ok {
		return ptr, nil
	}

	return nil, fmt.Errorf("Expected: '%v' to be of type types.DefaultApplication but it was not", in)
}

// ShouldBeApplication checks if two Applications resemble each other.
func ShouldBeApplication(actual interface{}, expected ...interface{}) string {
	if len(expected) != 1 {
		return fmt.Sprintf("Expected: one application to match but got %v", len(expected))
	}

	a, s := defaultApplication(actual)
	if s != nil {
		return s.Error()
	}

	b, s := defaultApplication(expected[0])
	if s != nil {
		return s.Error()
	}

	return all(
		ShouldBeApplicationIgnoringAutoFields(a, b),
		assertions.ShouldHappenWithin(a.Created, time.Millisecond, b.Created),
	)
}

// ShouldBeApplicationIgnoringAutoFields checks if two Applications resemble each other
// without looking at fields that are generated by the database: created.
func ShouldBeApplicationIgnoringAutoFields(actual interface{}, expected ...interface{}) string {
	if len(expected) != 1 {
		return fmt.Sprintf("Expected: one application to match but got %v", len(expected))
	}

	a, s := defaultApplication(actual)
	if s != nil {
		return s.Error()
	}

	b, s := defaultApplication(expected[0])
	if s != nil {
		return s.Error()
	}

	return all(
		assertions.ShouldEqual(a.ID, b.ID),
		assertions.ShouldResemble(a.Description, b.Description),
		assertions.ShouldResemble(a.EUIs, b.EUIs),
		assertions.ShouldResemble(a.APIKeys, b.APIKeys),
		assertions.ShouldResemble(a.Archived, b.Archived),
	)
}
