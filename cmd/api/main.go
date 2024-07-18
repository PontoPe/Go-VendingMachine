package main

import (
	"VendingMachineCLI/internal/server"
	"fmt"
    "strings"
    "os"
    "os/exec"
    "math"
)

type Item struct{
        id int
        Name string
        Price float64
        Stock int
}

var array []Item
var temp string


func roundFloat(val float64, precision uint) float64 {
    ratio := math.Pow(10, float64(precision))
    return math.Round(val*ratio) / ratio
}

func main() {
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

    cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
    cmd.Stdout = os.Stdout
    cmd.Run()

    array = []Item{
        {1, "coca", 3.2, 10},
        {2, "suco", 2.1, 15},
    }
    agua := Item {3, "agua", 1.6, 30}
    array = append(array, agua)

    menu()

}

func menu() {
    cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
    cmd.Stdout = os.Stdout
    cmd.Run()

    var u int

    fmt.Println("WELCOME TO PONTOPE's VENDING MACHINE!")
    fmt.Println("\n\nType 1 to buy!")

    fmt.Scan(&u)
    switch u {
        case 1:
            NormalMode()
        case 2:
            AdmMode()
        case 0:
            fmt.Println("Goodbye!")
    }
}
func NormalMode() {    
    var i, y int
    var price, payment float64
    cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
    cmd.Stdout = os.Stdout
    cmd.Run()
    
    fmt.Println("Heya! Here are the available items:")
    fmt.Printf("\n+%s+\n", strings.Repeat("-", 48))
    fmt.Println("| id:       | Name:     | Price:    | Stock:")
    fmt.Printf("+%s+\n", strings.Repeat("-", 48))

    for x := 0; x < len(array); x++{
        fmt.Printf("| %-10d| %-10s| %-10g| %-10d\n", array[x].id, array[x].Name, array[x].Price, array[x].Stock)
        fmt.Printf("+%s+\n", strings.Repeat("-", 48))
    }

    fmt.Println("What would you like to buy?")
    fmt.Scan(&i)

    for y = 0; y < len(array); y++{
        if i==array[y].id {
            price = array[y].Price
            fmt.Println(price)
            break
        }
    }

    fmt.Printf("That would be %g reais. Insert money amount:", price)
    fmt.Scan(&payment)
    change := payment - price
    if change < 0{
        fmt.Println("Payment has to be larger than price!")
    } else {
        fmt.Printf("Thank you! Your change is %g\n", roundFloat(change, 2))
        array[y].Stock--
        fmt.Println("Would you like to buy more?")
        fmt.Scan(&temp)
        if temp == "y" {
            NormalMode()
        } else {
            fmt.Println("Goodbye!")
        }
    }

}


func AdmMode() {
    var i int
    cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
    cmd.Stdout = os.Stdout
    cmd.Run()

    fmt.Println("Hello, admin!\n what would you like to do?")
    fmt.Println("1:Add item\n2:Edit item name\n3:Edit item stock\n4:Remove item")
    fmt.Scan(&i)
    
    switch i {
        case 1:
            Adm_add_item()
        case 2:
            Adm_edit_name()
        case 3:
            Adm_edit_stock()
        case 4:
            Adm_remove_item()
    }


}

func Adm_add_item(){
    var i, tempName string
    var tempPrice float64
    var tempID, tempStock int

    fmt.Println("Please type new item name:")
    fmt.Scan(&tempName)
    fmt.Println("Please type new item price:")
    fmt.Scan(&tempPrice)
    fmt.Println("Please type new item stock:")
    fmt.Scan(&tempStock)
    tempID = len(array) + 1

    fmt.Printf("id: %d\nName: %s\nPrice: %g\nStock:%d\n\n", tempID, tempName, tempPrice, tempStock)

    NewItem := Item {tempID, tempName, tempPrice, tempStock}

    array = append(array, NewItem)

    fmt.Println("add more?")
    fmt.Scan(&i)

    switch i {
        case "y":
            Adm_add_item()
    }
    menu()

}

