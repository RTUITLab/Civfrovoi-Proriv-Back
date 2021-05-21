package coordinates

import (
	"gorm.io/gorm"
	pb "github.com/RTUITLab/Civfrovoi-Proriv-Back/protos/coordservice"
	"time"
)

type Coordinates struct {
	ID			string		`gorm:"primaryKey"`
	DateTime	time.Time
	Resource 	pb.Resuource
	Lat			float64
	Long		float64
}

func (c *Coordinates) ToPB() *pb.Coordinates {
	return &pb.Coordinates{
		Id: c.ID,
		Dt: c.DateTime.String(),
		Type: c.Resource,
		Lat: c.Lat,
		Long: c.Long,
	}
}

func FromPB(coord *pb.Coordinates) *Coordinates {
	DateTime, _ := time.Parse(time.Stamp, coord.Dt)
	return &Coordinates{
		ID: coord.Id,
		DateTime: DateTime,
		Resource: coord.Type,
		Lat: coord.Lat,
		Long: coord.Long,
	}
}

func InitTable(tx *gorm.DB) error {
	return tx.AutoMigrate(&Coordinates{})
}