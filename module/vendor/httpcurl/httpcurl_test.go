package httpcurl

import "testing"

func Test_Request(t *testing.T) {

}
func Test_GetAllUserByOrgIDs(t *testing.T) {
	a := new(UMS)
	a.GetAllUserByOrgIDs([]uint64{1111})
}
