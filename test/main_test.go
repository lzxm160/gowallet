package test

import (
	"fmt"
	"math/big"
	"testing"
)

func TestXx(t *testing.T) {
	amount, ok := new(big.Float).SetString("10000000000000000000000000")
	if !ok {
		fmt.Println("!ok")
	}
	weightedAmount, acc := amount.Mul(amount, big.NewFloat(1.0000000022006)).Int(nil)
	fmt.Println(weightedAmount.String(), acc)
	//xx := make(map[string]*big.Int)
	//xx["xx"] = big.NewInt(10)
	//if val, ok := xx["xx"]; ok {
	//	val.Add(val, big.NewInt(10))
	//}
	//fmt.Println(xx["xx"].String())
	//val, _ := new(big.Int).SetString("13737312354955498326658513", 10)
	//votingPower := new(big.Float).SetInt(val)
	//val, _ = votingPower.Mul(votingPower, big.NewFloat(1.06)).Int(nil)
	//fmt.Println(val.Text(10))
	//name := "robotbp00021"
	//fmt.Println(hex.EncodeToString([]byte(name)))
	//
	//a, _ := big.NewInt(0).SetString("11447588583684614452276736", 10)
	//b, _ := big.NewInt(0).SetString("2289494821788705003536384", 10)
	//c, _ := big.NewInt(0).SetString("228949482178870500352", 10)
	////c, _ := big.NewInt(0).SetString("0", 10)
	//d := a.Add(a, b).Add(a, c)
	//fmt.Println(d.String())
	//m := make(map[string]bool)
	//m["xx"] = true
	//fmt.Println(m["xx"])
	//fmt.Println(m["yy"])
	//tt := &timestamp.Timestamp{Seconds: time.Now().Unix(), Nanos: 123}
	//bucketsBytes, _ := proto.Marshal(tt)
	//xx := &timestamp.Timestamp{}
	//if err := proto.Unmarshal(bucketsBytes, xx); err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(xx.Seconds, xx.Nanos)
	//tt := time.Unix(1362984425, 0)
	//nt := tt.Format("2006-01-02 15:04:05")
	//fmt.Println(nt)
	//fmt.Println(tt.String())
	//input := []byte("100")
	//encodeString := base64.StdEncoding.EncodeToString(input)
	//fmt.Println(encodeString)
	//data, err := base64.StdEncoding.DecodeString("NzYxNDYwMA==")
	//if err != nil {
	//	fmt.Println("error:", err)
	//	return
	//}
	//fmt.Printf("%q\n", data)
}

