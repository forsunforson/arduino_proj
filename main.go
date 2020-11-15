package main

import (
	"fmt"

	"github.com/forsunforson/arduino_proj/serial"

	mqtt "github.com/eclipse/paho.mqtt.golang"
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
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://forsun.icu:1883")
	opts.SetClientID("raspi")
	opts.SetUsername("forson")
	opts.SetPassword("99589958")
	opts.SetCleanSession(true)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	var ret []byte
	fmt.Println("------start-------")
	for i := 0; i < 1000; i++ {
		v := <-p.GetChan()
		if v == serial.EOF {
			fmt.Println("finished")
			return
		}
		ret = append(ret, v)
		length := len(ret)
		if length > 2 && ret[length-2] == byte(13) && ret[length-1] == byte(10) {
			fmt.Printf("%q\n", ret[:length-2])
			client.Publish("toilet", 0, false, ret)
			ret = []byte{}
		}
	}
	client.Disconnect(100)
	fmt.Println("disconnected")
}
