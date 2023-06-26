package utils

import (
	"bytes"
	"fmt"
	"github.com/ellioben/operator-hook/api/v1beta1"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"text/template"
)

func parseTemplate(templateName string, extPower *v1beta1.ExtPower) []byte {
	tmpl, err := template.ParseFiles("utils/template/" + templateName + ".yml")
	if err != nil {
		fmt.Printf("get template error：%v\n", err)
		return nil
	}
	b := new(bytes.Buffer)
	err = tmpl.Execute(b, extPower)
	if err != nil {
		fmt.Printf("Unmarshal template error：%v\n", err)
		return nil
	}
	return b.Bytes()
}

func NewDeployment(extPower *v1beta1.ExtPower) *appv1.Deployment {
	d := &appv1.Deployment{}
	err := yaml.Unmarshal(parseTemplate("deployment", extPower), d)
	if err != nil {
		fmt.Printf("Unmarshal deploy error：%v\n", err)
		return nil
	}
	return d
}

func NewIngress(extPower *v1beta1.ExtPower) *netv1.Ingress {
	i := &netv1.Ingress{}
	err := yaml.Unmarshal(parseTemplate("ingress", extPower), i)
	if err != nil {
		fmt.Printf("Unmarshal ingress error：%v\n", err)
		return nil
	}
	return i
}

func NewService(extPower *v1beta1.ExtPower) *corev1.Service {
	s := &corev1.Service{}
	err := yaml.Unmarshal(parseTemplate("service", extPower), s)
	if err != nil {
		fmt.Printf("Unmarshal service error：%v\n", err)
		return nil
	}
	return s
}
