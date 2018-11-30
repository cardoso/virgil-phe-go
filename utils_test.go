package phe

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/passw0rd/phe-go/swu"

	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	key := make([]byte, 32)
	rand.Read(key)

	data := make([]byte, 365)

	ciphertext, err := Encrypt(data, key)

	require.NoError(t, err)

	plaintext, err := Decrypt(ciphertext, key)
	require.NoError(t, err)

	require.Equal(t, plaintext, data)

}

func TestEncrypt_empty(t *testing.T) {
	key := make([]byte, 32)
	rand.Read(key)

	data := make([]byte, 0)

	ciphertext, err := Encrypt(data, key)

	require.NoError(t, err)

	plaintext, err := Decrypt(ciphertext, key)
	require.NoError(t, err)

	require.Equal(t, plaintext, data)

}

func TestEncrypt_badKey(t *testing.T) {
	key := make([]byte, 32)
	rand.Read(key)

	data := make([]byte, 365)

	ciphertext, err := Encrypt(data, key)

	require.NoError(t, err)

	key[0]++

	plaintext, err := Decrypt(ciphertext, key)
	require.Error(t, err)
	require.Nil(t, plaintext)
}

func TestDecrypt_badLength(t *testing.T) {
	ct := make([]byte, 32+15)
	key := make([]byte, 32)
	rand.Read(key)
	plaintext, err := Decrypt(ct, key)

	require.Error(t, err)
	require.Equal(t, err.Error(), "invalid ciphertext length")
	require.Nil(t, plaintext)
}

func TestHashZVector1(t *testing.T) {

	pub := []byte{
		0x04, 0x21, 0xc3, 0x71, 0x95, 0x74, 0xaf, 0xce,
		0xc6, 0x5e, 0x35, 0xbd, 0x77, 0x5a, 0x5b, 0xe3,
		0x6c, 0x77, 0xc0, 0xbe, 0x45, 0x01, 0xf5, 0xd7,
		0x0f, 0xf0, 0x70, 0xd5, 0x1a, 0x89, 0x3a, 0xd8,
		0xe0, 0x0c, 0xe6, 0xb8, 0x9b, 0x17, 0x88, 0xe6,
		0xc1, 0x27, 0xa0, 0xe1, 0x25, 0xd9, 0xde, 0x6a,
		0x71, 0x16, 0x46, 0xa0, 0x38, 0x0f, 0xc4, 0xe9,
		0x5a, 0x74, 0xe5, 0x2c, 0x89, 0xf1, 0x12, 0x2a,
		0x7c,
	}

	c0X := "97803661066250274657510595696566855164534492744724548093309723513248461995097"
	c0Y := "32563640650805051226489658838020042684659728733816530715089727234214066735908"
	c1X := "83901588226167680046300869772314554609808129217097458603677198943293551162597"
	c1Y := "69578797673242144759724361924884259223786981560985539034793627438888366836078"
	t1X := "34051691470374495568913340263568595354597873005782528499014802063444122859583"
	t1Y := "55902370943165854960816059167184401667567213725158022607170263924097403943290"
	t2X := "101861885104337123215820986653465602199317278936192518417111183141791463240617"
	t2Y := "40785451420258280256125533532563267231769863378114083364571107590767796025737"
	t3X := "79689595215343344259388135277552904427007069090288122793121340067386243614518"
	t3Y := "63043970895569149637126206639504503565389755448934804609068720159153015056302"
	chlng := "44284775164038922154509064072457018313778507095510488730681838539467538456334"

	z := hashZ(proofOk, pub, curveG, Point2Bytes(c0X, c0Y), Point2Bytes(c1X, c1Y), Point2Bytes(t1X, t1Y), Point2Bytes(t2X, t2Y), Point2Bytes(t3X, t3Y))
	require.Equal(t, chlng, z.String())
}

