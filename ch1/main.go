package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
}

// 문제1. 파일을 쓰기 전에 항상 파일을 비운다. 만약 동시에 파일을 읽으려고 하면 어떻게 될까?
// 문제2. 파일 사이즈에 따라 파일이 한 번에 쓰여지지 않을 수 있다. 중간에 읽으려고 시도하면 완전하지 않은 데이터가 읽힌다.
// 문제3. 언제 데이터가 디스크에 완전히 저장되는가? write 후에도 os의 캐시 데이터 읽힐 수 있다. 충돌이 나거나 시스템이 재시작되면 파일은 어떤 상태인가?
func saveData1(path string, data []byte) error {
	// os.O_WRONLY write-only로 파일을 연다
	// os.O_CREATE 파일이 없으면 만든다
	// os.O_TRUNC 기존에 파일이 있으면, 전체를 날리고 덮어 쓴다.
	// 0664 퍼미션 (맨 앞의 0은 8진수 octal 뜻함, hexadecimal 의 0x같은)
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}

	// defer는 해당 스코프가 종료되기 전에 실행된다. finally 의 역할과 비슷하다.
	// 여기서는 파일이 사용이 완료되든, 중간에 에러가 나든 항상 close 될 수 있도록 한다.
	defer fp.Close()

	_, err = fp.Write(data)
	return err
}
