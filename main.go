/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/13 09:36:55 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/15 00:21:06 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

/*
**** PUBLIC ********************************************************************
*/

func main() {
	headerCh := make(chan header)
	getHeader := initHeader()
	fmt.Println("Self ->", getHeader())
	fmt.Printf("\n")

	getHeader().Broadcast()
	go UDPServer(getHeader, headerCh)
	TCPServer(headerCh)
}
