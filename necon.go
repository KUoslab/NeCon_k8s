package cm

import(
	"k8s.io/api/core/v1"
)

type Necon struct {
	path string
	pod *v1.Pod
}

func NewNecon() *Necon{
	return &Necon{}
}

func (n *Necon) SetNeconPod(pod *v1.Pod) error{
	n.pod = pod

	return nil
}

func (n *Necon) 
