package currency

import (
	"slices"
	"strings"
)

type Currency string

const (
	AFN  Currency = "afn"
	EUR  Currency = "eur"
	ALL  Currency = "all"
	DZD  Currency = "dzd"
	USD  Currency = "usd"
	AOA  Currency = "aoa"
	XCD  Currency = "xcd"
	ARS  Currency = "ars"
	AMD  Currency = "amd"
	AWG  Currency = "awg"
	AUD  Currency = "aud"
	AZN  Currency = "azn"
	BSD  Currency = "bsd"
	BHD  Currency = "bhd"
	BDT  Currency = "bdt"
	BBD  Currency = "bbd"
	BYN  Currency = "byn"
	BZD  Currency = "bzd"
	XOF  Currency = "xof"
	BMD  Currency = "bmd"
	INR  Currency = "inr"
	BTN  Currency = "btn"
	BOB  Currency = "bob"
	BOV  Currency = "bov"
	BAM  Currency = "bam"
	BWP  Currency = "bwp"
	NOK  Currency = "nok"
	BRL  Currency = "brl"
	BND  Currency = "bnd"
	BGN  Currency = "bgn"
	BIF  Currency = "bif"
	CVE  Currency = "cve"
	KHR  Currency = "khr"
	XAF  Currency = "xaf"
	CAD  Currency = "cad"
	KYD  Currency = "kyd"
	CLP  Currency = "clp"
	CLF  Currency = "clf"
	CNY  Currency = "cny"
	COP  Currency = "cop"
	COU  Currency = "cou"
	KMF  Currency = "kmf"
	CDF  Currency = "cdf"
	NZD  Currency = "nzd"
	CRC  Currency = "crc"
	CUP  Currency = "cup"
	CUC  Currency = "cuc"
	ANG  Currency = "ang"
	CZK  Currency = "czk"
	DKK  Currency = "dkk"
	DJF  Currency = "djf"
	DOP  Currency = "dop"
	EGP  Currency = "egp"
	SVC  Currency = "svc"
	ERN  Currency = "ern"
	SZL  Currency = "szl"
	ETB  Currency = "etb"
	FKP  Currency = "fkp"
	FJD  Currency = "fjd"
	XPF  Currency = "xpf"
	GMD  Currency = "gmd"
	GEL  Currency = "gel"
	GHS  Currency = "ghs"
	GIP  Currency = "gip"
	GTQ  Currency = "gtq"
	GBP  Currency = "gbp"
	GNF  Currency = "gnf"
	GYD  Currency = "gyd"
	HTG  Currency = "htg"
	HNL  Currency = "hnl"
	HKD  Currency = "hkd"
	HUF  Currency = "huf"
	ISK  Currency = "isk"
	IDR  Currency = "idr"
	XDR  Currency = "xdr"
	IRR  Currency = "irr"
	IQD  Currency = "iqd"
	ILS  Currency = "ils"
	JMD  Currency = "jmd"
	JPY  Currency = "jpy"
	JOD  Currency = "jod"
	KZT  Currency = "kzt"
	KES  Currency = "kes"
	KPW  Currency = "kpw"
	KRW  Currency = "krw"
	KWD  Currency = "kwd"
	KGS  Currency = "kgs"
	LAK  Currency = "lak"
	LBP  Currency = "lbp"
	LSL  Currency = "lsl"
	ZAR  Currency = "zar"
	LRD  Currency = "lrd"
	LYD  Currency = "lyd"
	CHF  Currency = "chf"
	MOP  Currency = "mop"
	MKD  Currency = "mkd"
	MGA  Currency = "mga"
	MWK  Currency = "mwk"
	MYR  Currency = "myr"
	MVR  Currency = "mvr"
	MRU  Currency = "mru"
	MUR  Currency = "mur"
	XUA  Currency = "xua"
	MXN  Currency = "mxn"
	MXV  Currency = "mxv"
	MDL  Currency = "mdl"
	MNT  Currency = "mnt"
	MAD  Currency = "mad"
	MZN  Currency = "mzn"
	MMK  Currency = "mmk"
	NAD  Currency = "nad"
	NPR  Currency = "npr"
	NIO  Currency = "nio"
	NGN  Currency = "ngn"
	OMR  Currency = "omr"
	PKR  Currency = "pkr"
	PAB  Currency = "pab"
	PGK  Currency = "pgk"
	PYG  Currency = "pyg"
	PEN  Currency = "pen"
	PHP  Currency = "php"
	PLN  Currency = "pln"
	QAR  Currency = "qar"
	RON  Currency = "ron"
	RUB  Currency = "rub"
	RWF  Currency = "rwf"
	SHP  Currency = "shp"
	WST  Currency = "wst"
	STN  Currency = "stn"
	SAR  Currency = "sar"
	RSD  Currency = "rsd"
	SCR  Currency = "scr"
	SLL  Currency = "sll"
	SLE  Currency = "sle"
	SGD  Currency = "sgd"
	XSU  Currency = "xsu"
	SBD  Currency = "sbd"
	SOS  Currency = "sos"
	SSP  Currency = "ssp"
	LKR  Currency = "lkr"
	SDG  Currency = "sdg"
	SRD  Currency = "srd"
	SEK  Currency = "sek"
	CHE  Currency = "che"
	CHW  Currency = "chw"
	SYP  Currency = "syp"
	TWD  Currency = "twd"
	TJS  Currency = "tjs"
	TZS  Currency = "tzs"
	THB  Currency = "thb"
	TOP  Currency = "top"
	TTD  Currency = "ttd"
	TND  Currency = "tnd"
	TRY  Currency = "try"
	TMT  Currency = "tmt"
	UGX  Currency = "ugx"
	UAH  Currency = "uah"
	AED  Currency = "aed"
	USN  Currency = "usn"
	UYU  Currency = "uyu"
	UYI  Currency = "uyi"
	UYW  Currency = "uyw"
	UZS  Currency = "uzs"
	VUV  Currency = "vuv"
	VES  Currency = "ves"
	VED  Currency = "ved"
	VND  Currency = "vnd"
	YER  Currency = "yer"
	ZMW  Currency = "zmw"
	ZWL  Currency = "zwl"
	BCH  Currency = "bch"
	BTC  Currency = "btc"
	ETH  Currency = "eth"
	LTC  Currency = "ltc"
	TRX  Currency = "trx"
	USDT Currency = "usdt"
	XLM  Currency = "xlm"
	XRP  Currency = "xrp"
	XMR  Currency = "xmr"
	BNB  Currency = "bnb"
)

