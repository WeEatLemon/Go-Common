package redis

import "time"

const (
	TimeMinuteToOne     = time.Minute        //	一分钟
	TimeMinuteToFive    = time.Minute * 5    //	五分钟
	TimeMinuteToFIFTEEN = time.Minute * 15   //	十五分钟
	TimeHourToHalf      = time.Minute * 30   //	半小时
	TimeHourToOne       = time.Hour          //	一小时
	TimeDayToOne        = time.Hour * 24     //	一天
	TimeDayToSeven      = time.Hour * 24 * 7 //	七天
)
