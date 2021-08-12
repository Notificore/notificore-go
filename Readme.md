# com.notificore.rest

Библиотека позволяет работать с REST сервисом Notificore. Поддерживается автоматическая сериализация/десериализация

## Requirements

Go runtime

## Installation

go get github.com/Notificore/go/notificorerest
		
## Dependencies

* net/http
* log
* encoding/json
* strconv
* bytes

## Example

```go
//HLR
var hlrClient = notificorerest.NewHlrClient("https://app.notificore.com/rest", "YOUR_API_KEY")
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
	[]notificorerest.HlrRequest{notificorerest.HlrRequest{Tariff:nil, Reference:"ext_" + RandStringRunes(9), Msisdn:"380972920001", CallbackUrl:"http://someurl.com/callback"},
				 {Tariff:&tariff, Reference:"ext_" + RandStringRunes(9), Msisdn:"380972920002", CallbackUrl:"http://someurl.com/callback"}})
log.Println(hlr2)
//Hlr info by id
var hlrInfo = hlrClient.GetHlrInfoById(hlr.Hlrs[0].ID)
log.Println(hlrInfo)
//Hlr info by external id
var hlrInfo2 = hlrClient.GetHlrInfoByReference(hlr2.Hlrs[1].Reference)
log.Println(hlrInfo2)
//SMS
var smsClient = notificorerest.NewSmsClient("https://app.notificore.com/rest", "YOUR_API_KEY")
//Prices
var smsPrices = smsClient.GetSmsPrices()
log.Println(smsPrices)
//Single смс
var oneSms = smsClient.CreateSms("1", nil, "me", "123", &notificorerest.SmsPhone{Msisdn:"79999999999", Reference:"ext_" + RandStringRunes(9)})
log.Println(oneSms)
//Batch смс
var twoSms = smsClient.CreateMultipleSms(notificorerest.MultipleSmsRequest{
	Body:      "123",
	Originator:"me",
	Validity:  "1",
	Phones:    []notificorerest.SmsPhone{{Msisdn:"79999999999", Reference:"ext_" + RandStringRunes(9)},
					  {Msisdn:"79999999998", Reference:"ext_" + RandStringRunes(9)}}})
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
var viberClient = notificorerest.NewViberClient("https://app.notificore.com/rest", "YOUR_API_KEY")
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
			Recipients:                                   []notificorerest.Recipient{{Msisdn:"79999999999", Reference:"ext_" + RandStringRunes(9)}}}}})
log.Println(viber)
//Viber ingo by id
var viberInfo = viberClient.GetViberInfoById(viber.Messages[0].ID)
log.Println(viberInfo)
//Viber ingo by external id
var viberInfo2 = viberClient.GetViberInfoByReference(viber.Messages[0].Reference)
log.Println(viberInfo2)
```

## Support

Если у вас возникли сложности или вопросы по использованию пакета, создайте обсуждение в данном репозитории или напишите на электронную почту.

## Documentation

Документацию API можно получить из исходных кодов пакета