func (c Currency) UpperCaseCode() string {
	return strings.ToUpper(string(c))
}

func (c Currency) IsValid() bool {
	return slices.Contains(GetAllCurrencies(), c)
}

func (c Currency) IsCrypto() bool {
	switch c {
	case
		BCH,
		BTC,
		ETH,
		LTC,
		TRX,
		USDT,
		XLM,
		XRP,
		XMR,
		BNB:
		return true
	}

	return false
}

func (c Currency) IsFiat() bool {
	return !c.IsCrypto()
}

func (c Currency) GetPrecision() int {
	switch c {
	case XOF, BIF, XAF, CLP, KMF, DJF, XPF, GNF, ISK,
		JPY, KRW, PYG, RWF, UGX, UYI, VUV, VND:
		return 0
	case AFN, EUR, ALL, DZD, USD, AOA, XCD, ARS, AMD,
		AWG, AUD, AZN, BSD, BDT, BBD, BYN, BZD, BMD,
		INR, BTN, BOB, BOV, BAM, BWP, NOK, BRL, BND,
		BGN, CVE, KHR, CAD, KYD, CNY, COP, COU, CDF,
		NZD, CRC, CUP, CUC, ANG, CZK, DKK, DOP, EGP,
		SVC, ERN, SZL, ETB, FKP, FJD, GMD, GEL, GHS,
		GIP, GTQ, GBP, GYD, HTG, HNL, HKD, HUF, IDR,
		IRR, ILS, JMD, KZT, KES, KPW, KGS, LAK, LBP,
		LSL, ZAR, LRD, CHF, MOP, MKD, MGA, MWK, MYR,
		MVR, MRU, MUR, MXN, MXV, MDL, MNT, MAD, MZN,
		MMK, NAD, NPR, NIO, NGN, PKR, PAB, PGK, PEN,
		PHP, PLN, QAR, RON, RUB, SHP, WST, STN, SAR,
		RSD, SCR, SLL, SLE, SGD, SBD, SOS, SSP, LKR,
		SDG, SRD, SEK, CHE, CHW, SYP, TWD, TJS, TZS,
		THB, TOP, TTD, TRY, TMT, UAH, AED, USN, UYU,
		UZS, VES, VED, YER, ZMW, ZWL:
		return 2
	case BHD, IQD, JOD, KWD, LYD, OMR, TND:
		return 3
	case CLF, UYW:
		return 4
	case XDR, XUA, XSU:
		return 17
	case BCH, BTC, LTC, USDT:
		return 8
	case ETH, TRX, BNB:
		return 17
	case XLM:
		return 7
	case XRP:
		return 15
	case XMR:
		return 12
	}
	return -1
}

