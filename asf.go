package asf

import (
	"fmt"
	"strings"
)

var countryCode = map[string]int{
	"DZ": 143563,
	"AO": 143564,
	"AI": 143538,
	"AG": 143540,
	"AR": 143505,
	"AM": 143524,
	"AU": 143460,
	"AT": 143445,
	"AZ": 143568,
	"BH": 143559,
	"BD": 143490,
	"BB": 143541,
	"BY": 143565,
	"BE": 143446,
	"BZ": 143555,
	"BM": 143542,
	"BO": 143556,
	"BW": 143525,
	"BR": 143503,
	"VG": 143543,
	"BN": 143560,
	"BG": 143526,
	"CA": 143455,
	"KY": 143544,
	"CL": 143483,
	"CN": 143465,
	"CO": 143501,
	"CR": 143495,
	"CI": 143527,
	"HR": 143494,
	"CY": 143557,
	"CZ": 143489,
	"DK": 143458,
	"DM": 143545,
	"DO": 143508,
	"EC": 143509,
	"EG": 143516,
	"SV": 143506,
	"EE": 143518,
	"FI": 143447,
	"FR": 143442,
	"DE": 143443,
	"GH": 143573,
	"GR": 143448,
	"GD": 143546,
	"GT": 143504,
	"GY": 143553,
	"HN": 143510,
	"HK": 143463,
	"HU": 143482,
	"IS": 143558,
	"IN": 143467,
	"ID": 143476,
	"IE": 143449,
	"IL": 143491,
	"IT": 143450,
	"JM": 143511,
	"JP": 143462,
	"JO": 143528,
	"KZ": 143517,
	"KE": 143529,
	"KR": 143466,
	"KW": 143493,
	"LV": 143519,
	"LB": 143497,
	"LI": 143522,
	"LT": 143520,
	"LU": 143451,
	"MO": 143515,
	"MK": 143530,
	"MG": 143531,
	"MY": 143473,
	"MV": 143488,
	"ML": 143532,
	"MT": 143521,
	"MU": 143533,
	"MX": 143468,
	"MD": 143523,
	"MS": 143547,
	"NP": 143484,
	"NL": 143452,
	"NZ": 143461,
	"NI": 143512,
	"NE": 143534,
	"NG": 143561,
	"NO": 143457,
	"OM": 143562,
	"PK": 143477,
	"PA": 143485,
	"PY": 143513,
	"PE": 143507,
	"PH": 143474,
	"PL": 143478,
	"PT": 143453,
	"QA": 143498,
	"RO": 143487,
	"RU": 143469,
	"SA": 143479,
	"SN": 143535,
	"RS": 143500,
	"SG": 143464,
	"SK": 143496,
	"SI": 143499,
	"ZA": 143472,
	"ES": 143454,
	"LK": 143486,
	"KN": 143548,
	"LC": 143549,
	"VC": 143550,
	"SR": 143554,
	"SE": 143456,
	"CH": 143459,
	"TW": 143470,
	"TZ": 143572,
	"TH": 143475,
	"BS": 143539,
	"TT": 143551,
	"TN": 143536,
	"TR": 143480,
	"TC": 143552,
	"UG": 143537,
	"GB": 143444,
	"UA": 143492,
	"AE": 143481,
	"UY": 143514,
	"US": 143441,
	"UZ": 143566,
	"VE": 143502,
	"VN": 143471,
	"YE": 143571,
	"KG": 143586,
	"TJ": 143603,
}

var countryLanguageCode = map[string]int{
	"BR-pt": 15,
	"IT-it": 7,
	"IN-hi": 50,
	"ES-ca": 42,
	"ES-es": 8,
	"TW-zh": 18,
	"TH-th": 35,
	"DK-da": 11,
	"FI-fi": 12,
	"GB-en": 2,
	"NL-nl": 10,
	"PL-pl": 20,
	"RO-ro": 39,
	"CN-zh": 19,
	"NO-no": 14,
	"PT-pt": 24,
	"CH-de": 57,
	"UA-uk": 29,
	"US-en": 1,
	"VI-vi": 43,
	"SE-sv": 17,
	"CA-en": 6,
	"FR-fr": 3,
	"SK-sk": 40,
	"HK-zh": 45,
	"HR-hr": 41,
	"AU-en": 27,
	"KR-ko": 13,
	"TR-tr": 25,
	"GR-el": 23,
	"MY-ms": 38,
	"ID-id": 37,
	"CZ-cs": 22,
	"HU-hu": 21,
	"CA-fr": 5,
	"MX-es": 28,
	"DE-de": 4,
	"RU-ru": 16,
	"JP-ja": 9,
}

// CountryCode - return store front code
// for addtitional see: https://gist.github.com/hmml/8942940
func CountryCode(country string) (string, error) {
	res := ""

	if val, ok := countryCode[country]; ok {
		res = fmt.Sprintf("%d", val)
	}

	if res == "" {
		err := fmt.Errorf("Invalid country %s", country)

		return res, err
	}

	return res, nil
}

func CountryLanguageCode(country string, language string) string {
	key := fmt.Sprintf("%s-%s", country, language)

	res := ""

	if val, ok := countryLanguageCode[key]; ok {
		res = fmt.Sprintf("%d", val)
	}

	return res
}

func XAppleStoreFront(country string, language *string, platformCode int) (string, error) {
	country = strings.ToUpper(country)

	cc, err := CountryCode(country)

	if err != nil {
		return "", err
	}

	clc := ""

	if language != nil {
		lang := strings.ToLower(*language)

		clc = CountryLanguageCode(country, lang)
	}

	res := cc

	if clc != "" {
		res = fmt.Sprintf("%s-%s", cc, clc)
	}

	res = fmt.Sprintf("%s,%d", res, platformCode)

	return res, nil
}
