package alert

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/dropstore/src/services/logger"
	"github.com/isakgranqvist2021/dropstore/src/services/store"
)

func (alert *Alert) SetAlert(c *fiber.Ctx) error {
	sess, err := store.GetStore().Get(c)

	if err != nil {
		return err
	}

	sess.Set("ALERT", alert)

	return sess.Save()
}

func GetAlert(c *fiber.Ctx) *Alert {
	sess, err := store.GetStore().Get(c)

	if err != nil {
		go logger.Log(err)

		return nil
	}

	alert, ok := sess.Get("ALERT").(Alert)

	if !ok {
		go logger.Log(errors.New("Cast sess.Get(ALERT) -> *Alert failed"))

		return nil
	}

	return &alert
}