func FromIso4217(i int) Currency {
	switch i {
	case 971:
		return AFN
	case 978:
		return EUR
	case 8:
		return ALL
	case 12:
		return DZD
	case 840:
		return USD
	case 973:
		return AOA
	case 951:
		return XCD
	case 32:
		return ARS
	case 51:
		return AMD
	case 533:
		return AWG
	case 36:
		return AUD
	case 944:
		return AZN
	case 44:
		return BSD
	case 48:
		return BHD
	case 50:
		return BDT
	case 52:
		return BBD
	case 933:
		return BYN
	case 84:
		return BZD
	case 952:
		return XOF
	case 60:
		return BMD
	case 356:
		return INR
	case 64:
		return BTN
	case 68:
		return BOB
	case 984:
		return BOV
	case 977:
		return BAM
	case 72:
		return BWP
	case 578:
		return NOK
	case 986:
		return BRL
	case 96:
		return BND
	case 975:
		return BGN
	case 108:
		return BIF
	case 132:
		return CVE
	case 116:
		return KHR
	case 950:
		return XAF
	case 124:
		return CAD
	case 136:
		return KYD
	case 152:
		return CLP
	case 990:
		return CLF
	case 156:
		return CNY
	case 170:
		return COP
	case 970:
		return COU
	case 174:
		return KMF
	case 976:
		return CDF
	case 554:
		return NZD
	case 188:
		return CRC
	case 192:
		return CUP
	case 931:
		return CUC
	case 532:
		return ANG
	case 203:
		return CZK
	case 208:
		return DKK
	case 262:
		return DJF
	case 214:
		return DOP
	case 818:
		return EGP
	case 222:
		return SVC
	case 232:
		return ERN
	case 748:
		return SZL
	case 230:
		return ETB
	case 238:
		return FKP
	case 242:
		return FJD
	case 953:
		return XPF
	case 270:
		return GMD
	case 981:
		return GEL
	case 936:
		return GHS
	case 292:
		return GIP
	case 320:
		return GTQ
	case 826:
		return GBP
	case 324:
		return GNF
	case 328:
		return GYD
	case 332:
		return HTG
	case 340:
		return HNL
	case 344:
		return HKD
	case 348:
		return HUF
	case 352:
		return ISK
	case 360:
		return IDR
	case 960:
		return XDR
	case 364:
		return IRR
	case 368:
		return IQD
	case 376:
		return ILS
	case 388:
		return JMD
	case 392:
		return JPY
	case 400:
		return JOD
	case 398:
		return KZT
	case 404:
		return KES
	case 408:
		return KPW
	case 410:
		return KRW
	case 414:
		return KWD
	case 417:
		return KGS
	case 418:
		return LAK
	case 422:
		return LBP
	case 426:
		return LSL
	case 710:
		return ZAR
	case 430:
		return LRD
	case 434:
		return LYD
	case 756:
		return CHF
	case 446:
		return MOP
	case 807:
		return MKD
	case 969:
		return MGA
	case 454:
		return MWK
	case 458:
		return MYR
	case 462:
		return MVR
	case 929:
		return MRU
	case 480:
		return MUR
	case 965:
		return XUA
	case 484:
		return MXN
	case 979:
		return MXV
	case 498:
		return MDL
	case 496:
		return MNT
	case 504:
		return MAD
	case 943:
		return MZN
	case 104:
		return MMK
	case 516:
		return NAD
	case 524:
		return NPR
	case 558:
		return NIO
	case 566:
		return NGN
	case 512:
		return OMR
	case 586:
		return PKR
	case 590:
		return PAB
	case 598:
		return PGK
	case 600:
		return PYG
	case 604:
		return PEN
	case 608:
		return PHP
	case 985:
		return PLN
	case 634:
		return QAR
	case 946:
		return RON
	case 643:
		return RUB
	case 646:
		return RWF
	case 654:
		return SHP
	case 882:
		return WST
	case 930:
		return STN
	case 682:
		return SAR
	case 941:
		return RSD
	case 690:
		return SCR
	case 694:
		return SLL
	case 925:
		return SLE
	case 702:
		return SGD
	case 994:
		return XSU
	case 90:
		return SBD
	case 706:
		return SOS
	case 728:
		return SSP
	case 144:
		return LKR
	case 938:
		return SDG
	case 968:
		return SRD
	case 752:
		return SEK
	case 947:
		return CHE
	case 948:
		return CHW
	case 760:
		return SYP
	case 901:
		return TWD
	case 972:
		return TJS
	case 834:
		return TZS
	case 764:
		return THB
	case 776:
		return TOP
	case 780:
		return TTD
	case 788:
		return TND
	case 949:
		return TRY
	case 934:
		return TMT
	case 800:
		return UGX
	case 980:
		return UAH
	case 784:
		return AED
	case 997:
		return USN
	case 858:
		return UYU
	case 940:
		return UYI
	case 927:
		return UYW
	case 860:
		return UZS
	case 548:
		return VUV
	case 928:
		return VES
	case 926:
		return VED
	case 704:
		return VND
	case 886:
		return YER
	case 967:
		return ZMW
	case 932:
		return ZWL
	}

	return ""
}

