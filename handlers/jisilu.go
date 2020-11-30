package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"qtmd-php/model"
	"qtmd-php/util"
	"strconv"
	"strings"
	"time"
)

type Cell struct {
	BondID                string      `json:"bond_id"`
	BondNm                string      `json:"bond_nm"`
	StockID               string      `json:"stock_id"`
	StockNm               string      `json:"stock_nm"`
	Btype                 string      `json:"btype"`
	ConvertPrice          string      `json:"convert_price"`
	ConvertPriceValidFrom string      `json:"convert_price_valid_from"`
	ConvertDt             string      `json:"convert_dt"`
	MaturityDt            string      `json:"maturity_dt"`
	NextPutDt             string      `json:"next_put_dt"`
	PutDt                 interface{} `json:"put_dt"`
	PutNotes              interface{} `json:"put_notes"`
	PutPrice              string      `json:"put_price"`
	PutIncCpnFl           string      `json:"put_inc_cpn_fl"`
	PutConvertPriceRatio  string      `json:"put_convert_price_ratio"`
	PutCountDays          int         `json:"put_count_days"`
	PutTotalDays          int         `json:"put_total_days"`
	PutRealDays           int         `json:"put_real_days"`
	RepoDiscountRt        string      `json:"repo_discount_rt"`
	RepoValidFrom         interface{} `json:"repo_valid_from"`
	RepoValidTo           interface{} `json:"repo_valid_to"`
	TurnoverRt            string      `json:"turnover_rt"`
	RedeemPrice           string      `json:"redeem_price"`
	RedeemIncCpnFl        string      `json:"redeem_inc_cpn_fl"`
	RedeemPriceRatio      string      `json:"redeem_price_ratio"`
	RedeemCountDays       int         `json:"redeem_count_days"`
	RedeemTotalDays       int         `json:"redeem_total_days"`
	RedeemRealDays        int         `json:"redeem_real_days"`
	RedeemDt              interface{} `json:"redeem_dt"`
	RedeemFlag            string      `json:"redeem_flag"`
	OrigIssAmt            string      `json:"orig_iss_amt"`
	CurrIssAmt            string      `json:"curr_iss_amt"`
	RatingCd              string      `json:"rating_cd"`
	IssuerRatingCd        string      `json:"issuer_rating_cd"`
	Guarantor             string      `json:"guarantor"`
	SscDt                 interface{} `json:"ssc_dt"`
	EscDt                 interface{} `json:"esc_dt"`
	ScNotes               interface{} `json:"sc_notes"`
	ForceRedeem           interface{} `json:"force_redeem"`
	RealForceRedeemPrice  interface{} `json:"real_force_redeem_price"`
	ConvertCd             string      `json:"convert_cd"`
	RepoCd                interface{} `json:"repo_cd"`
	Ration                interface{} `json:"ration"`
	RationCd              string      `json:"ration_cd"`
	ApplyCd               string      `json:"apply_cd"`
	OnlineOfflineRatio    interface{} `json:"online_offline_ratio"`
	Qflag                 string      `json:"qflag"`
	Qflag2                string      `json:"qflag2"`
	RationRt              string      `json:"ration_rt"`
	FundRt                string      `json:"fund_rt"`
	MarginFlg             string      `json:"margin_flg"`
	Pb                    string      `json:"pb"`
	PbFlag                string      `json:"pb_flag"`
	TotalShares           string      `json:"total_shares"`
	Sqflg                 string      `json:"sqflg"`
	Sprice                string      `json:"sprice"`
	Svolume               string      `json:"svolume"`
	SincreaseRt           string      `json:"sincrease_rt"`
	Qstatus               string      `json:"qstatus"`
	BondValue             string      `json:"bond_value"`
	BondValue2            string      `json:"bond_value2"`
	LastTime              string      `json:"last_time"`
	ConvertValue          string      `json:"convert_value"`
	PremiumRt             string      `json:"premium_rt"`
	YearLeft              string      `json:"year_left"`
	YtmRt                 string      `json:"ytm_rt"`
	YtmRtTax              string      `json:"ytm_rt_tax"`
	Price                 string      `json:"price"`
	FullPrice             string      `json:"full_price"`
	IncreaseRt            string      `json:"increase_rt"`
	Volume                string      `json:"volume"`
	ConvertPriceValid     string      `json:"convert_price_valid"`
	AdjScnt               int         `json:"adj_scnt"`
	AdjCnt                int         `json:"adj_cnt"`
	RedeemIcon            string      `json:"redeem_icon"`
	RefYieldInfo          string      `json:"ref_yield_info"`
	AdjustTip             string      `json:"adjust_tip"`
	Adjusted              string      `json:"adjusted"`
	OptionTip             string      `json:"option_tip"`
	BondValue3            string      `json:"bond_value3"`
	LeftPutYear           string      `json:"left_put_year"`
	ShortMaturityDt       string      `json:"short_maturity_dt"`
	Dblow                 string      `json:"dblow"`
	ForceRedeemPrice      string      `json:"force_redeem_price"`
	PutConvertPrice       string      `json:"put_convert_price"`
	ConvertAmtRatio       string      `json:"convert_amt_ratio"`
	StockNetValue         string      `json:"stock_net_value"`
	StockCd               string      `json:"stock_cd"`
	PreBondID             string      `json:"pre_bond_id"`
	RepoValid             string      `json:"repo_valid"`
	ConvertCdTip          string      `json:"convert_cd_tip"`
	PriceTips             string      `json:"price_tips"`
}

