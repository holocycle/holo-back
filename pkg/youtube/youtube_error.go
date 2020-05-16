package youtube

import "github.com/pkg/errors"

func newErr(err error) error {
	return errors.WithStack(err)
}
