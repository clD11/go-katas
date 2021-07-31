package hackerrank

import (
	"github.com/clD11/go-katas/src/testsupport"
	"testing"
)

func TestStrongPassword_Numbers_Fail(t *testing.T) {
	p := Password{}
	p.addRule(numbers)
	actual := p.StrongPassword("aaaaaa")
	testsupport.AssertThatInt32(t, actual, 1)
}

func TestStrongPassword_Numbers_Pass(t *testing.T) {
	p := Password{}
	p.addRule(numbers)
	actual := p.StrongPassword("aaaaa1")
	testsupport.AssertThatInt32(t, actual, 0)
}

func TestStrongPassword_Lower_Fail(t *testing.T) {
	p := Password{}
	p.addRule(lower)
	actual := p.StrongPassword("AAA28940")
	testsupport.AssertThatInt32(t, actual, 1)
}

func TestStrongPassword_Lower_Pass(t *testing.T) {
	p := Password{}
	p.addRule(lower)
	actual := p.StrongPassword("aaaAAA")
	testsupport.AssertThatInt32(t, actual, 0)
}

func TestStrongPassword_Upper_Fail(t *testing.T) {
	p := Password{}
	p.addRule(upper)
	actual := p.StrongPassword("asascdasc2345345345")
	testsupport.AssertThatInt32(t, actual, 1)
}

func TestStrongPassword_Upper_Pass(t *testing.T) {
	p := Password{}
	p.addRule(upper)
	actual := p.StrongPassword("aaaaaA")
	testsupport.AssertThatInt32(t, actual, 0)
}

func TestStrongPassword_Special_Fail(t *testing.T) {
	p := Password{}
	p.addRule(special)
	actual := p.StrongPassword("sdvnsdon83439hnvjUIBNIU")
	testsupport.AssertThatInt32(t, actual, 1)
}

func TestStrongPassword_Special_Pass(t *testing.T) {
	p := Password{}
	p.addRule(special)
	actual := p.StrongPassword("IONVSIODV94304ufhwvkn*")
	testsupport.AssertThatInt32(t, actual, 0)
}

func TestStrongPassword_Successful(t *testing.T) {
	p := Password{}
	p.addRule(lower)
	p.addRule(upper)
	p.addRule(special)
	actual := p.StrongPassword("aA1!aA1!")
	testsupport.AssertThatInt32(t, actual, 0)
}

func TestStrongPassword_Fail(t *testing.T) {
	p := Password{}
	p.addRule(lower)
	p.addRule(upper)
	p.addRule(special)
	actual := p.StrongPassword("Ab1")
	testsupport.AssertThatInt32(t, actual, 3)
}

func TestStrongPassword_Min(t *testing.T) {
	p := Password{}
	p.addRule(lower)
	p.addRule(upper)
	p.addRule(special)
	actual := p.StrongPassword("4700")
	testsupport.AssertThatInt32(t, actual, 3)
}
