package Storage

import "testing"

func TestAdd(t *testing.T)  {
	if res:= add(1,2);res!=3{
		t.Errorf("add result is wrong %d + %d ",1,2)
	}
}
