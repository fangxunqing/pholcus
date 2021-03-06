// This file was generated by go generate; DO NOT EDIT

package number

import "golang.org/x/text/internal/stringset"

// CLDRVersion is the CLDR version from which the tables in this package are derived.
const CLDRVersion = "28"

var numSysData = []systemData{ // 55 elements
	0:  {id: 0x0, digitSize: 0x1, zero: [4]uint8{0x30, 0x0, 0x0, 0x0}},
	1:  {id: 0x1, digitSize: 0x4, zero: [4]uint8{0xf0, 0x91, 0x9c, 0xb0}},
	2:  {id: 0x2, digitSize: 0x2, zero: [4]uint8{0xd9, 0xa0, 0x0, 0x0}},
	3:  {id: 0x3, digitSize: 0x2, zero: [4]uint8{0xdb, 0xb0, 0x0, 0x0}},
	4:  {id: 0x4, digitSize: 0x3, zero: [4]uint8{0xe1, 0xad, 0x90, 0x0}},
	5:  {id: 0x5, digitSize: 0x3, zero: [4]uint8{0xe0, 0xa7, 0xa6, 0x0}},
	6:  {id: 0x6, digitSize: 0x4, zero: [4]uint8{0xf0, 0x91, 0x81, 0xa6}},
	7:  {id: 0x7, digitSize: 0x4, zero: [4]uint8{0xf0, 0x91, 0x84, 0xb6}},
	8:  {id: 0x8, digitSize: 0x3, zero: [4]uint8{0xea, 0xa9, 0x90, 0x0}},
	9:  {id: 0x9, digitSize: 0x3, zero: [4]uint8{0xe0, 0xa5, 0xa6, 0x0}},
	10: {id: 0xa, digitSize: 0x3, zero: [4]uint8{0xef, 0xbc, 0x90, 0x0}},
	11: {id: 0xb, digitSize: 0x3, zero: [4]uint8{0xe0, 0xab, 0xa6, 0x0}},
	12: {id: 0xc, digitSize: 0x3, zero: [4]uint8{0xe0, 0xa9, 0xa6, 0x0}},
	13: {id: 0xd, digitSize: 0x4, zero: [4]uint8{0xf0, 0x96, 0xad, 0x90}},
	14: {id: 0xe, digitSize: 0x3, zero: [4]uint8{0xea, 0xa7, 0x90, 0x0}},
	15: {id: 0xf, digitSize: 0x3, zero: [4]uint8{0xea, 0xa4, 0x80, 0x0}},
	16: {id: 0x10, digitSize: 0x3, zero: [4]uint8{0xe1, 0x9f, 0xa0, 0x0}},
	17: {id: 0x11, digitSize: 0x3, zero: [4]uint8{0xe0, 0xb3, 0xa6, 0x0}},
	18: {id: 0x12, digitSize: 0x3, zero: [4]uint8{0xe1, 0xaa, 0x80, 0x0}},
	19: {id: 0x13, digitSize: 0x3, zero: [4]uint8{0xe1, 0xaa, 0x90, 0x0}},
	20: {id: 0x14, digitSize: 0x3, zero: [4]uint8{0xe0, 0xbb, 0x90, 0x0}},
	21: {id: 0x15, digitSize: 0x3, zero: [4]uint8{0xe1, 0xb1, 0x80, 0x0}},
	22: {id: 0x16, digitSize: 0x3, zero: [4]uint8{0xe1, 0xa5, 0x86, 0x0}},
	23: {id: 0x17, digitSize: 0x4, zero: [4]uint8{0xf0, 0x9d, 0x9f, 0x8e}},
	24: {id: 0x18, digitSize: 0x4, zero: [4]uint8{0xf0, 0x9d, 0x9f, 0x98}},
	25: {id: 0x19, digitSize: 0x4, zero: [4]uint8{0xf0, 0x9d, 0x9f, 0xb6}},
	26: {id: 0x1a, digitSize: 0x4, zero: [4]uint8{0xf0, 0x9d, 0x9f, 0xac}},
	27: {id: 0x1b, digitSize: 0x4, zero: [4]uint8{0xf0, 0x9d, 0x9f, 0xa2}},
	28: {id: 0x1c, digitSize: 0x3, zero: [4]uint8{0xe0, 0xb5, 0xa6, 0x0}},
	29: {id: 0x1d, digitSize: 0x4, zero: [4]uint8{0xf0, 0x91, 0x99, 0x90}},
	30: {id: 0x1e, digitSize: 0x3, zero: [4]uint8{0xe1, 0xa0, 0x90, 0x0}},
	31: {id: 0x1f, digitSize: 0x4, zero: [4]uint8{0xf0, 0x96, 0xa9, 0xa0}},
	32: {id: 0x20, digitSize: 0x3, zero: [4]uint8{0xea, 0xaf, 0xb0, 0x0}},
	33: {id: 0x21, digitSize: 0x3, zero: [4]uint8{0xe1, 0x81, 0x80, 0x0}},
	34: {id: 0x22, digitSize: 0x3, zero: [4]uint8{0xe1, 0x82, 0x90, 0x0}},
	35: {id: 0x23, digitSize: 0x3, zero: [4]uint8{0xea, 0xa7, 0xb0, 0x0}},
	36: {id: 0x24, digitSize: 0x2, zero: [4]uint8{0xdf, 0x80, 0x0, 0x0}},
	37: {id: 0x25, digitSize: 0x3, zero: [4]uint8{0xe1, 0xb1, 0x90, 0x0}},
	38: {id: 0x26, digitSize: 0x3, zero: [4]uint8{0xe0, 0xad, 0xa6, 0x0}},
	39: {id: 0x27, digitSize: 0x4, zero: [4]uint8{0xf0, 0x90, 0x92, 0xa0}},
	40: {id: 0x28, digitSize: 0x3, zero: [4]uint8{0xea, 0xa3, 0x90, 0x0}},
	41: {id: 0x29, digitSize: 0x4, zero: [4]uint8{0xf0, 0x91, 0x87, 0x90}},
	42: {id: 0x2a, digitSize: 0x4, zero: [4]uint8{0xf0, 0x91, 0x8b, 0xb0}},
	43: {id: 0x2b, digitSize: 0x3, zero: [4]uint8{0xe0, 0xb7, 0xa6, 0x0}},
	44: {id: 0x2c, digitSize: 0x4, zero: [4]uint8{0xf0, 0x91, 0x83, 0xb0}},
	45: {id: 0x2d, digitSize: 0x3, zero: [4]uint8{0xe1, 0xae, 0xb0, 0x0}},
	46: {id: 0x2e, digitSize: 0x4, zero: [4]uint8{0xf0, 0x91, 0x9b, 0x80}},
	47: {id: 0x2f, digitSize: 0x3, zero: [4]uint8{0xe1, 0xa7, 0x90, 0x0}},
	48: {id: 0x30, digitSize: 0x3, zero: [4]uint8{0xe0, 0xaf, 0xa6, 0x0}},
	49: {id: 0x31, digitSize: 0x3, zero: [4]uint8{0xe0, 0xb1, 0xa6, 0x0}},
	50: {id: 0x32, digitSize: 0x3, zero: [4]uint8{0xe0, 0xb9, 0x90, 0x0}},
	51: {id: 0x33, digitSize: 0x3, zero: [4]uint8{0xe0, 0xbc, 0xa0, 0x0}},
	52: {id: 0x34, digitSize: 0x4, zero: [4]uint8{0xf0, 0x91, 0x93, 0x90}},
	53: {id: 0x35, digitSize: 0x3, zero: [4]uint8{0xea, 0x98, 0xa0, 0x0}},
	54: {id: 0x36, digitSize: 0x4, zero: [4]uint8{0xf0, 0x91, 0xa3, 0xa0}},
} // Size: 354 bytes

