package necon

import(
	"k8s.io/api/core/v1"
	"reflect"
	"fmt"
	"sync"
)

type necon struct {
	path string
	pod v1.Pod
	count int
}

var(
	n *necon
	once sync.Once
)


func GetInstance() *necon{
	once.Do(func() {
		n = &necon{}
	})

	return n
}

func (nec *necon) SetNeconPod(pod v1.Pod) error{
	namespace := pod.ObjectMeta.GetNamespace()
	if namespace != "kube-system"{
		nec.pod = pod
		nec.count++
		fmt.Println("count : ",nec.count)
		fmt.Println("not kubesystem pod ",nec.pod)
	}
	return nil
}
/*
func (nec *necon) GetNeconPod() *v1.Pod{
	return nec.pod
}
*/

func (nec *necon) SetSLO() error{

	fmt.Println("namespace : ",nec.pod.ObjectMeta.GetNamespace())
	fmt.Println("slo : ",reflect.TypeOf(nec.pod.Spec))

	//if nec.namespace != "kube-system" || nec.namespace != ""{
	//}

	return nil
}
//func (n *Necon) 
