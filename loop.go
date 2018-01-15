/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   loop.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/15 10:58:36 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/15 13:59:55 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"net"
)

/*
**** PRIVATE *******************************************************************
*/

func handleConn(conn *net.UDPConn) header {
	buf := make([]byte, HEADER_SIZE)

	_, err := conn.Read(buf)
	handleErr(err)
	peer := decode(buf)

	fmt.Println("New peer")
	fmt.Printf("Id [%s]\n", peer.Id)
	fmt.Printf("Addr [%s]\n", peer.Addr)
	fmt.Printf("Timestamp [%d]\n\n", peer.Timestamp)

	return peer
}

func debug(peers t_map) {
	fmt.Println("Routing Table")
	for _, value := range peers {
		fmt.Println(value)
	}
	fmt.Printf("\n")
}

/*
**** PUBLIC ********************************************************************
*/

func loop(getHeader func() header) {
	addPeer := initRoutingTable()
	conn := initUDPListen()

	for {
		fmt.Println("Listening...")
		peer := handleConn(conn)
		debug(addPeer(peer))
		go HeaderServer(addPeer)
		getHeader().Bootstrap(peer.Addr + TCP_PORT)
	}
	conn.Close()
}
