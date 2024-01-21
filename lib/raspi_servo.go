/*
package lib

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/drivers/pca9685"
)

func raspi_servo() {
	adaptor := i2c.NewAdafruitAdaptor()
	driver := pca9685.NewDriver(adaptor)

	work := func() {
		fmt.Println("Starting servo...")

		// サーボモーターが接続されているチャンネルを指定
		servoChannel := 0

		gobot.Every(1*time.Second, func() {
			// 0度から180度までの範囲でサーボモーターを回転
			for angle := 0; angle <= 180; angle += 30 {
				fmt.Printf("Moving servo to %d degrees\n", angle)
				driver.MoveServo(servoChannel, angle)
				time.Sleep(1 * time.Second)
			}

			// サーボモーターを初期位置に戻す
			fmt.Println("Moving servo to initial position")
			driver.MoveServo(servoChannel, 90)
			time.Sleep(1 * time.Second)
		})
	}

	robot := gobot.NewRobot("servoBot",
		[]gobot.Connection{adaptor},
		[]gobot.Device{driver},
		work,
	)

	robot.Start()
}
**