package req

import (
   "testing"
   "code.google.com/p/goprotobuf/proto"
)


func TestSqrt(t *testing.T) {
    r := &RegMessage{
        Id:       proto.Int32(10001),
        Username: proto.String("vicky"),
        Password: proto.String("123456"),
        Email:    proto.String("eclipser@163.com"),
    }
    buffer, err := proto.Marshal(r)
    if err != nil {
        t.Errorf("failed: %s\n", err)
        return
    }

    m := &RegMessage{}
    err = proto.Unmarshal(buffer, m)
    if m.GetUsername() != r.GetUsername() {
        t.Errorf("ERROR m.Username=%v Expected=%v", m.GetUsername(), r.GetUsername())
    }

}
