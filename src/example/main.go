package main

import (
	"notificorerest"
	"log"
	"time"
	"math/rand"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func main() {
	//HLR
	var hlrClient = notificorerest.NewHlrClient("https://app.notificore.com/rest", "test_JhIRDMsbo4prEsXCc3AW")
	//Balance
	var balance = hlrClient.GetBalance()
	log.Println(balance)
	//Prices
	var hlrPrices = hlrClient.GetHlrPricesById(9)
	log.Println(hlrPrices)
	//Single Hlr
	var hlr = hlrClient.CreateHlrWithParameters("380972920000", "ext_"+RandStringRunes(9), nil, "http://someurl.com/callback")
	log.Println(hlr)
	//Multiple Hlr
	var tariff = "9"
	var hlr2 = hlrClient.CreateHlrs(
		[]notificorerest.HlrRequest{
			notificorerest.HlrRequest{Tariff:nil, Reference:"ext_" + RandStringRunes(9),
				Msisdn:"380972920001", CallbackUrl:"http://someurl.com/callback"},
				     {Tariff:&tariff, Reference:"ext_" + RandStringRunes(9),
					     Msisdn:"380972920002", CallbackUrl:"http://someurl.com/callback"}})
	log.Println(hlr2)
	//Hlr info by id
	var hlrInfo = hlrClient.GetHlrInfoById(hlr.Hlrs[0].ID)
	log.Println(hlrInfo)
	//Hlr info by external id
	var hlrInfo2 = hlrClient.GetHlrInfoByReference(hlr2.Hlrs[1].Reference)
	log.Println(hlrInfo2)
	//SMS
	var smsClient = notificorerest.NewSmsClient("https://app.notificore.com/rest", "test_JhIRDMsbo4prEsXCc3AW")
	//Prices
	var smsPrices = smsClient.GetSmsPrices()
	log.Println(smsPrices)
	//Single смс
	var oneSms = smsClient.CreateSms("1", nil, "me", "123", &notificorerest.SmsPhone{Msisdn:"79774342488", Reference:"ext_" + RandStringRunes(9)})
	log.Println(oneSms)
	//Batch смс
	var twoSms = smsClient.CreateMultipleSms(notificorerest.MultipleSmsRequest{
		Body:      "123",
		Originator:"me",
		Validity:  "1",
		Phones:    []notificorerest.SmsPhone{{Msisdn:"79774342488", Reference:"ext_" + RandStringRunes(9)},
					      {Msisdn:"79774342489", Reference:"ext_" + RandStringRunes(9)}}})
	log.Println(twoSms)
	//Sms info by task id
	var taskInfo = smsClient.GetSmsInfoByTaskId(twoSms.TaskID)
	log.Println(taskInfo)
	//Sms info by id
	var smsInfo = smsClient.GetSmsInfoById(twoSms.Smses[0].ID)
	log.Println(smsInfo)
	//Sms info by external id
	var smsInfo2 = smsClient.GetSmsInfoByReference(oneSms.Sms.Reference)
	log.Println(smsInfo2)
	//Viber
	var viberClient = notificorerest.NewViberClient("https://app.notificore.com/rest", "test_JhIRDMsbo4prEsXCc3AW")
	//Prices
	var viberPrices = viberClient.GetViberPrices()
	log.Println(viberPrices)
	//Create viber message
	var viber = viberClient.CreateViber(
		notificorerest.ViberRequest{
			Validity:           "1",
			ViberMessages:      []notificorerest.ViberMessage{{AlphaName:"Notificore",
				IsPromotional:                                false,
				Text:                                         "123",
				Options:                                      notificorerest.ViberOptions{Viber: notificorerest.ViberOptionsContainer{Img: "123", Caption: "123", Action: "123"}},
				Recipients:                                   []notificorerest.Recipient{{Msisdn:"79774342488", Reference:"ext_" + RandStringRunes(9)}}}}})
	log.Println(viber)
	//Viber ingo by id
	var viberInfo = viberClient.GetViberInfoById(viber.Messages[0].ID)
	log.Println(viberInfo)
	//Viber ingo by external id
	var viberInfo2 = viberClient.GetViberInfoByReference(viber.Messages[0].Reference)
	log.Println(viberInfo2)

}
