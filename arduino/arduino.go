package arduino

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

const (
	// DefaultBaud default baud rate
	// Serial.begin(9600);
	DefaultBaud = 9600
)

// Arduino arduino struct
type Arduino struct {
	name     string
	port     string
	baud     int
	coretype int
}

// GetArduinos Use gort to scan serial, return arduino list
func GetArduinos() ([]*Arduino, error) {
	cmd := exec.Command("gort", "scan", "serial")
	var output bytes.Buffer
	cmd.Stdout = &output
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	var ret = make([]*Arduino, 0)
	outString := output.String()
	para := strings.Split(outString, "\n")
	var idx = 1
	for _, s := range para {
		prefix := fmt.Sprint(idx) + "."
		if strings.HasPrefix(s, prefix) {
			items := strings.Split(s, " ")
			arduino := &Arduino{
				name: items[3][1 : len(items[3])-1],
				port: items[1][1 : len(items[1])-1],
				baud: DefaultBaud,
			}
			ret = append(ret, arduino)
			idx++
		}
	}
	return ret, nil
}

// SetCoreType set your arduino type
func (a *Arduino) SetCoreType(coreType string) {
	a.coretype = GetCoreTypeInt(coreType)
}

// WriteHex use avrdude to write hex file to arduino
func (a *Arduino) WriteHex(file string) error {
	fmt.Printf("%v", a)
	var cmd *exec.Cmd
	if a.coretype == 0 {
		cmd = exec.Command("gort", "arduino", "upload", file, a.port)
	} else {
		cmd = exec.Command("gort", "arduino", "upload", file, a.port, "-b", GetCoreType(a.coretype))
	}
	var output bytes.Buffer
	cmd.Stdout = &output
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
