package main

import (
	"fmt"

	"github.com/forsun/arduino_proj/arduino"
	"github.com/forsun/arduino_proj/serial"
)

func main() {
	dir := "/dev/ttyACM0"
	p, err := serial.NewPort(dir, 9600)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("------ listen to port [%s] ------\n", dir)
	p.OpenChan()
	// soundFile := "./sound.log"
	// tempFile := "./temp.log"
	// lightFile := "./light.log"

	var ret []byte
	for i := 0; i < 100; i++ {
		v := <-p.GetChan()
		if v == serial.EOF {
			fmt.Println("finished")
			return
		}
		ret = append(ret, v)
		length := len(ret)
		if length > 2 && ret[length-2] == byte(13) && ret[length-1] == byte(10) {
			fmt.Printf("%q\n", ret[:length-2])
			ret = []byte{}
		}
	}
	a, _ := arduino.GetArduinos()
	a[0].SetCoreType("uno")
	_ = a[0].WriteHex(`/home/pi/sound_sensor.ino.with_bootloader.standard.hex`)
}
