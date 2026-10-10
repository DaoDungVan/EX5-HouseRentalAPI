package routers

import (
	"net/http"

	controllers "HouseRental/Controllers"
)

func BillRouter() {
	http.HandleFunc("/bills", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetAllBills(w, r)
		case http.MethodPost:
			controllers.CreateBill(w, r)
		default:
			controllers.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		}
	})

	http.HandleFunc("/bills/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetBillByID(w, r)
		case http.MethodPut:
			controllers.UpdateBillByID(w, r)
		case http.MethodDelete:
			controllers.DeleteBillByID(w, r)
		default:
			controllers.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		}
	})
}
