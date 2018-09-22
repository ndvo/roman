package roman

import (
  "testing"
  "fmt"
)

var testcases = []struct{
  roman string
  valid int
}{
  { "V", 5},
  { "VIIV", 0},
  { "Vamos la", 0},
  { "DDD", 0},
  { "IIIIV", 0},
  { "XXLV", 0 },
  { "XLV", 45 },
  { "CCLVII", 257 },
  { "XM", 990 },
  { "CMXC", 990 },
  { "DCCXLVIII", 748 },
  { "CCDXLVIII", 0 },
}

// Test converting roman chars
func TestToInt(t *testing.T){
  for _, testcase := range testcases {
    err := Errors(testcase.roman)
    if err != nil{
      println("Failed to convert to convert %s to int", testcase.roman)
      continue
    }
    arabic, _ := ToInt(testcase.roman)
    if arabic != testcase.valid {
      t.Errorf(fmt.Sprintf("%s should be %d, not %d", testcase.roman, testcase.valid, arabic))
    }else{
      println("Success for "+testcase.roman+"=>", arabic)
    }
  }
}

// Test validating roman numbers
func TestErros(t *testing.T){
  for _, testcase := range testcases {
    err := Errors(testcase.roman)
    if err != nil && testcase.valid != 0{
      t.Errorf(testcase.roman+" should be valid. Error reported "+err.Error())
    }else if err==nil && testcase.valid==0 {
      t.Errorf(testcase.roman+" should be invalid.")
    }else{
      println(testcase.roman+" passed the test.")
    }
  }
}
