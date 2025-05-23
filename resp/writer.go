package resp

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

// TODO: if slave then no write

func WriteSimpleString(conn net.Conn, res string) {
	conn.Write([]byte("+" + res + "\r\n"))
}

func WriteSimpleInt(conn net.Conn, res int) {
	conn.Write([]byte(fmt.Sprintf(":%d\r\n", res)))
}

func WriteSimpleError(conn net.Conn, errMsg string) {
	conn.Write([]byte("-" + errMsg + "\r\n"))
}

func WriteBulkString(conn net.Conn, str string) {
	if str == "" {
		conn.Write([]byte("$" + strconv.Itoa(-1) + "\r\n"))
		return
	}
	conn.Write([]byte("$" + strconv.Itoa(len(str)) + "\r\n" + str + "\r\n"))
}

func WriteArray(conn net.Conn, arr []string) {
	if len(arr) == 0 {
		conn.Write([]byte("$" + strconv.Itoa(-1) + "\r\n"))
		return
	}
	res := "*" + strconv.Itoa(len(arr)) + "\r\n"
	for _, key := range arr {
		res = res + "$" + strconv.Itoa(len(key)) + "\r\n" + key + "\r\n"
	}
	conn.Write([]byte(res))
}

func WriteRDB(conn net.Conn, file *os.File) error {
	stat, err := file.Stat()
	if err != nil {
		return err
	}

	conn.Write([]byte(fmt.Sprintf("$%d\r\n", stat.Size())))
	_, err = io.Copy(conn, file)
	if err != nil {
		return err
	}
	return nil
}