//type Vertex struct {
//	X, Y float64
//}
//
////传递了值接收器Vertex
//func (v Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
////传递了指针接收器*Vertex
//func (v *Vertex) Scale(f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}
//type People struct {
//	name string
//}
//
//func (p *People) set(p1 *People) {
//
//	p = p1
//
//	fmt.Println("set func :", p.name)
//
//}
//func TestXXX(t *testing.T) {
//	p := &People{"gjf"}
//
//	p1 := &People{"hyc"}
//
//	p.set(p1)
//
//	fmt.Println(p.name)
//	//v := Vertex{1, 2}
//	//v.Scale(10)
//	//fmt.Println(v.Abs())
//	//vp := &v
//	//fmt.Println(vp.Abs())
//}
//
////
////func Compress(data []byte) ([]byte, error) {
////	var bb bytes.Buffer
////	w, err := gzip.NewWriterLevel(&bb, gzip.BestCompression)
////	if err != nil {
////		return nil, err
////	}
////	_, err = w.Write(data)
////	if err != nil {
////		w.Close()
////		return nil, err
////	}
////	w.Close()
////	output := bb.Bytes()
////	return output, nil
////}
////
////// Decompress uses gzip to uncompress the input bytes
////func Decompress(data []byte) ([]byte, error) {
////	r, err := gzip.NewReader(bytes.NewBuffer(data))
////	if err != nil {
////		return nil, err
////	}
////	r.Close()
////	return ioutil.ReadAll(r)
////}
////func TestXXXXX(t *testing.T) {
////	data := []byte("11111111111111111111111111111111111111110000000000000000000000000000000000000000")
////	cd, err := Compress(data)
////	require.NoError(t, err)
////	fmt.Println(string(cd))
////	fmt.Println(len(cd))
////	dcd, err := Decompress(cd)
////	require.NoError(t, err)
////	fmt.Println(string(dcd))
////	fmt.Println(len(dcd))
////
////	//c := make(chan interface{})
////	//go func() {
////	//	for {
////	//		select {
////	//		case <-c:
////	//			fmt.Println("xx")
////	//			return
////	//		default:
////	//			fmt.Println("default")
////	//		}
////	//	}
////	//}()
////	//close(c)
////	//select {}
////}
//
////
////func testDijkstra() {
////	var name = make(map[string]int)
////	name["s"] = 0
////	name["a"] = 1
////	name["b"] = 2
////	name["c"] = 3
////	name["d"] = 4
////	name["e"] = 5
////	name["f"] = 6
////	name["g"] = 7
////	name["h"] = 8
////	var index = make(map[int]string)
////	index[0] = "s"
////	index[1] = "a"
////	index[2] = "b"
////	index[3] = "c"
////	index[4] = "d"
////	index[5] = "e"
////	index[6] = "f"
////	index[7] = "g"
////	index[8] = "h"
////
////	var weight [9][9]float64
////
////	for i := 0; i < 9; i++ {
////		for j := 0; j < 9; j++ {
////			if i == j {
////				weight[i][j] = 0
////			} else {
////				weight[i][j] = 1000
////			}
////		}
////	}
////	weight[name["s"]][name["a"]] = 0.5
////	weight[name["s"]][name["b"]] = 0.3
////	weight[name["s"]][name["c"]] = 0.2
////	weight[name["s"]][name["d"]] = 0.4
////
////	weight[name["a"]][name["e"]] = 0.3
////
////	weight[name["b"]][name["a"]] = 0.2
////	weight[name["b"]][name["f"]] = 0.1
////
////	weight[name["c"]][name["f"]] = 0.4
////	weight[name["c"]][name["h"]] = 0.8
////
////	weight[name["d"]][name["c"]] = 0.1
////	weight[name["d"]][name["h"]] = 0.6
////
////	weight[name["e"]][name["g"]] = 0.1
////
////	weight[name["f"]][name["e"]] = 0.1
////	weight[name["f"]][name["h"]] = 0.2
////
////	weight[name["h"]][name["g"]] = 0.4
////	//settled := make(map[int]int)
////	settledWei := make(map[int]float64)
////	mw := make(map[int]float64)
////	for i := 0; i < 9; i++ {
////		mw[i] = weight[name["s"]][i]
////	}
////	dijkstra(weight, mw, settledWei)
////	//for k, v := range settled {
////	//	fmt.Println(k, ":", v)
////	//}
////	for k, v := range settledWei {
////		fmt.Printf("%s:%0.2f ", index[k], v)
////	}
////	fmt.Println()
////}
////func dijkstra(weight [9][9]float64, mw map[int]float64, settledWei map[int]float64) {
////	//if len(settledWei) >= 9 {
////	//	return
////	//}
////	if len(mw) == 0 {
////		return
////	}
////	minDis := 1000.0
////	node := 0
////	// 找出距离最小的节点
////	for k, v := range mw {
////		if v < minDis {
////			minDis = v
////			node = k
////		}
////	}
////	// update min weight
////	for k, v := range mw {
////		if v > mw[node]+weight[node][k] {
////			mw[k] = mw[node] + weight[node][k]
////		}
////	}
////	settledWei[node] = minDis
////	delete(mw, node)
////
////	dijkstra(weight, mw, settledWei)
////
////}
////
////type TokenType string
////type Token interface {
////	Type() TokenType
////	Lexeme() string
////}
////
////type Match struct {
////	toktype TokenType
////	lexeme  string
////}
////
////type IntegerConstant struct {
////	Match
////	value uint64
////}
////
////func (m *Match) Type() TokenType {
////	return m.toktype
////}
////
////func (m *Match) Lexeme() string {
////	return m.lexeme
////}
////
////func (i *IntegerConstant) Value() uint64 {
////	return i.value
////}
////
////type Pet struct {
////	name string
////}
////
////type Dog struct {
////	Pet
////	Breed string
////}
////
////func (p *Pet) Speak() string {
////	return fmt.Sprintf("my name is %v", p.name)
////}
////
////func (p *Pet) Name() string {
////	return p.name
////}
////
////func (d *Dog) Speak2() string {
////	return fmt.Sprintf("%v and I am a %v", d.Speak(), d.Breed)
////}
////func testudp() {
////	go func() {
////		time.Sleep(time.Second * 2)
////		sip := net.ParseIP("127.0.0.1")
////		srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
////		dstAddr := &net.UDPAddr{IP: sip, Port: 9981}
////		conn, err := net.DialUDP("udp", srcAddr, dstAddr)
////		if err != nil {
////			fmt.Println(err)
////		}
////		defer conn.Close()
////		conn.Write([]byte("hello"))
////		fmt.Printf("<%s>\n", conn.RemoteAddr())
////
////		data := make([]byte, 1024)
////		n, err := conn.Read(data)
////		fmt.Printf("read %s from <%s>\n", data[:n], conn.RemoteAddr())
////	}()
////	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9981})
////	if err != nil {
////		fmt.Println(err)
////		return
////	}
////	fmt.Printf("Local: <%s> \n", listener.LocalAddr().String())
////	//data := make([]byte, 1024)
////	//remoteAddr := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 19982}
////	for {
////		//n, err := listener.Read(data)
////		//if err != nil {
////		//	fmt.Printf("error during read: %s", err)
////		//}
////		//fmt.Printf(" %s\n", data[:n])
////		//n, remoteAddr, err := listener.ReadFromUDP(data)
////		//if err != nil {
////		//	fmt.Printf("error during read: %s", err)
////		//}
////		//
////		//fmt.Printf("<%s> %s\n", remoteAddr, data[:n])
////		//_, err = listener.WriteToUDP([]byte("world"), remoteAddr)
////		//if err != nil {
////		//	fmt.Printf(err.Error())
////		//}
////		_, err = listener.Write([]byte("hello"))
////		if err != nil {
////			fmt.Println("server write:", err.Error())
////			break
////		}
////	}
////}
////func testmulti() {
////	go func() {
////		time.Sleep(time.Second * 3)
////		ip := net.ParseIP("224.0.0.250")
////		srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
////		dstAddr := &net.UDPAddr{IP: ip, Port: 9981}
////		conn, err := net.DialUDP("udp", srcAddr, dstAddr)
////		if err != nil {
////			fmt.Println(err)
////		}
////		defer conn.Close()
////		conn.Write([]byte("hello"))
////		fmt.Printf("<%s>\n", conn.RemoteAddr())
////	}()
////	//1. 得到一个interface
////	en4, err := net.InterfaceByName("eth0")
////	if err != nil {
////		fmt.Println(err)
////	}
////	group := net.IPv4(224, 0, 0, 250)
////	//2. bind一个本地地址
////	c, err := net.ListenPacket("udp4", "0.0.0.0:1024")
////	if err != nil {
////		fmt.Println(err)
////	}
////	defer c.Close()
////	//3.
////	p := ipv4.NewPacketConn(c)
////	if err := p.JoinGroup(en4, &net.UDPAddr{IP: group}); err != nil {
////		fmt.Println(err)
////	}
////	//4.更多的控制
////	if err := p.SetControlMessage(ipv4.FlagDst, true); err != nil {
////		fmt.Println(err)
////	}
////	//5.接收消息
////	b := make([]byte, 1500)
////	for {
////		n, cm, src, err := p.ReadFrom(b)
////		if err != nil {
////			fmt.Println(err)
////		}
////		if cm.Dst.IsMulticast() {
////			if cm.Dst.Equal(group) {
////				fmt.Printf("received: %s from <%s>\n", b[:n], src)
////				n, err = p.WriteTo([]byte("world"), cm, src)
////				if err != nil {
////					fmt.Println(err)
////				}
////			} else {
////				fmt.Println("Unknown group")
////				continue
////			}
////		}
////	}
////}
////func testopenfile() {
////	d := Dog{Pet: Pet{name: "spot"}, Breed: "pointer"}
////	fmt.Println(d.Name())
////	fmt.Println(d.Speak2())
////	fmt.Println(d.name)
////	//t := IntegerConstant{Match{"xx", "wizard"}, 2}
////	//fmt.Println(t.Type(), t.Lexeme(), t.Value())
////	//x := Token(t)
////	//fmt.Println(x.Type(), x.Lexeme())
////
////	//db, err := bolt.Open("my.db", 0600, nil)
////	//if err != nil {
////	//	fmt.Println(err)
////	//}
////	//fmt.Println(db)
////	//opt := bolt.Options{
////	//	ReadOnly: true,
////	//}
////	//db, err = bolt.Open("my.db", 0600, &opt)
////	//if err != nil {
////	//	fmt.Println(err)
////	//}
////	//fmt.Println(db)
////
////	//xx, err := os.OpenFile("./xx", os.O_RDWR|os.O_CREATE, 0666)
////	//fmt.Println(xx, ":", err)
////	//
////	////flag := syscall.LOCK_SH
////	//flag := syscall.LOCK_EX
////	//
////	//// Otherwise attempt to obtain an exclusive lock.
////	//err = syscall.Flock(int(xx.Fd()), flag|syscall.LOCK_NB)
////	//fmt.Println(err)
////	//
////	//xx2, err := os.OpenFile("./xx", os.O_RDWR, 0666)
////	//fmt.Println(xx2, ":", err)
////}
////func testbroard() {
////	go func() {
////		time.Sleep(time.Second * 4)
////		ip := net.ParseIP("172.17.255.255")
////		srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
////		dstAddr := &net.UDPAddr{IP: ip, Port: 9981}
////		conn, err := net.ListenUDP("udp", srcAddr)
////		if err != nil {
////			fmt.Println(err)
////		}
////		n, err := conn.WriteToUDP([]byte("hello"), dstAddr)
////		if err != nil {
////			fmt.Println(err)
////		}
////		data := make([]byte, 1024)
////		n, _, err = conn.ReadFrom(data)
////		if err != nil {
////			fmt.Println(err)
////		}
////		fmt.Printf("client read %s from <%v>\n", data[:n], conn.RemoteAddr())
////	}()
////	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 9981})
////	if err != nil {
////		fmt.Println(err)
////		return
////	}
////	fmt.Printf("Local: <%s> \n", listener.LocalAddr().String())
////	data := make([]byte, 1024)
////	for {
////		n, remoteAddr, err := listener.ReadFromUDP(data)
////		if err != nil {
////			fmt.Println("error during read: ", err)
////		}
////		fmt.Printf("server read: <%s> %s\n", remoteAddr, data[:n])
////		_, err = listener.WriteToUDP([]byte("world"), remoteAddr)
////		if err != nil {
////			fmt.Println(err.Error())
////		}
////	}
////}
////func testxx() {
////	//data := "0000000000000000000000006356908ace09268130dee2b7de643314bbeb36830000000000000000000000000000000000000000000000000000000000000004"
////	//hb, _ := hex.DecodeString(data)
////	//out := crypto.Keccak256(hb)
////	//fmt.Println(hex.EncodeToString(out))
////	//encodeString1 := base64.StdEncoding.EncodeToString(out)
////	//fmt.Println(encodeString1)
////	//
////	//data := "000000000000000000000000cf67d52eaf78b0fd89056f02b862d9ac43538c230000000000000000000000000000000000000000000000000000000000000000"
////	//hb, _ := hex.DecodeString(data)
////	//out2 := crypto.Keccak256(hb)
////	//fmt.Println(hex.EncodeToString(out2))
////	//encodeString2 := base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(out2)))
////	//fmt.Println(encodeString2)
////
////	input := []byte("18160ddd")
////	//input, _ := hex.DecodeString("0000000000000000000000006356908ace09268130dee2b7de643314bbeb36830000000000000000000000000000000000000000000000000000000000000004")
////	// 演示base64编码
////	encodeString := base64.StdEncoding.EncodeToString(input)
////	fmt.Println(encodeString)
////
////	decodeBytes, err := base64.StdEncoding.DecodeString("MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDA2NzY1Yzc5M2ZhMTAwNzlkMDAwMDAwMA==")
////	if err != nil {
////		fmt.Println(err)
////	}
////	fmt.Println(string(decodeBytes))
////	x, _ := big.NewInt(0).SetString(string(decodeBytes), 16)
////	fmt.Println(x.Text(10))
////	//
////	//h, _ := hex.DecodeString("000000000000000000000000000000000000000006765c7915fedc5ed9d40000")
////	//
////	//y := big.NewInt(0).SetBytes(h)
////	//fmt.Println(y.Text(10))
////	//
////	//CurrIndex := []byte{255, 255, 255, 255, 255, 255, 255, 255}
////	//maxuint64 := big.NewInt(0).SetBytes(CurrIndex)
////	//fmt.Println(maxuint64.Text(10))
////	//fmt.Println(^uint64(0))
////	////fmt.Println(uint64(0) - 1)
////	//for i := uint64(0); i >= 0; i-- {
////	//	fmt.Println(i)
////	//	if i != 0 {
////	//		fmt.Println(i)
////	//		break
////	//	}
////	//}
////	//for i := uint8(0); i < 0; i++ {
////	//	fmt.Println("wahttttat")
////	//}
////}
////func testttt() bool {
////	fmt.Println("testtttttt")
////	return true
////}
////func testdefer() int {
////	//defer fmt.Println(testttt())
////	//fmt.Println("testdefer")
////	xx := 1
////	defer func() { xx++ }()
////	return xx
////}
////func traceMemStats() {
////	var ms runtime.MemStats
////	runtime.ReadMemStats(&ms)
////	log.Printf("Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
////}
////func testmem() {
////	xx := make([]int, 8)
////	log.Println("start://///////////////")
////	traceMemStats()
////	fmt.Println(xx)
////	xx = nil
////	xx = make([]int, 8)
////	runtime.GC()
////	traceMemStats()
////	log.Println("end://///////////////")
////}
////
//func BenchmarkHello(b *testing.B) {
//	sk, err := crypto.GenerateKey()
//	if err != nil {
//		panic(err)
//	}
//	ecdsaSK, err := crypto.ToECDSA(sk.D.Bytes())
//	if err != nil {
//		panic(err)
//	}
//	pk := crypto.FromECDSAPub(&ecdsaSK.PublicKey)
//	hash := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
//	signed, err := crypto.Sign(hash, sk)
//	if err != nil {
//		panic(err)
//	}
//	//b.ResetTimer()
//	//for i := 0; i < b.N; i++ {
//	//	_, err = crypto.Sign(hash, sk)
//	//	if err != nil {
//	//		panic(err)
//	//	}
//	//}
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		for j := 0; j < 4; j++ {
//			go func() {
//				ret := crypto.VerifySignature(pk, hash, signed[:len(signed)-1])
//				if !ret {
//					panic(err)
//				}
//			}()
//		}
//
//	}
//}
//func TestXx(t *testing.T) {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*10))
//	defer cancel()
//
//	go func(ctx context.Context) {
//		//ctx.Done()
//
//		//time.Sleep(time.Minute)
//		cancel()
//	}(ctx)
//
//	select {
//	case <-ctx.Done():
//		fmt.Println("call successfully!!!")
//		return
//	case <-time.After(time.Duration(time.Millisecond * 900)):
//		fmt.Println("timeout!!!")
//		return
//	}
//	select {}
//	//ts := func(s *string, v *string) {
//	//	fmt.Println(*s, ":", *v)
//	//}
//	//x := []*string{}
//	//y := []*string{}
//	//m := map[string]string{"1": "1", "2": "2", "3": "3"}
//	//for k, v := range m {
//	//	// Make a copy of k because it will be reassigned with each loop.
//	//	//fmt.Println(k, ":", v)
//	//	//ts(&k, &v)
//	//	x = append(x, &k)
//	//	y = append(y, &v)
//	//}
//	//for i := 0; i < 3; i++ {
//	//	ts(x[i], y[i])
//	//}
//	//t1 := time.Now()
//	//t2 := time.Now().UTC()
//	//fmt.Println(t1)
//	//fmt.Println(t2)
//	//total := big.NewInt(0)
//	//fmt.Println(total.Add(total, big.NewInt(1)).Add(total, big.NewInt(2)).Add(total, big.NewInt(3)))
//	//if true {
//	//	defer fmt.Println("in")
//	//}
//	//
//	//fmt.Println("out")
//	//s := "sm2sk-io1tx4nf330r2ewdp8u4vt9gk5rmdnkxeajcl828m.pem"
//	//a := strings.TrimRight(strings.TrimLeft(s, "sm2sk-"), ".pem")
//	//fmt.Println(a)
//	//fmt.Println(strings.TrimLeft(s, "sm2sk-"))
//	//fmt.Println(strings.TrimRight("io1tx4nf330r2ewdp8u4vt9gk5rmdnkxeajcl828m.pem", ".pem"))
//	//fmt.Println(strings.TrimRight("ABbbbbbbbbbBA", "BA"))
//	//
//	//fmt.Println(strings.TrimSuffix(strings.TrimPrefix(s, "sm2sk-"), ".pem"))
//
//	//x := "xx"
//	//y := "xxx"
//	//fmt.Println(x == y)
//	//f, err := os.OpenFile("xx", os.O_CREATE, 777)
//	//fmt.Println("err:", err)
//	//f.Close()
//	//
//	//err = os.Remove("xx")
//	//fmt.Println("remove err:", err)
//	//xx := new(big.Int).SetBytes([]byte{})
//	//fmt.Println(xx)
//	//xx = new(big.Int).SetBytes([]byte{0})
//	//fmt.Println(xx)
//	//fmt.Println(xx.Bytes())
//	//ready := make(chan interface{})
//	//close(ready)
//	//<-ready
//	//fmt.Println("xxxxxxxxxxxxxxxxxxxxxxx")
//	//for {
//	//	//fmt.Println(time.Now().Unix())
//	//	if time.Now().Unix()%3600 == 0 {
//	//		fmt.Println("xxxx")
//	//	}
//	//}
//
//	//fmt.Println(testdefer())
//	//testmem()
//	//testyy()
//	//testDijkstra()
//	//testopenfile()
//	//testudp()
//	//testbroard()
//	//testxx()
//	//te := []byte{1, 2, 3}
//	//fmt.Println(fmt.Sprintf("%x", te))
//	//fmt.Println(hex.EncodeToString(te))
//	//tag := []byte("123")
//	//key := []byte("456")
//	//k := make([]byte, len(tag)+len(key))
//	//copy(k, tag)
//	//k = append(k, key...)
//	//fmt.Println(string(k))
//	//ctx := context.Background()
//	//d := time.Duration(10) * time.Second
//	//tt := time.NewTicker(d)
//	//defer tt.Stop()
//	//for {
//	//	select {
//	//	case <-tt.C:
//	//		fmt.Println("Prune run ", time.Now().String())
//	//		time.Sleep(15 * time.Second)
//	//	case <-ctx.Done():
//	//		fmt.Println("Prune exit")
//	//		return
//	//	}
//	//}
//	//for i := uint64(10); i >= 1; i-- {
//	//	fmt.Println(i)
//	//}
//	//actDetailList := make([]int, 0, 10)
//	//for i := 0; i < 10; i++ {
//	//	actDetailList = append(actDetailList, i)
//	//}
//	//fmt.Println(actDetailList)
//	//testfmt()
//	//testopenfile()
//	//fmt.Println(tt)
//	//time.Sleep(time.Second * 3)
//}

//func testopenfile() {
//	//db, err := bolt.Open("my.db", 0600, nil)
//	//if err != nil {
//	//	fmt.Println(err)
//	//}
//	//fmt.Println(db)
//	db, err := bolt.Open("my.db", 0600, nil)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(db)
//	err = db.Update(func(tx *bolt.Tx) error {
//		b, err := tx.CreateBucket([]byte("MyBucket11111"))
//
//		if err != nil {
//			return fmt.Errorf("create bucket: %s", err)
//		}
//		b, err = tx.CreateBucket([]byte("MyBucket22222"))
//		if err != nil {
//			return fmt.Errorf("create bucket: %s", err)
//		}
//		bytes := make([]byte, 8)
//		binary.BigEndian.PutUint64(bytes, 1)
//		err = b.Put(bytes, []byte("1"))
//		if err != nil {
//			return err
//		}
//		binary.BigEndian.PutUint64(bytes, 10)
//		err = b.Put(bytes, []byte("10"))
//		if err != nil {
//			return err
//		}
//		binary.BigEndian.PutUint64(bytes, 100)
//		//err = b.Put([]byte{100}, []byte("zhangsan2"))
//		//if err != nil {
//		//	return err
//		//}
//		err = b.Put(bytes, []byte("100"))
//		return err
//	})
//	err = db.View(func(tx *bolt.Tx) error {
//		if err := tx.ForEach(func(name []byte, b *bolt.Bucket) error {
//			fmt.Println(string(name))
//			return nil
//		}); err != nil {
//			fmt.Println(err)
//		}
//		return nil
//	})
//
//	fmt.Println(err)
//
//}
//
//	err = db.View(func(tx *bolt.Tx) error {
//		c := tx.Bucket([]byte("MyBucket")).Cursor()
//		//for k, v := c.Last(); k != nil; k, v = c.Prev() {
//		//	fmt.Println(string(k), ":", string(v))
//		//}
//		//return nil
//		k, v := c.Seek([]byte{1}) //1
//		fmt.Println(k, ":", v)
//		k, v = c.Seek([]byte{2}) //10
//		fmt.Println(k, ":", v)
//		k, v = c.Seek([]byte{3}) //10
//		fmt.Println(k, ":", v)
//		k, v = c.Seek([]byte{10}) //10
//		fmt.Println(k, ":", v)
//		k, v = c.Seek([]byte{12}) //100
//		fmt.Println(k, ":", v)
//		k, v = c.Seek([]byte{50}) //100
//		fmt.Println(k, ":", v)
//		k, v = c.Seek([]byte{100}) //100
//		fmt.Println(k, ":", v)
//		k, v = c.Seek([]byte{200}) //kong
//		fmt.Println(k, ":", v)
//		return nil
//	})
//
//	//fmt.Println(err)
//	//prefix := []byte("one")
//	//needDelete := [][]byte{}
//	//err = db.Batch(func(tx *bolt.Tx) error {
//	//	c := tx.Bucket([]byte("MyBucket")).Cursor()
//	//	for k, _ := c.Seek(prefix); bytes.HasPrefix(k, prefix); k, _ = c.Next() {
//	//		fmt.Println(string(k))
//	//		if bytes.Compare(k[3:], []byte("124")) > 0 {
//	//			//db.Delete(k)
//	//			//temp := make([]byte, len(k))
//	//			//copy(temp, k)
//	//			//needDelete = append(needDelete, temp)
//	//
//	//		}
//	//
//	//	}
//	//	return nil
//	//})
//	////fmt.Println(err)
//	//err = db.Update(func(tx *bolt.Tx) error {
//	//	b := tx.Bucket([]byte("MyBucket"))
//	//	for _, v := range needDelete {
//	//		fmt.Println("delete:", string(v))
//	//		if err := b.Delete(v); err != nil {
//	//			return err
//	//		}
//	//	}
//	//	return nil
//	//})
//
//	//err = db.Update(func(tx *bolt.Tx) error {
//	//	b := tx.Bucket([]byte("MyBucket")).Cursor()
//	//	for k, _ := b.Seek(prefix); bytes.HasPrefix(k, prefix); k, _ = b.Next() {
//	//		if bytes.Compare(k[3:], []byte("124")) > 0 {
//	//			fmt.Println(string(k))
//	//			tx.Bucket([]byte("MyBucket")).Delete(k)
//	//		}
//	//	}
//	//	return nil
//	//})
//	//fmt.Println(err)
//	//err = db.View(func(tx *bolt.Tx) error {
//	//	c := tx.Bucket([]byte("MyBucket")).Cursor()
//	//	for k, v := c.Last(); k != nil; k, v = c.Prev() {
//	//		fmt.Println(string(k), ":", string(v))
//	//	}
//	//	return nil
//	//})
//	//fmt.Println(err)
//}

//var tt bool

//func calledbyyy() {
//	time.Sleep(time.Second * 1)
//	fmt.Println("calledbyyy start:", tt)
//	tt = true
//	fmt.Println("calledbyyy end:", tt)
//
//}
//
//var y []byte
//
//func testx(in []byte) {
//	y = in
//	fmt.Println(y)
//}
//func testyy() {
//	x := []byte{1, 2, 3}
//	fmt.Println(x)
//	testx(x)
//	fmt.Println(x)
//	x[0] = 20
//	fmt.Println(x)
//	fmt.Println(y)

//fmt.Println("start:", tt)
//tt = false
//if !tt {
//	defer func() {
//		tt = true
//		fmt.Println("defer called")
//	}()
//}
//go calledbyyy()
//
//fmt.Println("end:", tt)

//p := make(PairList, 5)
//p[0] = Pair{"10", 10}
//p[1] = Pair{"4", 4}
//p[2] = Pair{"2", 2}
//p[3] = Pair{"3", 3}
//p[4] = Pair{"8", 8}
//fmt.Println(p)
//sort.Sort(p)
//fmt.Println(p)
//var xx []int
//xx = append(xx, 1)
//xx = append(xx, 2)
//fmt.Println(xx)
//xx := make([]string, 0)
//xx = append(xx, "11")
//xx = append(xx, "22")
//fmt.Println(xx)
//xx = append(xx[:1], xx[2:]...)
//fmt.Println(xx)
//indexHeightKey := append([]byte{1}, 2)
//fmt.Println(indexHeightKey)
//indexHeightKey2 := append(indexHeightKey, 3)
//fmt.Println(indexHeightKey)
//fmt.Println(indexHeightKey2)
//}

//type Pair struct {
//	Key   string
//	Value int
//}
//
//// A slice of Pairs that implements sort.Interface to sort by Value.
//type PairList []Pair
//
//func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
//func (p PairList) Len() int           { return len(p) }
//func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
//
//// A function to turn a map into a PairList, then sort and return it.
//func sortMapByValue(m map[string]int) PairList {
//	p := make(PairList, len(m))
//	i := 0
//	for k, v := range m {
//		p[i] = Pair{k, v}
//	}
//	sort.Sort(p)
//	return p
//}