type Rows struct {
	Id   int64 `json:"id"`
	Cell *Cell `json:"cell"`
}

type Rest struct {
	Page  int64   `json:"page"`
	Total int64   `json:"total"`
	Rows  []*Rows `json:"rows"`
}

func KzzDb(c *gin.Context) {
	t := time.Now()
	url := "https://www.jisilu.cn/data/cbnew/cb_list/?___jsl=LST___t=1605493453079&fprice=&tprice=&curr_iss_amt=&volume=&svolume=&premium_rt=&ytm_rt=&rating_cd=&is_search=N&market_cd%5B%5D=shmb&market_cd%5B%5D=szmb&market_cd%5B%5D=szzx&market_cd%5B%5D=szcy&btype=&listed=Y&sw_cd=&bond_ids=&rp=50&page=1"
	req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Http get err: ", err)
		c.JSON(http.StatusOK, err)
	}
	if resp.StatusCode != 200 {
		fmt.Println("Http StatusCode: ", resp.StatusCode)
		c.JSON(http.StatusOK, "")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read err: ", err)
		c.JSON(http.StatusOK, "")
	}
	var tempMap Rest
	json.Unmarshal(body, &tempMap)

	for _, item := range tempMap.Rows {
		//fmt.Println(item.Cell.PriceTips)
		if item.Cell.PriceTips == "待上市" || item.Cell.Btype == "E" { // 踢出未上市和交换债
			continue
		}

		//1.按照ytm(到期税后年化收益率)从大到小排序
		//2.踢出ytm小于0的
		//3.踢出正股PB大于2的
		//4.50%ytm+50%溢价率
		ytmt, _ := strconv.ParseFloat(strings.Replace(item.Cell.YtmRt, "%", "", 5), 10) // 税前到期收益
		ytm_premium_rt := 0.0
		if ytmt < 0.0 {
			ytm_premium_rt = 0.0
		} else {
			premium_rt, _ := strconv.ParseFloat(strings.Replace(item.Cell.PremiumRt, "%", "", 5), 10)
			ytmts := ytmt * 0.5
			premium_rts := premium_rt * 0.5
			ytm_premium_rt = ytmts + premium_rts
		}
		ytm_premium_rts := strconv.FormatFloat(ytm_premium_rt, 'e', -1, 64)

		//premium_rt, _ := strconv.ParseFloat(strings.Replace(item.Cell.PremiumRt, "%", "", 5), 10)
		//ytmts := ytmt * 0.5
		//premium_rts := premium_rt * 0.5
		//ytm_premium_rt := strconv.FormatFloat(ytmts + premium_rts, 'e', -1, 64)

		kzz := &model.Kzz{
			BondID:       item.Cell.BondID,
			BondNm:       item.Cell.BondNm,
			Pb:           item.Cell.Pb,
			PremiumRt:    item.Cell.PremiumRt,
			OrigIssAmt:   item.Cell.OrigIssAmt,
			Volume:       item.Cell.Volume,
			TurnoverRt:   item.Cell.TurnoverRt,
			Dblow:        item.Cell.Dblow,
			YtmRt:        item.Cell.YtmRt,
			Price:        item.Cell.Price,
			YtmPremiumRt: ytm_premium_rts,
		}
		go func() {
			kz, err := kzz.GetByBondID(kzz.BondID)
			if err == gorm.ErrRecordNotFound {
				_, err := kzz.Create()
				if err != nil {
					util.REST(c, gin.H{"err": err.Error()})
					return
				}
			} else {
				kzz.ID = kz.ID
				kzz.Update(kzz)
			}
		}()

		//_, err := kzz.Create()
	}
	elapsed := time.Since(t).Seconds()
	util.REST(c, gin.H{"data": util.SUCCESS, "elapsedTime": elapsed})
}

