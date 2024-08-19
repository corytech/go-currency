package currency

import (
	"slices"
	"strings"
	"testing"
)

var cryptoCurrencies = []string{
	"bch",
	"btc",
	"eth",
	"ltc",
	"trx",
	"usdt",
	"xlm",
	"xrp",
	"xmr",
	"bnb",
}

var allCurrencies = []string{
	"afn",
	"eur",
	"all",
	"dzd",
	"usd",
	"aoa",
	"xcd",
	"ars",
	"amd",
	"awg",
	"aud",
	"azn",
	"bsd",
	"bhd",
	"bdt",
	"bbd",
	"byn",
	"bzd",
	"xof",
	"bmd",
	"inr",
	"btn",
	"bob",
	"bov",
	"bam",
	"bwp",
	"nok",
	"brl",
	"bnd",
	"bgn",
	"bif",
	"cve",
	"khr",
	"xaf",
	"cad",
	"kyd",
	"clp",
	"clf",
	"cny",
	"cop",
	"cou",
	"kmf",
	"cdf",
	"nzd",
	"crc",
	"cup",
	"cuc",
	"ang",
	"czk",
	"dkk",
	"djf",
	"dop",
	"egp",
	"svc",
	"ern",
	"szl",
	"etb",
	"fkp",
	"fjd",
	"xpf",
	"gmd",
	"gel",
	"ghs",
	"gip",
	"gtq",
	"gbp",
	"gnf",
	"gyd",
	"htg",
	"hnl",
	"hkd",
	"huf",
	"isk",
	"idr",
	"xdr",
	"irr",
	"iqd",
	"ils",
	"jmd",
	"jpy",
	"jod",
	"kzt",
	"kes",
	"kpw",
	"krw",
	"kwd",
	"kgs",
	"lak",
	"lbp",
	"lsl",
	"zar",
	"lrd",
	"lyd",
	"chf",
	"mop",
	"mkd",
	"mga",
	"mwk",
	"myr",
	"mvr",
	"mru",
	"mur",
	"xua",
	"mxn",
	"mxv",
	"mdl",
	"mnt",
	"mad",
	"mzn",
	"mmk",
	"nad",
	"npr",
	"nio",
	"ngn",
	"omr",
	"pkr",
	"pab",
	"pgk",
	"pyg",
	"pen",
	"php",
	"pln",
	"qar",
	"ron",
	"rub",
	"rwf",
	"shp",
	"wst",
	"stn",
	"sar",
	"rsd",
	"scr",
	"sll",
	"sle",
	"sgd",
	"xsu",
	"sbd",
	"sos",
	"ssp",
	"lkr",
	"sdg",
	"srd",
	"sek",
	"che",
	"chw",
	"syp",
	"twd",
	"tjs",
	"tzs",
	"thb",
	"top",
	"ttd",
	"tnd",
	"try",
	"tmt",
	"ugx",
	"uah",
	"aed",
	"usn",
	"uyu",
	"uyi",
	"uyw",
	"uzs",
	"vuv",
	"ves",
	"ved",
	"vnd",
	"yer",
	"zmw",
	"zwl",
	"bch",
	"btc",
	"eth",
	"ltc",
	"trx",
	"usdt",
	"xlm",
	"xrp",
	"xmr",
	"bnb",
}

var iso4217Map = map[int]Currency{
	971: AFN,
}

func TestIsValidWillReturnFalseForNonCurrencyValues(t *testing.T) {
	if Currency("foo").IsValid() {
		t.Fatalf("Expected '%t' Got '%t' for non-currency value", false, true)
	}
}

func TestIsValidWillReturnTrueForCurrencies(t *testing.T) {
	for _, currency := range allCurrencies {
		if !Currency(currency).IsValid() {
			t.Fatalf("Expected '%t' Got '%t' for currency `%s`", true, false, currency)
		}
	}
}

func TestUpperCaseCodeWillReturnCorrectValue(t *testing.T) {
	for _, currency := range allCurrencies {
		if Currency(currency).UpperCaseCode() != strings.ToUpper(currency) {
			t.Fatalf("Expected '%s' Got '%s' for currency `%s`", strings.ToUpper(currency), Currency(currency).UpperCaseCode(), currency)
		}
	}
}

func TestIsCryptoWillReturnCorrectValue(t *testing.T) {
	for _, currency := range allCurrencies {
		if slices.Contains(cryptoCurrencies, currency) {
			if !Currency(currency).IsCrypto() {
				t.Fatalf("Expected '%t' Got '%t' for crypto currency `%s`", true, false, currency)
			}
		}
	}
}

func TestIsFiatWillReturnCorrectValue(t *testing.T) {
	for _, currency := range allCurrencies {
		if !slices.Contains(cryptoCurrencies, currency) {
			if !Currency(currency).IsFiat() {
				t.Fatalf("Expected '%t' Got '%t' for fiat currency `%s`", true, false, currency)
			}
		}
	}
}
