package main

import "fmt"
type Menu struct {
    items []int
}

type FoodItem struct {
    name string
    price int
    started bool
    completed bool
    eaten bool
}

type Order struct {
    id int
    items []FoodItem
}

type Customer struct {
    name string
    order Order
}


func createMenu(foodItems []FoodItem) []FoodItem {
    foodItems = append(foodItems, FoodItem{name : "Chicken Biryani", price : 120, started : false, completed : false, eaten : false})
    foodItems = append(foodItems, FoodItem{name : "Veg Biryani", price : 100, started : false, completed : false, eaten : false})
    foodItems = append(foodItems, FoodItem{name : "Veg Sandwich", price : 100, started : false, completed : false, eaten : false})
    foodItems = append(foodItems, FoodItem{name : "Chicken Sandwich", price : 120, started : false, completed : false, eaten : false})
    foodItems = append(foodItems, FoodItem{name : "Chicken Salad", price : 100, started : false, completed : false, eaten : false})
    return foodItems
}

func createOrder(id int) Order {
    foodItems := make([]FoodItem, 0)
    return Order{id : id, items : foodItems}
}

func createCustomers(customers []Customer) []Customer {
    customers = append(customers, Customer{name : "A1", order : createOrder(0)})
    customers = append(customers, Customer{name : "A2", order : createOrder(1)})
    customers = append(customers, Customer{name : "A3", order : createOrder(2)})
    customers = append(customers, Customer{name : "A4", order : createOrder(3)})
    customers = append(customers, Customer{name : "A5", order : createOrder(4)})
    customers = append(customers, Customer{name : "A6", order : createOrder(5)})
    return customers
}

func main() {
    customers := make([]Customer, 0)
    orders := make([]Order, 0)
    foodItems := make([]FoodItem, 0)
    processingFoodItems := make([]FoodItem, 0)
    foodItems = createMenu(foodItems)
    customers = createCustomers(customers)
    fmt.Println(len(customers))
    fmt.Println(len(orders))
    fmt.Println(len(processingFoodItems))
    fmt.Println(len(foodItems))
}
