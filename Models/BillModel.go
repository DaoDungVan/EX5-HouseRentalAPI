package models

import (
	"errors"
	"strings"
)

type Bill struct {
	ID               int     `json:"id"`
	RoomName         string  `json:"room_name"`
	TenantName       string  `json:"tenant_name"`
	Month            string  `json:"month"`
	RoomPrice        float64 `json:"room_price"`
	ElectricityOld   int     `json:"electricity_old"`
	ElectricityNew   int     `json:"electricity_new"`
	ElectricityPrice float64 `json:"electricity_price"`
	WaterOld         int     `json:"water_old"`
	WaterNew         int     `json:"water_new"`
	WaterPrice       float64 `json:"water_price"`
	InternetFee      float64 `json:"internet_fee"`
	GarbageFee       float64 `json:"garbage_fee"`
	OtherFee         float64 `json:"other_fee"`
	ElectricityCost  float64 `json:"electricity_cost"`
	WaterCost        float64 `json:"water_cost"`
	Total            float64 `json:"total"`
}

// kiểm tra dữ liệu theo bảng validate trong đề, sai chỗ nào thì trả lỗi chỗ đó
func (b Bill) Validate() error {
	if strings.TrimSpace(b.RoomName) == "" {
		return errors.New("room_name must not be empty")
	}
	if strings.TrimSpace(b.TenantName) == "" {
		return errors.New("tenant_name must not be empty")
	}
	if strings.TrimSpace(b.Month) == "" {
		return errors.New("month must not be empty")
	}
	if b.RoomPrice < 0 {
		return errors.New("room_price must be >= 0")
	}
	if b.ElectricityNew < b.ElectricityOld {
		return errors.New("electricity_new must be >= electricity_old")
	}
	if b.WaterNew < b.WaterOld {
		return errors.New("water_new must be >= water_old")
	}
	if b.ElectricityPrice < 0 {
		return errors.New("electricity_price must be >= 0")
	}
	if b.WaterPrice < 0 {
		return errors.New("water_price must be >= 0")
	}
	if b.InternetFee < 0 {
		return errors.New("internet_fee must be >= 0")
	}
	if b.GarbageFee < 0 {
		return errors.New("garbage_fee must be >= 0")
	}
	if b.OtherFee < 0 {
		return errors.New("other_fee must be >= 0")
	}
	return nil
}

// tính tiền điện, tiền nước, tổng tiền theo công thức trong đề
// phải dùng pointer receiver thì mới sửa được bill gốc
func (b *Bill) Calculate() {
	b.ElectricityCost = float64(b.ElectricityNew-b.ElectricityOld) * b.ElectricityPrice
	b.WaterCost = float64(b.WaterNew-b.WaterOld) * b.WaterPrice
	b.Total = b.RoomPrice + b.ElectricityCost + b.WaterCost +
		b.InternetFee + b.GarbageFee + b.OtherFee
}
