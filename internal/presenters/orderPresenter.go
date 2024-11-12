package presenters

// import (
// 	"time"
// 	"trinity-be/internal/entities"
// )

// // OrderPresenter prepares order data for presentation.
// type OrderPresenter struct{}

// // NewOrderPresenter creates a new instance of OrderPresenter.
// func NewOrderPresenter() *OrderPresenter {
//     return &OrderPresenter{}
// }

// // Present formats the order data for output.
// func (p *OrderPresenter) Present(order entities.Order) map[string]interface{} {
//     items := make([]map[string]interface{}, len(order.Items))
//     for i, item := range order.Items {
//         items[i] = map[string]interface{}{
//             "product_id": item.ProductID,
//             "quantity":   item.Quantity,
//         }
//     }

//     return map[string]interface{}{
//         "id":         order.ID,
//         "customer_id": order.UserID,
//         "order_date":  order.CreatedAt.Add(7 * time.Hour).Format("2006-01-02 15:04:05"),
//         "items":       items,
//     }
// }
