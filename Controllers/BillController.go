package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	models "HouseRental/Models"
	services "HouseRental/Services"
)

// hàm dùng chung để trả JSON kèm status code
func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, map[string]string{"error": message})
}

// lấy id trong URL /bills/{id}
func getID(r *http.Request) (int, error) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/bills/"))
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetAllBills(w http.ResponseWriter, r *http.Request) {
	bills, err := services.GetAllBills()
	if err != nil {
		fmt.Println("GetAllBills error:", err)
		WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	WriteJSON(w, http.StatusOK, bills)
}

func GetBillByID(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "id must be an integer")
		return
	}

	bill, found, err := services.GetBillByID(id)
	if err != nil {
		fmt.Println("GetBillByID error:", err)
		WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	if !found {
		WriteError(w, http.StatusNotFound, "bill not found")
		return
	}
	WriteJSON(w, http.StatusOK, bill)
}

func CreateBill(w http.ResponseWriter, r *http.Request) {
	var bill models.Bill
	if err := json.NewDecoder(r.Body).Decode(&bill); err != nil {
		WriteError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	if err := bill.Validate(); err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	// tiền điện, nước, tổng tiền server tự tính, không lấy số client gửi lên
	bill.Calculate()

	created, err := services.CreateBill(bill)
	if err != nil {
		fmt.Println("CreateBill error:", err)
		WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	WriteJSON(w, http.StatusCreated, created)
}

func UpdateBillByID(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "id must be an integer")
		return
	}

	var bill models.Bill
	if err := json.NewDecoder(r.Body).Decode(&bill); err != nil {
		WriteError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	if err := bill.Validate(); err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	_, found, err := services.GetBillByID(id)
	if err != nil {
		fmt.Println("UpdateBillByID error:", err)
		WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	if !found {
		WriteError(w, http.StatusNotFound, "bill not found")
		return
	}

	// sửa xong phải tính lại tiền
	bill.Calculate()

	updated, err := services.UpdateBill(id, bill)
	if err != nil {
		fmt.Println("UpdateBillByID error:", err)
		WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	WriteJSON(w, http.StatusOK, updated)
}

func DeleteBillByID(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "id must be an integer")
		return
	}

	_, found, err := services.GetBillByID(id)
	if err != nil {
		fmt.Println("DeleteBillByID error:", err)
		WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	if !found {
		WriteError(w, http.StatusNotFound, "bill not found")
		return
	}

	if err := services.DeleteBill(id); err != nil {
		fmt.Println("DeleteBillByID error:", err)
		WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	WriteJSON(w, http.StatusOK, map[string]string{
		"message": fmt.Sprintf("delete bill with id: %d success", id),
	})
}