const (
	numAhom     = 0x1
	numArab     = 0x2
	numArabext  = 0x3
	numArmn     = 0x37
	numArmnlow  = 0x38
	numBali     = 0x4
	numBeng     = 0x5
	numBrah     = 0x6
	numCakm     = 0x7
	numCham     = 0x8
	numCyrl     = 0x39
	numDeva     = 0x9
	numEthi     = 0x3a
	numFullwide = 0xa
	numGeor     = 0x3b
	numGrek     = 0x3c
	numGreklow  = 0x3d
	numGujr     = 0xb
	numGuru     = 0xc
	numHanidays = 0x3e
	numHanidec  = 0x3f
	numHans     = 0x40
	numHansfin  = 0x41
	numHant     = 0x42
	numHantfin  = 0x43
	numHebr     = 0x44
	numHmng     = 0xd
	numJava     = 0xe
	numJpan     = 0x45
	numJpanfin  = 0x46
	numKali     = 0xf
	numKhmr     = 0x10
	numKnda     = 0x11
	numLana     = 0x12
	numLanatham = 0x13
	numLaoo     = 0x14
	numLatn     = 0x0
	numLepc     = 0x15
	numLimb     = 0x16
	numMathbold = 0x17
	numMathdbl  = 0x18
	numMathmono = 0x19
	numMathsanb = 0x1a
	numMathsans = 0x1b
	numMlym     = 0x1c
	numModi     = 0x1d
	numMong     = 0x1e
	numMroo     = 0x1f
	numMtei     = 0x20
	numMymr     = 0x21
	numMymrshan = 0x22
	numMymrtlng = 0x23
	numNkoo     = 0x24
	numOlck     = 0x25
	numOrya     = 0x26
	numOsma     = 0x27
	numRoman    = 0x47
	numRomanlow = 0x48
	numSaur     = 0x28
	numShrd     = 0x29
	numSind     = 0x2a
	numSinh     = 0x2b
	numSora     = 0x2c
	numSund     = 0x2d
	numTakr     = 0x2e
	numTalu     = 0x2f
	numTaml     = 0x49
	numTamldec  = 0x30
	numTelu     = 0x31
	numThai     = 0x32
	numTibt     = 0x33
	numTirh     = 0x34
	numVaii     = 0x35
	numWara     = 0x36
	numNumberSystems
)

