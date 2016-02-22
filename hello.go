package main

import (
	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/joystick"
)

func main() {
	gbot := gobot.NewGobot()

	joystickAdaptor := joystick.NewJoystickAdaptor("logitech")
	joystick := joystick.NewJoystickDriver(joystickAdaptor,
		"logitech",
		"./controllerconfig/logitech.json",
	)

	work := func() {
		gobot.On(joystick.Event("square_press"), func(data interface{}) {
			fmt.Println("square_press")
		})
		gobot.On(joystick.Event("square_release"), func(data interface{}) {
			fmt.Println("square_release")
		})
		gobot.On(joystick.Event("triangle_press"), func(data interface{}) {
			fmt.Println("triangle_press")
		})
		gobot.On(joystick.Event("triangle_release"), func(data interface{}) {
			fmt.Println("triangle_release")
		})
		gobot.On(joystick.Event("left_x"), func(data interface{}) {
			val, found := data.(int16)
			if found {
				//if val == 32767 || val == -32768 {
					fmt.Println("left_x", val >> 7)
				//}
			}
		})
		gobot.On(joystick.Event("left_y"), func(data interface{}) {
			val, found := data.(int16)
			if found {
				//if val == 32767 || val == -32768 {
					fmt.Println("left_y", val >> 7)
				//}
			}
		})
		gobot.On(joystick.Event("right_x"), func(data interface{}) {
			val, found := data.(int16)
			if found {
				//if val == 32767 || val == -32768 {
					fmt.Println("right_x", val >> 7)
				//}
			}
		})
		gobot.On(joystick.Event("right_y"), func(data interface{}) {
			val, found := data.(int16)
			if found {
				//if val == 32767 || val == -32768 {
					fmt.Println("right_y", val >> 7)
				//}
			}
		})
	}

	robot := gobot.NewRobot("joystickBot",
		[]gobot.Connection{joystickAdaptor},
		[]gobot.Device{joystick},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}