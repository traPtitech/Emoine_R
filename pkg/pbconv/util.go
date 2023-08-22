package pbconv

import "database/sql/driver"

func mustValue[T any](v driver.Valuer) (value T) {
	if v == nil {
		return
	}

	vv, err := v.Value()
	if err != nil {
		return
	}

	value, _ = vv.(T)

	return value
}
