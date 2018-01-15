/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   server.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/13 11:42:46 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/15 13:49:29 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"time"
)

/*
**** PUBLIC ********************************************************************
*/

func HeaderServer(addPeer func(peer header) t_map) {
	buf := make([]byte, HEADER_SIZE)
	listener := initTCPListen()
	listener.SetDeadline(time.Now().Add(1 * time.Second))

	for {
		conn, err := listener.Accept()
		switch err {
		case nil:
			conn.Read(buf)
			addPeer(decode(buf))
			conn.Close()
		default:
			return
		}
	}
	listener.Close()
}
