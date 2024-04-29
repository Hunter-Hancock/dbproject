package handler

import (
	"fmt"
	"net/http"

	"github.com/Hunter-Hancock/dbproject/db"
	mw "github.com/Hunter-Hancock/dbproject/middleware"

	"github.com/Hunter-Hancock/dbproject/view/checkout"
	"github.com/Hunter-Hancock/dbproject/view/order"
)

type OrderHandler struct {
	CartStore  db.CartStore
	OrderStore db.OrderStore
}

func NewOrderHandler(db db.OrderStore, cdb db.CartStore) *OrderHandler {
	return &OrderHandler{OrderStore: db, CartStore: cdb}
}

func (oh *OrderHandler) HandleShow(w http.ResponseWriter, r *http.Request) {
	methods := oh.HandleGetPaymentMethods(w, r)

	checkout.CheckoutPage(methods).Render(r.Context(), w)
}

func (oh *OrderHandler) ShowOrders(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(mw.UserIDKey).(*db.User)
	orderIDs, err := oh.OrderStore.GetOrders(user.Customer.ID)
	if err != nil {
		fmt.Println(err)
	}

	orders, err2 := oh.OrderStore.GetOrderDetail(orderIDs)
	if err2 != nil {
		fmt.Println(err2)
	}

	for _, order := range orders {
		foodItem, err := oh.OrderStore.GetFoodByID(order.ItemID)
		if err != nil {
			fmt.Println(err)
		}
		order.Item = foodItem
	}

	order.OrdersPage(orders).Render(r.Context(), w)
}

func (oh *OrderHandler) HandleProcessPayment(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(mw.UserIDKey).(*db.User)
	r.ParseForm()

	paymentMethod := r.FormValue("paymentMethod")

	order, err := oh.OrderStore.CreateOrder(user.Customer.ID, paymentMethod)
	if err != nil {
		fmt.Println(err)
		return
	}

	items, err2 := oh.CartStore.GetItems(user.Customer.ID)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	_, err3 := oh.OrderStore.CreateDetail(order.ID, items)
	if err3 != nil {
		fmt.Println(err3)
		return
	}

	http.Redirect(w, r, "/orders", http.StatusSeeOther)
}

func (oh *OrderHandler) HandleGetPaymentMethods(w http.ResponseWriter, r *http.Request) []db.PaymentMethod {
	methods, err := oh.OrderStore.GetPaymentMethods()
	if err != nil {
		return nil
	}

	return methods
}