func TestHashZVector2(t *testing.T) {

	pub := []byte{
		0x04, 0x39, 0x01, 0x9b, 0x9e, 0x2f, 0x1b, 0xae,
		0x60, 0x65, 0xcd, 0x9b, 0x85, 0x94, 0xfe, 0xa6,
		0xe3, 0x5a, 0x9a, 0xfd, 0xd3, 0x15, 0x96, 0xca,
		0xd8, 0xf8, 0xa4, 0xb1, 0xbd, 0xcd, 0x9b, 0x24,
		0x40, 0x5b, 0x8b, 0x13, 0x23, 0xf2, 0xdd, 0x6b,
		0x1b, 0x1d, 0x3f, 0x57, 0x5d, 0x00, 0xf4, 0xa8,
		0x5f, 0xb8, 0x67, 0x90, 0x69, 0x74, 0xea, 0x16,
		0x4b, 0x41, 0x9e, 0x93, 0x66, 0x47, 0xd8, 0xfb,
		0x7b,
	}

	c0X := "66305582120524875023859689648303664817335268054431490163250455437389177295478"
	c0Y := "19615011428787373705295950431517815162915845805720956004550495681707511034851"
	c1X := "11237049376971579382843942757546874380042467137583453135179008882019225463739"
	c1Y := "80961525191994723690800208523971748057046695876178833586656397502847317233228"
	t1X := "39244241269455735193598520026736537476566784866134072628798326598844377151651"
	t1Y := "10612278657611837393693400625940452527356993857624739575347941960949401758261"
	t2X := "108016526337105983792792579967716341976396349948643843073602635679441433077833"
	t2Y := "90379537067318020066230942533439624193620174277378193732900885672181004096656"
	t3X := "36913295823787819500630010367019659122715720420780370192192548665300728488299"
	t3Y := "36547572032269541322937508337036635249923361457001752921238955135105574250650"
	t4X := "49166285642990312777312778351013119878896537776050488997315166935690363463787"
	t4Y := "66983832439067043864623691503721372978034854603698954939248898067109763920732"
	chlng := "70372334159548608970477463377226881596996243878451645928327743798394814050835"

	z := hashZ(proofError, pub, curveG, Point2Bytes(c0X, c0Y), Point2Bytes(c1X, c1Y), Point2Bytes(t1X, t1Y), Point2Bytes(t2X, t2Y), Point2Bytes(t3X, t3Y), Point2Bytes(t4X, t4Y))
	require.Equal(t, chlng, z.String())
}

func Point2Bytes(xs, ys string) []byte {
	x, _ := new(big.Int).SetString(xs, 10)
	y, _ := new(big.Int).SetString(ys, 10)

	p := &Point{
		X: x,
		Y: y,
	}

	return p.Marshal()
}

func TestSimpleHashZ(t *testing.T) {
	require.Equal(t, "97888341710369812510024597077129852329763301580926521329107926771848618239575", hashZ(proofOk, curveG).String())
}

func TestData2Hash(t *testing.T) {

	data := []byte{
		0x02, 0x6c, 0x68, 0xba, 0x79, 0x9b, 0x95, 0x8d,
		0xa1, 0xdd, 0xec, 0x47, 0xcf, 0x77, 0xb6, 0x1a,
		0x68, 0xe3, 0x27, 0xbb, 0x16, 0xdd, 0x04, 0x6f,
		0x90, 0xfe, 0x2d, 0x7e, 0x46, 0xc7, 0x86, 0x1b,
		0xf9, 0x7a, 0xdb, 0xda, 0x15, 0xef, 0x5c, 0x13,
		0x63, 0xe7, 0x0d, 0x7c, 0xfa, 0x78, 0x24, 0xca,
		0xb9, 0x29, 0x74, 0x96, 0x09, 0x47, 0x15, 0x4d,
		0x34, 0xc4, 0x38, 0xe3, 0xeb, 0xcf, 0xfc, 0xbc,
	}

	x, y := swu.DataToPoint(data)

	require.Equal(t, "41644486759784367771047752285976210905566569374059610763941558650382638987514", x.String())
	require.Equal(t, "47123545766650584118634862924645280635136629360149764686957339607865971771956", y.String())
}

