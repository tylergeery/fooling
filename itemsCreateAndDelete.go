package main

import (
    "errors"
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "net/url"
    "strconv"
)

const itemsJSONFile = "items.json"

func main() {
    http.HandleFunc("/shirts/create", addShirt)
    http.HandleFunc("/shirts/all", allShirts)
    http.HandleFunc("/shirts/delete", deleteShirt)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func getBodyParam(r *http.Request, key string) (string, error) {
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        return "", errors.New(fmt.Sprintf("Error reading body: %v", err))
    }

    jurl := ParseUrl(r.URL.String()+ "?" + string(body))
    params := jurl.params

    if params[key] != nil {
        return params[key][0], nil
    }

    return "", nil
}

func deleteShirt(w http.ResponseWriter, r *http.Request) {
    idString, parseErr := getBodyParam(r, "id")
    if parseErr != nil {
        http.Error(w, "can't read body", http.StatusBadRequest)
        return
    }

    id, _ := strconv.Atoi(idString)
    items, err := loadItems()

    if err != nil || id < 0 || id >= len(items) {
        http.Error(w, "cannot determine which shirt to delete", http.StatusBadRequest)
        return
    }

    items = append(items[:id], items[id+1:]...)

    err = saveItems(items)
    if err != nil {
        http.Error(w, "items could not be saved for deletion", http.StatusBadRequest)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

// Begin -------- Jons code ---------------------
func addShirt(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Printf("Error reading body: %v", err)
        http.Error(w, "can't read body", http.StatusBadRequest)
        return
    }
    jurl := ParseUrl(r.URL.String()+ "?" + string(body))
    params := jurl.params
    price, e := strconv.ParseFloat(safeGetFirstElem(params["price"]), 64)
    if e != nil {
        log.Printf("Error parsing price: %v", e)
        http.Error(w, "can't parse price", http.StatusBadRequest)
        return
    }
    addItem(safeGetFirstElem(params["brand"]), safeGetFirstElem(params["size"]), safeGetFirstElem(params["color"]), price)
    w.WriteHeader(201)
}

func allShirts(w http.ResponseWriter, r *http.Request) {
    items, err := loadItems()
    if err != nil {
        log.Printf("Error loading all shirts: %v", err)
        http.Error(w, "can't load all items", http.StatusBadRequest)
        return
    }

    if err != nil {
        log.Printf("Error loading all shirts: %v", err)
        http.Error(w, "can't load all items", http.StatusBadRequest)
        return
    }

    response := ""
    for _, item := range items {
        shirtJsonBytes, e := json.Marshal(item)
        if err != nil {
            log.Printf("Error serializing shirt: %v", e)
            http.Error(w, "can't load all items", http.StatusBadRequest)
            return
        }
        response += string(shirtJsonBytes) + "\n"
    }
    w.Write([]byte(response))
    w.WriteHeader(200)
}

func safeGetFirstElem(s []string) string {
    if len(s) > 0 {
        return s[0]
    }
    return ""
}
// End --------- Peppys code ---------------------

// Begin -------- Peppys code ---------------------
type Item struct {
    Brand string  `json:"brand"`
    Size  string  `json:"size"`
    Color string  `json:"color"`
    Price float64 `json:"price"`
}

func loadItems() ([]Item, error) {
    var items []Item

    content, err := ioutil.ReadFile(itemsJSONFile)
    if err != nil {
        return items, nil
    }

    err = json.Unmarshal(content, &items)
    if err != nil {
        return items, err
    }

    return items, nil
}

func saveItems(items []Item) error {
    bytes, err := json.Marshal(items)
    if err != nil {
        return err
    }

    return ioutil.WriteFile(itemsJSONFile, bytes, 0644)
}

func addItem(brand, size, color string, price float64) ([]Item, error) {
    items, err := loadItems()
    if err != nil {
        return nil, err
    }

    items = append(items, Item{brand, size, color, price})

    err = saveItems(items)
    if err != nil {
        return nil, err
    }

    return items, nil
}

// End --------- Peppys code ---------------------

// Begin --------------- Gustavos code ----------------
type JURL struct {
    protocol, host, path string
    params               url.Values
}

func ParseUrl(inputURL string) *JURL {
    u, err := url.Parse(inputURL)
    if err != nil {
        log.Fatal(err)
    }

    return &JURL{
        protocol: u.Scheme,
        host:     u.Host,
        path:     u.Path,
        params:   u.Query(),
    }
}

// End --------------- Gustavos code ------------
