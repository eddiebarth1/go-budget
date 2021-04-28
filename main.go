package main

import (
  "bufio"
  "flag"
  "fmt"
  "encoding/json"
  "math"
  "log"
  "os"
  "strconv"
  "strings"
)

// Helper method to check for errors, simple enough but will
// save us some duplicate code
func errCheck(err error) {
    if err != nil {
        log.Fatal(err)
        panic(err)
    }
}


func readExpenses(expensesPointer string) map[string]float64 {
    file, openErr := os.Open(expensesPointer)
    expenseMap := make(map[string]float64)

    errCheck(openErr)

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
      lineSplit := strings.Split(scanner.Text(), ",")

      expenseName := lineSplit[0]
      // We need to get rid of the whitespace from having spaces in the
      // expense file
      expenseAmtStr := strings.ReplaceAll(lineSplit[1], " ", "")
      expenseAmt, strconvErr := strconv.ParseFloat(expenseAmtStr, 64)
      errCheck(strconvErr)

      expenseMap[expenseName] = expenseAmt

    }

    return expenseMap
}


func main() {
    var availBal   float64
    var billsTotal float64
    var paycheck   float64
    var file       string

    flag.Float64Var(&paycheck, "paycheck", 3671.00, "Amount of `paycheck` to use for calculation")
    flag.StringVar(&file, "file", "/home/eddie/repos/py-budget/start_month", "Expense file `path`")

    flag.Parse()

    expenses := readExpenses(file)

    for _, amt := range expenses {
        billsTotal += amt
    }

    billsTotal = math.Round(billsTotal*100/100)
    availBal = math.Round(availBal*100/100)

    prettyExpenses, err := json.MarshalIndent(expenses, "", "  ")
    errCheck(err)

    availBal = paycheck - billsTotal
    fmt.Println("Total pay for this period is: $", paycheck, "\n")
    fmt.Println("Expenses for this period are: \n\n", string(prettyExpenses), "\n ============<>============\n")
    fmt.Println("Total bill payments: $", billsTotal)
    fmt.Printf("Your remaining balance after bills is: $%v \n", availBal)

}
