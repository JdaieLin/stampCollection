package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDB(t *testing.T) {
	var user = User{
		Name:          "张麻子",
		LoginID:       "balabala",
		LoginPassword: "aabbcc",
		Mail:          "balabala@stamp.com",
		Phone:         "13813888138",
		IDCard:        "439509199902024040",
		Ingots:        134,
		Coins:         28,
		Books: []StampBook{
			{
				StampID:    202,
				Status:     100,
				CreateTime: time.Now().Unix(),
			},
			{
				StampID:    203,
				Status:     101,
				CreateTime: time.Now().Unix() + 10,
			},
		},
		Chests: Chests{
			Date: 20180606,
			List: [24]Chest{
				{
					Hour:   0,
					Ingot:  3,
					Opened: true,
				},
				{
					Hour:   1,
					Ingot:  2,
					Opened: false,
				},
			},
		},
	}
	assert.Equal(t, user.Insert(), nil)
}
