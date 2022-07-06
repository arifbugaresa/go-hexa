package dto

type GetListAbsensi struct {
	Date         string `json:"date"`
	CheckInTime  string `json:"check_in_time"`
	CheckOutTime string `json:"check_out_time"`
}
