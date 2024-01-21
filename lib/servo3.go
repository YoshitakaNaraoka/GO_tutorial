package lib

import (
	"fmt"
	"log"
	"time"
	"gobot.io/x/gobot"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/experimental/devices/pca9685"
	"periph.io/x/periph/host"
)

func ard_servo2() {
	// periphホストを初期化
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// I2C バスを開く
	bus, err := i2creg.Open("")
	if err != nil {
		log.Fatal(err)
	}
	defer bus.Close()

	// PCA9685 デバイスを初期化
	pca, err := pca9685.NewI2C(bus, &pca9685.DefaultOpts)
	if err != nil {
		log.Fatal(err)
	}

	// サーボモーターが接続されているチャンネルを指定
	servoChannel := 0

	fmt.Println("Starting servo...")

	gap := 100 * time.Millisecond
	pca.SetPulseDuration(servoChannel, 1000*gap)

	gobot.Every(1*time.Second, func() {
		// 0度から180度までの範囲でサーボモーターを回転
		for angle := 0; angle <= 180; angle += 30 {
			fmt.Printf("Moving servo to %d degrees\n", angle)
			pca.SetPulseDuration(servoChannel, uint16(500+angle*(2500/180)))
			time.Sleep(gap)
		}

		// サーボモーターを初期位置に戻す
		fmt.Println("Moving servo to initial position")
		pca.SetPulseDuration(servoChannel, 1000*gap)
		time.Sleep(gap)
	})
}
