package main

import "time"

type Order struct {
    OrderID         string    `json:"orderId" bson:"orderId"`
    CustomerID      string    `json:"customerId" bson:"customerId"`
    CustomerName    string    `json:"customerName" bson:"customerName"`
    CustomerEmail   string    `json:"customerEmail" bson:"customerEmail"`
    RewardsNumber   string    `json:"rewardsNumber,omitempty" bson:"rewardsNumber,omitempty"`
    Items           []Item    `json:"items" bson:"items"`
    ShippingAddress Address   `json:"shippingAddress" bson:"shippingAddress"`
    PaymentMethod   string    `json:"paymentMethod" bson:"paymentMethod"`
    Subtotal        float64   `json:"subtotal" bson:"subtotal"`
    Tax             float64   `json:"tax" bson:"tax"`
    ShippingCost    float64   `json:"shippingCost" bson:"shippingCost"`
    ProtectionCost  float64   `json:"protectionCost" bson:"protectionCost"`
    RewardsDiscount float64   `json:"rewardsDiscount" bson:"rewardsDiscount"`
    Total           float64   `json:"total" bson:"total"`
    AddProtection   bool      `json:"addProtection" bson:"addProtection"`
    UseRewards      bool      `json:"useRewards" bson:"useRewards"`
    Status          Status    `json:"status" bson:"status"`
    CreatedAt       time.Time `json:"createdAt" bson:"createdAt"`
    UpdatedAt       time.Time `json:"updatedAt" bson:"updatedAt"`
    ProcessedAt     *time.Time `json:"processedAt,omitempty" bson:"processedAt,omitempty"`
}

type Status int

const (
    Pending Status = iota
    Processing
    Completed
    Failed
    Cancelled
)

func (s Status) String() string {
    return [...]string{"Pending", "Processing", "Completed", "Failed", "Cancelled"}[s]
}

type Item struct {
    ProductID int     `json:"productId" bson:"productId"`
    Name      string  `json:"name" bson:"name"`
    Brand     string  `json:"brand" bson:"brand"`
    Category  string  `json:"category" bson:"category"`
    Quantity  int     `json:"quantity" bson:"quantity"`
    Price     float64 `json:"price" bson:"price"`
    Total     float64 `json:"total" bson:"total"`
}

type Address struct {
    Street  string `json:"street" bson:"street"`
    City    string `json:"city" bson:"city"`
    State   string `json:"state" bson:"state"`
    ZipCode string `json:"zipCode" bson:"zipCode"`
}

type OrderRepo interface {
    GetPendingOrders() ([]Order, error)
    GetOrder(id string) (Order, error)
    InsertOrders(orders []Order) error
    UpdateOrder(order Order) error
}

type OrderService struct {
    repo OrderRepo
}

func NewOrderService(repo OrderRepo) *OrderService {
    return &OrderService{repo}
}