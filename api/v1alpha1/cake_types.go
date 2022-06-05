/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Topping string

const (
	Strawberry   Topping = "strawberry"
	Choclate     Topping = "choclate"
	PeanutButter Topping = "peanut butter"
	BlueBerry    Topping = "blueberry"
)

type DeliveryType string

const (
	Pickup  DeliveryType = "pickup"
	Deliver DeliveryType = "deliver"
)

type Delivery struct {
	DeliveryType    DeliveryType `json:"deliveryType,omitempty"`
	DeliveryAddress string       `json:"deliveryAddress,omitempty"`
}

// CakeSpec defines the desired state of Cake
type CakeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Flavor        string    `json:"flavor,omitempty"`
	Occasion      string    `json:"occasion,omitempty"`
	MessageOnCake string    `json:"messageOnCake,omitempty"`
	Toppings      []Topping `json:"toppings,omitempty"`
	Delivery      Delivery  `json:"delivery,omitempty"`
}

// CakeStatus defines the observed state of Cake
type CakeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Ready         bool `json:"ready,omitempty"`
	Delivered     bool `json:"delivered,omitempty"`
	InProgress    bool `json:"inProgress,omitempty"`
	OrderReceived bool `json:"orderReceived,omitempty"`
	PickedUp      bool `json:"pickedUp,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Cake is the Schema for the cakes API
type Cake struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CakeSpec   `json:"spec,omitempty"`
	Status CakeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CakeList contains a list of Cake
type CakeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cake `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cake{}, &CakeList{})
}
