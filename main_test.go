package main

import (
	"testing"
	"fmt"
)

func TestValidateValues(t *testing.T) {
	// case 1: all values are ok
	var (
		validXValue = 1
		validVValue = 1
		invalidXValue1 = xMax + 1
		invalidXValue2 = xMin - 1
		invalidXValue3 = "a"

		invalidVValue1 = vMax + 1
		invalidVValue2 = vMin - 1
		invalidVValue3 = "a"

	)

	// case 1: all values are ok
	inputString := fmt.Sprintf("%d %d %d %d", validVValue, validVValue, validXValue, validVValue)
	x1, v1, x2, v2, err := validateValues(inputString)
	if err != nil {
		t.Errorf("[case 1] expected no error, got %v", err)
	}
	if x1 != validXValue {
		t.Errorf("[case 1] expected x1 equal %d, got %d", validXValue, x1)
	}
	if x2 != validXValue {
		t.Errorf("[case 1] expected x2 equal %d, got %d", validXValue, x2)
	}
	if v1 != validVValue {
		t.Errorf("[case 1] expected v1 equal %d, got %d", validVValue, v1)
	}
	if v2 != validVValue {
		t.Errorf("[case 1] expected v2 equal %d, got %d", validVValue, v2)
	}

	// case 2: empty input string
	inputString = ""
	_, _, _, _, err = validateValues(inputString)
	if err == nil {
		t.Error("[case 1] expected error, got nothing")
	}

	// case 3: receive more than four values
	inputString = fmt.Sprintf("%d %d %d", validXValue, validVValue, validXValue)
	_, _, _, _, err = validateValues(inputString)
	if err == nil {
		t.Error("[case 1] expected error, got nothing")
	}

	// case 4: invalid values

	// fist value:
	inputString = fmt.Sprintf("%d %d %d %d", invalidXValue1, validVValue, validXValue, validVValue)
	_, _, _, _, err = validateValues(inputString)
	if err == nil {
		t.Error("[case 1] expected error, got nothing")
	}
	inputString = fmt.Sprintf("%d %d %d %d", invalidXValue2, validVValue, validXValue, validVValue)
	_, _, _, _, err = validateValues(inputString)
	if err == nil {
		t.Error("[case 1] expected error, got nothing")
	}
	inputString = fmt.Sprintf("%s %d %d %d", invalidXValue3, validVValue, validXValue, validVValue)
	_, _, _, _, err = validateValues(inputString)
	if err == nil {
		t.Error("[case 1] expected error, got nothing")
	}

	// second value:
	inputString = fmt.Sprintf("%d %d %d %d", validXValue, invalidVValue1, validXValue, validVValue)
	_, _, _, _, err = validateValues(inputString)
	if err == nil {
		t.Error("[case 1] expected error, got nothing")
	}
	inputString = fmt.Sprintf("%d %d %d %d", validXValue, invalidVValue2, validXValue, validVValue)
	_, _, _, _, err = validateValues(inputString)
	if err == nil {
		t.Error("[case 1] expected error, got nothing")
	}
	inputString = fmt.Sprintf("%d %s %d %d", validXValue, invalidVValue3, validXValue, validVValue)
	_, _, _, _, err = validateValues(inputString)
	if err == nil {
		t.Error("[case 1] expected error, got nothing")
	}

	// third value:
	inputString = fmt.Sprintf("%d %d %d %d", validXValue, validVValue, invalidXValue1, validVValue)
	_, _, _, _, err = validateValues(inputString)
	if err == nil {
		t.Error("[case 1] expected error, got nothing")
	}
	inputString = fmt.Sprintf("%d %d %d %d", validXValue, validVValue, invalidXValue2, validVValue)
	_, _, _, _, err = validateValues(inputString)
	if err == nil {
		t.Error("[case 1] expected error, got nothing")
	}
	inputString = fmt.Sprintf("%d %d %s %d", validXValue, validVValue, invalidXValue3, validVValue)
	_, _, _, _, err = validateValues(inputString)
	if err == nil {
		t.Error("[case 1] expected error, got nothing")
	}

	// fourth value:
	inputString = fmt.Sprintf("%d %d %d %d", validXValue, validVValue, validXValue, invalidVValue1)
	_, _, _, _, err = validateValues(inputString)
	if err == nil {
		t.Error("[case 1] expected error, got nothing")
	}
	inputString = fmt.Sprintf("%d %d %d %d", validXValue, validVValue, validXValue, invalidVValue2)
	_, _, _, _, err = validateValues(inputString)
	if err == nil {
		t.Error("[case 1] expected error, got nothing")
	}
	inputString = fmt.Sprintf("%d %d %d %s", validXValue, validVValue, validXValue, invalidVValue3)
	_, _, _, _, err = validateValues(inputString)
	if err == nil {
		t.Error("[case 1] expected error, got nothing")
	}

}

func TestKenGo(t *testing.T) {
	// case 1: same start coordinates, different speeds
	var (
		x = 0
		v11, v12 = 1,2
	)
	if result := kenGo(x, v11, x, v12); result != false  {
		t.Errorf("[case 1] has to got %t, got %t", false, result)
	}

	// case 2: same speeds, different start coordinates
	var (
		x21, x22 = 1,2
		v = 1
	)
	if result := kenGo(x21, v, x22, v); result != false  {
		t.Errorf("[case 2] has to got %t, got %t", false, result)
	}

	// case 3: same speeds, same coordinates
	var (
		x3 =  1
		v3 = 1
	)
	if result := kenGo(x3, v3, x3, v3); result != true  {
		t.Errorf("[case 3] has to got %t, got %t", true, result)
	}

	// case 4: one is standing still, the other get away
	var (
		x41, x42 =  1, 2
		v41, v42 = 0, 1
	)
	if result := kenGo(x41, v41, x42, v42); result != false  {
		t.Errorf("[case 4] has to got %t, got %t", false, result)
	}

	// case 5: one is standing still, the other is coming towards him
	var (
		x51, x52 =  0, -10
		v51, v52 = 0, 1
	)
	if result := kenGo(x51, v51, x52, v52); result != true  {
		t.Errorf("[case 5] has to got %t, got %t", true, result)
	}

	// case 6: different values, have to met
	var (
		x61, x62 =  10, 30
		v61, v62 = 15, 5
	)
	if result := kenGo(x61, v61, x62, v62); result != true  {
		t.Errorf("[case 6] has to got %t, got %t", true, result)
	}

	// case 7: different values, don't have to met
	var (
		x71, x72 =  -1, 1
		v71, v72 = -1, 1
	)
	if result := kenGo(x71, v71, x72, v72); result != false  {
		t.Errorf("[case 7] has to got %t, got %t", false, result)
	}
}