package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"shippingAddr/internal/config"
	"shippingAddr/internal/svc"
	pb "shippingAddr/shippingAddrs"
)

type Address struct {
	gorm.Model
	UserId          int64  `gorm:"index;comment('用户id')"`
	RecipientName   string `gorm:"varchar(255);comment('收货人姓名')"`
	PhoneNumber     string `gorm:"varchar(20);comment('联系电话')"`
	Location        string `gorm:"varchar(255);comment('所在地址,省市区县')"`
	DetailedAddress string `gorm:"varchar(255);comment('详细地址,街道门牌号')"`
	Tabs            string `gorm:"varchar(50);comment('标签')"`
	IsDefault       int64  `gorm:"int;comment('1开启默认收货地址,0否')"`
}

func withDB(c config.Config, fun func(db *gorm.DB) error) error {
	open, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.MysqlConf.Username, c.MysqlConf.Password, c.MysqlConf.Host, c.MysqlConf.Port, c.MysqlConf.DatabasesName)), &gorm.Config{})
	if err != nil {
		return err
	}
	defer func() {
		db, _ := open.DB()
		db.Close()
	}()
	return fun(open)
}

func Sync(c config.Config) error {
	log.Println(c.MysqlConf)
	return withDB(c, func(db *gorm.DB) error {
		return db.AutoMigrate(new(Address))
	})
}

func NewAddress() *Address {
	return new(Address)
}

func (a *Address) Create(c *svc.ServiceContext) error {
	return withDB(c.Config, func(db *gorm.DB) error {
		return db.Create(a).Error
	})
}

func (a *Address) Update(c *svc.ServiceContext) error {
	return withDB(c.Config, func(db *gorm.DB) error {
		return db.Where("id = ?", a.ID).Updates(a).Error
	})
}

func (a *Address) Delete(c *svc.ServiceContext) error {
	return withDB(c.Config, func(db *gorm.DB) error {
		return db.Delete(a).Error
	})
}

func (a *Address) Get(c *svc.ServiceContext, where map[string]interface{}) error {
	return withDB(c.Config, func(db *gorm.DB) error {
		return db.Where(where).First(a).Error
	})
}

func (a *Address) Gets(c *svc.ServiceContext, where map[string]interface{}) ([]*Address, error) {
	addresses := make([]*Address, 0)
	err := withDB(c.Config, func(db *gorm.DB) error {
		return db.Where(where).Find(&addresses).Error
	})
	if err != nil {
		return nil, err
	}
	return addresses, nil
}
func SqlToPb(a *Address) *pb.ShippingAddress {
	return &pb.ShippingAddress{
		ID:              int64(a.ID),
		UserID:          a.UserId,
		RecipientName:   a.RecipientName,
		PhoneNumber:     a.PhoneNumber,
		Location:        a.Location,
		DetailedAddress: a.DetailedAddress,
		Tabs:            a.Tabs,
		IsDefault:       a.IsDefault,
	}
}

func PbToSql(a *pb.ShippingAddress) *Address {
	return &Address{
		Model:           gorm.Model{ID: uint(a.ID)},
		UserId:          a.UserID,
		RecipientName:   a.RecipientName,
		PhoneNumber:     a.PhoneNumber,
		Location:        a.Location,
		DetailedAddress: a.DetailedAddress,
		Tabs:            a.Tabs,
		IsDefault:       a.IsDefault,
	}
}
