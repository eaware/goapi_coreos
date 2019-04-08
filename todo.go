package main

/*
domain=web.gen.local
frontip=10.11.23.41
frontmask=24
defaultgw=10.11.23.1
bckip=10.13.86.177
bckmask=23
bckgw=10.13.86.1
mgtip=10.13.218.177
mgtmask=23
mgtgw=10.13.218.1
*/

type todoT struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Domain    string   `json:"domain"`
	Defaultgw string   `json:"defaultgw"`
	Front     *network `json:"front"`
	Backup    *network `json:"backup"`
	Managment *network `json:"managment"`
}

type network struct {
	Ipaddress string `json:"ipaddress"`
	Netmask   string `json:"netmask"`
	Gateway   string `json:"gateway"`
}

type todosT []todoT
