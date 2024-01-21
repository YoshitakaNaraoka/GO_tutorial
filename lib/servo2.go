package lib

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/drivers/pca9685"
	"gobot.io/x/gobot/platforms/firmata"
)

func ard_servo() {
	firmataAdaptor := firmata.NewAdaptor("/dev/ttyACM0") // Arduinoのポートに適切に変更してください
	pca9685Driver := pca9685.NewDriver(i2c.NewAdafruitAdaptor(firmataAdaptor))

	work := func() {
		fmt.Println("Starting servo...")

		// サーボモーターが接続されているチャンネルを指定
		servoChannel := 0

		gobot.Every(1*time.Second, func() {
			// 0度から180度までの範囲でサーボモーターを回転
			for angle := 0; angle <= 180; angle += 30 {
				fmt.Printf("Moving servo to %d degrees\n", angle)
				pca9685Driver.MoveServo(servoChannel, angle)
				time.Sleep(1 * time.Second)
			}

			// サーボモーターを初期位置に戻す
			fmt.Println("Moving servo to initial position")
			pca9685Driver.MoveServo(servoChannel, 90)
			time.Sleep(1 * time.Second)
		})
	}

	robot := gobot.NewRobot("servoBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{pca9685Driver},
		work,
	)

	robot.Start()
}
