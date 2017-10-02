package simple

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Message struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   MessageSpec
	Status MessageStatus
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MessageList struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Items []Message
}

type MessageSpec struct {
	Header string
	Body   string
}

type MessageStatus struct {
	Sent bool
}