var systemMap = map[string]system{
	"ahom":     numAhom,
	"arab":     numArab,
	"arabext":  numArabext,
	"armn":     numArmn,
	"armnlow":  numArmnlow,
	"bali":     numBali,
	"beng":     numBeng,
	"brah":     numBrah,
	"cakm":     numCakm,
	"cham":     numCham,
	"cyrl":     numCyrl,
	"deva":     numDeva,
	"ethi":     numEthi,
	"fullwide": numFullwide,
	"geor":     numGeor,
	"grek":     numGrek,
	"greklow":  numGreklow,
	"gujr":     numGujr,
	"guru":     numGuru,
	"hanidays": numHanidays,
	"hanidec":  numHanidec,
	"hans":     numHans,
	"hansfin":  numHansfin,
	"hant":     numHant,
	"hantfin":  numHantfin,
	"hebr":     numHebr,
	"hmng":     numHmng,
	"java":     numJava,
	"jpan":     numJpan,
	"jpanfin":  numJpanfin,
	"kali":     numKali,
	"khmr":     numKhmr,
	"knda":     numKnda,
	"lana":     numLana,
	"lanatham": numLanatham,
	"laoo":     numLaoo,
	"latn":     numLatn,
	"lepc":     numLepc,
	"limb":     numLimb,
	"mathbold": numMathbold,
	"mathdbl":  numMathdbl,
	"mathmono": numMathmono,
	"mathsanb": numMathsanb,
	"mathsans": numMathsans,
	"mlym":     numMlym,
	"modi":     numModi,
	"mong":     numMong,
	"mroo":     numMroo,
	"mtei":     numMtei,
	"mymr":     numMymr,
	"mymrshan": numMymrshan,
	"mymrtlng": numMymrtlng,
	"nkoo":     numNkoo,
	"olck":     numOlck,
	"orya":     numOrya,
	"osma":     numOsma,
	"roman":    numRoman,
	"romanlow": numRomanlow,
	"saur":     numSaur,
	"shrd":     numShrd,
	"sind":     numSind,
	"sinh":     numSinh,
	"sora":     numSora,
	"sund":     numSund,
	"takr":     numTakr,
	"talu":     numTalu,
	"taml":     numTaml,
	"tamldec":  numTamldec,
	"telu":     numTelu,
	"thai":     numThai,
	"tibt":     numTibt,
	"tirh":     numTirh,
	"vaii":     numVaii,
	"wara":     numWara,
}

