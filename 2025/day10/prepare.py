import re
import unicodedata
import time
from playwright.sync_api import sync_playwright

def replace_unicode_digit(match):
    char = match.group(0)
    try:
        return str(unicodedata.digit(char))
    except ValueError:
        return char

header = """
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

"""

function_footer = """
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

"""

with open("input.txt", 'r') as file:
    lines = file.readlines()

go_file = open("main.go", "w")
print(header, file=go_file)

i = 1
with sync_playwright() as p:
    browser = p.chromium.launch()
    

    for line in lines:
        elements = line.split(' ')
        target = elements[-1][1:-2].split(',')
        buttons = elements[1:-1]
        m = []
        m = [0 for i in range(len(buttons) * len(target))]
        k = 0
        for button in buttons:
            for b in button[1:-1].split(','):
                m[int(b)*len(buttons) + k] = 1
            k = k + 1
        url = 'https://www.matrixcalc.org/slu.html#solve-using-Gaussian-elimination%28%7B%7B'
        for k in range(len(m)):
            q, r = divmod(k, len(buttons))
            if k>0 and r==0:
                url = url + target[q-1] + '%7D'
                if k < len(m)-1:
                    url = url + ',%7B'
            url = url + str(m[k])
            if k < len(m)-1:
                url = url + ','
        url = url + ',' + target[-1]
        url = url + '%7D%7D%29'

        print(url)

        page = browser.new_page()
        page.goto(url)
        time.sleep(5)
        
        equations = page.locator('ul.list-unstyled[data-i="1"] li').all_inner_texts()

        numVars = len(equations)
        freeVars = []
        assignments = []
        for eq in equations:
            #eq = re.sub(r'\d', replace_unicode_digit, eq, flags=re.UNICODE)
            #eq = unicodedata.normalize('NFKD', eq).encode('ascii', 'ignore').decode('ascii')
            eq = eq.replace('⁢', '*')
            eq = eq.replace('𝑥', 'x')
            eq = eq.replace('⋅', '*')
            eq = re.sub(r'(\d+)\n(\d+)', r'(\1.0/\2.0)', eq)
            eq = eq.replace("\n", "")
            eq = eq.replace('−', '-')
            eq = eq.replace(',', '.')
            s1, s2 = eq.split('=')
            if s1 == s2:
                s1 = s1[1:]
                freeVars.append(int(s1))
                continue
            eq = re.sub(r'x(\d+)',r'x[\1]', eq)
            assignments.append(eq)

        print("func solve" + str(i) + " () int {", file=go_file)
        print("  var x [" + str(numVars+1) + "]float64", file=go_file)
        print('  minSum := 99999.0', file=go_file)
        print('  minSol := ""', file=go_file)
        print('', file=go_file)

        for freeVar in freeVars:
            print(f"  for x[{freeVar}]=0.0;x[{freeVar}]<maxValue;x[{freeVar}]++ {{", file=go_file)
    
        for assignment in assignments:
            print("  " + assignment, file=go_file)

        print(function_footer, file=go_file)
    
        for freeVar in freeVars:
            print("  }", file=go_file)

        print('  fmt.Println(int(minSum), ": ", minSol)', file=go_file)
        print('  return int(minSum)', file=go_file)
        print("}", file=go_file)
        print("", file=go_file)

        page.close()

        i = i + 1

print('func main() {', file=go_file)
print('  answer := 0', file=go_file)
print('', file=go_file)
for k in range(1, i):
    print('  answer += solve' + str(k) + '()', file=go_file)
print('', file=go_file)
print('  fmt.Println("Answer: ", answer)', file=go_file)
print('}', file=go_file)

go_file.close()
