package tutorial_16

func TypeFunc() {

}

type getRowFn func(row int) (string, string)

func drawTable(rows int, getRow getRowFn) {}

type kelvin float64

type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}
