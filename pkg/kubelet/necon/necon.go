package necon

import(
	"k8s.io/api/core/v1"
//	"github.com/prometheus/procfs"
//	"os"
	"io/ioutil"
//	"os/exec"
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
		//fmt.Println("slo : ",nec.pod.Spec.Containers[0].Resources.Limits["example.com/SLO"])
		//fmt.Println("count : ",nec.count)
		//fmt.Println("not kubesystem pod ",nec.pod)
	}
	return nil
}
/*
func (nec *necon) GetNeconPod() *v1.Pod{
	return nec.pod
}
*/

func (nec *necon) SetSLO() error{
	if nec.pod.ObjectMeta.GetNamespace() != ""{
		//fmt.Println("namespace : ",nec.pod.ObjectMeta.GetNamespace())
		//w := exec.Command("brctl","show")
		//w,err := ioutil.ReadFile(fmt.Sprintf("/proc/oslab/vif%d/goal",1))
		err := ioutil.WriteFile(fmt.Sprintf("/proc/oslab/vif%d/goal",nec.count),[]byte("1022"),0)
		if err != nil{
			fmt.Println("write errr!")
		}

	}
	nec.pod.Reset()

	return nil
}
