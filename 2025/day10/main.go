
package main

import (
	"fmt"
	"math"
    "strconv"
)
func round(x float64, decimals int) float64 {
    pow := math.Pow(10, float64(decimals))
    return math.Round(x * pow) / pow
}

const maxValue = 250


func solve1 () int {
  var x [8]float64
  minSum := 99999.0
  minSol := ""

  for x[6]=0.0;x[6]<maxValue;x[6]++ {
  for x[7]=0.0;x[7]<maxValue;x[7]++ {
  x[1]=-10+x[6]+2*x[7]
  x[2]=-9+x[6]+x[7]
  x[3]=53-2*x[6]-3*x[7]
  x[4]=25-x[6]-2*x[7]
  x[5]=2

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve2 () int {
  var x [6]float64
  minSum := 99999.0
  minSol := ""

  for x[5]=0.0;x[5]<maxValue;x[5]++ {
  x[1]=20-x[5]
  x[2]=12
  x[3]=17-x[5]
  x[4]=11+x[5]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve3 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  x[1]=0
  x[2]=7
  x[3]=0
  x[4]=20
  x[5]=11
  x[6]=2
  x[7]=0
  x[8]=0

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve4 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  x[1]=14
  x[2]=5
  x[3]=1
  x[4]=4
  x[5]=9
  x[6]=16
  x[7]=10
  x[8]=15

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve5 () int {
  var x [12]float64
  minSum := 99999.0
  minSol := ""

  for x[8]=0.0;x[8]<maxValue;x[8]++ {
  x[1]=4
  x[2]=11+x[8]
  x[3]=2-x[8]
  x[4]=17
  x[5]=5
  x[6]=9
  x[7]=9+x[8]
  x[9]=16
  x[10]=133
  x[11]=9

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve6 () int {
  var x [4]float64
  minSum := 99999.0
  minSol := ""

  x[1]=20
  x[2]=12
  x[3]=18

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve7 () int {
  var x [8]float64
  minSum := 99999.0
  minSol := ""

  for x[6]=0.0;x[6]<maxValue;x[6]++ {
  x[1]=6-x[6]
  x[2]=15+x[6]
  x[3]=15-x[6]
  x[4]=13
  x[5]=14
  x[7]=0

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve8 () int {
  var x [3]float64
  minSum := 99999.0
  minSol := ""

  x[1]=10
  x[2]=13

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve9 () int {
  var x [8]float64
  minSum := 99999.0
  minSol := ""

  x[1]=5
  x[2]=0
  x[3]=20
  x[4]=18
  x[5]=9
  x[6]=2
  x[7]=3

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve10 () int {
  var x [8]float64
  minSum := 99999.0
  minSol := ""

  x[1]=2
  x[2]=11
  x[3]=9
  x[4]=3
  x[5]=2
  x[6]=14
  x[7]=13

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve11 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  x[1]=18
  x[2]=2
  x[3]=18
  x[4]=129
  x[5]=13
  x[6]=17
  x[7]=17
  x[8]=0

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve12 () int {
  var x [13]float64
  minSum := 99999.0
  minSol := ""

  for x[10]=0.0;x[10]<maxValue;x[10]++ {
  for x[12]=0.0;x[12]<maxValue;x[12]++ {
  x[1]=(47.0/3.0)+(1.0/3.0)*x[10]-(2.0/3.0)*x[12]
  x[2]=(41.0/3.0)+(1.0/3.0)*x[10]+(1.0/3.0)*x[12]
  x[3]=(26.0/3.0)+(1.0/3.0)*x[10]+(4.0/3.0)*x[12]
  x[4]=(58.0/3.0)-(1.0/3.0)*x[10]-(1.0/3.0)*x[12]
  x[5]=2+x[10]
  x[6]=(76.0/3.0)-(4.0/3.0)*x[10]-(1.0/3.0)*x[12]
  x[7]=16+x[12]
  x[8]=-(10.0/3.0)+(1.0/3.0)*x[10]+(1.0/3.0)*x[12]
  x[9]=40-2*x[10]+x[12]
  x[11]=14-x[12]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve13 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  x[1]=17
  x[2]=16
  x[3]=1
  x[4]=17
  x[5]=2
  x[6]=19
  x[7]=199
  x[8]=8

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve14 () int {
  var x [6]float64
  minSum := 99999.0
  minSol := ""

  for x[5]=0.0;x[5]<maxValue;x[5]++ {
  x[1]=13
  x[2]=157-x[5]
  x[3]=-14+x[5]
  x[4]=33-x[5]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve15 () int {
  var x [11]float64
  minSum := 99999.0
  minSol := ""

  x[1]=9
  x[2]=16
  x[3]=19
  x[4]=6
  x[5]=7
  x[6]=5
  x[7]=13
  x[8]=12
  x[9]=2
  x[10]=13

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve16 () int {
  var x [6]float64
  minSum := 99999.0
  minSol := ""

  x[1]=18
  x[2]=109
  x[3]=20
  x[4]=6
  x[5]=1

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve17 () int {
  var x [6]float64
  minSum := 99999.0
  minSol := ""

  for x[5]=0.0;x[5]<maxValue;x[5]++ {
  x[1]=31-x[5]
  x[2]=27-x[5]
  x[3]=23-x[5]
  x[4]=-2+x[5]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve18 () int {
  var x [7]float64
  minSum := 99999.0
  minSol := ""

  x[1]=12
  x[2]=16
  x[3]=11
  x[4]=11
  x[5]=7
  x[6]=2

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve19 () int {
  var x [7]float64
  minSum := 99999.0
  minSol := ""

  for x[5]=0.0;x[5]<maxValue;x[5]++ {
  x[1]=2
  x[2]=1
  x[3]=27-x[5]
  x[4]=12-x[5]
  x[6]=14

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve20 () int {
  var x [10]float64
  minSum := 99999.0
  minSol := ""

  for x[9]=0.0;x[9]<maxValue;x[9]++ {
  x[1]=18-x[9]
  x[2]=5-x[9]
  x[3]=-3+2*x[9]
  x[4]=15-2*x[9]
  x[5]=8-x[9]
  x[6]=132+x[9]
  x[7]=14-x[9]
  x[8]=14+x[9]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve21 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  for x[7]=0.0;x[7]<maxValue;x[7]++ {
  for x[8]=0.0;x[8]<maxValue;x[8]++ {
  x[1]=25+x[7]-x[8]
  x[2]=15-x[8]
  x[3]=-11-x[7]+3*x[8]
  x[4]=13-x[7]+x[8]
  x[5]=31+x[7]-3*x[8]
  x[6]=22-x[7]-x[8]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve22 () int {
  var x [8]float64
  minSum := 99999.0
  minSol := ""

  for x[6]=0.0;x[6]<maxValue;x[6]++ {
  for x[7]=0.0;x[7]<maxValue;x[7]++ {
  x[1]=13
  x[2]=27-x[7]
  x[3]=41-x[6]-x[7]
  x[4]=-8+2*x[6]
  x[5]=9-x[6]+x[7]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve23 () int {
  var x [10]float64
  minSum := 99999.0
  minSol := ""

  x[1]=14
  x[2]=2
  x[3]=16
  x[4]=4
  x[5]=18
  x[6]=7
  x[7]=6
  x[8]=16
  x[9]=3

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve24 () int {
  var x [8]float64
  minSum := 99999.0
  minSol := ""

  x[1]=2
  x[2]=19
  x[3]=6
  x[4]=20
  x[5]=11
  x[6]=3
  x[7]=8

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve25 () int {
  var x [8]float64
  minSum := 99999.0
  minSol := ""

  x[1]=6
  x[2]=15
  x[3]=20
  x[4]=11
  x[5]=11
  x[6]=19
  x[7]=17

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve26 () int {
  var x [10]float64
  minSum := 99999.0
  minSol := ""

  for x[8]=0.0;x[8]<maxValue;x[8]++ {
  for x[9]=0.0;x[9]<maxValue;x[9]++ {
  x[1]=-40+2*x[8]+2*x[9]
  x[2]=-17+x[8]
  x[3]=2
  x[4]=131-x[8]-x[9]
  x[5]=14
  x[6]=81-3*x[8]-3*x[9]
  x[7]=41-x[8]-x[9]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve27 () int {
  var x [11]float64
  minSum := 99999.0
  minSol := ""

  for x[6]=0.0;x[6]<maxValue;x[6]++ {
  for x[9]=0.0;x[9]<maxValue;x[9]++ {
  x[1]=-1-x[6]+x[9]
  x[2]=25+x[6]-x[9]
  x[3]=35-x[9]
  x[4]=16
  x[5]=8-x[6]
  x[7]=10
  x[8]=3
  x[10]=17

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve28 () int {
  var x [8]float64
  minSum := 99999.0
  minSol := ""

  for x[7]=0.0;x[7]<maxValue;x[7]++ {
  x[1]=18-x[7]
  x[2]=11
  x[3]=8+x[7]
  x[4]=15-x[7]
  x[5]=11-x[7]
  x[6]=-3+x[7]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve29 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  for x[7]=0.0;x[7]<maxValue;x[7]++ {
  x[1]=5
  x[2]=0
  x[3]=-2+x[7]
  x[4]=21-x[7]
  x[5]=7
  x[6]=8
  x[8]=15

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve30 () int {
  var x [8]float64
  minSum := 99999.0
  minSol := ""

  for x[7]=0.0;x[7]<maxValue;x[7]++ {
  x[1]=0
  x[2]=15-x[7]
  x[3]=3
  x[4]=0
  x[5]=25-x[7]
  x[6]=197

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve31 () int {
  var x [6]float64
  minSum := 99999.0
  minSol := ""

  for x[5]=0.0;x[5]<maxValue;x[5]++ {
  x[1]=2+x[5]
  x[2]=18
  x[3]=-3+x[5]
  x[4]=23-2*x[5]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve32 () int {
  var x [11]float64
  minSum := 99999.0
  minSol := ""

  for x[7]=0.0;x[7]<maxValue;x[7]++ {
  for x[9]=0.0;x[9]<maxValue;x[9]++ {
  x[1]=8
  x[2]=37-x[7]-x[9]
  x[3]=19
  x[4]=17-x[7]
  x[5]=-4+x[7]
  x[6]=-13+x[7]+x[9]
  x[8]=5
  x[10]=20

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve33 () int {
  var x [14]float64
  minSum := 99999.0
  minSol := ""

  for x[9]=0.0;x[9]<maxValue;x[9]++ {
  for x[12]=0.0;x[12]<maxValue;x[12]++ {
  for x[13]=0.0;x[13]<maxValue;x[13]++ {
  x[1]=(149.0/2.0)+x[9]-4*x[12]-(3.0/2.0)*x[13]
  x[2]=124+2*x[9]-8*x[12]-3*x[13]
  x[3]=-73-2*x[9]+6*x[12]+x[13]
  x[4]=-(23.0/2.0)+x[12]-(1.0/2.0)*x[13]
  x[5]=(125.0/2.0)+2*x[9]-4*x[12]-(3.0/2.0)*x[13]
  x[6]=-(63.0/2.0)-x[9]+3*x[12]+(5.0/2.0)*x[13]
  x[7]=56-3*x[12]-x[13]
  x[8]=-(207.0/2.0)-2*x[9]+8*x[12]+(5.0/2.0)*x[13]
  x[10]=(135.0/2.0)-3*x[12]-(1.0/2.0)*x[13]
  x[11]=-(29.0/2.0)+x[12]+(1.0/2.0)*x[13]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve34 () int {
  var x [10]float64
  minSum := 99999.0
  minSol := ""

  x[1]=18
  x[2]=11
  x[3]=102
  x[4]=7
  x[5]=11
  x[6]=15
  x[7]=7
  x[8]=13
  x[9]=9

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve35 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  x[1]=17
  x[2]=3
  x[3]=14
  x[4]=2
  x[5]=7
  x[6]=1
  x[7]=5
  x[8]=17

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve36 () int {
  var x [10]float64
  minSum := 99999.0
  minSol := ""

  for x[9]=0.0;x[9]<maxValue;x[9]++ {
  x[1]=42-2*x[9]
  x[2]=20
  x[3]=14-x[9]
  x[4]=28-x[9]
  x[5]=-17+2*x[9]
  x[6]=16
  x[7]=-11+2*x[9]
  x[8]=26-2*x[9]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve37 () int {
  var x [10]float64
  minSum := 99999.0
  minSol := ""

  x[1]=5
  x[2]=10
  x[3]=9
  x[4]=6
  x[5]=11
  x[6]=13
  x[7]=1
  x[8]=11
  x[9]=2

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve38 () int {
  var x [4]float64
  minSum := 99999.0
  minSol := ""

  x[1]=12
  x[2]=1
  x[3]=0

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve39 () int {
  var x [8]float64
  minSum := 99999.0
  minSol := ""

  for x[6]=0.0;x[6]<maxValue;x[6]++ {
  for x[7]=0.0;x[7]<maxValue;x[7]++ {
  x[1]=-129+x[7]
  x[2]=15
  x[3]=149-x[6]-x[7]
  x[4]=135-x[6]-x[7]
  x[5]=12+x[6]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve40 () int {
  var x [10]float64
  minSum := 99999.0
  minSol := ""

  x[1]=9
  x[2]=20
  x[3]=19
  x[4]=0
  x[5]=140
  x[6]=6
  x[7]=13
  x[8]=1
  x[9]=2

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve41 () int {
  var x [7]float64
  minSum := 99999.0
  minSol := ""

  for x[5]=0.0;x[5]<maxValue;x[5]++ {
  for x[6]=0.0;x[6]<maxValue;x[6]++ {
  x[1]=(31.0/2.0)+(1.0/2.0)*x[6]
  x[2]=(9.0/2.0)-x[5]-(1.0/2.0)*x[6]
  x[3]=4+x[5]-x[6]
  x[4]=(29.0/2.0)-(1.0/2.0)*x[6]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve42 () int {
  var x [11]float64
  minSum := 99999.0
  minSol := ""

  for x[9]=0.0;x[9]<maxValue;x[9]++ {
  x[1]=10
  x[2]=38-x[9]
  x[3]=21-x[9]
  x[4]=-18+x[9]
  x[5]=6
  x[6]=-3+x[9]
  x[7]=24-x[9]
  x[8]=17
  x[10]=6

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve43 () int {
  var x [12]float64
  minSum := 99999.0
  minSol := ""

  for x[9]=0.0;x[9]<maxValue;x[9]++ {
  x[1]=12-x[9]
  x[2]=13
  x[3]=15+x[9]
  x[4]=6
  x[5]=16-x[9]
  x[6]=8
  x[7]=14
  x[8]=8
  x[10]=4
  x[11]=3

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve44 () int {
  var x [3]float64
  minSum := 99999.0
  minSol := ""

  x[1]=9
  x[2]=185

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve45 () int {
  var x [14]float64
  minSum := 99999.0
  minSol := ""

  for x[10]=0.0;x[10]<maxValue;x[10]++ {
  for x[12]=0.0;x[12]<maxValue;x[12]++ {
  for x[13]=0.0;x[13]<maxValue;x[13]++ {
  x[1]=54-x[10]-3*x[12]+2*x[13]
  x[2]=22-x[12]+x[13]
  x[3]=36-2*x[10]-x[12]+x[13]
  x[4]=-22+x[10]+x[12]
  x[5]=-29+x[10]+2*x[12]-x[13]
  x[6]=40-2*x[12]+x[13]
  x[7]=2-x[10]+x[12]-x[13]
  x[8]=20-x[12]
  x[9]=5+x[12]-x[13]
  x[11]=33-x[12]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve46 () int {
  var x [5]float64
  minSum := 99999.0
  minSol := ""

  x[1]=20
  x[2]=17
  x[3]=3
  x[4]=4

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve47 () int {
  var x [6]float64
  minSum := 99999.0
  minSol := ""

  x[1]=3
  x[2]=0
  x[3]=13
  x[4]=7
  x[5]=0

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve48 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  for x[7]=0.0;x[7]<maxValue;x[7]++ {
  for x[8]=0.0;x[8]<maxValue;x[8]++ {
  x[1]=19
  x[2]=x[8]
  x[3]=5+x[7]
  x[4]=30-x[7]-x[8]
  x[5]=20-x[8]
  x[6]=37-x[7]-x[8]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve49 () int {
  var x [5]float64
  minSum := 99999.0
  minSol := ""

  x[1]=182
  x[2]=1
  x[3]=5
  x[4]=10

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve50 () int {
  var x [8]float64
  minSum := 99999.0
  minSol := ""

  x[1]=18
  x[2]=20
  x[3]=20
  x[4]=17
  x[5]=2
  x[6]=18
  x[7]=7

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve51 () int {
  var x [7]float64
  minSum := 99999.0
  minSol := ""

  x[1]=109
  x[2]=17
  x[3]=9
  x[4]=15
  x[5]=3
  x[6]=20

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve52 () int {
  var x [7]float64
  minSum := 99999.0
  minSol := ""

  for x[5]=0.0;x[5]<maxValue;x[5]++ {
  for x[6]=0.0;x[6]<maxValue;x[6]++ {
  x[1]=-7+x[5]-x[6]
  x[2]=4-x[6]
  x[3]=22-x[5]
  x[4]=29-x[5]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve53 () int {
  var x [4]float64
  minSum := 99999.0
  minSol := ""

  x[1]=19
  x[2]=10
  x[3]=8

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve54 () int {
  var x [14]float64
  minSum := 99999.0
  minSol := ""

  for x[11]=0.0;x[11]<maxValue;x[11]++ {
  for x[12]=0.0;x[12]<maxValue;x[12]++ {
  for x[13]=0.0;x[13]<maxValue;x[13]++ {
  x[1]=14-(1.0/2.0)*x[11]+(1.0/4.0)*x[12]-(1.0/4.0)*x[13]
  x[2]=-3+x[11]+x[12]-x[13]
  x[3]=15-(1.0/2.0)*x[11]+(1.0/4.0)*x[12]-(5.0/4.0)*x[13]
  x[4]=15-x[12]
  x[5]=15-(1.0/2.0)*x[11]-(3.0/4.0)*x[12]+(3.0/4.0)*x[13]
  x[6]=11+(3.0/2.0)*x[11]-(1.0/4.0)*x[12]-(3.0/4.0)*x[13]
  x[7]=9-(3.0/2.0)*x[11]-(1.0/4.0)*x[12]+(1.0/4.0)*x[13]
  x[8]=10-(1.0/2.0)*x[11]+(1.0/4.0)*x[12]+(3.0/4.0)*x[13]
  x[9]=14+(1.0/2.0)*x[11]+(3.0/4.0)*x[12]-(3.0/4.0)*x[13]
  x[10]=15+(1.0/2.0)*x[11]+(1.0/4.0)*x[12]-(1.0/4.0)*x[13]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve55 () int {
  var x [8]float64
  minSum := 99999.0
  minSol := ""

  x[1]=12
  x[2]=13
  x[3]=4
  x[4]=20
  x[5]=12
  x[6]=15
  x[7]=7

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve56 () int {
  var x [7]float64
  minSum := 99999.0
  minSol := ""

  x[1]=0
  x[2]=11
  x[3]=17
  x[4]=3
  x[5]=12
  x[6]=12

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve57 () int {
  var x [11]float64
  minSum := 99999.0
  minSol := ""

  x[1]=18
  x[2]=4
  x[3]=12
  x[4]=2
  x[5]=12
  x[6]=3
  x[7]=11
  x[8]=1
  x[9]=16
  x[10]=18

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve58 () int {
  var x [10]float64
  minSum := 99999.0
  minSol := ""

  for x[5]=0.0;x[5]<maxValue;x[5]++ {
  x[1]=x[5]
  x[2]=16
  x[3]=15-x[5]
  x[4]=6
  x[6]=0
  x[7]=5
  x[8]=12
  x[9]=20

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve59 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  for x[5]=0.0;x[5]<maxValue;x[5]++ {
  for x[6]=0.0;x[6]<maxValue;x[6]++ {
  x[1]=-12+x[5]+x[6]
  x[2]=34-x[5]-x[6]
  x[3]=6
  x[4]=29-x[6]
  x[7]=13
  x[8]=6

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve60 () int {
  var x [6]float64
  minSum := 99999.0
  minSol := ""

  x[1]=3
  x[2]=14
  x[3]=5
  x[4]=10
  x[5]=11

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve61 () int {
  var x [8]float64
  minSum := 99999.0
  minSol := ""

  x[1]=4
  x[2]=18
  x[3]=19
  x[4]=8
  x[5]=129
  x[6]=12
  x[7]=19

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve62 () int {
  var x [12]float64
  minSum := 99999.0
  minSol := ""

  for x[11]=0.0;x[11]<maxValue;x[11]++ {
  x[1]=-(12.0/5.0)+(4.0/5.0)*x[11]
  x[2]=4+(1.0/3.0)*x[11]
  x[3]=(102.0/5.0)-(4.0/5.0)*x[11]
  x[4]=(62.0/5.0)-(7.0/15.0)*x[11]
  x[5]=(71.0/5.0)-(11.0/15.0)*x[11]
  x[6]=(46.0/5.0)+(3.0/5.0)*x[11]
  x[7]=(164.0/5.0)-(14.0/15.0)*x[11]
  x[8]=(124.0/5.0)-(4.0/15.0)*x[11]
  x[9]=(41.0/5.0)-(1.0/15.0)*x[11]
  x[10]=(849.0/5.0)-(3.0/5.0)*x[11]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve63 () int {
  var x [8]float64
  minSum := 99999.0
  minSol := ""

  x[1]=7
  x[2]=178
  x[3]=18
  x[4]=8
  x[5]=16
  x[6]=19
  x[7]=3

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve64 () int {
  var x [7]float64
  minSum := 99999.0
  minSol := ""

  for x[4]=0.0;x[4]<maxValue;x[4]++ {
  x[1]=11-x[4]
  x[2]=12-x[4]
  x[3]=7+x[4]
  x[5]=6
  x[6]=7

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve65 () int {
  var x [6]float64
  minSum := 99999.0
  minSol := ""

  x[1]=108
  x[2]=4
  x[3]=6
  x[4]=4
  x[5]=10

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve66 () int {
  var x [4]float64
  minSum := 99999.0
  minSol := ""

  x[1]=14
  x[2]=12
  x[3]=10

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve67 () int {
  var x [8]float64
  minSum := 99999.0
  minSol := ""

  x[1]=8
  x[2]=11
  x[3]=9
  x[4]=3
  x[5]=13
  x[6]=18
  x[7]=16

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve68 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  x[1]=0
  x[2]=12
  x[3]=9
  x[4]=12
  x[5]=20
  x[6]=0
  x[7]=17
  x[8]=7

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve69 () int {
  var x [7]float64
  minSum := 99999.0
  minSol := ""

  for x[5]=0.0;x[5]<maxValue;x[5]++ {
  x[1]=5-2*x[5]
  x[2]=17-x[5]
  x[3]=21-x[5]
  x[4]=7+2*x[5]
  x[6]=162

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve70 () int {
  var x [13]float64
  minSum := 99999.0
  minSol := ""

  for x[11]=0.0;x[11]<maxValue;x[11]++ {
  for x[12]=0.0;x[12]<maxValue;x[12]++ {
  x[1]=15-x[11]+x[12]
  x[2]=-5+x[12]
  x[3]=7-x[11]+2*x[12]
  x[4]=19-x[12]
  x[5]=190-x[11]+x[12]
  x[6]=2+x[11]-2*x[12]
  x[7]=16+2*x[11]-3*x[12]
  x[8]=9
  x[9]=2
  x[10]=3

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve71 () int {
  var x [5]float64
  minSum := 99999.0
  minSol := ""

  x[1]=1
  x[2]=132
  x[3]=19
  x[4]=20

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve72 () int {
  var x [6]float64
  minSum := 99999.0
  minSol := ""

  for x[5]=0.0;x[5]<maxValue;x[5]++ {
  x[1]=17
  x[2]=17-x[5]
  x[3]=25-x[5]
  x[4]=12

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve73 () int {
  var x [7]float64
  minSum := 99999.0
  minSol := ""

  for x[6]=0.0;x[6]<maxValue;x[6]++ {
  x[1]=14
  x[2]=18-x[6]
  x[3]=5+x[6]
  x[4]=24-x[6]
  x[5]=2

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve74 () int {
  var x [8]float64
  minSum := 99999.0
  minSol := ""

  x[1]=16
  x[2]=20
  x[3]=7
  x[4]=1
  x[5]=8
  x[6]=14
  x[7]=187

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve75 () int {
  var x [5]float64
  minSum := 99999.0
  minSol := ""

  x[1]=20
  x[2]=124
  x[3]=10
  x[4]=5

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve76 () int {
  var x [8]float64
  minSum := 99999.0
  minSol := ""

  for x[7]=0.0;x[7]<maxValue;x[7]++ {
  x[1]=-38+3*x[7]
  x[2]=-3+x[7]
  x[3]=11
  x[4]=-9+x[7]
  x[5]=47-2*x[7]
  x[6]=53-2*x[7]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve77 () int {
  var x [12]float64
  minSum := 99999.0
  minSol := ""

  for x[9]=0.0;x[9]<maxValue;x[9]++ {
  x[1]=-14+x[9]
  x[2]=23-x[9]
  x[3]=58-2*x[9]
  x[4]=165+x[9]
  x[5]=23-x[9]
  x[6]=50-2*x[9]
  x[7]=-12+x[9]
  x[8]=-1+x[9]
  x[10]=13
  x[11]=13

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve78 () int {
  var x [10]float64
  minSum := 99999.0
  minSol := ""

  x[1]=6
  x[2]=7
  x[3]=17
  x[4]=18
  x[5]=19
  x[6]=3
  x[7]=15
  x[8]=1
  x[9]=16

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve79 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  for x[6]=0.0;x[6]<maxValue;x[6]++ {
  for x[8]=0.0;x[8]<maxValue;x[8]++ {
  x[1]=8
  x[2]=-12+x[8]
  x[3]=8
  x[4]=11-x[6]
  x[5]=34-x[6]-x[8]
  x[7]=1

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve80 () int {
  var x [8]float64
  minSum := 99999.0
  minSol := ""

  x[1]=11
  x[2]=9
  x[3]=17
  x[4]=15
  x[5]=12
  x[6]=2
  x[7]=8

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve81 () int {
  var x [13]float64
  minSum := 99999.0
  minSol := ""

  for x[10]=0.0;x[10]<maxValue;x[10]++ {
  for x[11]=0.0;x[11]<maxValue;x[11]++ {
  x[1]=128-x[10]-x[11]
  x[2]=-2-(1.0/2.0)*x[10]+x[11]
  x[3]=14+(1.0/2.0)*x[10]
  x[4]=3+x[10]
  x[5]=23-x[11]
  x[6]=8-(1.0/2.0)*x[10]
  x[7]=-3+x[11]
  x[8]=16-x[10]
  x[9]=21+(1.0/2.0)*x[10]-x[11]
  x[12]=11

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve82 () int {
  var x [8]float64
  minSum := 99999.0
  minSol := ""

  for x[6]=0.0;x[6]<maxValue;x[6]++ {
  for x[7]=0.0;x[7]<maxValue;x[7]++ {
  x[1]=38-x[6]-x[7]
  x[2]=(47.0/2.0)-(1.0/2.0)*x[7]
  x[3]=-24+x[6]+x[7]
  x[4]=(43.0/2.0)-(1.0/2.0)*x[7]
  x[5]=(77.0/2.0)-x[6]-(1.0/2.0)*x[7]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve83 () int {
  var x [7]float64
  minSum := 99999.0
  minSol := ""

  for x[4]=0.0;x[4]<maxValue;x[4]++ {
  for x[6]=0.0;x[6]<maxValue;x[6]++ {
  x[1]=-94+x[4]-x[6]
  x[2]=114-x[4]
  x[3]=123-x[4]+x[6]
  x[5]=12-x[6]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve84 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  for x[7]=0.0;x[7]<maxValue;x[7]++ {
  for x[8]=0.0;x[8]<maxValue;x[8]++ {
  x[1]=-12+x[8]
  x[2]=-44+x[7]+3*x[8]
  x[3]=5+x[8]
  x[4]=-30+x[7]+x[8]
  x[5]=44-x[7]-2*x[8]
  x[6]=58-x[7]-3*x[8]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve85 () int {
  var x [11]float64
  minSum := 99999.0
  minSol := ""

  for x[10]=0.0;x[10]<maxValue;x[10]++ {
  x[1]=17
  x[2]=-10+x[10]
  x[3]=1+x[10]
  x[4]=4+x[10]
  x[5]=30-x[10]
  x[6]=24-x[10]
  x[7]=26-x[10]
  x[8]=x[10]
  x[9]=23-x[10]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve86 () int {
  var x [5]float64
  minSum := 99999.0
  minSol := ""

  x[1]=9
  x[2]=11
  x[3]=14
  x[4]=5

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve87 () int {
  var x [7]float64
  minSum := 99999.0
  minSol := ""

  for x[5]=0.0;x[5]<maxValue;x[5]++ {
  for x[6]=0.0;x[6]<maxValue;x[6]++ {
  x[1]=30-x[5]-x[6]
  x[2]=-3+x[6]
  x[3]=-3+x[5]+x[6]
  x[4]=32-x[5]-x[6]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve88 () int {
  var x [10]float64
  minSum := 99999.0
  minSol := ""

  x[1]=5
  x[2]=16
  x[3]=6
  x[4]=8
  x[5]=3
  x[6]=4
  x[7]=10
  x[8]=18
  x[9]=1

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve89 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  x[1]=5
  x[2]=1
  x[3]=16
  x[4]=13
  x[5]=5
  x[6]=20
  x[7]=7
  x[8]=1

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve90 () int {
  var x [10]float64
  minSum := 99999.0
  minSol := ""

  for x[8]=0.0;x[8]<maxValue;x[8]++ {
  for x[9]=0.0;x[9]<maxValue;x[9]++ {
  x[1]=18-x[8]+x[9]
  x[2]=4
  x[3]=19-x[9]
  x[4]=17-x[9]
  x[5]=139+x[9]
  x[6]=20+x[8]-2*x[9]
  x[7]=-4-x[8]+2*x[9]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve91 () int {
  var x [5]float64
  minSum := 99999.0
  minSol := ""

  x[1]=20
  x[2]=14
  x[3]=6
  x[4]=9

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve92 () int {
  var x [6]float64
  minSum := 99999.0
  minSol := ""

  x[1]=17
  x[2]=181
  x[3]=0
  x[4]=17
  x[5]=6

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve93 () int {
  var x [5]float64
  minSum := 99999.0
  minSol := ""

  x[1]=7
  x[2]=13
  x[3]=5
  x[4]=15

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve94 () int {
  var x [10]float64
  minSum := 99999.0
  minSol := ""

  for x[8]=0.0;x[8]<maxValue;x[8]++ {
  for x[9]=0.0;x[9]<maxValue;x[9]++ {
  x[1]=44-x[8]-x[9]
  x[2]=35-x[8]
  x[3]=10
  x[4]=53-x[8]-x[9]
  x[5]=-20+x[8]+x[9]
  x[6]=-15+x[8]
  x[7]=25-x[8]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve95 () int {
  var x [4]float64
  minSum := 99999.0
  minSol := ""

  x[1]=9
  x[2]=19
  x[3]=124

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve96 () int {
  var x [7]float64
  minSum := 99999.0
  minSol := ""

  x[1]=10
  x[2]=8
  x[3]=13
  x[4]=4
  x[5]=6
  x[6]=13

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve97 () int {
  var x [4]float64
  minSum := 99999.0
  minSol := ""

  x[1]=5
  x[2]=18
  x[3]=2

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve98 () int {
  var x [5]float64
  minSum := 99999.0
  minSol := ""

  x[1]=151
  x[2]=4
  x[3]=13
  x[4]=3

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve99 () int {
  var x [8]float64
  minSum := 99999.0
  minSol := ""

  x[1]=13
  x[2]=133
  x[3]=4
  x[4]=0
  x[5]=1
  x[6]=16
  x[7]=8

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve100 () int {
  var x [13]float64
  minSum := 99999.0
  minSol := ""

  for x[11]=0.0;x[11]<maxValue;x[11]++ {
  for x[12]=0.0;x[12]<maxValue;x[12]++ {
  x[1]=(825.0/4.0)-(3.0/4.0)*x[11]-(1.0/4.0)*x[12]
  x[2]=5+x[11]
  x[3]=(15.0/2.0)-(1.0/2.0)*x[11]+(1.0/2.0)*x[12]
  x[4]=-(7.0/2.0)+(1.0/2.0)*x[11]+(1.0/2.0)*x[12]
  x[5]=25-x[12]
  x[6]=(115.0/4.0)-(1.0/4.0)*x[11]-(3.0/4.0)*x[12]
  x[7]=(5.0/2.0)+(1.0/2.0)*x[11]+(1.0/2.0)*x[12]
  x[8]=(29.0/4.0)+(1.0/4.0)*x[11]+(3.0/4.0)*x[12]
  x[9]=(25.0/4.0)+(1.0/4.0)*x[11]-(1.0/4.0)*x[12]
  x[10]=(25.0/4.0)+(1.0/4.0)*x[11]-(1.0/4.0)*x[12]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve101 () int {
  var x [10]float64
  minSum := 99999.0
  minSol := ""

  for x[8]=0.0;x[8]<maxValue;x[8]++ {
  for x[9]=0.0;x[9]<maxValue;x[9]++ {
  x[1]=8-x[8]
  x[2]=15
  x[3]=14-x[8]-x[9]
  x[4]=13
  x[5]=7+x[8]
  x[6]=19-x[8]+x[9]
  x[7]=3-x[9]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve102 () int {
  var x [10]float64
  minSum := 99999.0
  minSol := ""

  for x[8]=0.0;x[8]<maxValue;x[8]++ {
  x[1]=24-x[8]
  x[2]=-1+x[8]
  x[3]=19
  x[4]=-5+x[8]
  x[5]=162-x[8]
  x[6]=1
  x[7]=34-x[8]
  x[9]=10

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve103 () int {
  var x [4]float64
  minSum := 99999.0
  minSol := ""

  x[1]=7
  x[2]=2
  x[3]=20

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve104 () int {
  var x [3]float64
  minSum := 99999.0
  minSol := ""

  x[1]=1
  x[2]=9

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve105 () int {
  var x [10]float64
  minSum := 99999.0
  minSol := ""

  for x[8]=0.0;x[8]<maxValue;x[8]++ {
  x[1]=17-x[8]
  x[2]=4+x[8]
  x[3]=13+x[8]
  x[4]=13
  x[5]=2+2*x[8]
  x[6]=12-3*x[8]
  x[7]=17-2*x[8]
  x[9]=15

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve106 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  for x[3]=0.0;x[3]<maxValue;x[3]++ {
  for x[8]=0.0;x[8]<maxValue;x[8]++ {
  x[1]=-(9.0/2.0)+x[3]+(1.0/2.0)*x[8]
  x[2]=27-x[3]
  x[4]=19
  x[5]=(1.0/2.0)-(1.0/2.0)*x[8]
  x[6]=12-x[8]
  x[7]=(15.0/2.0)-(1.0/2.0)*x[8]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve107 () int {
  var x [6]float64
  minSum := 99999.0
  minSol := ""

  x[1]=14
  x[2]=0
  x[3]=17
  x[4]=11
  x[5]=0

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve108 () int {
  var x [11]float64
  minSum := 99999.0
  minSol := ""

  x[1]=3
  x[2]=13
  x[3]=5
  x[4]=10
  x[5]=14
  x[6]=18
  x[7]=19
  x[8]=10
  x[9]=0
  x[10]=15

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve109 () int {
  var x [7]float64
  minSum := 99999.0
  minSol := ""

  for x[6]=0.0;x[6]<maxValue;x[6]++ {
  x[1]=32-x[6]
  x[2]=-14+x[6]
  x[3]=15
  x[4]=5
  x[5]=36-x[6]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve110 () int {
  var x [6]float64
  minSum := 99999.0
  minSol := ""

  x[1]=1
  x[2]=148
  x[3]=9
  x[4]=19
  x[5]=12

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve111 () int {
  var x [8]float64
  minSum := 99999.0
  minSol := ""

  x[1]=6
  x[2]=4
  x[3]=5
  x[4]=4
  x[5]=9
  x[6]=14
  x[7]=8

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve112 () int {
  var x [11]float64
  minSum := 99999.0
  minSol := ""

  x[1]=12
  x[2]=15
  x[3]=13
  x[4]=9
  x[5]=18
  x[6]=3
  x[7]=4
  x[8]=17
  x[9]=11
  x[10]=9

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve113 () int {
  var x [11]float64
  minSum := 99999.0
  minSol := ""

  for x[8]=0.0;x[8]<maxValue;x[8]++ {
  for x[9]=0.0;x[9]<maxValue;x[9]++ {
  x[1]=21-x[8]
  x[2]=46-2*x[8]-x[9]
  x[3]=131+x[8]
  x[4]=40-x[8]-x[9]
  x[5]=x[8]
  x[6]=19
  x[7]=12
  x[10]=11

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve114 () int {
  var x [11]float64
  minSum := 99999.0
  minSol := ""

  for x[10]=0.0;x[10]<maxValue;x[10]++ {
  x[1]=11+4*x[10]
  x[2]=27-4*x[10]
  x[3]=28-5*x[10]
  x[4]=-1+2*x[10]
  x[5]=-2+x[10]
  x[6]=21-3*x[10]
  x[7]=12+2*x[10]
  x[8]=8-x[10]
  x[9]=177+4*x[10]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve115 () int {
  var x [11]float64
  minSum := 99999.0
  minSol := ""

  x[1]=18
  x[2]=12
  x[3]=11
  x[4]=9
  x[5]=14
  x[6]=16
  x[7]=7
  x[8]=14
  x[9]=7
  x[10]=12

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve116 () int {
  var x [7]float64
  minSum := 99999.0
  minSol := ""

  for x[5]=0.0;x[5]<maxValue;x[5]++ {
  for x[6]=0.0;x[6]<maxValue;x[6]++ {
  x[1]=22-x[6]
  x[2]=21+x[5]-x[6]
  x[3]=20-x[5]
  x[4]=16-x[5]+x[6]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve117 () int {
  var x [6]float64
  minSum := 99999.0
  minSol := ""

  x[1]=10
  x[2]=15
  x[3]=2
  x[4]=8
  x[5]=15

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve118 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  for x[8]=0.0;x[8]<maxValue;x[8]++ {
  x[1]=4-x[8]
  x[2]=6+2*x[8]
  x[3]=189
  x[4]=19
  x[5]=6
  x[6]=11
  x[7]=18-x[8]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve119 () int {
  var x [13]float64
  minSum := 99999.0
  minSol := ""

  for x[9]=0.0;x[9]<maxValue;x[9]++ {
  for x[12]=0.0;x[12]<maxValue;x[12]++ {
  x[1]=7
  x[2]=17+x[9]-x[12]
  x[3]=21-x[12]
  x[4]=5
  x[5]=10
  x[6]=24-x[9]
  x[7]=11-x[9]
  x[8]=11-x[9]
  x[10]=49-2*x[12]
  x[11]=3+x[12]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve120 () int {
  var x [14]float64
  minSum := 99999.0
  minSol := ""

  for x[10]=0.0;x[10]<maxValue;x[10]++ {
  for x[12]=0.0;x[12]<maxValue;x[12]++ {
  for x[13]=0.0;x[13]<maxValue;x[13]++ {
  x[1]=-2+x[10]-x[12]+x[13]
  x[2]=-9-x[10]+7*x[12]-2*x[13]
  x[3]=2-x[10]+x[12]
  x[4]=177+x[10]-3*x[12]+x[13]
  x[5]=1-x[10]+3*x[12]-x[13]
  x[6]=-2*x[10]+6*x[12]-2*x[13]
  x[7]=11-3*x[12]+x[13]
  x[8]=14-x[12]-x[13]
  x[9]=39-5*x[12]+x[13]
  x[11]=-23+5*x[12]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve121 () int {
  var x [7]float64
  minSum := 99999.0
  minSol := ""

  x[1]=0
  x[2]=8
  x[3]=19
  x[4]=19
  x[5]=2
  x[6]=10

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve122 () int {
  var x [11]float64
  minSum := 99999.0
  minSol := ""

  for x[9]=0.0;x[9]<maxValue;x[9]++ {
  x[1]=32-2*x[9]
  x[2]=3+x[9]
  x[3]=-2+x[9]
  x[4]=x[9]
  x[5]=17
  x[6]=26-2*x[9]
  x[7]=36-3*x[9]
  x[8]=-4+x[9]
  x[10]=3

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve123 () int {
  var x [4]float64
  minSum := 99999.0
  minSol := ""

  x[1]=2
  x[2]=2
  x[3]=9

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve124 () int {
  var x [13]float64
  minSum := 99999.0
  minSol := ""

  for x[11]=0.0;x[11]<maxValue;x[11]++ {
  for x[12]=0.0;x[12]<maxValue;x[12]++ {
  x[1]=(44.0/3.0)+x[11]-(1.0/3.0)*x[12]
  x[2]=(89.0/3.0)-x[11]-(1.0/3.0)*x[12]
  x[3]=3
  x[4]=-(113.0/3.0)+(10.0/3.0)*x[12]
  x[5]=(101.0/3.0)-(4.0/3.0)*x[12]
  x[6]=-(50.0/3.0)+(4.0/3.0)*x[12]
  x[7]=49-2*x[12]
  x[8]=27-x[11]-x[12]
  x[9]=-5+x[12]
  x[10]=10+x[11]-x[12]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve125 () int {
  var x [4]float64
  minSum := 99999.0
  minSol := ""

  x[1]=7
  x[2]=17
  x[3]=6

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve126 () int {
  var x [7]float64
  minSum := 99999.0
  minSol := ""

  x[1]=0
  x[2]=12
  x[3]=12
  x[4]=1
  x[5]=7
  x[6]=4

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve127 () int {
  var x [11]float64
  minSum := 99999.0
  minSol := ""

  for x[8]=0.0;x[8]<maxValue;x[8]++ {
  for x[10]=0.0;x[10]<maxValue;x[10]++ {
  x[1]=(23.0/8.0)+x[8]+(3.0/8.0)*x[10]
  x[2]=(55.0/8.0)+(3.0/8.0)*x[10]
  x[3]=(49.0/8.0)+(5.0/8.0)*x[10]
  x[4]=(17.0/2.0)-(1.0/2.0)*x[10]
  x[5]=(37.0/4.0)-(3.0/4.0)*x[10]
  x[6]=(55.0/4.0)-x[8]-(1.0/4.0)*x[10]
  x[7]=(9.0/8.0)-x[8]+(5.0/8.0)*x[10]
  x[9]=(193.0/8.0)-(3.0/8.0)*x[10]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve128 () int {
  var x [12]float64
  minSum := 99999.0
  minSol := ""

  for x[11]=0.0;x[11]<maxValue;x[11]++ {
  x[1]=(19.0/4.0)-(1.0/4.0)*x[11]
  x[2]=-(1.0/2.0)+(1.0/2.0)*x[11]
  x[3]=-(45.0/2.0)+(5.0/2.0)*x[11]
  x[4]=(175.0/4.0)-(9.0/4.0)*x[11]
  x[5]=(11.0/2.0)+(1.0/2.0)*x[11]
  x[6]=(63.0/2.0)-(3.0/2.0)*x[11]
  x[7]=(233.0/4.0)-(11.0/4.0)*x[11]
  x[8]=(45.0/4.0)+(1.0/4.0)*x[11]
  x[9]=(27.0/4.0)+(3.0/4.0)*x[11]
  x[10]=(151.0/4.0)-(5.0/4.0)*x[11]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve129 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  for x[8]=0.0;x[8]<maxValue;x[8]++ {
  x[1]=6
  x[2]=30-x[8]
  x[3]=7
  x[4]=34-x[8]
  x[5]=-13+x[8]
  x[6]=17-x[8]
  x[7]=-17+x[8]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve130 () int {
  var x [5]float64
  minSum := 99999.0
  minSol := ""

  x[1]=11
  x[2]=19
  x[3]=18
  x[4]=1

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve131 () int {
  var x [5]float64
  minSum := 99999.0
  minSol := ""

  x[1]=5
  x[2]=5
  x[3]=175
  x[4]=11

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve132 () int {
  var x [12]float64
  minSum := 99999.0
  minSol := ""

  for x[9]=0.0;x[9]<maxValue;x[9]++ {
  for x[11]=0.0;x[11]<maxValue;x[11]++ {
  x[1]=8
  x[2]=4+x[9]-x[11]
  x[3]=-2+x[9]
  x[4]=20-x[11]
  x[5]=31-x[11]
  x[6]=19-x[9]+x[11]
  x[7]=22-x[9]
  x[8]=7-x[9]+x[11]
  x[10]=3

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve133 () int {
  var x [8]float64
  minSum := 99999.0
  minSol := ""

  for x[7]=0.0;x[7]<maxValue;x[7]++ {
  x[1]=(7.0/2.0)+(1.0/2.0)*x[7]
  x[2]=23-x[7]
  x[3]=(21.0/2.0)+(1.0/2.0)*x[7]
  x[4]=(35.0/2.0)-(1.0/2.0)*x[7]
  x[5]=(51.0/2.0)-(1.0/2.0)*x[7]
  x[6]=-1+x[7]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve134 () int {
  var x [12]float64
  minSum := 99999.0
  minSol := ""

  for x[10]=0.0;x[10]<maxValue;x[10]++ {
  for x[11]=0.0;x[11]<maxValue;x[11]++ {
  x[1]=43-3*x[10]-x[11]
  x[2]=-62+5*x[10]+4*x[11]
  x[3]=80+6*x[10]+3*x[11]
  x[4]=59-3*x[10]-3*x[11]
  x[5]=-54+4*x[10]+4*x[11]
  x[6]=6+x[10]
  x[7]=89-6*x[10]-6*x[11]
  x[8]=-41+3*x[10]+3*x[11]
  x[9]=-3+x[10]+2*x[11]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve135 () int {
  var x [11]float64
  minSum := 99999.0
  minSol := ""

  for x[7]=0.0;x[7]<maxValue;x[7]++ {
  for x[10]=0.0;x[10]<maxValue;x[10]++ {
  x[1]=-5+x[7]+x[10]
  x[2]=20-x[10]
  x[3]=-16-x[7]+x[10]
  x[4]=104-5*x[10]
  x[5]=26-x[10]
  x[6]=-53+4*x[10]
  x[8]=54-3*x[10]
  x[9]=-18+2*x[10]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve136 () int {
  var x [6]float64
  minSum := 99999.0
  minSol := ""

  for x[4]=0.0;x[4]<maxValue;x[4]++ {
  x[1]=38-x[4]
  x[2]=15
  x[3]=x[4]
  x[5]=15

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve137 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  x[1]=20
  x[2]=8
  x[3]=14
  x[4]=10
  x[5]=14
  x[6]=4
  x[7]=2
  x[8]=188

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve138 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  x[1]=3
  x[2]=10
  x[3]=17
  x[4]=1
  x[5]=1
  x[6]=12
  x[7]=4
  x[8]=149

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve139 () int {
  var x [7]float64
  minSum := 99999.0
  minSol := ""

  for x[6]=0.0;x[6]<maxValue;x[6]++ {
  x[1]=1+x[6]
  x[2]=9
  x[3]=8-x[6]
  x[4]=15
  x[5]=200

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve140 () int {
  var x [5]float64
  minSum := 99999.0
  minSol := ""

  x[1]=14
  x[2]=6
  x[3]=17
  x[4]=19

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve141 () int {
  var x [8]float64
  minSum := 99999.0
  minSol := ""

  x[1]=12
  x[2]=7
  x[3]=0
  x[4]=18
  x[5]=9
  x[6]=1
  x[7]=19

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve142 () int {
  var x [7]float64
  minSum := 99999.0
  minSol := ""

  x[1]=15
  x[2]=4
  x[3]=16
  x[4]=10
  x[5]=18
  x[6]=14

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve143 () int {
  var x [14]float64
  minSum := 99999.0
  minSol := ""

  for x[7]=0.0;x[7]<maxValue;x[7]++ {
  for x[12]=0.0;x[12]<maxValue;x[12]++ {
  for x[13]=0.0;x[13]<maxValue;x[13]++ {
  x[1]=11-(1.0/3.0)*x[13]
  x[2]=8-(2.0/3.0)*x[13]
  x[3]=24-x[7]-2*x[12]+2*x[13]
  x[4]=20-2*x[7]-2*x[12]+2*x[13]
  x[5]=15+x[7]-x[12]+(2.0/3.0)*x[13]
  x[6]=15+(1.0/3.0)*x[13]
  x[8]=x[12]-(5.0/3.0)*x[13]
  x[9]=25-x[12]+x[13]
  x[10]=-1+x[12]-(5.0/3.0)*x[13]
  x[11]=118+x[12]-(4.0/3.0)*x[13]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve144 () int {
  var x [7]float64
  minSum := 99999.0
  minSol := ""

  for x[4]=0.0;x[4]<maxValue;x[4]++ {
  x[1]=16
  x[2]=39-x[4]
  x[3]=-11+x[4]
  x[5]=14
  x[6]=0

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve145 () int {
  var x [3]float64
  minSum := 99999.0
  minSol := ""

  x[1]=7
  x[2]=0

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve146 () int {
  var x [7]float64
  minSum := 99999.0
  minSol := ""

  x[1]=20
  x[2]=14
  x[3]=16
  x[4]=4
  x[5]=20
  x[6]=5

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve147 () int {
  var x [3]float64
  minSum := 99999.0
  minSol := ""

  x[1]=18
  x[2]=2

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve148 () int {
  var x [4]float64
  minSum := 99999.0
  minSol := ""

  x[1]=157
  x[2]=9
  x[3]=6

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve149 () int {
  var x [6]float64
  minSum := 99999.0
  minSol := ""

  x[1]=4
  x[2]=16
  x[3]=1
  x[4]=153
  x[5]=0

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve150 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  x[1]=124
  x[2]=17
  x[3]=17
  x[4]=18
  x[5]=4
  x[6]=9
  x[7]=4
  x[8]=7

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve151 () int {
  var x [7]float64
  minSum := 99999.0
  minSol := ""

  x[1]=11
  x[2]=2
  x[3]=11
  x[4]=7
  x[5]=2
  x[6]=11

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve152 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  x[1]=13
  x[2]=14
  x[3]=8
  x[4]=10
  x[5]=15
  x[6]=17
  x[7]=6
  x[8]=13

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve153 () int {
  var x [3]float64
  minSum := 99999.0
  minSol := ""

  x[1]=7
  x[2]=2

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve154 () int {
  var x [6]float64
  minSum := 99999.0
  minSol := ""

  x[1]=6
  x[2]=19
  x[3]=12
  x[4]=16
  x[5]=3

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve155 () int {
  var x [6]float64
  minSum := 99999.0
  minSol := ""

  x[1]=13
  x[2]=9
  x[3]=2
  x[4]=6
  x[5]=18

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve156 () int {
  var x [10]float64
  minSum := 99999.0
  minSol := ""

  x[1]=14
  x[2]=7
  x[3]=11
  x[4]=20
  x[5]=17
  x[6]=15
  x[7]=9
  x[8]=12
  x[9]=10

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve157 () int {
  var x [5]float64
  minSum := 99999.0
  minSol := ""

  x[1]=4
  x[2]=0
  x[3]=3
  x[4]=9

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve158 () int {
  var x [10]float64
  minSum := 99999.0
  minSol := ""

  for x[7]=0.0;x[7]<maxValue;x[7]++ {
  for x[9]=0.0;x[9]<maxValue;x[9]++ {
  x[1]=-3+x[9]
  x[2]=4-x[7]
  x[3]=17+x[7]
  x[4]=7-x[7]
  x[5]=26-x[7]-x[9]
  x[6]=20+x[7]
  x[8]=129+x[9]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve159 () int {
  var x [12]float64
  minSum := 99999.0
  minSol := ""

  for x[10]=0.0;x[10]<maxValue;x[10]++ {
  for x[11]=0.0;x[11]<maxValue;x[11]++ {
  x[1]=(1.0/2.0)+x[10]+(1.0/2.0)*x[11]
  x[2]=1-x[10]+x[11]
  x[3]=(7.0/2.0)+(1.0/2.0)*x[11]
  x[4]=18-x[11]
  x[5]=1+x[11]
  x[6]=31-2*x[11]
  x[7]=(45.0/2.0)-x[10]-(1.0/2.0)*x[11]
  x[8]=17+x[10]-x[11]
  x[9]=11-x[10]+x[11]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve160 () int {
  var x [8]float64
  minSum := 99999.0
  minSol := ""

  x[1]=19
  x[2]=14
  x[3]=20
  x[4]=16
  x[5]=5
  x[6]=12
  x[7]=163

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve161 () int {
  var x [5]float64
  minSum := 99999.0
  minSol := ""

  x[1]=17
  x[2]=14
  x[3]=16
  x[4]=4

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve162 () int {
  var x [12]float64
  minSum := 99999.0
  minSol := ""

  for x[10]=0.0;x[10]<maxValue;x[10]++ {
  for x[11]=0.0;x[11]<maxValue;x[11]++ {
  x[1]=11
  x[2]=13
  x[3]=22-x[11]
  x[4]=36-x[10]-x[11]
  x[5]=20
  x[6]=22-x[10]
  x[7]=-18+x[11]
  x[8]=-19+x[11]
  x[9]=22-x[11]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve163 () int {
  var x [11]float64
  minSum := 99999.0
  minSol := ""

  x[1]=7
  x[2]=8
  x[3]=162
  x[4]=7
  x[5]=4
  x[6]=2
  x[7]=15
  x[8]=16
  x[9]=20
  x[10]=4

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve164 () int {
  var x [5]float64
  minSum := 99999.0
  minSol := ""

  x[1]=4
  x[2]=17
  x[3]=8
  x[4]=16

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve165 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  x[1]=6
  x[2]=2
  x[3]=13
  x[4]=5
  x[5]=11
  x[6]=3
  x[7]=4
  x[8]=2

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve166 () int {
  var x [12]float64
  minSum := 99999.0
  minSol := ""

  for x[9]=0.0;x[9]<maxValue;x[9]++ {
  for x[10]=0.0;x[10]<maxValue;x[10]++ {
  x[1]=0
  x[2]=59-x[9]-2*x[10]
  x[3]=25-x[9]
  x[4]=14
  x[5]=-9+x[9]+x[10]
  x[6]=-26+x[9]+2*x[10]
  x[7]=-15+x[9]+x[10]
  x[8]=50-x[9]-2*x[10]
  x[11]=3

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve167 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  x[1]=4
  x[2]=13
  x[3]=18
  x[4]=18
  x[5]=15
  x[6]=20
  x[7]=19
  x[8]=1

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve168 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  x[1]=20
  x[2]=10
  x[3]=11
  x[4]=18
  x[5]=13
  x[6]=19
  x[7]=15
  x[8]=13

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve169 () int {
  var x [6]float64
  minSum := 99999.0
  minSol := ""

  for x[5]=0.0;x[5]<maxValue;x[5]++ {
  x[1]=21-x[5]
  x[2]=30-2*x[5]
  x[3]=12-x[5]
  x[4]=-19+2*x[5]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve170 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  x[1]=4
  x[2]=14
  x[3]=19
  x[4]=14
  x[5]=8
  x[6]=20
  x[7]=14
  x[8]=20

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve171 () int {
  var x [7]float64
  minSum := 99999.0
  minSol := ""

  x[1]=12
  x[2]=6
  x[3]=13
  x[4]=14
  x[5]=5
  x[6]=5

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve172 () int {
  var x [11]float64
  minSum := 99999.0
  minSol := ""

  for x[5]=0.0;x[5]<maxValue;x[5]++ {
  for x[8]=0.0;x[8]<maxValue;x[8]++ {
  for x[9]=0.0;x[9]<maxValue;x[9]++ {
  x[1]=36-x[8]-x[9]
  x[2]=31-x[8]
  x[3]=34-x[5]-x[8]
  x[4]=-2+x[5]
  x[6]=x[8]
  x[7]=12+x[9]
  x[10]=8

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve173 () int {
  var x [6]float64
  minSum := 99999.0
  minSol := ""

  x[1]=16
  x[2]=16
  x[3]=17
  x[4]=5
  x[5]=4

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func solve174 () int {
  var x [9]float64
  minSum := 99999.0
  minSol := ""

  for x[7]=0.0;x[7]<maxValue;x[7]++ {
  for x[8]=0.0;x[8]<maxValue;x[8]++ {
  x[1]=20-x[8]
  x[2]=-5+x[8]
  x[3]=13+x[7]-2*x[8]
  x[4]=40-2*x[8]
  x[5]=-24-x[7]+4*x[8]
  x[6]=17-x[7]+x[8]

  skip := false
  sum := 0.0
  sol := ""
  for k:=1;k<len(x);k++ {
    x[k] = round(x[k], 4)
    if (x[k] < 0) || (math.Abs(x[k] - math.Round(x[k])) > 0.00001) { 
      skip = true
      break
    }
    sum += x[k]
    sol += strconv.Itoa(int(x[k])) + " "
  }
  if !skip {
    if sum < minSum {
      minSum = sum
      minSol = sol
    }
  }


  }
  }
  fmt.Println(int(minSum), ": ", minSol)
  return int(minSum)
}

func main() {
  answer := 0

  answer += solve1()
  answer += solve2()
  answer += solve3()
  answer += solve4()
  answer += solve5()
  answer += solve6()
  answer += solve7()
  answer += solve8()
  answer += solve9()
  answer += solve10()
  answer += solve11()
  answer += solve12()
  answer += solve13()
  answer += solve14()
  answer += solve15()
  answer += solve16()
  answer += solve17()
  answer += solve18()
  answer += solve19()
  answer += solve20()
  answer += solve21()
  answer += solve22()
  answer += solve23()
  answer += solve24()
  answer += solve25()
  answer += solve26()
  answer += solve27()
  answer += solve28()
  answer += solve29()
  answer += solve30()
  answer += solve31()
  answer += solve32()
  answer += solve33()
  answer += solve34()
  answer += solve35()
  answer += solve36()
  answer += solve37()
  answer += solve38()
  answer += solve39()
  answer += solve40()
  answer += solve41()
  answer += solve42()
  answer += solve43()
  answer += solve44()
  answer += solve45()
  answer += solve46()
  answer += solve47()
  answer += solve48()
  answer += solve49()
  answer += solve50()
  answer += solve51()
  answer += solve52()
  answer += solve53()
  answer += solve54()
  answer += solve55()
  answer += solve56()
  answer += solve57()
  answer += solve58()
  answer += solve59()
  answer += solve60()
  answer += solve61()
  answer += solve62()
  answer += solve63()
  answer += solve64()
  answer += solve65()
  answer += solve66()
  answer += solve67()
  answer += solve68()
  answer += solve69()
  answer += solve70()
  answer += solve71()
  answer += solve72()
  answer += solve73()
  answer += solve74()
  answer += solve75()
  answer += solve76()
  answer += solve77()
  answer += solve78()
  answer += solve79()
  answer += solve80()
  answer += solve81()
  answer += solve82()
  answer += solve83()
  answer += solve84()
  answer += solve85()
  answer += solve86()
  answer += solve87()
  answer += solve88()
  answer += solve89()
  answer += solve90()
  answer += solve91()
  answer += solve92()
  answer += solve93()
  answer += solve94()
  answer += solve95()
  answer += solve96()
  answer += solve97()
  answer += solve98()
  answer += solve99()
  answer += solve100()
  answer += solve101()
  answer += solve102()
  answer += solve103()
  answer += solve104()
  answer += solve105()
  answer += solve106()
  answer += solve107()
  answer += solve108()
  answer += solve109()
  answer += solve110()
  answer += solve111()
  answer += solve112()
  answer += solve113()
  answer += solve114()
  answer += solve115()
  answer += solve116()
  answer += solve117()
  answer += solve118()
  answer += solve119()
  answer += solve120()
  answer += solve121()
  answer += solve122()
  answer += solve123()
  answer += solve124()
  answer += solve125()
  answer += solve126()
  answer += solve127()
  answer += solve128()
  answer += solve129()
  answer += solve130()
  answer += solve131()
  answer += solve132()
  answer += solve133()
  answer += solve134()
  answer += solve135()
  answer += solve136()
  answer += solve137()
  answer += solve138()
  answer += solve139()
  answer += solve140()
  answer += solve141()
  answer += solve142()
  answer += solve143()
  answer += solve144()
  answer += solve145()
  answer += solve146()
  answer += solve147()
  answer += solve148()
  answer += solve149()
  answer += solve150()
  answer += solve151()
  answer += solve152()
  answer += solve153()
  answer += solve154()
  answer += solve155()
  answer += solve156()
  answer += solve157()
  answer += solve158()
  answer += solve159()
  answer += solve160()
  answer += solve161()
  answer += solve162()
  answer += solve163()
  answer += solve164()
  answer += solve165()
  answer += solve166()
  answer += solve167()
  answer += solve168()
  answer += solve169()
  answer += solve170()
  answer += solve171()
  answer += solve172()
  answer += solve173()
  answer += solve174()

  fmt.Println("Answer: ", answer)
}
