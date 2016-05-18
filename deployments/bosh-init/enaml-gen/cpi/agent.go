package cpi 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Agent struct {

	/*Nats - Descr: Address of the nats server Default: <nil>
*/
	Nats *Nats `yaml:"nats,omitempty"`

	/*User - Descr: Username agent uses to connect to blobstore used by simple blobstore plugin Default: <nil>
*/
	User interface{} `yaml:"user,omitempty"`

	/*Password - Descr: Password agent uses to connect to blobstore used by simple blobstore plugin Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

	/*Blobstore - Descr: Address for agent to connect to blobstore server used by simple blobstore plugin Default: <nil>
*/
	Blobstore *Blobstore `yaml:"blobstore,omitempty"`

	/*Mbus - Descr: Agent mbus Default: <nil>
*/
	Mbus interface{} `yaml:"mbus,omitempty"`

}