func KzzTest() {
	t := time.Now()
	url := "https://www.jisilu.cn/data/cbnew/cb_list/?___jsl=LST___t=1605493453079&fprice=&tprice=&curr_iss_amt=&volume=&svolume=&premium_rt=&ytm_rt=&rating_cd=&is_search=N&market_cd%5B%5D=shmb&market_cd%5B%5D=szmb&market_cd%5B%5D=szzx&market_cd%5B%5D=szcy&btype=&listed=Y&sw_cd=&bond_ids=&rp=50&page=1"
	req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Http get err: ", err)
		return
	}
	if resp.StatusCode != 200 {
		fmt.Println("Http StatusCode: ", resp.StatusCode)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read err: ", err)
		return
	}
	var tempMap Rest
	json.Unmarshal(body, &tempMap)

	for _, item := range tempMap.Rows {
		//fmt.Println(item.Cell.PriceTips)
		if item.Cell.PriceTips == "待上市" || item.Cell.Btype == "E" { // 踢出未上市和交换债
			continue
		}

		//1.按照ytm(到期税后年化收益率)从大到小排序
		//2.踢出ytm小于0的
		//3.踢出正股PB大于2的
		//4.50%ytm+50%溢价率
		ytmt, _ := strconv.ParseFloat(strings.Replace(item.Cell.YtmRt, "%", "", 5), 10) // 税前到期收益
		ytm_premium_rt := 0.0
		if ytmt < 0.0 {
			ytm_premium_rt = 0.0
		} else {
			premium_rt, _ := strconv.ParseFloat(strings.Replace(item.Cell.PremiumRt, "%", "", 5), 10)
			ytmts := ytmt * 0.5
			premium_rts := premium_rt * 0.5
			ytm_premium_rt = ytmts + premium_rts
		}
		ytm_premium_rts := strconv.FormatFloat(ytm_premium_rt, 'e', -1, 64)

		//premium_rt, _ := strconv.ParseFloat(strings.Replace(item.Cell.PremiumRt, "%", "", 5), 10)
		//ytmts := ytmt * 0.5
		//premium_rts := premium_rt * 0.5
		//ytm_premium_rt := strconv.FormatFloat(ytmts + premium_rts, 'e', -1, 64)

		kzz := &model.Kzz{
			BondID:       item.Cell.BondID,
			BondNm:       item.Cell.BondNm,
			Pb:           item.Cell.Pb,
			PremiumRt:    item.Cell.PremiumRt,
			OrigIssAmt:   item.Cell.OrigIssAmt,
			Volume:       item.Cell.Volume,
			TurnoverRt:   item.Cell.TurnoverRt,
			Dblow:        item.Cell.Dblow,
			YtmRt:        item.Cell.YtmRt,
			Price:        item.Cell.Price,
			YtmPremiumRt: ytm_premium_rts,
		}
		go func() {
			kz, err := kzz.GetByBondID(kzz.BondID)
			if err == gorm.ErrRecordNotFound {
				_, err := kzz.Create()
				if err != nil {
					fmt.Println("create err: ", err.Error())
					return
				}
			} else {
				kzz.ID = kz.ID
				kzz.Update(kzz)
			}
		}()

		//_, err := kzz.Create()
	}
	elapsed := time.Since(t)
	fmt.Println("elapsedTime: ", elapsed)
}