var symIndex = [][12]uint8{ // 68 elements
	0:  [12]uint8{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb},
	1:  [12]uint8{0x1, 0xc, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb},
	2:  [12]uint8{0x0, 0x1, 0x2, 0xd, 0xe, 0xf, 0x6, 0x7, 0x8, 0x9, 0x10, 0xb},
	3:  [12]uint8{0x1, 0x0, 0x2, 0xd, 0xe, 0xf, 0x6, 0x7, 0x8, 0x9, 0x10, 0xb},
	4:  [12]uint8{0x1, 0x0, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x11, 0xb},
	5:  [12]uint8{0x1, 0x0, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb},
	6:  [12]uint8{0x1, 0xc, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0x0},
	7:  [12]uint8{0x1, 0xc, 0x2, 0x3, 0x4, 0x5, 0x6, 0x0, 0x8, 0x9, 0xa, 0xb},
	8:  [12]uint8{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x12, 0xb},
	9:  [12]uint8{0x0, 0x1, 0x2, 0xd, 0xe, 0xf, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb},
	10: [12]uint8{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x13, 0x8, 0x9, 0xa, 0xb},
	11: [12]uint8{0x1, 0x0, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0x0},
	12: [12]uint8{0x1, 0x0, 0x2, 0x3, 0x4, 0x5, 0x6, 0x14, 0x8, 0x9, 0xa, 0xb},
	13: [12]uint8{0x1, 0xc, 0x2, 0x3, 0x4, 0x5, 0x6, 0x14, 0x8, 0x9, 0xa, 0xb},
	14: [12]uint8{0x0, 0x15, 0x2, 0x3, 0x4, 0x5, 0x6, 0x14, 0x8, 0x9, 0xa, 0xb},
	15: [12]uint8{0x0, 0xc, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb},
	16: [12]uint8{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x16, 0xb},
	17: [12]uint8{0x1, 0x0, 0x2, 0x3, 0x4, 0x5, 0x17, 0x7, 0x8, 0x9, 0xa, 0xb},
	18: [12]uint8{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x17, 0x7, 0x8, 0x9, 0xa, 0xb},
	19: [12]uint8{0x1, 0xc, 0x2, 0x3, 0x4, 0x5, 0x18, 0x7, 0x8, 0x9, 0xa, 0xb},
	20: [12]uint8{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x19, 0x1a, 0xa, 0xb},
	21: [12]uint8{0x1, 0xc, 0x2, 0x3, 0x4, 0x1b, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb},
	22: [12]uint8{0x1, 0xc, 0x2, 0x3, 0x4, 0x1b, 0x18, 0x7, 0x8, 0x9, 0xa, 0xb},
	23: [12]uint8{0x0, 0x1, 0x2, 0x3, 0xe, 0x1c, 0x6, 0x7, 0x8, 0x9, 0x1d, 0xb},
	24: [12]uint8{0x1, 0xc, 0x2, 0x3, 0x4, 0x1b, 0x6, 0x7, 0x8, 0x9, 0x1e, 0x0},
	25: [12]uint8{0x1, 0x0, 0x2, 0x3, 0x4, 0x1b, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb},
	26: [12]uint8{0x0, 0x1f, 0x2, 0x3, 0x4, 0x1b, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb},
	27: [12]uint8{0x0, 0x1, 0x2, 0x3, 0xe, 0xf, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb},
	28: [12]uint8{0x0, 0x15, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb},
	29: [12]uint8{0x1, 0xc, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x20, 0xb},
	30: [12]uint8{0x1, 0x0, 0x2, 0x3, 0x4, 0x5, 0x18, 0x14, 0x8, 0x9, 0x21, 0xb},
	31: [12]uint8{0x0, 0x1, 0x2, 0x3, 0xe, 0xf, 0x6, 0x7, 0x8, 0x22, 0xa, 0xb},
	32: [12]uint8{0x1, 0xc, 0x2, 0x3, 0x4, 0x1b, 0x18, 0x7, 0x8, 0x9, 0x21, 0xb},
	33: [12]uint8{0x1, 0xc, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x23, 0xb},
	34: [12]uint8{0x1, 0x0, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x24, 0xb},
	35: [12]uint8{0x1, 0xc, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x25, 0xb},
	36: [12]uint8{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x26, 0xb},
	37: [12]uint8{0x1, 0xc, 0x2, 0x3, 0x4, 0x1b, 0x6, 0x7, 0x8, 0x9, 0xa, 0x0},
	38: [12]uint8{0x1, 0xc, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x27, 0xb},
	39: [12]uint8{0x1, 0x0, 0x2, 0x3, 0xe, 0x1c, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb},
	40: [12]uint8{0x1, 0xc, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x28, 0xb},
	41: [12]uint8{0x1, 0xc, 0x2, 0x3, 0x4, 0x1b, 0x18, 0x14, 0x8, 0x9, 0x21, 0xb},
	42: [12]uint8{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0x0},
	43: [12]uint8{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x29, 0xb},
	44: [12]uint8{0x1, 0xc, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x2a, 0xb},
	45: [12]uint8{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x2b, 0xb},
	46: [12]uint8{0x1, 0xc, 0x2, 0x3, 0x4, 0x5, 0x2c, 0x7, 0x8, 0x9, 0xa, 0xb},
	47: [12]uint8{0x1, 0xc, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x2d, 0xb},
	48: [12]uint8{0x1, 0x1f, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb},
	49: [12]uint8{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x2e, 0xb},
	50: [12]uint8{0x2f, 0x30, 0x31, 0xd, 0x32, 0x33, 0x34, 0x7, 0x35, 0x9, 0xa, 0xb},
	51: [12]uint8{0x2f, 0x30, 0x31, 0xd, 0x32, 0x33, 0x34, 0x7, 0x35, 0x9, 0x36, 0xb},
	52: [12]uint8{0x2f, 0x30, 0x31, 0xd, 0x32, 0x37, 0x34, 0x7, 0x35, 0x9, 0xa, 0xb},
	53: [12]uint8{0x2f, 0xc, 0x31, 0xd, 0x4, 0x1b, 0x34, 0x7, 0x35, 0x9, 0xa, 0x0},
	54: [12]uint8{0x2f, 0x30, 0x31, 0xd, 0x38, 0x39, 0x3a, 0x7, 0x35, 0x9, 0xa, 0x2f},
	55: [12]uint8{0x2f, 0x30, 0x31, 0xd, 0x38, 0x1c, 0x3a, 0x7, 0x35, 0x9, 0x1d, 0xb},
	56: [12]uint8{0x2f, 0x30, 0x31, 0xd, 0xe, 0x1c, 0x3a, 0x7, 0x35, 0x9, 0xa, 0x2f},
	57: [12]uint8{0x1, 0xc, 0x31, 0xd, 0x4, 0x1b, 0x3a, 0x7, 0x35, 0x9, 0xa, 0x0},
	58: [12]uint8{0x2f, 0x1, 0x31, 0xd, 0x4, 0x5, 0x3a, 0x7, 0x35, 0x9, 0xa, 0x2f},
	59: [12]uint8{0x2f, 0x30, 0x31, 0xd, 0x38, 0x37, 0x3a, 0x7, 0x35, 0x9, 0xa, 0xb},
	60: [12]uint8{0x0, 0x1, 0x2, 0x3, 0x38, 0x39, 0x3b, 0x7, 0x8, 0x9, 0x3c, 0x2f},
	61: [12]uint8{0x2f, 0x30, 0x31, 0xd, 0x4, 0x5, 0x3a, 0x7, 0x35, 0x9, 0x2d, 0x2f},
	62: [12]uint8{0x2f, 0x30, 0x31, 0xd, 0x4, 0x5, 0x3a, 0x7, 0x35, 0x9, 0xa, 0x2f},
	63: [12]uint8{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x3d, 0xb},
	64: [12]uint8{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x3e, 0xb},
	65: [12]uint8{0x0, 0x1, 0x3f, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb},
	66: [12]uint8{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0x40, 0xb},
	67: [12]uint8{0x0, 0x1, 0x41, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x42, 0x43, 0xb},
} // Size: 840 bytes

