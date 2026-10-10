package services

import (
	configs "HouseRental/Configs"
	models "HouseRental/Models"
)

func GetAllBills() ([]models.Bill, error) {
	rows, err := configs.MySQLDB.Query(
		`SELECT id, room_name, tenant_name, month, room_price,
			electricity_old, electricity_new, electricity_price,
			water_old, water_new, water_price,
			internet_fee, garbage_fee, other_fee,
			electricity_cost, water_cost, total
		FROM bills ORDER BY id`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// tạo sẵn slice rỗng, không thì lúc chưa có hóa đơn JSON sẽ ra null chứ không phải []
	bills := []models.Bill{}
	for rows.Next() {
		var b models.Bill
		// Scan phải đúng thứ tự cột ở câu SELECT
		err := rows.Scan(
			&b.ID, &b.RoomName, &b.TenantName, &b.Month, &b.RoomPrice,
			&b.ElectricityOld, &b.ElectricityNew, &b.ElectricityPrice,
			&b.WaterOld, &b.WaterNew, &b.WaterPrice,
			&b.InternetFee, &b.GarbageFee, &b.OtherFee,
			&b.ElectricityCost, &b.WaterCost, &b.Total,
		)
		if err != nil {
			return nil, err
		}
		bills = append(bills, b)
	}
	return bills, nil
}

// trả thêm bool để biết có tìm thấy hóa đơn không,
// controller dựa vào đó trả 404, còn có err thì trả 500
func GetBillByID(id int) (models.Bill, bool, error) {
	var b models.Bill

	rows, err := configs.MySQLDB.Query(
		`SELECT id, room_name, tenant_name, month, room_price,
			electricity_old, electricity_new, electricity_price,
			water_old, water_new, water_price,
			internet_fee, garbage_fee, other_fee,
			electricity_cost, water_cost, total
		FROM bills WHERE id = ?`,
		id,
	)
	if err != nil {
		return b, false, err
	}
	defer rows.Close()

	if !rows.Next() {
		return b, false, nil
	}

	err = rows.Scan(
		&b.ID, &b.RoomName, &b.TenantName, &b.Month, &b.RoomPrice,
		&b.ElectricityOld, &b.ElectricityNew, &b.ElectricityPrice,
		&b.WaterOld, &b.WaterNew, &b.WaterPrice,
		&b.InternetFee, &b.GarbageFee, &b.OtherFee,
		&b.ElectricityCost, &b.WaterCost, &b.Total,
	)
	if err != nil {
		return b, false, err
	}
	return b, true, nil
}

func CreateBill(bill models.Bill) (models.Bill, error) {
	result, err := configs.MySQLDB.Exec(
		`INSERT INTO bills (
			room_name, tenant_name, month, room_price,
			electricity_old, electricity_new, electricity_price,
			water_old, water_new, water_price,
			internet_fee, garbage_fee, other_fee,
			electricity_cost, water_cost, total
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		bill.RoomName, bill.TenantName, bill.Month, bill.RoomPrice,
		bill.ElectricityOld, bill.ElectricityNew, bill.ElectricityPrice,
		bill.WaterOld, bill.WaterNew, bill.WaterPrice,
		bill.InternetFee, bill.GarbageFee, bill.OtherFee,
		bill.ElectricityCost, bill.WaterCost, bill.Total,
	)
	if err != nil {
		return models.Bill{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Bill{}, err
	}

	bill.ID = int(id)
	return bill, nil
}

func UpdateBill(id int, bill models.Bill) (models.Bill, error) {
	_, err := configs.MySQLDB.Exec(
		`UPDATE bills SET
			room_name = ?, tenant_name = ?, month = ?, room_price = ?,
			electricity_old = ?, electricity_new = ?, electricity_price = ?,
			water_old = ?, water_new = ?, water_price = ?,
			internet_fee = ?, garbage_fee = ?, other_fee = ?,
			electricity_cost = ?, water_cost = ?, total = ?
		WHERE id = ?`,
		bill.RoomName, bill.TenantName, bill.Month, bill.RoomPrice,
		bill.ElectricityOld, bill.ElectricityNew, bill.ElectricityPrice,
		bill.WaterOld, bill.WaterNew, bill.WaterPrice,
		bill.InternetFee, bill.GarbageFee, bill.OtherFee,
		bill.ElectricityCost, bill.WaterCost, bill.Total,
		id,
	)
	if err != nil {
		return models.Bill{}, err
	}

	bill.ID = id
	return bill, nil
}

func DeleteBill(id int) error {
	_, err := configs.MySQLDB.Exec("DELETE FROM bills WHERE id = ?", id)
	return err
}
