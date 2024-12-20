package configs

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

var SnapClient snap.Client

func InitMidtrans() {
	midtrans.ServerKey = "SB-Mid-server-tEhXc7DgvrhK9vysgHwMU-bF"
	midtrans.Environment = midtrans.Production // Gunakan Production untuk mode produksi

	SnapClient.New(midtrans.ServerKey, midtrans.Environment)
}