var symData = stringset.Set{
	Data: "" + // Size: 560 bytes
		".,;%+-E×‰∞NaN:\u00a0٪\u200e+\u200e-ليس\u00a0رقمًاNDТерхьаш\u00a0дацx·'mn" +
		"ne×10^0/00INF−\u200e−ناعددepäluku’არ\u00a0არის\u00a0რიცხვი¤¤¤?сан\u00a0э" +
		"месບໍ່\u200bແມ່ນ\u200bໂຕ\u200bເລກnav\u00a0skaitlisဂဏန်းမဟုတ်သောННне" +
		"\u00a0числоepilohosan\u00a0dälTFЕhaqiqiy\u00a0son\u00a0emas非數值٫٬؛\u200f+" +
		"\u200f-اس؉ليس\u00a0رقم\u200f−\u200e+\u200e\u200e-\u200e×۱۰^قیہ\u00a0عدد" +
		"\u00a0نہیںসংখ্যা\u00a0নাസംഖ്യയല്ല၊ཨང་མེན་དང་གྲངས་མེདཨང་མད",
	Index: []uint16{ // 69 elements
		// Entry 0 - 3F
		0x0000, 0x0001, 0x0002, 0x0003, 0x0004, 0x0005, 0x0006, 0x0007,
		0x0009, 0x000c, 0x000f, 0x0012, 0x0013, 0x0015, 0x0017, 0x001b,
		0x001f, 0x0031, 0x0033, 0x0049, 0x004a, 0x004c, 0x004d, 0x0050,
		0x0051, 0x0056, 0x005a, 0x005d, 0x0060, 0x0066, 0x0070, 0x0078,
		0x007b, 0x00a3, 0x00a9, 0x00aa, 0x00ba, 0x00e7, 0x00f4, 0x011b,
		0x011f, 0x012f, 0x0136, 0x013f, 0x0141, 0x0143, 0x0155, 0x015e,
		0x0160, 0x0162, 0x0164, 0x0168, 0x016c, 0x0170, 0x0172, 0x0180,
		0x0186, 0x018d, 0x0194, 0x019b, 0x019d, 0x01b3, 0x01cd, 0x01e8,
		// Entry 40 - 7F
		0x01eb, 0x0200, 0x0209, 0x0221, 0x0230,
	},
} // Size: 738 bytes

