package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"

	"strings"

	"gopkg.in/ini.v1"
)

// isIPV4, checks string is valid IP
func isIPV4(i string) bool {
	parts := strings.Split(i, ".")

	if len(parts) < 4 {
		return false
	}

	for _, x := range parts {
		if i, err := strconv.Atoi(x); err == nil {
			if i < 0 || i > 255 {
				return false

			}
		} else {
			return false
		}

	}
	return true
}

// isNetmaskPrefix, checks string is valid NETMASK
func isNetmaskPrefix(i string) bool {
	parts := strings.Split(i, ".")
	if len(parts) < 4 {
		return false
	}
	return true
}

// netmask2prefix, converts netmask to prefix
func netmask2prefix(i string) string {
	mask := net.IPMask(net.ParseIP(i).To4()) // If you have the mask as a string
	prefixSize, _ := mask.Size()

	return strconv.Itoa(prefixSize)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

// Test with this curl command:
// curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos
func ignitionCreate(w http.ResponseWriter, r *http.Request) {
	var todo todoT
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	body = bytes.Replace(body, []byte("'"), []byte("\""), -1)
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := repoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}

	cfg := ini.Empty()
	ini.PrettyFormat = false
	parts := strings.Split(todo.Name, ".")
	cfg.Section("").Key("name").SetValue(parts[0])

	cfg.Section("").Key("domain").SetValue(todo.Domain)
	if todo.Defaultgw != "" {
		cfg.Section("").Key("provided_defaultgw").SetValue(todo.Defaultgw)
	}

	if todo.Front != nil {
		if todo.Front.Gateway != "" {
			cfg.Section("").Key("defaultgw").SetValue(todo.Front.Gateway)
		}
		if isIPV4(todo.Front.Ipaddress) {
			cfg.Section("").Key("frontip").SetValue(todo.Front.Ipaddress)
		}
		if todo.Front.Netmask != "" {
			if isNetmaskPrefix(todo.Front.Netmask) {
				cfg.Section("").Key("frontmask").SetValue(netmask2prefix(todo.Front.Netmask))
			} else {
				cfg.Section("").Key("frontmask").SetValue(todo.Front.Netmask)
			}
		}
		if todo.Front.Gateway != "" {
			cfg.Section("").Key("frontgw").SetValue(todo.Front.Gateway)
		}
	}

	if todo.Backup != nil {
		if isIPV4(todo.Backup.Ipaddress) {
			cfg.Section("").Key("bckip").SetValue(todo.Backup.Ipaddress)
		}
		if todo.Backup.Netmask != "" {
			if isNetmaskPrefix(todo.Backup.Netmask) {
				cfg.Section("").Key("bckmask").SetValue(netmask2prefix(todo.Backup.Netmask))
			} else {
				cfg.Section("").Key("bckmask").SetValue(todo.Backup.Netmask)
			}
		}
		if todo.Backup.Gateway != "" {
			cfg.Section("").Key("bckgw").SetValue(todo.Backup.Gateway)
		}
	}

	if todo.Managment != nil {
		if isIPV4(todo.Managment.Ipaddress) {
			cfg.Section("").Key("mgtip").SetValue(todo.Managment.Ipaddress)
		}
		if todo.Managment.Netmask != "" {
			mask := net.IPMask(net.ParseIP(todo.Managment.Netmask).To4()) // If you have the mask as a string
			prefixSize, _ := mask.Size()

			cfg.Section("").Key("mgtmask").SetValue(strconv.Itoa(prefixSize))
			//cfg.Section("").Key("mgtmask").SetValue(todo.Managment.Netmask)
		}
		if todo.Managment.Gateway != "" {
			cfg.Section("").Key("mgtgw").SetValue(todo.Managment.Gateway)
		}
	}

	//        cfg.Section("").Key("completed").SetValue(strconv.FormatBool(todo.Completed))
	path := "/coreosini"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
	cfg.SaveTo("/coreosini/" + parts[0] + ".config")
}
