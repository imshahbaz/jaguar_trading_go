package constants

import (
	"encoding/json"
	"fmt"
)

type Margin struct {
	Name    string  `json:"name"`
	Symbol  string  `json:"symbol"`
	Percent float64 `json:"percent"`
}

var data map[string]Margin

func Init() {
	err := json.Unmarshal([]byte(mtf), &data)
	if err != nil {
		fmt.Println("Error in parsing the JSON file: ", err)
		return
	}
}

func GetMarginData() map[string]Margin {
	return data
}

var mtf = `{
    "ALKYLAMINE": {
      "name": "ALKYLAMINESCHEM.LTD",
      "symbol": "ALKYLAMINE",
      "percent": 60
    },
    "UNOMINDA": {
      "name": "UNOMINDALIMITED",
      "symbol": "UNOMINDA",
      "percent": 60
    },
    "RELAXO": {
      "name": "RELAXOFOOTLTD.",
      "symbol": "RELAXO",
      "percent": 65
    },
    "IRCTC": {
      "name": "INDIANRAILTOURCORPLTD",
      "symbol": "IRCTC",
      "percent": 65
    },
    "EICHERMOT": {
      "name": "EICHERMOTORSLTD",
      "symbol": "EICHERMOT",
      "percent": 75
    },
    "TANLA": {
      "name": "TANLAPLATFORMSLIMITED",
      "symbol": "TANLA",
      "percent": 60
    },
    "RBA": {
      "name": "RESTAURANTBRANDASIALTD",
      "symbol": "RBA",
      "percent": 60
    },
    "INDUSINDBK": {
      "name": "INDUSINDBANKLIMITED",
      "symbol": "INDUSINDBK",
      "percent": 70
    },
    "PFIZER": {
      "name": "PFIZERLTD",
      "symbol": "PFIZER",
      "percent": 60
    },
    "ABBOTINDIA": {
      "name": "ABBOTTINDIALIMITED",
      "symbol": "ABBOTINDIA",
      "percent": 70
    },
    "SUNDARMFIN": {
      "name": "SUNDARAMFINANCELTD",
      "symbol": "SUNDARMFIN",
      "percent": 60
    },
    "ASHOKLEY": {
      "name": "ASHOKLEYLANDLTD",
      "symbol": "ASHOKLEY",
      "percent": 70
    },
    "GRAPHITE": {
      "name": "GRAPHITEINDIALTD",
      "symbol": "GRAPHITE",
      "percent": 60
    },
    "JBCHEPHARM": {
      "name": "JBCHEMICALSANDPHARMA",
      "symbol": "JBCHEPHARM",
      "percent": 60
    },
    "POLICYBZR": {
      "name": "PBFINTECHLIMITED",
      "symbol": "POLICYBZR",
      "percent": 60
    },
    "EXIDEIND": {
      "name": "EXIDEINDUSTRIESLTD",
      "symbol": "EXIDEIND",
      "percent": 70
    },
    "TITAN": {
      "name": "TITANCOMPANYLIMITED",
      "symbol": "TITAN",
      "percent": 75
    },
    "KANSAINER": {
      "name": "KANSAINEROLACPAINTSLTD",
      "symbol": "KANSAINER",
      "percent": 65
    },
    "BAJAJ-AUTO": {
      "name": "BAJAJAUTOLIMITED",
      "symbol": "BAJAJ-AUTO",
      "percent": 75
    },
    "RCF": {
      "name": "RASHTRIYACHEMICALS&FER",
      "symbol": "RCF",
      "percent": 60
    },
    "PAGEIND": {
      "name": "PAGEINDUSTRIESLTD",
      "symbol": "PAGEIND",
      "percent": 75
    },
    "WABAG": {
      "name": "VATECHWABAGLTD",
      "symbol": "WABAG",
      "percent": 60
    },
    "CANBK": {
      "name": "CANARABANK",
      "symbol": "CANBK",
      "percent": 65
    },
    "L&TFH": {
      "name": "L&TFINANCEHOLDINGSLTD",
      "symbol": "L&TFH",
      "percent": 65
    },
    "BANKINDIA": {
      "name": "BANKOFINDIA",
      "symbol": "BANKINDIA",
      "percent": 60
    },
    "ABB": {
      "name": "ABBINDIALIMITED",
      "symbol": "ABB",
      "percent": 75
    },
    "AVANTIFEED": {
      "name": "AVANTIFEEDSLIMITED",
      "symbol": "AVANTIFEED",
      "percent": 60
    },
    "VIPIND": {
      "name": "VIPINDUSTRIESLTD",
      "symbol": "VIPIND",
      "percent": 60
    },
    "TATAELXSI": {
      "name": "TATAELXSILIMITED",
      "symbol": "TATAELXSI",
      "percent": 60
    },
    "BDL": {
      "name": "BHARATDYNAMICSLIMITED",
      "symbol": "BDL",
      "percent": 60
    },
    "CIPLA": {
      "name": "CIPLALTD",
      "symbol": "CIPLA",
      "percent": 75
    },
    "ASHOKA": {
      "name": "ASHOKABUILDCONLTD",
      "symbol": "ASHOKA",
      "percent": 60
    },
    "OBEROIRLTY": {
      "name": "OBEROIREALTYLIMITED",
      "symbol": "OBEROIRLTY",
      "percent": 70
    },
    "LAOPALA": {
      "name": "LAOPALARGLIMITED",
      "symbol": "LAOPALA",
      "percent": 60
    },
    "BEL": {
      "name": "BHARATELECTRONICSLTD",
      "symbol": "BEL",
      "percent": 70
    },
    "SOBHA": {
      "name": "SOBHALIMITED",
      "symbol": "SOBHA",
      "percent": 60
    },
    "LEMONTREE": {
      "name": "LEMONTREEHOTELSLTD",
      "symbol": "LEMONTREE",
      "percent": 60
    },
    "NYKAA": {
      "name": "FSNECOMMERCEVENTURES",
      "symbol": "NYKAA",
      "percent": 60
    },
    "AEGISCHEM": {
      "name": "AEGISLOGISTICSLIMITED",
      "symbol": "AEGISCHEM",
      "percent": 60
    },
    "LALPATHLAB": {
      "name": "DR.LALPATHLABSLTD.",
      "symbol": "LALPATHLAB",
      "percent": 70
    },
    "ITC": {
      "name": "ITCLTD",
      "symbol": "ITC",
      "percent": 75
    },
    "VOLTAMP": {
      "name": "VOLTAMPTRANSFORMERSLTD",
      "symbol": "VOLTAMP",
      "percent": 60
    },
    "NLCINDIA": {
      "name": "NLCINDIALIMITED",
      "symbol": "NLCINDIA",
      "percent": 60
    },
    "GPPL": {
      "name": "GUJARATPIPAVAVPORTLTD",
      "symbol": "GPPL",
      "percent": 60
    },
    "CREDITACC": {
      "name": "CREDITACCESSGRAMEENLTD",
      "symbol": "CREDITACC",
      "percent": 60
    },
    "LTIM": {
      "name": "LTIMINDTREELIMITED",
      "symbol": "LTIM",
      "percent": 70
    },
    "HINDCOPPER": {
      "name": "HINDUSTANCOPPERLTD",
      "symbol": "HINDCOPPER",
      "percent": 60
    },
    "MARUTI": {
      "name": "MARUTISUZUKIINDIALTD.",
      "symbol": "MARUTI",
      "percent": 75
    },
    "BALRAMCHIN": {
      "name": "BALRAMPURCHINIMILLSLTD",
      "symbol": "BALRAMCHIN",
      "percent": 65
    },
    "BAJAJHLDNG": {
      "name": "BAJAJHOLDINGS&INVSLTD",
      "symbol": "BAJAJHLDNG",
      "percent": 60
    },
    "INFY": {
      "name": "INFOSYSLIMITED",
      "symbol": "INFY",
      "percent": 75
    },
    "PSUBNKBEES": {
      "name": "NIPINDETFPSUBANKBEES",
      "symbol": "PSUBNKBEES",
      "percent": 70
    },
    "BSOFT": {
      "name": "BIRLASOFTLIMITED",
      "symbol": "BSOFT",
      "percent": 70
    },
    "CIEINDIA": {
      "name": "CIEAUTOMOTIVEINDIALTD",
      "symbol": "CIEINDIA",
      "percent": 60
    },
    "TORNTPOWER": {
      "name": "TORRENTPOWERLTD",
      "symbol": "TORNTPOWER",
      "percent": 65
    },
    "RATNAMANI": {
      "name": "RATNAMANIMET&TUBLTD.",
      "symbol": "RATNAMANI",
      "percent": 60
    },
    "INDHOTEL": {
      "name": "THEINDIANHOTELSCO.LTD",
      "symbol": "INDHOTEL",
      "percent": 70
    },
    "TECHM": {
      "name": "TECHMAHINDRALIMITED",
      "symbol": "TECHM",
      "percent": 75
    },
    "MPHASIS": {
      "name": "MPHASISLIMITED",
      "symbol": "MPHASIS",
      "percent": 70
    },
    "DLF": {
      "name": "DLFLIMITED",
      "symbol": "DLF",
      "percent": 70
    },
    "TATACOMM": {
      "name": "TATACOMMUNICATIONSLTD",
      "symbol": "TATACOMM",
      "percent": 70
    },
    "FORTIS": {
      "name": "FORTISHEALTHCARELTD",
      "symbol": "FORTIS",
      "percent": 60
    },
    "JUBLFOOD": {
      "name": "JUBILANTFOODWORKSLTD",
      "symbol": "JUBLFOOD",
      "percent": 70
    },
    "RAMCOCEM": {
      "name": "THERAMCOCEMENTSLIMITED",
      "symbol": "RAMCOCEM",
      "percent": 75
    },
    "MUTHOOTFIN": {
      "name": "MUTHOOTFINANCELIMITED",
      "symbol": "MUTHOOTFIN",
      "percent": 70
    },
    "JSWSTEEL": {
      "name": "JSWSTEELLIMITED",
      "symbol": "JSWSTEEL",
      "percent": 70
    },
    "M&MFIN": {
      "name": "M&MFIN.SERVICESLTD",
      "symbol": "M&MFIN",
      "percent": 70
    },
    "TATAINVEST": {
      "name": "TATAINVESTMENTCORPLTD",
      "symbol": "TATAINVEST",
      "percent": 65
    },
    "APTUS": {
      "name": "APTUSVALUEHSGFINILTD",
      "symbol": "APTUS",
      "percent": 60
    },
    "HAVELLS": {
      "name": "HAVELLSINDIALIMITED",
      "symbol": "HAVELLS",
      "percent": 75
    },
    "PETRONET": {
      "name": "PETRONETLNGLIMITED",
      "symbol": "PETRONET",
      "percent": 75
    },
    "LICHSGFIN": {
      "name": "LICHOUSINGFINANCELTD",
      "symbol": "LICHSGFIN",
      "percent": 65
    },
    "POWERINDIA": {
      "name": "HITACHIENERGYINDIALTD",
      "symbol": "POWERINDIA",
      "percent": 60
    },
    "CPSEETF": {
      "name": "CPSEETF",
      "symbol": "CPSEETF",
      "percent": 70
    },
    "PARAS": {
      "name": "PARASDEFANDSPCETECHL",
      "symbol": "PARAS",
      "percent": 60
    },
    "FLUOROCHEM": {
      "name": "GUJARATFLUOROCHEMLTD",
      "symbol": "FLUOROCHEM",
      "percent": 60
    },
    "PEL": {
      "name": "PIRAMALENTERPRISESLTD",
      "symbol": "PEL",
      "percent": 65
    },
    "HDFCAMC": {
      "name": "HDFCAMCLIMITED",
      "symbol": "HDFCAMC",
      "percent": 75
    },
    "SUVENPHAR": {
      "name": "SUVENPHARMACEUTICALSLTD",
      "symbol": "SUVENPHAR",
      "percent": 60
    },
    "GESHIP": {
      "name": "THEGESHPG.LTD",
      "symbol": "GESHIP",
      "percent": 60
    },
    "SOLARINDS": {
      "name": "SOLARINDUSTRIES(I)LTD",
      "symbol": "SOLARINDS",
      "percent": 60
    },
    "UJJIVAN": {
      "name": "UJJIVANFIN.SERVC.LTD.",
      "symbol": "UJJIVAN",
      "percent": 60
    },
    "GODREJAGRO": {
      "name": "GODREJAGROVETLIMITED",
      "symbol": "GODREJAGRO",
      "percent": 60
    },
    "PNBHOUSING": {
      "name": "PNBHOUSINGFINLTD.",
      "symbol": "PNBHOUSING",
      "percent": 60
    },
    "LAURUSLABS": {
      "name": "LAURUSLABSLIMITED",
      "symbol": "LAURUSLABS",
      "percent": 70
    },
    "NHPC": {
      "name": "NHPCLTD",
      "symbol": "NHPC",
      "percent": 60
    },
    "BAYERCROP": {
      "name": "BAYERCROPSCIENCELTD",
      "symbol": "BAYERCROP",
      "percent": 60
    },
    "NATCOPHARM": {
      "name": "NATCOPHARMALTD.",
      "symbol": "NATCOPHARM",
      "percent": 60
    },
    "PFC": {
      "name": "POWERFINCORPLTD.",
      "symbol": "PFC",
      "percent": 75
    },
    "CMSINFO": {
      "name": "CMSINFOSYSTEMSLIMITED",
      "symbol": "CMSINFO",
      "percent": 70
    },
    "TATAMOTORS": {
      "name": "TATAMOTORSLIMITED",
      "symbol": "TATAMOTORS",
      "percent": 70
    },
    "UPL": {
      "name": "UPLLIMITED",
      "symbol": "UPL",
      "percent": 75
    },
    "ZFCVINDIA": {
      "name": "ZFCOMVECTRSYSINDLTD",
      "symbol": "ZFCVINDIA",
      "percent": 60
    },
    "SUNTV": {
      "name": "SUNTVNETWORKLIMITED",
      "symbol": "SUNTV",
      "percent": 65
    },
    "FSL": {
      "name": "FIRSTSOURCESOLU.LTD.",
      "symbol": "FSL",
      "percent": 65
    },
    "MHRIL": {
      "name": "MAHINDRAHOLIDAYSLTD",
      "symbol": "MHRIL",
      "percent": 60
    },
    "BPCL": {
      "name": "BHARATPETROLEUMCORPLT",
      "symbol": "BPCL",
      "percent": 75
    },
    "POWERGRID": {
      "name": "POWERGRIDCORP.LTD.",
      "symbol": "POWERGRID",
      "percent": 75
    },
    "GHCL": {
      "name": "GHCLLIMITED",
      "symbol": "GHCL",
      "percent": 60
    },
    "NCC": {
      "name": "NCCLIMITED",
      "symbol": "NCC",
      "percent": 60
    },
    "NBCC": {
      "name": "NBCC(INDIA)LIMITED",
      "symbol": "NBCC",
      "percent": 60
    },
    "MAPMYINDIA": {
      "name": "C.E.INFOSYSTEMSLIMITED",
      "symbol": "MAPMYINDIA",
      "percent": 65
    },
    "ASTEC": {
      "name": "ASTECLIFESCIENCESLTD",
      "symbol": "ASTEC",
      "percent": 60
    },
    "ZEEL": {
      "name": "ZEEENTERTAINMENTENTLTD",
      "symbol": "ZEEL",
      "percent": 60
    },
    "FINEORG": {
      "name": "FINEORGANICIND.LTD.",
      "symbol": "FINEORG",
      "percent": 60
    },
    "METROPOLIS": {
      "name": "METROPOLISHEALTHCARELTD",
      "symbol": "METROPOLIS",
      "percent": 65
    },
    "HINDALCO": {
      "name": "HINDALCOINDUSTRIESLTD",
      "symbol": "HINDALCO",
      "percent": 70
    },
    "BEML": {
      "name": "BEMLLIMITED",
      "symbol": "BEML",
      "percent": 60
    },
    "VGUARD": {
      "name": "V-GUARDINDLTD.",
      "symbol": "VGUARD",
      "percent": 60
    },
    "CASTROLIND": {
      "name": "CASTROLINDIALIMITED",
      "symbol": "CASTROLIND",
      "percent": 60
    },
    "GODREJIND": {
      "name": "GODREJINDUSTRIESLTD",
      "symbol": "GODREJIND",
      "percent": 60
    },
    "ICICIBANK": {
      "name": "ICICIBANKLTD.",
      "symbol": "ICICIBANK",
      "percent": 75
    },
    "SAPPHIRE": {
      "name": "SAPPHIREFOODSINDIALTD",
      "symbol": "SAPPHIRE",
      "percent": 60
    },
    "DRREDDY": {
      "name": "DR.REDDYSLABORATORIES",
      "symbol": "DRREDDY",
      "percent": 75
    },
    "NATIONALUM": {
      "name": "NATIONALALUMINIUMCOLTD",
      "symbol": "NATIONALUM",
      "percent": 65
    },
    "GREENPANEL": {
      "name": "GREENPANELINDUSTRIESLTD",
      "symbol": "GREENPANEL",
      "percent": 60
    },
    "JMFINANCIL": {
      "name": "JMFINANCIALLIMITED",
      "symbol": "JMFINANCIL",
      "percent": 60
    },
    "TRENT": {
      "name": "TRENTLTD",
      "symbol": "TRENT",
      "percent": 70
    },
    "GABRIEL": {
      "name": "GABRIELINDIALTD",
      "symbol": "GABRIEL",
      "percent": 60
    },
    "TTKPRESTIG": {
      "name": "TTKPRESTIGELTD",
      "symbol": "TTKPRESTIG",
      "percent": 60
    },
    "VTL": {
      "name": "VARDHMANTEXTILESLIMITED",
      "symbol": "VTL",
      "percent": 60
    },
    "DEEPAKNTR": {
      "name": "DEEPAKNITRITELTD",
      "symbol": "DEEPAKNTR",
      "percent": 65
    },
    "TATAMTRDVR": {
      "name": "TATAMOTORSDVRAORD",
      "symbol": "TATAMTRDVR",
      "percent": 60
    },
    "COFORGE": {
      "name": "COFORGELIMITED",
      "symbol": "COFORGE",
      "percent": 70
    },
    "COALINDIA": {
      "name": "COALINDIALTD",
      "symbol": "COALINDIA",
      "percent": 70
    },
    "PCBL": {
      "name": "PCBLLIMITED",
      "symbol": "PCBL",
      "percent": 60
    },
    "CENTURYTEX": {
      "name": "CENTURYTEXTILESLTD",
      "symbol": "CENTURYTEX",
      "percent": 60
    },
    "BARBEQUE": {
      "name": "BARBEQUENATIONHOSP.LTD",
      "symbol": "BARBEQUE",
      "percent": 60
    },
    "JAMNAAUTO": {
      "name": "JAMNAAUTOINDLTD",
      "symbol": "JAMNAAUTO",
      "percent": 60
    },
    "AAVAS": {
      "name": "AAVASFINANCIERSLIMITED",
      "symbol": "AAVAS",
      "percent": 60
    },
    "POLYPLEX": {
      "name": "POLYPLEXCORPORATIONLTD",
      "symbol": "POLYPLEX",
      "percent": 60
    },
    "UJJIVANSFB": {
      "name": "UJJIVANSMALLFINANCBANK",
      "symbol": "UJJIVANSFB",
      "percent": 60
    },
    "GALAXYSURF": {
      "name": "GALAXYSURFACTANTSLTD",
      "symbol": "GALAXYSURF",
      "percent": 65
    },
    "SONATSOFTW": {
      "name": "SONATASOFTWARELTD",
      "symbol": "SONATSOFTW",
      "percent": 60
    },
    "MAHSEAMLES": {
      "name": "MAHARASHTRASEAMLESSLTD",
      "symbol": "MAHSEAMLES",
      "percent": 60
    },
    "HEROMOTOCO": {
      "name": "HEROMOTOCORPLIMITED",
      "symbol": "HEROMOTOCO",
      "percent": 75
    },
    "STLTECH": {
      "name": "STERLITETECHNOLOGIESLTD",
      "symbol": "STLTECH",
      "percent": 60
    },
    "TATACOFFEE": {
      "name": "TATACOFFEELIMITED",
      "symbol": "TATACOFFEE",
      "percent": 60
    },
    "SONACOMS": {
      "name": "SONABLWPRECISIONFRGSL",
      "symbol": "SONACOMS",
      "percent": 60
    },
    "MASTEK": {
      "name": "MASTEKLTD",
      "symbol": "MASTEK",
      "percent": 60
    },
    "GODREJPROP": {
      "name": "GODREJPROPERTIESLTD",
      "symbol": "GODREJPROP",
      "percent": 70
    },
    "TATASTEEL": {
      "name": "TATASTEELLIMITED",
      "symbol": "TATASTEEL",
      "percent": 70
    },
    "PVRINOX": {
      "name": "PVRINOXLIMITED",
      "symbol": "PVRINOX",
      "percent": 65
    },
    "MCX": {
      "name": "MULTICOMMODITYEXCHANGE",
      "symbol": "MCX",
      "percent": 70
    },
    "BRITANNIA": {
      "name": "BRITANNIAINDUSTRIESLTD",
      "symbol": "BRITANNIA",
      "percent": 75
    },
    "LT": {
      "name": "LARSEN&TOUBROLTD.",
      "symbol": "LT",
      "percent": 75
    },
    "GRASIM": {
      "name": "GRASIMINDUSTRIESLTD",
      "symbol": "GRASIM",
      "percent": 75
    },
    "JKCEMENT": {
      "name": "JKCEMENTLIMITED",
      "symbol": "JKCEMENT",
      "percent": 70
    },
    "CHALET": {
      "name": "CHALETHOTELSLIMITED",
      "symbol": "CHALET",
      "percent": 60
    },
    "MFSL": {
      "name": "MAXFINANCIALSERVLTD",
      "symbol": "MFSL",
      "percent": 60
    },
    "BALAMINES": {
      "name": "BALAJIAMINESLIMITED",
      "symbol": "BALAMINES",
      "percent": 60
    },
    "FINCABLES": {
      "name": "FINOLEXCABLESLTD",
      "symbol": "FINCABLES",
      "percent": 60
    },
    "GICRE": {
      "name": "GENERALINSCORPOFINDIA",
      "symbol": "GICRE",
      "percent": 60
    },
    "ALKEM": {
      "name": "ALKEMLABORATORIESLTD.",
      "symbol": "ALKEM",
      "percent": 75
    },
    "SCHAEFFLER": {
      "name": "SCHAEFFLERINDIALIMITED",
      "symbol": "SCHAEFFLER",
      "percent": 60
    },
    "PRAJIND": {
      "name": "PRAJINDUSTRIESLTD",
      "symbol": "PRAJIND",
      "percent": 60
    },
    "BAJAJELEC": {
      "name": "BAJAJELECT.LTD",
      "symbol": "BAJAJELEC",
      "percent": 60
    },
    "HOMEFIRST": {
      "name": "HOMEFIRSTFINCOINDLTD",
      "symbol": "HOMEFIRST",
      "percent": 60
    },
    "SIEMENS": {
      "name": "SIEMENSLTD",
      "symbol": "SIEMENS",
      "percent": 75
    },
    "ANGELONE": {
      "name": "ANGELONELIMITED",
      "symbol": "ANGELONE",
      "percent": 60
    },
    "MCDOWELL-N": {
      "name": "UNITEDSPIRITSLIMITED",
      "symbol": "MCDOWELL-N",
      "percent": 70
    },
    "TEAMLEASE": {
      "name": "TEAMLEASESERVICESLTD.",
      "symbol": "TEAMLEASE",
      "percent": 60
    },
    "OIL": {
      "name": "OILINDIALTD",
      "symbol": "OIL",
      "percent": 60
    },
    "CAMLINFINE": {
      "name": "CAMLINFINESCIENCESLTD",
      "symbol": "CAMLINFINE",
      "percent": 60
    },
    "SUMICHEM": {
      "name": "SUMITOMOCHEMINDIALTD",
      "symbol": "SUMICHEM",
      "percent": 60
    },
    "KPRMILL": {
      "name": "KPRMILLLTD.",
      "symbol": "KPRMILL",
      "percent": 60
    },
    "NH": {
      "name": "NARAYANAHRUDAYALAYALTD.",
      "symbol": "NH",
      "percent": 60
    },
    "NAM-INDIA": {
      "name": "NIPPONLIAMLTD",
      "symbol": "NAM-INDIA",
      "percent": 65
    },
    "SUDARSCHEM": {
      "name": "SUDARSHANCHEMICALINDSL",
      "symbol": "SUDARSCHEM",
      "percent": 60
    },
    "RADICO": {
      "name": "RADICOKHAITANLTD",
      "symbol": "RADICO",
      "percent": 60
    },
    "TIMKEN": {
      "name": "TIMKENINDIALTD.",
      "symbol": "TIMKEN",
      "percent": 60
    },
    "SRF": {
      "name": "SRFLTD",
      "symbol": "SRF",
      "percent": 70
    },
    "APOLLOTYRE": {
      "name": "APOLLOTYRESLTD",
      "symbol": "APOLLOTYRE",
      "percent": 70
    },
    "COLPAL": {
      "name": "COLGATEPALMOLIVELTD.",
      "symbol": "COLPAL",
      "percent": 75
    },
    "AXISBANK": {
      "name": "AXISBANKLIMITED",
      "symbol": "AXISBANK",
      "percent": 75
    },
    "CHAMBLFERT": {
      "name": "CHAMBALFERTILIZERSLTD",
      "symbol": "CHAMBLFERT",
      "percent": 65
    },
    "ZYDUSWELL": {
      "name": "ZYDUSWELLNESSLIMITED",
      "symbol": "ZYDUSWELL",
      "percent": 60
    },
    "THERMAX": {
      "name": "THERMAXLTD",
      "symbol": "THERMAX",
      "percent": 60
    },
    "JUBLINGREA": {
      "name": "JUBILANTINGREVIALIMITED",
      "symbol": "JUBLINGREA",
      "percent": 60
    },
    "VINATIORGA": {
      "name": "VINATIORGANICSLTD",
      "symbol": "VINATIORGA",
      "percent": 60
    },
    "ASIANPAINT": {
      "name": "ASIANPAINTSLIMITED",
      "symbol": "ASIANPAINT",
      "percent": 75
    },
    "BATAINDIA": {
      "name": "BATAINDIALTD",
      "symbol": "BATAINDIA",
      "percent": 75
    },
    "TVSMOTOR": {
      "name": "TVSMOTORCOMPANYLTD",
      "symbol": "TVSMOTOR",
      "percent": 75
    },
    "GLENMARK": {
      "name": "GLENMARKPHARMACEUTICALS",
      "symbol": "GLENMARK",
      "percent": 70
    },
    "PERSISTENT": {
      "name": "PERSISTENTSYSTEMSLTD",
      "symbol": "PERSISTENT",
      "percent": 70
    },
    "BIRLACORPN": {
      "name": "BIRLACORPORATIONLTD",
      "symbol": "BIRLACORPN",
      "percent": 60
    },
    "KOTAKBANK": {
      "name": "KOTAKMAHINDRABANKLTD",
      "symbol": "KOTAKBANK",
      "percent": 75
    },
    "FEDERALBNK": {
      "name": "FEDERALBANKLTD",
      "symbol": "FEDERALBNK",
      "percent": 70
    },
    "ESCORTS": {
      "name": "ESCORTSKUBOTALIMITED",
      "symbol": "ESCORTS",
      "percent": 65
    },
    "MGL": {
      "name": "MAHANAGARGASLTD.",
      "symbol": "MGL",
      "percent": 70
    },
    "M&M": {
      "name": "MAHINDRA&MAHINDRALTD",
      "symbol": "M&M",
      "percent": 75
    },
    "SKFINDIA": {
      "name": "SKFINDIALTD",
      "symbol": "SKFINDIA",
      "percent": 60
    },
    "ASAHIINDIA": {
      "name": "ASAHIINDIAGLASSLIMITED",
      "symbol": "ASAHIINDIA",
      "percent": 60
    },
    "COROMANDEL": {
      "name": "COROMANDELINTERNTL.LTD",
      "symbol": "COROMANDEL",
      "percent": 75
    },
    "LXCHEM": {
      "name": "LAXMIORGANICINDUSLTD",
      "symbol": "LXCHEM",
      "percent": 60
    },
    "ASTRAL": {
      "name": "ASTRALLIMITED",
      "symbol": "ASTRAL",
      "percent": 70
    },
    "UTIAMC": {
      "name": "UTIASSETMNGMTCOLTD",
      "symbol": "UTIAMC",
      "percent": 60
    },
    "GRINDWELL": {
      "name": "GRINDWELLNORTONLIMITED",
      "symbol": "GRINDWELL",
      "percent": 60
    },
    "DELTACORP": {
      "name": "DELTACORPLIMITED",
      "symbol": "DELTACORP",
      "percent": 60
    },
    "ENDURANCE": {
      "name": "ENDURANCETECHNO.LTD.",
      "symbol": "ENDURANCE",
      "percent": 60
    },
    "CYIENT": {
      "name": "CYIENTLIMITED",
      "symbol": "CYIENT",
      "percent": 60
    },
    "PNB": {
      "name": "PUNJABNATIONALBANK",
      "symbol": "PNB",
      "percent": 65
    },
    "JKIL": {
      "name": "JKUMARINFR.LTD.",
      "symbol": "JKIL",
      "percent": 60
    },
    "AIAENG": {
      "name": "AIAENGINEERINGLIMITED",
      "symbol": "AIAENG",
      "percent": 65
    },
    "CERA": {
      "name": "CERASANITARYWARELTD",
      "symbol": "CERA",
      "percent": 60
    },
    "CUB": {
      "name": "CITYUNIONBANKLTD",
      "symbol": "CUB",
      "percent": 70
    },
    "ATUL": {
      "name": "ATULLTD",
      "symbol": "ATUL",
      "percent": 70
    },
    "CONCOR": {
      "name": "CONTAINERCORPOFINDLTD",
      "symbol": "CONCOR",
      "percent": 70
    },
    "SUNPHARMA": {
      "name": "SUNPHARMACEUTICALINDL",
      "symbol": "SUNPHARMA",
      "percent": 75
    },
    "NTPC": {
      "name": "NTPCLTD",
      "symbol": "NTPC",
      "percent": 75
    },
    "INDUSTOWER": {
      "name": "INDUSTOWERSLIMITED",
      "symbol": "INDUSTOWER",
      "percent": 70
    },
    "REDINGTON": {
      "name": "REDINGTONLIMITED",
      "symbol": "REDINGTON",
      "percent": 60
    },
    "LTTS": {
      "name": "L&TTECHNOLOGYSER.LTD.",
      "symbol": "LTTS",
      "percent": 70
    },
    "SUPREMEIND": {
      "name": "SUPREMEINDUSTRIESLTD",
      "symbol": "SUPREMEIND",
      "percent": 60
    },
    "INDIGOPNTS": {
      "name": "INDIGOPAINTSLIMITED",
      "symbol": "INDIGOPNTS",
      "percent": 65
    },
    "TATAPOWER": {
      "name": "TATAPOWERCOLTD",
      "symbol": "TATAPOWER",
      "percent": 70
    },
    "ICICIPRULI": {
      "name": "ICICIPRULIFEINSCOLTD",
      "symbol": "ICICIPRULI",
      "percent": 75
    },
    "CESC": {
      "name": "CESCLTD",
      "symbol": "CESC",
      "percent": 65
    },
    "NAVINFLUOR": {
      "name": "NAVINFLUORINEINT.LTD",
      "symbol": "NAVINFLUOR",
      "percent": 70
    },
    "INDIGO": {
      "name": "INTERGLOBEAVIATIONLTD",
      "symbol": "INDIGO",
      "percent": 70
    },
    "KEI": {
      "name": "KEIINDUSTRIESLTD.",
      "symbol": "KEI",
      "percent": 60
    },
    "CROMPTON": {
      "name": "CROMPTGREACONELECLTD",
      "symbol": "CROMPTON",
      "percent": 75
    },
    "KTKBANK": {
      "name": "KARNATAKABANKLIMITED",
      "symbol": "KTKBANK",
      "percent": 60
    },
    "HEIDELBERG": {
      "name": "HEIDELBERGCEMENT(I)LTD",
      "symbol": "HEIDELBERG",
      "percent": 60
    },
    "RAJESHEXPO": {
      "name": "RAJESHEXPORTSLTD",
      "symbol": "RAJESHEXPO",
      "percent": 60
    },
    "BAJFINANCE": {
      "name": "BAJAJFINANCELIMITED",
      "symbol": "BAJFINANCE",
      "percent": 70
    },
    "INDIACEM": {
      "name": "THEINDIACEMENTSLIMITED",
      "symbol": "INDIACEM",
      "percent": 65
    },
    "RELIANCE": {
      "name": "RELIANCEINDUSTRIESLTD",
      "symbol": "RELIANCE",
      "percent": 75
    },
    "HONAUT": {
      "name": "HONEYWELLAUTOMATIONIND",
      "symbol": "HONAUT",
      "percent": 70
    },
    "CRISIL": {
      "name": "CRISILLTD",
      "symbol": "CRISIL",
      "percent": 60
    },
    "HINDPETRO": {
      "name": "HINDUSTANPETROLEUMCORP",
      "symbol": "HINDPETRO",
      "percent": 70
    },
    "AFFLE": {
      "name": "AFFLE(INDIA)LIMITED",
      "symbol": "AFFLE",
      "percent": 60
    },
    "IFBIND": {
      "name": "IFBINDUSTRIESLTD",
      "symbol": "IFBIND",
      "percent": 60
    },
    "GUJGASLTD": {
      "name": "GUJARATGASLIMITED",
      "symbol": "GUJGASLTD",
      "percent": 65
    },
    "DIVISLAB": {
      "name": "DIVISLABORATORIESLTD",
      "symbol": "DIVISLAB",
      "percent": 75
    },
    "OFSS": {
      "name": "ORACLEFINSERVSOFTLTD.",
      "symbol": "OFSS",
      "percent": 70
    },
    "NMDC": {
      "name": "NMDCLTD.",
      "symbol": "NMDC",
      "percent": 70
    },
    "IDFCFIRSTB": {
      "name": "IDFCFIRSTBANKLIMITED",
      "symbol": "IDFCFIRSTB",
      "percent": 70
    },
    "ANURAS": {
      "name": "ANUPAMRASAYANINDIALTD",
      "symbol": "ANURAS",
      "percent": 60
    },
    "BLUESTARCO": {
      "name": "BLUESTARLIMITED",
      "symbol": "BLUESTARCO",
      "percent": 65
    },
    "PIDILITIND": {
      "name": "PIDILITEINDUSTRIESLTD",
      "symbol": "PIDILITIND",
      "percent": 75
    },
    "IPCALAB": {
      "name": "IPCALABORATORIESLTD",
      "symbol": "IPCALAB",
      "percent": 75
    },
    "HAL": {
      "name": "HINDUSTANAERONAUTICSLTD",
      "symbol": "HAL",
      "percent": 65
    },
    "DIXON": {
      "name": "DIXONTECHNO(INDIA)LTD",
      "symbol": "DIXON",
      "percent": 70
    },
    "MAHLIFE": {
      "name": "MAHINDRALIFESPACEDEVLTD",
      "symbol": "MAHLIFE",
      "percent": 60
    },
    "BALKRISIND": {
      "name": "BALKRISHNAIND.LTD",
      "symbol": "BALKRISIND",
      "percent": 75
    },
    "IDBI": {
      "name": "IDBIBANKLIMITED",
      "symbol": "IDBI",
      "percent": 60
    },
    "LICI": {
      "name": "LIFEINSURACORPOFINDIA",
      "symbol": "LICI",
      "percent": 70
    },
    "ROUTE": {
      "name": "ROUTEMOBILELIMITED",
      "symbol": "ROUTE",
      "percent": 60
    },
    "NIFTYBEES": {
      "name": "NIPINDETFNIFTYBEES",
      "symbol": "NIFTYBEES",
      "percent": 75
    },
    "SBICARD": {
      "name": "SBICARDS&PAYSERLTD",
      "symbol": "SBICARD",
      "percent": 75
    },
    "NESCO": {
      "name": "NESCOLTD.",
      "symbol": "NESCO",
      "percent": 60
    },
    "APLAPOLLO": {
      "name": "APLAPOLLOTUBESLTD",
      "symbol": "APLAPOLLO",
      "percent": 60
    },
    "GSFC": {
      "name": "GUJSTATEFERT&CHEMLTD",
      "symbol": "GSFC",
      "percent": 60
    },
    "JKLAKSHMI": {
      "name": "JKLAKSHMICEMENTLTD",
      "symbol": "JKLAKSHMI",
      "percent": 60
    },
    "ABFRL": {
      "name": "ADITYABIRLAFASHION&RT",
      "symbol": "ABFRL",
      "percent": 65
    },
    "ECLERX": {
      "name": "ECLERXSERVICESLTD",
      "symbol": "ECLERX",
      "percent": 60
    },
    "LATENTVIEW": {
      "name": "LATENTVIEWANALYTICSLTD",
      "symbol": "LATENTVIEW",
      "percent": 60
    },
    "WHIRLPOOL": {
      "name": "WHIRLPOOLOFINDIALTD",
      "symbol": "WHIRLPOOL",
      "percent": 70
    },
    "SUPRIYA": {
      "name": "SUPRIYALIFESCIENCELTD",
      "symbol": "SUPRIYA",
      "percent": 60
    },
    "DEEPAKFERT": {
      "name": "DEEPAKFERTILIZERS&PETR",
      "symbol": "DEEPAKFERT",
      "percent": 60
    },
    "RECLTD": {
      "name": "RECLIMITED",
      "symbol": "RECLTD",
      "percent": 75
    },
    "WIPRO": {
      "name": "WIPROLTD",
      "symbol": "WIPRO",
      "percent": 75
    },
    "NESTLEIND": {
      "name": "NESTLEINDIALIMITED",
      "symbol": "NESTLEIND",
      "percent": 75
    },
    "CAMS": {
      "name": "COMPUTERAGEMNGTSERLTD",
      "symbol": "CAMS",
      "percent": 65
    },
    "ROLEXRINGS": {
      "name": "ROLEXRINGSLIMITED",
      "symbol": "ROLEXRINGS",
      "percent": 60
    },
    "MOTHERSON": {
      "name": "SAMVRDHNAMTHRSNINTLLTD",
      "symbol": "MOTHERSON",
      "percent": 70
    },
    "ENGINERSIN": {
      "name": "ENGINEERSINDIALTD",
      "symbol": "ENGINERSIN",
      "percent": 60
    },
    "BAJAJFINSV": {
      "name": "BAJAJFINSERVLTD.",
      "symbol": "BAJAJFINSV",
      "percent": 70
    },
    "ICICIGI": {
      "name": "ICICILOMBARDGICLIMITED",
      "symbol": "ICICIGI",
      "percent": 75
    },
    "SBILIFE": {
      "name": "SBILIFEINSURANCECOLTD",
      "symbol": "SBILIFE",
      "percent": 75
    },
    "APARINDS": {
      "name": "APARINDUSTRIESLTD.",
      "symbol": "APARINDS",
      "percent": 60
    },
    "CANFINHOME": {
      "name": "CANFINHOMESLTD",
      "symbol": "CANFINHOME",
      "percent": 65
    },
    "IEX": {
      "name": "INDIANENERGYEXCLTD",
      "symbol": "IEX",
      "percent": 65
    },
    "BASF": {
      "name": "BASFINDIALTD",
      "symbol": "BASF",
      "percent": 60
    },
    "VOLTAS": {
      "name": "VOLTASLTD",
      "symbol": "VOLTAS",
      "percent": 75
    },
    "RITES": {
      "name": "RITESLIMITED",
      "symbol": "RITES",
      "percent": 65
    },
    "HEMIPROP": {
      "name": "HEMISPHEREPROPINDLTD",
      "symbol": "HEMIPROP",
      "percent": 60
    },
    "CAMPUS": {
      "name": "CAMPUSACTIVEWEARLIMITED",
      "symbol": "CAMPUS",
      "percent": 65
    },
    "KAJARIACER": {
      "name": "KAJARIACERAMICSLTD",
      "symbol": "KAJARIACER",
      "percent": 60
    },
    "SAFARI": {
      "name": "SAFARIIND(INDIA)LTD",
      "symbol": "SAFARI",
      "percent": 65
    },
    "PRESTIGE": {
      "name": "PRESTIGEESTATELTD",
      "symbol": "PRESTIGE",
      "percent": 60
    },
    "ABSLAMC": {
      "name": "ADITBIRLSUNLIFAMCLTD",
      "symbol": "ABSLAMC",
      "percent": 60
    },
    "ULTRACEMCO": {
      "name": "ULTRATECHCEMENTLIMITED",
      "symbol": "ULTRACEMCO",
      "percent": 75
    },
    "INTELLECT": {
      "name": "INTELLECTDESIGNARENA",
      "symbol": "INTELLECT",
      "percent": 60
    },
    "NUVOCO": {
      "name": "NUVOCOVISTASCORPLTD",
      "symbol": "NUVOCO",
      "percent": 60
    },
    "ONGC": {
      "name": "OILANDNATURALGASCORP.",
      "symbol": "ONGC",
      "percent": 70
    },
    "SUNTECK": {
      "name": "SUNTECKREALTYLIMITED",
      "symbol": "SUNTECK",
      "percent": 60
    },
    "MOL": {
      "name": "MEGHMANIORGANICSLIMITED",
      "symbol": "MOL",
      "percent": 60
    },
    "BHARATFORG": {
      "name": "BHARATFORGELTD",
      "symbol": "BHARATFORG",
      "percent": 70
    },
    "IGL": {
      "name": "INDRAPRASTHAGASLTD",
      "symbol": "IGL",
      "percent": 70
    },
    "DALBHARAT": {
      "name": "DALMIABHARATLIMITED",
      "symbol": "DALBHARAT",
      "percent": 70
    },
    "RBLBANK": {
      "name": "RBLBANKLIMITED",
      "symbol": "RBLBANK",
      "percent": 60
    },
    "STOVEKRAFT": {
      "name": "STOVEKRAFTLIMITED",
      "symbol": "STOVEKRAFT",
      "percent": 60
    },
    "PRINCEPIPE": {
      "name": "PRINCEPIPESFITTINGSLTD",
      "symbol": "PRINCEPIPE",
      "percent": 60
    },
    "MAHABANK": {
      "name": "BANKOFMAHARASHTRA",
      "symbol": "MAHABANK",
      "percent": 60
    },
    "HDFCLIFE": {
      "name": "HDFCLIFEINSCOLTD",
      "symbol": "HDFCLIFE",
      "percent": 75
    },
    "DMART": {
      "name": "AVENUESUPERMARTSLIMITED",
      "symbol": "DMART",
      "percent": 65
    },
    "ORIENTELEC": {
      "name": "ORIENTELECTRICLIMITED",
      "symbol": "ORIENTELEC",
      "percent": 60
    },
    "MOTILALOFS": {
      "name": "MOTILALOSWALFINLTD",
      "symbol": "MOTILALOFS",
      "percent": 60
    },
    "RALLIS": {
      "name": "RALLISINDIALTD",
      "symbol": "RALLIS",
      "percent": 60
    },
    "TATACONSUM": {
      "name": "TATACONSUMERPRODUCTLTD",
      "symbol": "TATACONSUM",
      "percent": 75
    },
    "SHRIRAMFIN": {
      "name": "SHRIRAMFINANCELIMITED",
      "symbol": "SHRIRAMFIN",
      "percent": 70
    },
    "TATACHEM": {
      "name": "TATACHEMICALSLTD",
      "symbol": "TATACHEM",
      "percent": 65
    },
    "HINDUNILVR": {
      "name": "HINDUSTANUNILEVERLTD.",
      "symbol": "HINDUNILVR",
      "percent": 75
    },
    "AUROPHARMA": {
      "name": "AUROBINDOPHARMALTD",
      "symbol": "AUROPHARMA",
      "percent": 70
    },
    "MANAPPURAM": {
      "name": "MANAPPURAMFINANCELTD",
      "symbol": "MANAPPURAM",
      "percent": 65
    },
    "GMMPFAUDLR": {
      "name": "GMMPFAUDLERLIMITED",
      "symbol": "GMMPFAUDLR",
      "percent": 60
    },
    "AMARAJABAT": {
      "name": "AMARARAJABATTERIESLTD.",
      "symbol": "AMARAJABAT",
      "percent": 70
    },
    "LUPIN": {
      "name": "LUPINLIMITED",
      "symbol": "LUPIN",
      "percent": 75
    },
    "DABUR": {
      "name": "DABURINDIALTD",
      "symbol": "DABUR",
      "percent": 80
    },
    "VBL": {
      "name": "VARUNBEVERAGESLIMITED",
      "symbol": "VBL",
      "percent": 60
    },
    "CENTURYPLY": {
      "name": "CENTURYPLYBOARDS(I)LTD",
      "symbol": "CENTURYPLY",
      "percent": 60
    },
    "ABCAPITAL": {
      "name": "ADITYABIRLACAPITALLTD.",
      "symbol": "ABCAPITAL",
      "percent": 70
    },
    "EVEREADY": {
      "name": "EVEREADYINDS.IND.LTD.",
      "symbol": "EVEREADY",
      "percent": 60
    },
    "JINDALSTEL": {
      "name": "JINDALSTEEL&POWERLTD",
      "symbol": "JINDALSTEL",
      "percent": 65
    },
    "MIDHANI": {
      "name": "MISHRADHATUNIGAMLTD",
      "symbol": "MIDHANI",
      "percent": 60
    },
    "TINPLATE": {
      "name": "THETINPLATECO.(I)LTD",
      "symbol": "TINPLATE",
      "percent": 60
    },
    "NOCIL": {
      "name": "NOCILLIMITED",
      "symbol": "NOCIL",
      "percent": 60
    },
    "NIITLTD": {
      "name": "NIITLIMITED",
      "symbol": "NIITLTD",
      "percent": 60
    },
    "IDFC": {
      "name": "IDFCLIMITED",
      "symbol": "IDFC",
      "percent": 65
    },
    "BRIGADE": {
      "name": "BRIGADEENTER.LTD",
      "symbol": "BRIGADE",
      "percent": 60
    },
    "MRF": {
      "name": "MRFLTD",
      "symbol": "MRF",
      "percent": 75
    },
    "POLYCAB": {
      "name": "POLYCABINDIALIMITED",
      "symbol": "POLYCAB",
      "percent": 65
    },
    "SHOPERSTOP": {
      "name": "SHOPPERSSTOPLIMITED",
      "symbol": "SHOPERSTOP",
      "percent": 60
    },
    "FDC": {
      "name": "FDCLIMITED",
      "symbol": "FDC",
      "percent": 60
    },
    "SYNGENE": {
      "name": "SYNGENEINTERNATIONALLTD",
      "symbol": "SYNGENE",
      "percent": 70
    },
    "GOODYEAR": {
      "name": "GOODYEARINDIALIMITED",
      "symbol": "GOODYEAR",
      "percent": 60
    },
    "GRANULES": {
      "name": "GRANULESINDIALIMITED",
      "symbol": "GRANULES",
      "percent": 65
    },
    "AMIORG": {
      "name": "AMIORGANICSLIMITED",
      "symbol": "AMIORG",
      "percent": 60
    },
    "NAUKRI": {
      "name": "INFOEDGE(I)LTD",
      "symbol": "NAUKRI",
      "percent": 70
    },
    "LINDEINDIA": {
      "name": "LINDEINDIALIMITED",
      "symbol": "LINDEINDIA",
      "percent": 60
    },
    "GAIL": {
      "name": "GAIL(INDIA)LTD",
      "symbol": "GAIL",
      "percent": 75
    },
    "UBL": {
      "name": "UNITEDBREWERIESLTD",
      "symbol": "UBL",
      "percent": 75
    },
    "VEDL": {
      "name": "VEDANTALIMITED",
      "symbol": "VEDL",
      "percent": 60
    },
    "UNIONBANK": {
      "name": "UNIONBANKOFINDIA",
      "symbol": "UNIONBANK",
      "percent": 60
    },
    "DCBBANK": {
      "name": "DCBBANKLIMITED",
      "symbol": "DCBBANK",
      "percent": 60
    },
    "GODREJCP": {
      "name": "GODREJCONSUMERPRODUCTS",
      "symbol": "GODREJCP",
      "percent": 75
    },
    "MARICO": {
      "name": "MARICOLIMITED",
      "symbol": "MARICO",
      "percent": 75
    },
    "AARTIIND": {
      "name": "AARTIINDUSTRIESLTD",
      "symbol": "AARTIIND",
      "percent": 70
    },
    "ISEC": {
      "name": "ICICISECURITIESLIMITED",
      "symbol": "ISEC",
      "percent": 65
    },
    "BERGEPAINT": {
      "name": "BERGERPAINTS(I)LTD",
      "symbol": "BERGEPAINT",
      "percent": 75
    },
    "BHEL": {
      "name": "BHEL",
      "symbol": "BHEL",
      "percent": 65
    },
    "HAPPSTMNDS": {
      "name": "HAPPIESTMINDSTECHNOLTD",
      "symbol": "HAPPSTMNDS",
      "percent": 60
    },
    "AMBER": {
      "name": "AMBERENTERPRISES(I)LTD",
      "symbol": "AMBER",
      "percent": 60
    },
    "MTARTECH": {
      "name": "MTARTECHNOLOGIESLIMITED",
      "symbol": "MTARTECH",
      "percent": 60
    },
    "TORNTPHARM": {
      "name": "TORRENTPHARMACEUTICALSL",
      "symbol": "TORNTPHARM",
      "percent": 75
    },
    "VRLLOG": {
      "name": "VRLLOGISTICSLIMITED",
      "symbol": "VRLLOG",
      "percent": 60
    },
    "CARERATING": {
      "name": "CARERATINGSLIMITED",
      "symbol": "CARERATING",
      "percent": 60
    },
    "BHARTIARTL": {
      "name": "BHARTIAIRTELLIMITED",
      "symbol": "BHARTIARTL",
      "percent": 75
    },
    "SHILPAMED": {
      "name": "SHILPAMEDICARELTD",
      "symbol": "SHILPAMED",
      "percent": 60
    },
    "TIINDIA": {
      "name": "TUBEINVESTOFINDIALTD",
      "symbol": "TIINDIA",
      "percent": 60
    },
    "IRFC": {
      "name": "INDIANRAILWAYFINCORPL",
      "symbol": "IRFC",
      "percent": 60
    },
    "RATEGAIN": {
      "name": "RATEGAINTRAVELTECHNLTD",
      "symbol": "RATEGAIN",
      "percent": 60
    },
    "BLUEDART": {
      "name": "BLUEDARTEXPRESSLTD",
      "symbol": "BLUEDART",
      "percent": 60
    },
    "DCMSHRIRAM": {
      "name": "DCMSHRIRAMLIMITED",
      "symbol": "DCMSHRIRAM",
      "percent": 60
    },
    "ASTERDM": {
      "name": "ASTERDMHEALTHCARELTD.",
      "symbol": "ASTERDM",
      "percent": 60
    },
    "EIDPARRY": {
      "name": "EIDPARRYINDIALTD",
      "symbol": "EIDPARRY",
      "percent": 60
    },
    "SBIN": {
      "name": "STATEBANKOFINDIA",
      "symbol": "SBIN",
      "percent": 75
    },
    "SHREECEM": {
      "name": "SHREECEMENTLIMITED",
      "symbol": "SHREECEM",
      "percent": 75
    },
    "EQUITASBNK": {
      "name": "EQUITASSMALLFINBNKLTD",
      "symbol": "EQUITASBNK",
      "percent": 60
    },
    "EMAMILTD": {
      "name": "EMAMILIMITED",
      "symbol": "EMAMILTD",
      "percent": 60
    },
    "GRAUWEIL": {
      "name": "GRAUER&WEILINDLTD",
      "symbol": "GRAUWEIL",
      "percent": 60
    },
    "WELCORP": {
      "name": "WELSPUNCORPLIMITED",
      "symbol": "WELCORP",
      "percent": 60
    },
    "KPITTECH": {
      "name": "KPITTECHNOLOGIESLIMITED",
      "symbol": "KPITTECH",
      "percent": 60
    },
    "AJANTPHARM": {
      "name": "AJANTAPHARMALIMITED",
      "symbol": "AJANTPHARM",
      "percent": 70
    },
    "MEDPLUS": {
      "name": "MEDPLUSHEALTHSERVLTD",
      "symbol": "MEDPLUS",
      "percent": 60
    },
    "KIMS": {
      "name": "KRISHNAINSTOFMEDSCIL",
      "symbol": "KIMS",
      "percent": 60
    },
    "INDIANB": {
      "name": "INDIANBANK",
      "symbol": "INDIANB",
      "percent": 60
    },
    "CUMMINSIND": {
      "name": "CUMMINSINDIALTD",
      "symbol": "CUMMINSIND",
      "percent": 75
    },
    "HCLTECH": {
      "name": "HCLTECHNOLOGIESLTD",
      "symbol": "HCLTECH",
      "percent": 75
    },
    "PIIND": {
      "name": "PIINDUSTRIESLTD",
      "symbol": "PIIND",
      "percent": 75
    },
    "BANKBEES": {
      "name": "NIPINDETFBANKBEES",
      "symbol": "BANKBEES",
      "percent": 70
    },
    "3MINDIA": {
      "name": "3MINDIALIMITED",
      "symbol": "3MINDIA",
      "percent": 60
    },
    "RAIN": {
      "name": "RAININDUSTRIESLIMITED",
      "symbol": "RAIN",
      "percent": 65
    },
    "BIOCON": {
      "name": "BIOCONLIMITED.",
      "symbol": "BIOCON",
      "percent": 65
    },
    "JUSTDIAL": {
      "name": "JUSTDIALLTD.",
      "symbol": "JUSTDIAL",
      "percent": 60
    },
    "TCS": {
      "name": "TATACONSULTANCYSERVLT",
      "symbol": "TCS",
      "percent": 75
    },
    "JKPAPER": {
      "name": "JKPAPERLIMITED",
      "symbol": "JKPAPER",
      "percent": 60
    },
    "GNFC": {
      "name": "GUJNARVALFER&CHEML",
      "symbol": "GNFC",
      "percent": 65
    },
    "AUBANK": {
      "name": "AUSMALLFINANCEBANKLTD",
      "symbol": "AUBANK",
      "percent": 70
    },
    "CHOLAFIN": {
      "name": "CHOLAMANDALAMIN&FINCO",
      "symbol": "CHOLAFIN",
      "percent": 70
    },
    "GSPL": {
      "name": "GUJARATSTATEPETROLTD",
      "symbol": "GSPL",
      "percent": 65
    },
    "HDFCBANK": {
      "name": "HDFCBANKLTD",
      "symbol": "HDFCBANK",
      "percent": 75
    },
    "BANKBARODA": {
      "name": "BANKOFBARODA",
      "symbol": "BANKBARODA",
      "percent": 70
    },
    "CGCL": {
      "name": "CAPRIGLOBALCAPITALLTD",
      "symbol": "CGCL",
      "percent": 60
    },
    "KRBL": {
      "name": "KRBLLIMITED",
      "symbol": "KRBL",
      "percent": 60
    },
    "SAIL": {
      "name": "STEELAUTHORITYOFINDIA",
      "symbol": "SAIL",
      "percent": 65
    },
    "AMRUTANJAN": {
      "name": "AMRUTAJANHEALTHLTD",
      "symbol": "AMRUTANJAN",
      "percent": 60
    },
    "IOC": {
      "name": "INDIANOILCORPLTD",
      "symbol": "IOC",
      "percent": 75
    },
    "BOSCHLTD": {
      "name": "BOSCHLIMITED",
      "symbol": "BOSCHLTD",
      "percent": 70
    },
    "PHOENIXLTD": {
      "name": "THEPHOENIXMILLSLTD",
      "symbol": "PHOENIXLTD",
      "percent": 60
    },
    "APLLTD": {
      "name": "ALEMBICPHARMALTD",
      "symbol": "APLLTD",
      "percent": 60
    },
    "CCL": {
      "name": "CCLPRODUCTS(I)LTD",
      "symbol": "CCL",
      "percent": 60
    },
    "DEVYANI": {
      "name": "DEVYANIINTERNATIONALLTD",
      "symbol": "DEVYANI",
      "percent": 60
    },
    "BANDHANBNK": {
      "name": "BANDHANBANKLIMITED",
      "symbol": "BANDHANBNK",
      "percent": 65
    },
    "MINDACORP": {
      "name": "MINDACORPORATIONLTD",
      "symbol": "MINDACORP",
      "percent": 60
    },
    "LUXIND": {
      "name": "LUXINDUSTRIESLIMITED",
      "symbol": "LUXIND",
      "percent": 60
    },
    "SCI": {
      "name": "SHIPPINGCORPOFINDIALT",
      "symbol": "SCI",
      "percent": 60
    },
    "COCHINSHIP": {
      "name": "COCHINSHIPYARDLIMITED",
      "symbol": "COCHINSHIP",
      "percent": 60
    },
    "APOLLOHOSP": {
      "name": "APOLLOHOSPITALSENTER.L",
      "symbol": "APOLLOHOSP",
      "percent": 70
    },
    "CARBORUNIV": {
      "name": "CARBORUNDUMUNIVERSALLTD",
      "symbol": "CARBORUNIV",
      "percent": 60
    },
    "RAILTEL": {
      "name": "RAILTELCORPOFINDLTD",
      "symbol": "RAILTEL",
      "percent": 60
    },
    "INDIAMART": {
      "name": "INDIAMARTINTERMESHLTD",
      "symbol": "INDIAMART",
      "percent": 70
    },
    "JUNIORBEES": {
      "name": "NIPINDETFJUNIORBEES",
      "symbol": "JUNIORBEES",
      "percent": 60
    },
    "ZYDUSLIFE": {
      "name": "ZYDUSLIFESCIENCESLTD",
      "symbol": "ZYDUSLIFE",
      "percent": 75
    },
    "GLAXO": {
      "name": "GLAXOSMITHKLINEPHARMALT",
      "symbol": "GLAXO",
      "percent": 65
    }
  }
`