// langToDefaults maps a compact language index to the default numbering system
// and default symbol set
var langToDefaults = [742]uint8{
	// Entry 0 - 3F
	0x80, 0x05, 0x14, 0x01, 0x01, 0x01, 0x01, 0x01,
	0x00, 0x00, 0x00, 0x00, 0x83, 0x02, 0x02, 0x02,
	0x02, 0x03, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
	0x02, 0x02, 0x03, 0x03, 0x03, 0x03, 0x02, 0x02,
	0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x03,
	0x02, 0x85, 0x00, 0x00, 0x00, 0x86, 0x04, 0x05,
	0x05, 0x05, 0x05, 0x05, 0x01, 0x01, 0x06, 0x06,
	0x00, 0x00, 0x00, 0x00, 0x07, 0x07, 0x00, 0x00,
	// Entry 40 - 7F
	0x00, 0x89, 0x00, 0x00, 0x8b, 0x00, 0x00, 0x8d,
	0x01, 0x00, 0x00, 0x05, 0x05, 0x05, 0x05, 0x05,
	0x05, 0x05, 0x05, 0x05, 0x05, 0x08, 0x08, 0x00,
	0x00, 0x00, 0x00, 0x8f, 0x09, 0x09, 0x91, 0x01,
	0x01, 0x01, 0x93, 0x0a, 0x0b, 0x0b, 0x0b, 0x00,
	0x00, 0x0c, 0x0d, 0x0c, 0x0e, 0x0c, 0x0e, 0x0c,
	0x0f, 0x0f, 0x0c, 0x0c, 0x01, 0x01, 0x00, 0x01,
	0x01, 0x95, 0x00, 0x00, 0x00, 0x10, 0x10, 0x10,
	// Entry 80 - BF
	0x11, 0x11, 0x11, 0x00, 0x00, 0x05, 0x00, 0x00,
	0x00, 0x0c, 0x12, 0x00, 0x05, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00,
	0x00, 0x0c, 0x00, 0x0b, 0x00, 0x00, 0x06, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	// Entry C0 - FF
	0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x13, 0x00, 0x00, 0x11, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x15, 0x15, 0x05, 0x00,
	0x00, 0x05, 0x05, 0x05, 0x05, 0x01, 0x00, 0x00,
	0x05, 0x05, 0x05, 0x05, 0x00, 0x00, 0x05, 0x00,
	// Entry 100 - 13F
	0x00, 0x00, 0x00, 0x05, 0x00, 0x05, 0x00, 0x00,
	0x05, 0x05, 0x16, 0x16, 0x05, 0x05, 0x01, 0x01,
	0x97, 0x17, 0x17, 0x01, 0x01, 0x01, 0x01, 0x01,
	0x18, 0x18, 0x00, 0x00, 0x19, 0x19, 0x19, 0x9a,
	0x05, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
	0x01, 0x0f, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
	0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x05, 0x05,
	0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
	// Entry 140 - 17F
	0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
	0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x05, 0x05,
	0x05, 0x05, 0x00, 0x00, 0x9d, 0x00, 0x05, 0x05,
	0x1a, 0x1a, 0x1a, 0x1a, 0xa0, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x1b, 0x1b, 0x00, 0x00, 0x05, 0x05, 0x05,
	0x0c, 0x0c, 0x01, 0x01, 0x05, 0x05, 0x0b, 0x0b,
	0x00, 0x00, 0x00, 0x00, 0x05, 0x05, 0x05, 0x1c,
	// Entry 180 - 1BF
	0x05, 0x05, 0x00, 0x00, 0x00, 0x00, 0x05, 0x05,
	0x00, 0x00, 0x00, 0x1d, 0x1d, 0x01, 0x01, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x0f,
	0x0f, 0x00, 0x00, 0x01, 0x01, 0x05, 0x05, 0x1e,
	0x1e, 0x00, 0x00, 0x05, 0x05, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0xa2, 0x1f, 0x00, 0x00,
	0x01, 0x01, 0x20, 0x20, 0x00, 0x00, 0x00, 0x21,
	0x21, 0x00, 0x00, 0x05, 0x05, 0x00, 0x00, 0x00,
	// Entry 1C0 - 1FF
	0x00, 0x05, 0x05, 0x05, 0x05, 0x05, 0x22, 0x22,
	0xa4, 0x00, 0x00, 0x16, 0x16, 0x05, 0x05, 0x00,
	0x00, 0x00, 0x00, 0x23, 0x23, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x0f, 0x0f, 0x00, 0x00, 0x05, 0x05,
	0x00, 0x00, 0x05, 0x05, 0xa6, 0x00, 0x00, 0x00,
	0xa8, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00,
	0x05, 0x05, 0xa9, 0x24, 0xab, 0x00, 0x00, 0x00,
	0x00, 0xac, 0x25, 0x25, 0x00, 0x00, 0xaf, 0x00,
	// Entry 200 - 23F
	0x00, 0xb0, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05,
	0x05, 0x01, 0x01, 0x15, 0x15, 0x05, 0x05, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x26, 0x26, 0x26,
	0xb2, 0xb4, 0x1b, 0x00, 0x00, 0x00, 0x01, 0x01,
	0x01, 0x01, 0xb6, 0x27, 0x05, 0x01, 0x05, 0x01,
	0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x00, 0x05,
	0x00, 0x00, 0x1a, 0x1a, 0x05, 0x05, 0x05, 0x05,
	// Entry 240 - 27F
	0x05, 0x00, 0x00, 0x28, 0x28, 0x28, 0x28, 0x28,
	0x28, 0x28, 0x05, 0x05, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x29, 0x29, 0x29,
	0x29, 0x05, 0x05, 0x0f, 0x0f, 0x05, 0x05, 0x01,
	0x01, 0x01, 0x01, 0x01, 0x2a, 0x2a, 0x01, 0x01,
	0x11, 0x11, 0x00, 0x00, 0x00, 0x2b, 0x2b, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
	0x01, 0x01, 0x01, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b,
	// Entry 280 - 2BF
	0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x00, 0x00,
	0x00, 0xb8, 0x20, 0x20, 0x20, 0x00, 0x05, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x2c, 0x2c, 0x00, 0x2d, 0x2d,
	0x05, 0x05, 0x05, 0x00, 0x0f, 0x0f, 0x01, 0x01,
	0x00, 0x00, 0x2e, 0x2e, 0xbb, 0xbd, 0x1b, 0xbe,
	0xc0, 0x27, 0xc2, 0x01, 0x2f, 0x2f, 0x00, 0x00,
	// Entry 2C0 - 2FF
	0x00, 0x00, 0x00, 0x00, 0x05, 0x05, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x30, 0x30, 0x00, 0x00, 0x00,
	0x00, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x31, 0x31, 0x31, 0x31, 0x00, 0x00,
} // Size: 742 bytes

