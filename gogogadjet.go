package gogogadjet

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/twinj/uuid"
	"io/ioutil"
	"os"
)

const (
	stdLongMonth      = "January"
	stdMonth          = "Jan"
	stdNumMonth       = "1"
	stdZeroMonth      = "01"
	stdLongWeekDay    = "Monday"
	stdWeekDay        = "Mon"
	stdDay            = "2"
	stdUnderDay       = "_2"
	stdZeroDay        = "02"
	stdHour           = "15"
	stdHour12         = "3"
	stdZeroHour12     = "03"
	stdMinute         = "4"
	stdZeroMinute     = "04"
	stdSecond         = "5"
	stdZeroSecond     = "05"
	stdLongYear       = "2006"
	stdYear           = "06"
	stdPM             = "PM"
	stdpm             = "pm"
	stdTZ             = "MST"
	stdISO8601TZ      = "Z0700"  //prints Z for UTC
	stdISO8601ColonTZ = "Z07:00" //prints Z for UTC
	stdNumTZ          = "-0700"  //always numeric
	stdNumShortTZ     = "-07"    //always numeric
	stdNumColonTZ     = "-07:00" //always numeric
)

func NewGuid() uuid.UUID {
	return uuid.NewV4()
}

func StringToByteSlice(s string) []byte {
	return []byte(s)
}

func ByteSliceToString(b []byte) string {
	return string(b)
}

func ReadLinesOfString(s string) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(bytes.NewReader([]byte(s)))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ReadLinesOfByteSlice(b []byte) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func MakeDir(path string) error {
	b, err := FolderExists(path)
	if err != nil {
		return err
	} else if b == true {
		return errors.New("Path Exists")
	} else {
		err := os.Mkdir(path, 0644)
		return err
	}
}

func FileExists(path string) (bool, error) {
	f, err := os.Stat(path)
	if err == nil {
		if f.IsDir() == true {
			return true, errors.New("This Is A Dir")
		} else {
			return true, nil
		}
	} else {
		if os.IsNotExist(err) {
			return false, nil
		}
	}

	return true, err
}

func FolderExists(path string) (bool, error) {
	f, err := os.Stat(path)
	if err == nil {
		if f.IsDir() == false {
			return true, errors.New("This Is A File")
		} else {
			return true, nil
		}
	} else {
		if os.IsNotExist(err) {
			return false, nil
		}
	}

	return true, err
}

func FileDelete(path string) error {
	f, err := os.Stat(path)
	if err != nil {
		return err
	} else {
		if f.IsDir() == true {
			return errors.New("Path Is Dir")
		} else {
			err := os.Remove(path)
			return err
		}
	}
}

func FolderDelete(path string, NonEmpty bool) error {
	f, err := os.Stat(path)
	if err != nil {
		return err
	} else {
		if f.IsDir() == false {
			return errors.New("Path Is A File")
		} else {
			if NonEmpty == true {
				err := os.RemoveAll(path)
				return err
			} else {
				err := os.Remove(path)
				return err
			}
		}
	}
}

func WriteStringToFile(path string, data string) error {
	err := ioutil.WriteFile(path, StringToByteSlice(data), 0644)
	return err
}

func WriteAllLinesToFile(path string, data []string) error {
	s := ""
	for i, line := range data {
		if i != 0 {
			s += "\r\n" + line
		} else {
			s += line
		}
	}
	err := ioutil.WriteFile(path, StringToByteSlice(s), 0644)
	return err
}

func WriteByteSliceToFile(path string, data []byte) error {
	err := ioutil.WriteFile(path, data, 0644)
	return err
}

func ReadLinesOfFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ReadByteSliceOfFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	} else {
		return b, nil
	}
}

func ReadStringOfFile(path string) (string, error) {
	b, err := ReadByteSliceOfFile(path)
	if err != nil {
		return "", err
	} else {
		return ByteSliceToString(b), nil
	}
}
