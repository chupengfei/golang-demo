package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Mail _
type Mail struct {
	Value string
}

// UnmarshalJSON _
func (m *Mail) UnmarshalJSON(data []byte) error {
	// 这里简单演示一下，简单判断即可
	if bytes.Contains(data, []byte("@")) {
		return fmt.Errorf("mail format error")
	}
	m.Value = string(data)
	return nil
}

// UnmarshalJSON _
func (m *Mail) MarshalJSON() (data []byte, err error) {
	if m != nil {
		data = []byte(m.Value)
	}
	return
}

// Phone _
type Phone struct {
	Value string
}

// UnmarshalJSON _
func (p *Phone) UnmarshalJSON(data []byte) error {
	// 这里简单演示一下，简单判断即可
	if len(data) != 11 {
		return fmt.Errorf("phone format error")
	}
	p.Value = string(data)
	return nil
}

// UnmarshalJSON _
func (p *Phone) MarshalJSON() (data []byte, err error) {
	if p != nil {
		data = []byte(p.Value)
	}
	return
}

// UserRequest _
type UserRequest struct {
	Name  string
	Mail  Mail
	Phone Phone
}

func main() {
	user := UserRequest{}
	user.Name = "ysy"
	user.Mail.Value = "yangshiyu@x.com"
	//user.Mail.UnmarshalJSON("yangshiyu@x.com")
	user.Phone.Value = "18900001111"
	data,_:= json.Marshal(user)
	fmt.Println(string(data))
}