// langToAlt is a list of numbering system and symbol set pairs, sorted and
// marked by compact language index.
var langToAlt = []altSymData{ // 68 elements
	1:  {compactTag: 0x0, system: 0x2, symIndex: 0x32},
	2:  {compactTag: 0x0, system: 0x3, symIndex: 0x36},
	3:  {compactTag: 0xc, system: 0x2, symIndex: 0x33},
	4:  {compactTag: 0xc, system: 0x0, symIndex: 0x2},
	5:  {compactTag: 0x29, system: 0x5, symIndex: 0x0},
	6:  {compactTag: 0x2d, system: 0x0, symIndex: 0x4},
	7:  {compactTag: 0x2d, system: 0x2, symIndex: 0x32},
	8:  {compactTag: 0x2d, system: 0x3, symIndex: 0x36},
	9:  {compactTag: 0x41, system: 0x5, symIndex: 0x3f},
	10: {compactTag: 0x41, system: 0x0, symIndex: 0x0},
	11: {compactTag: 0x44, system: 0x0, symIndex: 0x0},
	12: {compactTag: 0x44, system: 0x33, symIndex: 0x42},
	13: {compactTag: 0x47, system: 0x0, symIndex: 0x1},
	14: {compactTag: 0x47, system: 0x2, symIndex: 0x32},
	15: {compactTag: 0x5b, system: 0x2, symIndex: 0x32},
	16: {compactTag: 0x5b, system: 0x0, symIndex: 0x9},
	17: {compactTag: 0x5e, system: 0x0, symIndex: 0x1},
	18: {compactTag: 0x5e, system: 0x2, symIndex: 0x32},
	19: {compactTag: 0x62, system: 0x0, symIndex: 0xa},
	20: {compactTag: 0x62, system: 0x2, symIndex: 0x32},
	21: {compactTag: 0x79, system: 0x33, symIndex: 0x43},
	22: {compactTag: 0x79, system: 0x0, symIndex: 0x0},
	23: {compactTag: 0x110, system: 0x3, symIndex: 0x37},
	24: {compactTag: 0x110, system: 0x0, symIndex: 0x17},
	25: {compactTag: 0x110, system: 0x2, symIndex: 0x32},
	26: {compactTag: 0x11f, system: 0x0, symIndex: 0x1},
	27: {compactTag: 0x11f, system: 0x2, symIndex: 0x34},
	28: {compactTag: 0x11f, system: 0x3, symIndex: 0x38},
	29: {compactTag: 0x154, system: 0x0, symIndex: 0x0},
	30: {compactTag: 0x154, system: 0x2, symIndex: 0x32},
	31: {compactTag: 0x154, system: 0x3, symIndex: 0x36},
	32: {compactTag: 0x15c, system: 0x0, symIndex: 0x0},
	33: {compactTag: 0x15c, system: 0x2, symIndex: 0x32},
	34: {compactTag: 0x1ac, system: 0x3, symIndex: 0x36},
	35: {compactTag: 0x1ac, system: 0x0, symIndex: 0x1f},
	36: {compactTag: 0x1c8, system: 0x3, symIndex: 0x36},
	37: {compactTag: 0x1c8, system: 0x0, symIndex: 0x0},
	38: {compactTag: 0x1e4, system: 0x0, symIndex: 0x0},
	39: {compactTag: 0x1e4, system: 0x1c, symIndex: 0x40},
	40: {compactTag: 0x1e8, system: 0x9, symIndex: 0x0},
	41: {compactTag: 0x1f2, system: 0x21, symIndex: 0x41},
	42: {compactTag: 0x1f2, system: 0x0, symIndex: 0x24},
	43: {compactTag: 0x1f4, system: 0x3, symIndex: 0x36},
	44: {compactTag: 0x1f9, system: 0x0, symIndex: 0x25},
	45: {compactTag: 0x1f9, system: 0x2, symIndex: 0x35},
	46: {compactTag: 0x1f9, system: 0x3, symIndex: 0x39},
	47: {compactTag: 0x1fe, system: 0x9, symIndex: 0x0},
	48: {compactTag: 0x201, system: 0x0, symIndex: 0x5},
	49: {compactTag: 0x201, system: 0x2, symIndex: 0x32},
	50: {compactTag: 0x220, system: 0x0, symIndex: 0x0},
	51: {compactTag: 0x220, system: 0x3, symIndex: 0x3a},
	52: {compactTag: 0x221, system: 0x3, symIndex: 0x36},
	53: {compactTag: 0x221, system: 0x0, symIndex: 0x1b},
	54: {compactTag: 0x22a, system: 0x3, symIndex: 0x36},
	55: {compactTag: 0x22a, system: 0x0, symIndex: 0x27},
	56: {compactTag: 0x289, system: 0x0, symIndex: 0x20},
	57: {compactTag: 0x289, system: 0x2, symIndex: 0x34},
	58: {compactTag: 0x289, system: 0x3, symIndex: 0x3b},
	59: {compactTag: 0x2b4, system: 0x0, symIndex: 0x1b},
	60: {compactTag: 0x2b4, system: 0x3, symIndex: 0x3c},
	61: {compactTag: 0x2b5, system: 0x3, symIndex: 0x3c},
	62: {compactTag: 0x2b7, system: 0x0, symIndex: 0x2f},
	63: {compactTag: 0x2b7, system: 0x3, symIndex: 0x3d},
	64: {compactTag: 0x2b8, system: 0x3, symIndex: 0x36},
	65: {compactTag: 0x2b8, system: 0x0, symIndex: 0x27},
	66: {compactTag: 0x2ba, system: 0x0, symIndex: 0x1},
	67: {compactTag: 0x2ba, system: 0x3, symIndex: 0x3e},
} // Size: 296 bytes

// Total table size 4592 bytes (4KiB); checksum: FCCA88FA
