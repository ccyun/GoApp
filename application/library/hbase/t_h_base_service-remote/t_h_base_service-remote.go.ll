// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"flag"
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"

	"github.com/ccyun/GoApp/application/library/hbase"

	//"git.apache.org/thrift.git/lib/go/thrift"
	//"hbase"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  bool exists(string table, TGet tget)")
	fmt.Fprintln(os.Stderr, "  TResult get(string table, TGet tget)")
	fmt.Fprintln(os.Stderr, "   getMultiple(string table,  tgets)")
	fmt.Fprintln(os.Stderr, "  void put(string table, TPut tput)")
	fmt.Fprintln(os.Stderr, "  bool checkAndPut(string table, string row, string family, string qualifier, string value, TPut tput)")
	fmt.Fprintln(os.Stderr, "  void putMultiple(string table,  tputs)")
	fmt.Fprintln(os.Stderr, "  void deleteSingle(string table, TDelete tdelete)")
	fmt.Fprintln(os.Stderr, "   deleteMultiple(string table,  tdeletes)")
	fmt.Fprintln(os.Stderr, "  bool checkAndDelete(string table, string row, string family, string qualifier, string value, TDelete tdelete)")
	fmt.Fprintln(os.Stderr, "  TResult increment(string table, TIncrement tincrement)")
	fmt.Fprintln(os.Stderr, "  TResult append(string table, TAppend tappend)")
	fmt.Fprintln(os.Stderr, "  i32 openScanner(string table, TScan tscan)")
	fmt.Fprintln(os.Stderr, "   getScannerRows(i32 scannerId, i32 numRows)")
	fmt.Fprintln(os.Stderr, "  void closeScanner(i32 scannerId)")
	fmt.Fprintln(os.Stderr, "  void mutateRow(string table, TRowMutations trowMutations)")
	fmt.Fprintln(os.Stderr, "   getScannerResults(string table, TScan tscan, i32 numRows)")
	fmt.Fprintln(os.Stderr, "  THRegionLocation getRegionLocation(string table, string row, bool reload)")
	fmt.Fprintln(os.Stderr, "   getAllRegionLocations(string table)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := hbase.NewTHBaseServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "exists":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "Exists requires 2 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		arg68 := flag.Arg(2)
		mbTrans69 := thrift.NewTMemoryBufferLen(len(arg68))
		defer mbTrans69.Close()
		_, err70 := mbTrans69.WriteString(arg68)
		if err70 != nil {
			Usage()
			return
		}
		factory71 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt72 := factory71.GetProtocol(mbTrans69)
		argvalue1 := hbase.NewTGet()
		err73 := argvalue1.Read(jsProt72)
		if err73 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.Exists(value0, value1))
		fmt.Print("\n")
		break
	case "get":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "Get requires 2 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		arg75 := flag.Arg(2)
		mbTrans76 := thrift.NewTMemoryBufferLen(len(arg75))
		defer mbTrans76.Close()
		_, err77 := mbTrans76.WriteString(arg75)
		if err77 != nil {
			Usage()
			return
		}
		factory78 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt79 := factory78.GetProtocol(mbTrans76)
		argvalue1 := hbase.NewTGet()
		err80 := argvalue1.Read(jsProt79)
		if err80 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.Get(value0, value1))
		fmt.Print("\n")
		break
	case "getMultiple":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "GetMultiple requires 2 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		arg82 := flag.Arg(2)
		mbTrans83 := thrift.NewTMemoryBufferLen(len(arg82))
		defer mbTrans83.Close()
		_, err84 := mbTrans83.WriteString(arg82)
		if err84 != nil {
			Usage()
			return
		}
		factory85 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt86 := factory85.GetProtocol(mbTrans83)
		containerStruct1 := hbase.NewTHBaseServiceGetMultipleArgs()
		err87 := containerStruct1.ReadField2(jsProt86)
		if err87 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.Tgets
		value1 := argvalue1
		fmt.Print(client.GetMultiple(value0, value1))
		fmt.Print("\n")
		break
	case "put":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "Put requires 2 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		arg89 := flag.Arg(2)
		mbTrans90 := thrift.NewTMemoryBufferLen(len(arg89))
		defer mbTrans90.Close()
		_, err91 := mbTrans90.WriteString(arg89)
		if err91 != nil {
			Usage()
			return
		}
		factory92 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt93 := factory92.GetProtocol(mbTrans90)
		argvalue1 := hbase.NewTPut()
		err94 := argvalue1.Read(jsProt93)
		if err94 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.Put(value0, value1))
		fmt.Print("\n")
		break
	case "checkAndPut":
		if flag.NArg()-1 != 6 {
			fmt.Fprintln(os.Stderr, "CheckAndPut requires 6 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		argvalue1 := []byte(flag.Arg(2))
		value1 := argvalue1
		argvalue2 := []byte(flag.Arg(3))
		value2 := argvalue2
		argvalue3 := []byte(flag.Arg(4))
		value3 := argvalue3
		argvalue4 := []byte(flag.Arg(5))
		value4 := argvalue4
		arg100 := flag.Arg(6)
		mbTrans101 := thrift.NewTMemoryBufferLen(len(arg100))
		defer mbTrans101.Close()
		_, err102 := mbTrans101.WriteString(arg100)
		if err102 != nil {
			Usage()
			return
		}
		factory103 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt104 := factory103.GetProtocol(mbTrans101)
		argvalue5 := hbase.NewTPut()
		err105 := argvalue5.Read(jsProt104)
		if err105 != nil {
			Usage()
			return
		}
		value5 := argvalue5
		fmt.Print(client.CheckAndPut(value0, value1, value2, value3, value4, value5))
		fmt.Print("\n")
		break
	case "putMultiple":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "PutMultiple requires 2 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		arg107 := flag.Arg(2)
		mbTrans108 := thrift.NewTMemoryBufferLen(len(arg107))
		defer mbTrans108.Close()
		_, err109 := mbTrans108.WriteString(arg107)
		if err109 != nil {
			Usage()
			return
		}
		factory110 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt111 := factory110.GetProtocol(mbTrans108)
		containerStruct1 := hbase.NewTHBaseServicePutMultipleArgs()
		err112 := containerStruct1.ReadField2(jsProt111)
		if err112 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.Tputs
		value1 := argvalue1
		fmt.Print(client.PutMultiple(value0, value1))
		fmt.Print("\n")
		break
	case "deleteSingle":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "DeleteSingle requires 2 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		arg114 := flag.Arg(2)
		mbTrans115 := thrift.NewTMemoryBufferLen(len(arg114))
		defer mbTrans115.Close()
		_, err116 := mbTrans115.WriteString(arg114)
		if err116 != nil {
			Usage()
			return
		}
		factory117 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt118 := factory117.GetProtocol(mbTrans115)
		argvalue1 := hbase.NewTDelete()
		err119 := argvalue1.Read(jsProt118)
		if err119 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.DeleteSingle(value0, value1))
		fmt.Print("\n")
		break
	case "deleteMultiple":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "DeleteMultiple requires 2 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		arg121 := flag.Arg(2)
		mbTrans122 := thrift.NewTMemoryBufferLen(len(arg121))
		defer mbTrans122.Close()
		_, err123 := mbTrans122.WriteString(arg121)
		if err123 != nil {
			Usage()
			return
		}
		factory124 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt125 := factory124.GetProtocol(mbTrans122)
		containerStruct1 := hbase.NewTHBaseServiceDeleteMultipleArgs()
		err126 := containerStruct1.ReadField2(jsProt125)
		if err126 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.Tdeletes
		value1 := argvalue1
		fmt.Print(client.DeleteMultiple(value0, value1))
		fmt.Print("\n")
		break
	case "checkAndDelete":
		if flag.NArg()-1 != 6 {
			fmt.Fprintln(os.Stderr, "CheckAndDelete requires 6 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		argvalue1 := []byte(flag.Arg(2))
		value1 := argvalue1
		argvalue2 := []byte(flag.Arg(3))
		value2 := argvalue2
		argvalue3 := []byte(flag.Arg(4))
		value3 := argvalue3
		argvalue4 := []byte(flag.Arg(5))
		value4 := argvalue4
		arg132 := flag.Arg(6)
		mbTrans133 := thrift.NewTMemoryBufferLen(len(arg132))
		defer mbTrans133.Close()
		_, err134 := mbTrans133.WriteString(arg132)
		if err134 != nil {
			Usage()
			return
		}
		factory135 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt136 := factory135.GetProtocol(mbTrans133)
		argvalue5 := hbase.NewTDelete()
		err137 := argvalue5.Read(jsProt136)
		if err137 != nil {
			Usage()
			return
		}
		value5 := argvalue5
		fmt.Print(client.CheckAndDelete(value0, value1, value2, value3, value4, value5))
		fmt.Print("\n")
		break
	case "increment":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "Increment requires 2 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		arg139 := flag.Arg(2)
		mbTrans140 := thrift.NewTMemoryBufferLen(len(arg139))
		defer mbTrans140.Close()
		_, err141 := mbTrans140.WriteString(arg139)
		if err141 != nil {
			Usage()
			return
		}
		factory142 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt143 := factory142.GetProtocol(mbTrans140)
		argvalue1 := hbase.NewTIncrement()
		err144 := argvalue1.Read(jsProt143)
		if err144 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.Increment(value0, value1))
		fmt.Print("\n")
		break
	case "append":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "Append requires 2 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		arg146 := flag.Arg(2)
		mbTrans147 := thrift.NewTMemoryBufferLen(len(arg146))
		defer mbTrans147.Close()
		_, err148 := mbTrans147.WriteString(arg146)
		if err148 != nil {
			Usage()
			return
		}
		factory149 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt150 := factory149.GetProtocol(mbTrans147)
		argvalue1 := hbase.NewTAppend()
		err151 := argvalue1.Read(jsProt150)
		if err151 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.Append(value0, value1))
		fmt.Print("\n")
		break
	case "openScanner":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "OpenScanner requires 2 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		arg153 := flag.Arg(2)
		mbTrans154 := thrift.NewTMemoryBufferLen(len(arg153))
		defer mbTrans154.Close()
		_, err155 := mbTrans154.WriteString(arg153)
		if err155 != nil {
			Usage()
			return
		}
		factory156 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt157 := factory156.GetProtocol(mbTrans154)
		argvalue1 := hbase.NewTScan()
		err158 := argvalue1.Read(jsProt157)
		if err158 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.OpenScanner(value0, value1))
		fmt.Print("\n")
		break
	case "getScannerRows":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "GetScannerRows requires 2 args")
			flag.Usage()
		}
		tmp0, err159 := (strconv.Atoi(flag.Arg(1)))
		if err159 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		tmp1, err160 := (strconv.Atoi(flag.Arg(2)))
		if err160 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		fmt.Print(client.GetScannerRows(value0, value1))
		fmt.Print("\n")
		break
	case "closeScanner":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "CloseScanner requires 1 args")
			flag.Usage()
		}
		tmp0, err161 := (strconv.Atoi(flag.Arg(1)))
		if err161 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.CloseScanner(value0))
		fmt.Print("\n")
		break
	case "mutateRow":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "MutateRow requires 2 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		arg163 := flag.Arg(2)
		mbTrans164 := thrift.NewTMemoryBufferLen(len(arg163))
		defer mbTrans164.Close()
		_, err165 := mbTrans164.WriteString(arg163)
		if err165 != nil {
			Usage()
			return
		}
		factory166 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt167 := factory166.GetProtocol(mbTrans164)
		argvalue1 := hbase.NewTRowMutations()
		err168 := argvalue1.Read(jsProt167)
		if err168 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.MutateRow(value0, value1))
		fmt.Print("\n")
		break
	case "getScannerResults":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "GetScannerResults requires 3 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		arg170 := flag.Arg(2)
		mbTrans171 := thrift.NewTMemoryBufferLen(len(arg170))
		defer mbTrans171.Close()
		_, err172 := mbTrans171.WriteString(arg170)
		if err172 != nil {
			Usage()
			return
		}
		factory173 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt174 := factory173.GetProtocol(mbTrans171)
		argvalue1 := hbase.NewTScan()
		err175 := argvalue1.Read(jsProt174)
		if err175 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		tmp2, err176 := (strconv.Atoi(flag.Arg(3)))
		if err176 != nil {
			Usage()
			return
		}
		argvalue2 := int32(tmp2)
		value2 := argvalue2
		fmt.Print(client.GetScannerResults(value0, value1, value2))
		fmt.Print("\n")
		break
	case "getRegionLocation":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "GetRegionLocation requires 3 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		argvalue1 := []byte(flag.Arg(2))
		value1 := argvalue1
		argvalue2 := flag.Arg(3) == "true"
		value2 := argvalue2
		fmt.Print(client.GetRegionLocation(value0, value1, value2))
		fmt.Print("\n")
		break
	case "getAllRegionLocations":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetAllRegionLocations requires 1 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		fmt.Print(client.GetAllRegionLocations(value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}