package main

import (
    "bufio"
    "errors"
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
    "strings"
)

const itemFile = "items.txt"

type Item struct {
    Brand string
    Size string
    Color string
    Price float64
}

func (i Item) Output() string {
    return fmt.Sprintf("%v %v %v $%v", i.Brand, i.Size, i.Color, i.Price)
}

func createItem(itemString string) (Item, error) {
    var err error
    var item Item
    itemParts := strings.Split(itemString, " ")

    if len(itemParts) != 4 {
        err = errors.New("Items must contain exactly these 4 parts <brand> <size> <color> <price>")
        return item, err
    }

    if len(itemParts[3]) > 0 && itemParts[3][0] == '$' {
        itemParts[3] = itemParts[3][1:]
    }

    val, parseErr := strconv.ParseFloat(itemParts[3], 64);
    if parseErr != nil {
        return item, parseErr
    }

    return Item{ itemParts[0], itemParts[1], itemParts[2], val }, nil
}

func outputItems(items []Item) {
    for _, i := range items {
        fmt.Println(i.Output())
    }
}

func readInitialItems() []Item {
    items := []Item{}

    if len(os.Args) > 1 && string(os.Args[1]) == "new" {
        return items
    }

    contents, err := ioutil.ReadFile(itemFile)
    if err != nil {
        fmt.Printf("Read file error: %v\n", err)
        return items
    }

    itemLines := strings.Split(string(contents), "\n")

    for _, itemLine := range itemLines {
        item, createErr := createItem(itemLine)

        if createErr == nil {
            items = append(items, item)
        }
    }

    return items
}

func writeItemsToFile(items []Item) {
    output := ""
    for i, item := range items {
        output += item.Output()
        if i != (len(items) - 1) {
            output += "\n"
        }
    }

    err := ioutil.WriteFile(itemFile, []byte(output), 0644)
    if err != nil {
        fmt.Printf("File could not be written to: %v", err)
    }
}

func main() {
    items := readInitialItems()
    reader := bufio.NewReader(os.Stdin)
    // wait until input is "exit"

    for true {
        fmt.Print("Add item (<brand> <size> <color> <price>): ")
        itemBytes, err := reader.ReadString('\n')

        if err != nil {
            fmt.Println("Could not read input, try again")
            continue;
        }

        itemString := string(itemBytes)
        itemString = itemString[:len(itemString) - 1]

        if itemString == "quit" || itemString == "exit" {
            fmt.Println("Saving your items for next time")
            break;
        }

        item, createErr := createItem(itemString)
        if createErr != nil {
            continue
        }

        items = append(items, item)
        outputItems(items)
    }

    writeItemsToFile(items)
}
