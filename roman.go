package roman

import (
  "regexp"
  "errors"
  "strings"
)


var romanDict = map[string]int {
    "I" : 1,
    "V" : 5,
    "X" : 10,
    "L" : 50,
    "C" : 100,
    "D" : 500,
    "M" : 1000,
}

// ToInt Converts string with roman numeral to int
func ToInt(roman string) (n int, err error){
  n = 0
  err =  Errors(roman)
  roman = strings.TrimSpace(roman)
  r, _ := regexp.Compile(`^N$`)
  if r.MatchString(roman){
    return
  }
  oldValue := 0
  for i:=len(roman)-1; i>=0; i--{
    currentLetter := roman[i]
    currentValue := romanDict[string(currentLetter)]
    if currentValue < oldValue{
      n -= currentValue
    }else{
      n += currentValue
    }
    oldValue = currentValue
  }
  return 
}

// Errors validates a roman numeral string
func Errors(roman string) (err error){
  roman = strings.TrimSpace(roman)
  r, _ := regexp.Compile(`[^IVXLCDM]`)
  if r.MatchString(roman){
    return errors.New("There are invalide characters in the roman string.")
  }
  r, _ = regexp.Compile(`^.*(V.*V|L.*L|D.*D).*`)
  if r.MatchString(roman){
    return errors.New("VLD should not repeat")
  }
  r, _ = regexp.Compile(`^.*(I{4}|X{4}|C{4}|M{4}).*`)
  if r.MatchString(roman){
    return errors.New("No character should be repeated more than three times in a row")
  }
  r, _ = regexp.Compile(`^.*(II[^I]|XX[^IX]|CC[DM]|V[^I]|L[CDM]|DM).*`)
  if r.MatchString(roman){
    return errors.New("No character should be repeated more than three times in a row")
  }
  return nil
}