func GetFiatCurrencies() []Currency {
	var result []Currency

	for _, currency := range GetAllCurrencies() {
		if currency.IsFiat() {
			result = append(result, currency)
		}
	}

	return result
}

func GetCryptoCurrencies() []Currency {
	var result []Currency

	for _, currency := range GetAllCurrencies() {
		if currency.IsCrypto() {
			result = append(result, currency)
		}
	}

	return result
}

func GetAllCurrencies() []Currency {
	return []Currency{
		AFN,
		EUR,
		ALL,
		DZD,
		USD,
		AOA,
		XCD,
		ARS,
		AMD,
		AWG,
		AUD,
		AZN,
		BSD,
		BHD,
		BDT,
		BBD,
		BYN,
		BZD,
		XOF,
		BMD,
		INR,
		BTN,
		BOB,
		BOV,
		BAM,
		BWP,
		NOK,
		BRL,
		BND,
		BGN,
		BIF,
		CVE,
		KHR,
		XAF,
		CAD,
		KYD,
		CLP,
		CLF,
		CNY,
		COP,
		COU,
		KMF,
		CDF,
		NZD,
		CRC,
		CUP,
		CUC,
		ANG,
		CZK,
		DKK,
		DJF,
		DOP,
		EGP,
		SVC,
		ERN,
		SZL,
		ETB,
		FKP,
		FJD,
		XPF,
		GMD,
		GEL,
		GHS,
		GIP,
		GTQ,
		GBP,
		GNF,
		GYD,
		HTG,
		HNL,
		HKD,
		HUF,
		ISK,
		IDR,
		XDR,
		IRR,
		IQD,
		ILS,
		JMD,
		JPY,
		JOD,
		KZT,
		KES,
		KPW,
		KRW,
		KWD,
		KGS,
		LAK,
		LBP,
		LSL,
		ZAR,
		LRD,
		LYD,
		CHF,
		MOP,
		MKD,
		MGA,
		MWK,
		MYR,
		MVR,
		MRU,
		MUR,
		XUA,
		MXN,
		MXV,
		MDL,
		MNT,
		MAD,
		MZN,
		MMK,
		NAD,
		NPR,
		NIO,
		NGN,
		OMR,
		PKR,
		PAB,
		PGK,
		PYG,
		PEN,
		PHP,
		PLN,
		QAR,
		RON,
		RUB,
		RWF,
		SHP,
		WST,
		STN,
		SAR,
		RSD,
		SCR,
		SLL,
		SLE,
		SGD,
		XSU,
		SBD,
		SOS,
		SSP,
		LKR,
		SDG,
		SRD,
		SEK,
		CHE,
		CHW,
		SYP,
		TWD,
		TJS,
		TZS,
		THB,
		TOP,
		TTD,
		TND,
		TRY,
		TMT,
		UGX,
		UAH,
		AED,
		USN,
		UYU,
		UYI,
		UYW,
		UZS,
		VUV,
		VES,
		VED,
		VND,
		YER,
		ZMW,
		ZWL,
		BCH,
		BTC,
		ETH,
		LTC,
		TRX,
		USDT,
		XLM,
		XRP,
		XMR,
		BNB,
	}
}
