package jwt

import (
	"testing"

	"github.com/golang-jwt/jwt/v4"
)

func TestGenerateToken(t *testing.T) {
	data := []byte(`-----BEGIN RSA PRIVATE KEY-----
	MIIEpAIBAAKCAQEAuUE6MQYpLSYTTWPFgqlGhjQuvvM7BPJknnRM6i0M/QxkPjV7
	wVhUmO0m/OQH0jpyLRfCDtIe+6k9Em/R5f+tItwJdX7YnRQj739BDLS6OtVe1U6A
	zaIxPXUNp9whw53wDi46Npu9Oj7W4opqLqcebeCTGsTs2ZZXMTeuEqdjH5qIpAp8
	9wiAcfkJVDXXzLn53ua2UF8UBW6UB6WZUiyDOnbZiyM8aEyPFZOS+dD4J6GzhWWO
	Hk6LSlSS3J0QTwxWP76mX+Z0hEFN2y7aYkYM887G0vEHxgGR1gRYL9W2d6ofGOpC
	9rntglM/lFjTLuyTt82OvG/zNpITqMpl2f7mIQIDAQABAoIBADSTtphDT1xo+gF0
	15GFyWY9WScIkS+OO9Bc4Zd0wcLvyWN+emkdUKU2aWJMtYuM2urBW2opaHG0Toj1
	A3VrYxAJswWNwpySIsDONQAOKuyXshqH9TZnWeI/XmjNqSeDo6XiqNLm0IW/X+1b
	q1ciRxiYLCfky/PVH1nwm47fwtos3f7Tco3IvWoJ5cLx81G1RjGy40g4JQgu6Y4d
	gemg6gxNxVkUlIhKogmbxw8lQ52HQj5MA0oIizhBZ5lGfvQweZg9uAAxD7B/o9FN
	m1GMAVbnE+37KdOe+IpSqwXSM4z1O/TstlVPVzETkgpi6EH1r39cVmjMrYjS3P/p
	o2YzN4ECgYEA0y/Nq92slp+ox1/SGO40eIHxsSkeu8c+2hU/5NFTxthTZkhsuc/8
	D176WmcuP9f4cH9GQNttciHXENLQQxr7wqyQmXr6LH+xKuIEG6Azniyva2+AStXZ
	ClWVDfenKZ6vdl6fQrBQZBm4ZIMzMGUOHNDN6KFqjTCfeBvFmCFGSx0CgYEA4JC8
	TWS8w5NzmtrZQRN5nXRlJrF/zfF/T2dYKVbbU+36P04/aBokiK8H+t9SMYBKBZ4Y
	O6vCTlQ/2Ylw28+LYHmk9pI3tOSH7eN0yK3uf3lhfhzfuqSA42x2CwLRYV3ePR94
	+y6pOiAduw3Gk6MHg7cov+j9YjQv96wyrEBIU9UCgYAo4E7R3kVCnkKf3Icg9jJH
	xPcQEGJ3fBhJQ+WrF208xrBV0tzglkY2f267ws7r4pSeybPCZ/ZmKfCju/o9kuJ6
	CsauPmRf6pNTNo4O/tIe4o1NoTZWxc+NtYT1QCMl5iYWJjhzQwc8aNIM+mwVCOpT
	X4MLbPrfuzKSn+V+Daq6bQKBgQCwOwV7Ww4EEahMo2iReF4dYeWrmupoTRgS0VBP
	iNkwoXaczaRW5wU6utlxWGZhcIMmq23rDhy7acR9zC1kYxjHCk+aLR9Id3vxBWp1
	voGHSlNnK/OrvkNjto9gjsKQPaEVqKhBvtmRQvjx/ZqT7AIOrONuz8JhmgD2xg0y
	PnRETQKBgQC4d6f7Z6TELOOHBYSNCdVT0DKo6PKw5h0B2nrBMGbFLeqsp+B9tvm+
	bAbrNnZHeYJsc4LGIWDg0z+9VxIDIb3Hq67M+TCplD1b4ZitU45L2GwNbZFvw+OQ
	UvH3j7t051JzNcJqiidgrTBHAsytUnctThD+6KXkhVjqsoycXdFr9w==
	-----END RSA PRIVATE KEY-----
	`)

	pk, err := jwt.ParseRSAPrivateKeyFromPEM(data)
	if err != nil {
		t.Fatal(err)
	}

	tokenManager := &TokenManager{
		privateKey: pk,
		publicKey:  &pk.PublicKey,
	}

	_, err = tokenManager.sign("xxarchexx", "ss@mail.ru")
	if err != nil {
		t.Fatal(err)
	}

}