// yyb 双低
func KzzYyb(c *gin.Context) {
	t := time.Now()
	url := "https://www.jisilu.cn/data/cbnew/cb_list/?___jsl=LST___t=1605493453079&fprice=&tprice=&curr_iss_amt=&volume=&svolume=&premium_rt=&ytm_rt=&rating_cd=&is_search=N&market_cd%5B%5D=shmb&market_cd%5B%5D=szmb&market_cd%5B%5D=szzx&market_cd%5B%5D=szcy&btype=&listed=Y&sw_cd=&bond_ids=&rp=50&page=1"
	req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Http get err: ", err)
		return
	}
	if resp.StatusCode != 200 {
		fmt.Println("Http StatusCode: ", resp.StatusCode)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read err: ", err)
		return
	}
	var tempMap Rest
	json.Unmarshal(body, &tempMap)

	for _, item := range tempMap.Rows {
		//fmt.Println(item.Cell.PriceTips)
		if item.Cell.PriceTips == "待上市" || item.Cell.Btype == "E" { // 踢出未上市和交换债
			continue
		}

		// 双低 低价+溢价率*100
		premium_rt, _ := strconv.ParseFloat(strings.Replace(item.Cell.PremiumRt, "%", "", 5), 10) // 溢价率
		price, _ := strconv.ParseFloat(item.Cell.Price, 10)
		dbslow := price + premium_rt

		kzzyyb := &model.Kzzyyb{
			BondID:     item.Cell.BondID,
			BondNm:     item.Cell.BondNm,
			Pb:         item.Cell.Pb,
			PremiumRt:  item.Cell.PremiumRt,
			OrigIssAmt: item.Cell.OrigIssAmt,
			Volume:     item.Cell.Volume,
			TurnoverRt: item.Cell.TurnoverRt,
			Dblow:      dbslow,
			YtmRt:      item.Cell.YtmRt,
			Price:      item.Cell.Price,
		}
		//kz, err := kzzyyb.GetByBondID(kzzyyb.BondID)
		//if err == gorm.ErrRecordNotFound {
		//	_, err := kzzyyb.Create()
		//	if err != nil {
		//		fmt.Println("create err: ", err.Error())
		//		return
		//	}
		//} else {
		//	kzzyyb.ID = kz.ID
		//	kzzyyb.Update(kzzyyb)
		//}

		go func() {
			kz, err := kzzyyb.GetByBondID(kzzyyb.BondID)
			if err == gorm.ErrRecordNotFound {
				_, err := kzzyyb.Create()
				if err != nil {
					fmt.Println("create err: ", err.Error())
					return
				}
			} else {
				kzzyyb.ID = kz.ID
				kzzyyb.Update(kzzyyb)
			}
		}()
	}
	elapsed := time.Since(t)
	fmt.Println("elapsedTime: ", elapsed)
}

// todo 持有封基
