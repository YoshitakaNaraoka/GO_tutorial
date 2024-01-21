package lib

import (
  "fmt"
  "time"

  "gobot.io/x/gobot" // go get -d -u gobot.io/x/gobot で install
  "gobot.io/x/gobot/platforms/firmata" // go get gobot.io/x/gobot/platforms/firmata でインストール
  "github.com/warthog618/gpio" // go get github.com/warthog618/gpio でインストール
)

func main() {
  gobot.NewRobot()

  firmataAdaptor := firmata.NewAdaptor("firmata", "/dev/tty.usbmodem1411")
  servo := firmataAdaptor.ServoConfig("3",0,180)

  work := func() {
    gobot.Every(1*time.Second, func() {
      i := uint8(gobot.Rand(180))
      fmt.Println("Turning", i)
      servo.Move(i)
    })
  }

  robot := gobot.NewRobot("servoBot",
    []gobot.Connection{firmataAdaptor},
    []gobot.Device{servo},
    work,
  )
  robot.Start()
}