func TestHs0(t *testing.T) {
	ns1 := []byte{
		0x8e, 0x48, 0xac, 0x4b, 0x4a, 0x0c, 0x3f, 0x87,
		0x83, 0x69, 0x6f, 0x5d, 0x1f, 0x77, 0xd4, 0x25,
		0x64, 0x84, 0xd5, 0xb0, 0x7f, 0xd3, 0x8a, 0xf6,
		0xb2, 0xbf, 0x2d, 0x7b, 0x34, 0x57, 0x8a, 0x24,
	}

	p := hashToPoint(dhs0, ns1)

	require.Equal(t, "31738960577604984452512668768987751127290810426600335784941688353782437934958", p.X.String())
	require.Equal(t, "62583992082059176732972391810458341110572440440301779394415482355227454068077", p.Y.String())
}

func TestHs1(t *testing.T) {
	ns2 := []byte{
		0x04, 0x60, 0x41, 0x90, 0xea, 0xe3, 0x03, 0x48,
		0xc4, 0x67, 0xa2, 0x56, 0xaa, 0x20, 0xf0, 0xe1,
		0x22, 0xfd, 0x4c, 0x54, 0xb0, 0x2a, 0x03, 0x26,
		0x84, 0xf1, 0x22, 0x11, 0xfc, 0x9a, 0x8e, 0xe3,
	}

	p := hashToPoint(dhs1, ns2)

	require.Equal(t, "103792586023238657505718799913093621561253722738830297361241446168196965927733", p.X.String())
	require.Equal(t, "22377663630657926188215004914868715599610999655918661796946500276604457667995", p.Y.String())
}

func TestHc0(t *testing.T) {
	nc1 := []byte{
		0xdb, 0x59, 0x4e, 0x9a, 0x53, 0xeb, 0x35, 0x39,
		0x84, 0x63, 0x67, 0xf1, 0x4c, 0x15, 0xa1, 0x9b,
		0x4b, 0xee, 0x1d, 0x27, 0x13, 0xf3, 0xaa, 0xb5,
		0x3b, 0x11, 0x72, 0xd6, 0x02, 0x51, 0x63, 0x36,
	}

	pwd := []byte{
		0x5a, 0xf6, 0xf9, 0x9a, 0xc2, 0x0d, 0x0d, 0x54,
		0x52, 0xa2,
	}

	p := hashToPoint(dhc0, nc1, pwd)

	require.Equal(t, "67320834235162488735491952753309979921968147294363168750656052130501668262456", p.X.String())
	require.Equal(t, "6247579557587633046026752924567715390619882862197125195715158100861582898874", p.Y.String())
}

func TestHc1(t *testing.T) {
	nc1 := []byte{
		0x91, 0xd2, 0x04, 0x0b, 0x8e, 0x52, 0x7e, 0x8a,
		0xe3, 0x40, 0xf6, 0x89, 0xda, 0x01, 0x7c, 0xd6,
		0x1e, 0x20, 0x25, 0xd0, 0xbc, 0xc4, 0xd1, 0x24,
		0x92, 0x5c, 0x87, 0xc3, 0xe9, 0x59, 0xc7, 0x54,
	}

	pwd := []byte{
		0xb8, 0xce, 0xc3, 0xde, 0xfd, 0xfc, 0x80, 0x3c, 0x18,
		0x5d,
	}

	p := hashToPoint(dhc1, nc1, pwd)

	require.Equal(t, "33088581634824153508416124572141426244994685558673428197533469477988085629502", p.X.String())
	require.Equal(t, "30779922907513090428991463735414325156908518926812825488164363412500447467118", p.Y.String())
}
