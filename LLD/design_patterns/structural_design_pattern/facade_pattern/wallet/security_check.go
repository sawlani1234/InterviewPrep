package wallet 


type SecurityCheck struct {
	code int 

}

func NewSecurityCheck(code int) *SecurityCheck {
	return &SecurityCheck{code}
}


func (s SecurityCheck) checkCode(n int) bool {
	return s.code == n

}