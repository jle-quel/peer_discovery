/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jle-quel <jle-quel@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/01/05 10:49:05 by jle-quel          #+#    #+#             */
/*   Updated: 2018/01/05 11:05:12 by jle-quel         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	guid := getGuid()
	fmt.Println(guid)
}