func Adm_edit_name() {
    
    fmt.Println("Heya! Here are the available items:")
    fmt.Printf("\n+%s+\n", strings.Repeat("-", 48))
    fmt.Println("| id:       | Name:     | Price:    | Stock:")
    fmt.Printf("+%s+\n", strings.Repeat("-", 48))

    for x := 0; x < len(array); x++{
        fmt.Printf("| %-10d| %-10s| %-10g| %-10d\n", array[x].id, array[x].Name, array[x].Price, array[x].Stock)
        fmt.Printf("+%s+\n", strings.Repeat("-", 48))
    }

    fmt.Println("\nWhich item would you like to edit?")
    var i, y int
    var temp string
    fmt.Scan(&i)
    fmt.Println("What is the new name?")
    fmt.Scan(&temp)


    for y = 0; y < len(array); y++{
        if i==array[y].id {
            array[y].Name = temp            
            fmt.Println("Name changed!")
            fmt.Println(array[y].Name)
            break
        }
    }

    fmt.Println("New list of items:")
    fmt.Printf("\n+%s+\n", strings.Repeat("-", 48))
    fmt.Println("| id:       | Name:     | Price:    | Stock:")
    fmt.Printf("+%s+\n", strings.Repeat("-", 48))

    for x := 0; x < len(array); x++{
        fmt.Printf("| %-10d| %-10s| %-10g| %-10d\n", array[x].id, array[x].Name, array[x].Price, array[x].Stock)
        fmt.Printf("+%s+\n", strings.Repeat("-", 48))
    }
    fmt.Println("Would you like to edit another item?")
    fmt.Scan(&temp)
    if temp == "y" {
        Adm_edit_name()
    } else {
        menu()
    }

}

func Adm_edit_stock() {
    
    fmt.Println("Heya! Here are the available items:")
    fmt.Printf("\n+%s+\n", strings.Repeat("-", 48))
    fmt.Println("| id:       | Name:     | Price:    | Stock:")
    fmt.Printf("+%s+\n", strings.Repeat("-", 48))

    for x := 0; x < len(array); x++{
        fmt.Printf("| %-10d| %-10s| %-10g| %-10d\n", array[x].id, array[x].Name, array[x].Price, array[x].Stock)
        fmt.Printf("+%s+\n", strings.Repeat("-", 48))
    }

    fmt.Println("\nWhich item would you like to edit?")
    var i, y, temp int
    var temps string
    fmt.Scan(&i)
    fmt.Println("What is the new stock?")
    fmt.Scan(&temp)


    for y = 0; y < len(array); y++{
        if i==array[y].id {
            array[y].Stock = temp            
            fmt.Println("Stock changed!")
            fmt.Println(array[y].Stock)
            break
        }
    }

    fmt.Println("New list of items:")
    fmt.Printf("\n+%s+\n", strings.Repeat("-", 48))
    fmt.Println("| id:       | Name:     | Price:    | Stock:")
    fmt.Printf("+%s+\n", strings.Repeat("-", 48))

    for x := 0; x < len(array); x++{
        fmt.Printf("| %-10d| %-10s| %-10g| %-10d\n", array[x].id, array[x].Name, array[x].Price, array[x].Stock)
        fmt.Printf("+%s+\n", strings.Repeat("-", 48))
    }
    fmt.Println("Would you like to edit another item?")
    fmt.Scan(&temps)
    if temps == "y" {
        Adm_edit_stock()
    } else {
        menu()
    }


}

func Adm_remove_item() {
    
    fmt.Println("Heya! Here are the available items:")
    fmt.Printf("\n+%s+\n", strings.Repeat("-", 48))
    fmt.Println("| id:       | Name:     | Price:    | Stock:")
    fmt.Printf("+%s+\n", strings.Repeat("-", 48))

    for x := 0; x < len(array); x++{
        fmt.Printf("| %-10d| %-10s| %-10g| %-10d\n", array[x].id, array[x].Name, array[x].Price, array[x].Stock)
        fmt.Printf("+%s+\n", strings.Repeat("-", 48))
    }

    fmt.Println("\nWhich item would you like to remove?")
    var i, y int
    var tempS string
    fmt.Scan(&i)
    if i > len(array) {
        fmt.Println("Invalid id!")
        Adm_remove_item()
    } else {
        for y = 0; y < len(array); y++{
            if i==array[y].id {
                array = append(array[:i-1], array[i:]...)
                fmt.Println("Item removed!")
            }
            }
        }

        fmt.Println("New list of items:")
        fmt.Printf("\n+%s+\n", strings.Repeat("-", 48))
        fmt.Println("| id:       | Name:     | Price:    | Stock:")
        fmt.Printf("+%s+\n", strings.Repeat("-", 48))

        for x := 0; x < len(array); x++{
            fmt.Printf("| %-10d| %-10s| %-10g| %-10d\n", array[x].id, array[x].Name, array[x].Price, array[x].Stock)
            fmt.Printf("+%s+\n", strings.Repeat("-", 48))
        }
        fmt.Println("Would you like to remove another item?")
        fmt.Scan(&tempS)
        if tempS == "y" {
            Adm_remove_item()
        } else {
            menu()
        }
}

