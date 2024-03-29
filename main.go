package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/firmata"
	"gobot.io/x/gobot/drivers/gpio"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor("COM3") // Arduinoのポートに適切に変更してください windows : COM3
	servo := gpio.NewServoDriver(firmataAdaptor, "9") // この行を修正する pin が 9 なので線を差したピン番号に
	led := gpio.NewLedDriver(firmataAdaptor, "12")         // LEDのピン番号に適切に変更してください


	work := func() {
		fmt.Println("Starting servo and LED...")

		gobot.Every(1*time.Second, func() {
			// 0度から180度までの範囲でサーボモーターを回転
			for angle := 0; angle <= 180; angle += 30 {
				fmt.Printf("Moving servo to %d degrees\n", angle)
				servo.Move(uint8(angle))

				// LEDを点滅させる
				led.Toggle()
				time.Sleep(500 * time.Millisecond)
				led.Toggle()
				time.Sleep(500 * time.Millisecond)
			}

			// サーボモーターを初期位置に戻す
			fmt.Println("Moving servo to initial position")
			servo.Move(90)

			// LEDを点滅させる
			led.On()
			time.Sleep(500 * time.Millisecond)
			led.Off()
			time.Sleep(500 * time.Millisecond)
		})
	}

	robot := gobot.NewRobot("servoBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{servo, led},
		work,
	)

	robot.Start()
}
// terminal に Ctrl + C で停止