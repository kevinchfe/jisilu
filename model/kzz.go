package model

import (
	"qtmd-php/database"
)

type Kzz struct {
	ID int64 `gorm:"column:id; PRIMARY_KEY; autoIncrement"`
	//*gorm.Model
	BondID       string `gorm:"column:bond_id"`
	BondNm       string `gorm:"column:bond_nm"`
	Pb           string `gorm:"column:pb"`
	PremiumRt    string `gorm:"column:premium_rt"`
	OrigIssAmt   string `gorm:"column:orig_iss_amt"`
	Volume       string `gorm:"column:volume"`
	TurnoverRt   string `gorm:"column:turnover_rt; default:0"`
	Price        string `gorm:"column:price"`
	YtmRt        string `gorm:"column:ytm_rt"`
	Dblow        string `gorm:"column:dblow"`
	YtmPremiumRt string `gorm:"ytm_premium_rt"`
}

type Kzzyyb struct {
	ID int64 `gorm:"column:id; PRIMARY_KEY; autoIncrement"`
	//*gorm.Model
	BondID       string `gorm:"column:bond_id"`
	BondNm       string `gorm:"column:bond_nm"`
	Pb           string `gorm:"column:pb"`
	PremiumRt    string `gorm:"column:premium_rt"`
	OrigIssAmt   string `gorm:"column:orig_iss_amt"`
	Volume       string `gorm:"column:volume"`
	TurnoverRt   string `gorm:"column:turnover_rt; default:0"`
	Price        string `gorm:"column:price"`
	YtmRt        string `gorm:"column:ytm_rt"`
	Dblow        float64 `gorm:"column:dblow"`
}

func (k Kzz) Create() (*Kzz, error) {
	if err := database.DB.Create(&k).Error; err != nil {
		return nil, err
	}
	return &k, nil
}

func (k Kzz) GetByBondID(bond_id string) (*Kzz, error) {
	if err := database.DB.Where("bond_id=?", bond_id).First(&k).Error; err != nil {
		return nil, err
	}
	return &k, nil
}

func (k Kzz) Update(values interface{}) error {
	if err := database.DB.Model(&k).Updates(values).Where("id=?", k.ID).Error; err != nil {
		return err
	}
	return nil
}

// yyb
func (k Kzzyyb) Create() (*Kzzyyb, error) {
	if err := database.DB.Create(&k).Error; err != nil {
		return nil, err
	}
	return &k, nil
}

func (k Kzzyyb) GetByBondID(bond_id string) (*Kzzyyb, error) {
	if err := database.DB.Where("bond_id=?", bond_id).First(&k).Error; err != nil {
		return nil, err
	}
	return &k, nil
}

func (k Kzzyyb) Update(values interface{}) error {
	if err := database.DB.Model(&k).Updates(values).Where("id=?", k.ID).Error; err != nil {
		return err
	}
	return nil
}

