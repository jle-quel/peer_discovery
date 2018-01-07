/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   get.go                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/05 10:52:27 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/05 17:08:41 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"os/exec"
	"net"
	"strings"
	"math/rand"
	"time"
)

/*
**** PUBLIC ********************************************************************
*/

func getGuid() string {
	cmd := exec.Command("/bin/sh", "-c", "/usr/bin/base64 /dev/urandom | /usr/bin/head -c 64")
	ret, err := cmd.Output()
	handleErr(err)
	return string(ret)
}

func getAddr(addr net.Addr) string {
	ip := addr.String()
	return ip[:strings.LastIndex(ip, ":")]
}

func getRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
    return rand.Intn(42)
}

func getHeader() s_header {
	return s_header{getGuid(), uint32(getRandomNumber())}
}
