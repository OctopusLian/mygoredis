/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-06-18 22:35:17
 * @LastEditors: neozhang
 * @LastEditTime: 2022-06-18 22:35:21
 */
package tcp

import (
	"context"
	"net"
)

// Handler represents application server over tcp
type Handler interface {
	Handle(ctx context.Context, conn net.Conn)
	Close() error
}
