package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet= "abcdefghijklmnopqrstuvwxyz"
func init() {
	rand.Seed(time.Now().UnixNano())

}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string{
	var sb strings.Builder // what is strings.builder
	k := len(alphabet)

	for i:=0;i<n;i++ {
		c:= alphabet[rand.Intn(k)];
		sb.WriteByte(c)
	}

	return sb.String();
}


func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0,1000)
}


func RandomCurrency() string {
	currencies := []string{EUR,USD,INR}
	// ["EUR","USD","INR"];
	return currencies[RandomInt(0,int64(len(currencies)-1))]
}

func RandomEmail() string{
	return fmt.Sprintf("%s@email.com",RandomString(int(RandomInt(10,20))))
}