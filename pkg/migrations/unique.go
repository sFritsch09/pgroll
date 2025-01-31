// SPDX-License-Identifier: Apache-2.0

package migrations

type UniqueConstraint struct {
	Name string `json:"name"`
}

func (c *UniqueConstraint) Validate() error {
	if c.Name == "" {
		return FieldRequiredError{Name: "name"}
	}

	return nil
}
