package dedicated_node 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Broker struct {

	/*Ssl - Descr: SSL Certificate for broker (PEM encoded) Default: <nil>
*/
	Ssl *Ssl `yaml:"ssl,omitempty"`

	/*StartRedisTimeout - Descr: Maximum wait time in seconds for Redis to start up Default: 600
*/
	StartRedisTimeout interface{} `yaml:"start_redis_timeout,omitempty"`

	/*DedicatedPort - Descr: The port used by the redis instances Default: 6379
*/
	DedicatedPort interface{} `yaml:"dedicated_port,omitempty"`

	/*Backups - Descr: Path within the above bucket to which backups will be uploaded Default: 
*/
	Backups *Backups `yaml:"backups,omitempty"`

	/*BackendHost - Descr: The port for the broker unicorn process to run on Default: <nil>
*/
	BackendHost interface{} `yaml:"backend_host,omitempty"`

	/*Auth - Descr: The username for HTTP Basic Auth on the agent, also used for the broker Default: admin
*/
	Auth *Auth `yaml:"auth,omitempty"`

}