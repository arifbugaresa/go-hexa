package absensi

import "time"

type Absensi struct {
	ID           int64
	UserInfoID   int64
	CheckInTime  time.Time
	CheckIn      bool
	CheckOutTime time.Time
	CheckOut     bool
	CreatedBy    string
	CreatedAt    time.Time
	UpdatedBy    string
	UpdatedAt    time.Time
	Deleted      bool
}
