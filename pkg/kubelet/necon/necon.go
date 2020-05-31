package necon

import(
	"k8s.io/api/core/v1"
	"fmt"
	"sync"
)

type necon struct {
	path string
	pod *v1.Pod
}

var(
	n *necon
	once sync.Once
)


func GetInstance() *necon{
	once.Do(func() {
		n = new(necon)
	})

	return n
}

func (nec *necon) SetNeconPod(pod *v1.Pod) error{
	nec.pod = pod

	fmt.Println("nec.pod : ",nec.pod)
	return nil
}

func (nec *necon) GetNeconPod() *v1.Pod{
	return nec.pod
}
/*
func (n *Necon) SetSLO() error{
	

	return nil
}*/
//func (n *Necon) 
