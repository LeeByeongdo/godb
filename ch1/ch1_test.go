package ch1

import (
	"os"
	"testing"
)

// 코드로만 보면 정상 처리되는 것 같으나, sample.txt파일이 클 경우 (4억줄 텍스트 파일로 테스트)
// 테스트 프로그램이 종료 된 후 test.txt파일을 열어보면 기존 데이터가 나오는 시간이 존재한다.
// 아마 os 캐시 때문에 그런 현상이 나타나는 것 같다.
func TestSave1(t *testing.T) {
	sample, err := os.ReadFile("test/sample.txt")
	if err != nil {
		t.Fatalf("Reading sample error, %v", err)
		return
	}

	saveErr := saveData2("test/test.txt", sample)

	if saveErr != nil {
		t.Fatalf("Saving data error, %v", err)
		return
	}

	readingData, readingError := os.ReadFile("test/test.txt")

	a := string(readingData)
	last := a[len(a)-50:]

	println(last)

	if readingError != nil {
		t.Fatalf("Reading test file error, %v", err)
		return
	}
}
