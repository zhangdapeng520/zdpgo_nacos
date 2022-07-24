/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package naming_client

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net"
	"strconv"
	"time"

	"github.com/zhangdapeng520/zdpgo_nacos/nacos/util"
)

type PushReceiver struct {
	port        int
	host        string
	hostReactor *HostReactor
}

type PushData struct {
	PushType    string `json:"type"`
	Data        string `json:"data"`
	LastRefTime int64  `json:"lastRefTime"`
}

var (
	GZIP_MAGIC = []byte("\x1F\x8B")
)

func NewPushReceiver(hostReactor *HostReactor) *PushReceiver {
	pr := PushReceiver{
		hostReactor: hostReactor,
	}
	pr.startServer()
	return &pr
}

func (us *PushReceiver) tryListen() (*net.UDPConn, bool) {
	addr, err := net.ResolveUDPAddr("udp", us.host+":"+strconv.Itoa(us.port))
	if err != nil {

		return nil, false
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {

		return nil, false
	}

	return conn, true
}

func (us *PushReceiver) getConn() *net.UDPConn {
	var conn *net.UDPConn
	for i := 0; i < 3; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		port := r.Intn(1000) + 54951
		us.port = port
		conn1, ok := us.tryListen()

		if ok {
			conn = conn1

			return conn
		}

		if !ok && i == 2 {

		}
	}
	return nil
}

func (us *PushReceiver) startServer() {
	conn := us.getConn()
	go func() {
		defer func() {
			if conn != nil {
				conn.Close()
			}
		}()
		for {
			us.handleClient(conn)
		}
	}()
}

func (us *PushReceiver) handleClient(conn *net.UDPConn) {

	if conn == nil {
		time.Sleep(time.Second * 5)
		conn = us.getConn()
		if conn == nil {
			return
		}
	}

	data := make([]byte, 4024)
	n, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {

		return
	}

	s := TryDecompressData(data[:n])

	var pushData PushData
	err1 := json.Unmarshal([]byte(s), &pushData)
	if err1 != nil {

		return
	}
	ack := make(map[string]string)

	if pushData.PushType == "dom" || pushData.PushType == "service" {
		us.hostReactor.ProcessServiceJson(pushData.Data)

		ack["type"] = "push-ack"
		ack["lastRefTime"] = strconv.FormatInt(pushData.LastRefTime, 10)
		ack["data"] = ""

	} else if pushData.PushType == "dump" {
		ack["type"] = "dump-ack"
		ack["lastRefTime"] = strconv.FormatInt(pushData.LastRefTime, 10)
		ack["data"] = util.ToJsonString(us.hostReactor.serviceInfoMap)
	} else {
		ack["type"] = "unknow-ack"
		ack["lastRefTime"] = strconv.FormatInt(pushData.LastRefTime, 10)
		ack["data"] = ""
	}

	bs, _ := json.Marshal(ack)
	_, err = conn.WriteToUDP(bs, remoteAddr)
	if err != nil {

	}
}

func TryDecompressData(data []byte) string {

	if !IsGzipFile(data) {
		return string(data)
	}
	reader, err := gzip.NewReader(bytes.NewReader(data))

	if err != nil {

		return ""
	}

	defer reader.Close()
	bs, err := ioutil.ReadAll(reader)

	if err != nil {

		return ""
	}

	return string(bs)
}

func IsGzipFile(data []byte) bool {
	if len(data) < 2 {
		return false
	}

	return bytes.HasPrefix(data, GZIP_MAGIC)
}
