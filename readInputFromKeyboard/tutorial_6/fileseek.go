package tutorial_6

import (
	"errors"
	"fmt"
	"os"
)

const lineLength = 25

func FileSeek() {
	file, err := os.OpenFile("flatfile.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("1")
	}
	defer file.Close()

	fmt.Println(readRecords(2, "last", file))
	if err := writeRecord(2, "first", "Radomir", file); err != nil {
		panic(err)
	}
	fmt.Println(readRecords(2, "first", file))
	if err := writeRecord(10, "first", "Andrew", file); err != nil {
		panic(err)
	}
	fmt.Println(readRecords(10, "first", file))
	fmt.Println(readLine(2, file))
}

func readLine(line int, f *os.File) (string, error) {
	lineBuffer := make([]byte, 24)
	f.Seek(int64(line*lineLength), 0)
	_, err := f.Read(lineBuffer)

	return string(lineBuffer), err
}

func writeRecord(line int, column, dataStr string, f *os.File) error {
	definedLen := 10
	position := int64(line * lineLength)
	switch column {
	case "id":
		definedLen = 4
	case "first":
		definedLen += 4
	case "last":
		position += 14
	default:
		return errors.New("столбец неопределен.")
	}

	if len([]byte(dataStr)) > definedLen {
		return fmt.Errorf("Максимальная длина для %s это %d", column, definedLen)
	}

	data := make([]byte, definedLen)
	for i := range data {
		data[i] = '.'
	}
	copy(data, []byte(dataStr))
	_, err := f.WriteAt(data, position)

	return err
}

func readRecords(line int, column string, f *os.File) (string, error) {
	lineBuffer := make([]byte, 24)
	f.ReadAt(lineBuffer, int64(line*lineLength))

	var retVal string
	switch column {
	case "id":
		return string(lineBuffer[:3]), nil
	case "first":
		return string(lineBuffer[4:13]), nil
	case "last":
		return string(lineBuffer[14:23]), nil
	}

	return retVal, errors.New("cтолбец неопределен")
}
