package utils

import (
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/sony/sonyflake"
)

var generator *sonyflake.Sonyflake

func init() {
	generator = sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Now().UTC(),
	})
}

func GenerateID() (string, error) {
	uid, err := generator.NextID()
	if err != nil {
		return "", errors.Wrap(err, "failed to generate id")
	}

	return strconv.FormatUint(uid, 10), nil

}
