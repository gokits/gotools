package ratelimit

import (
	"context"
	"io"
	"net"
	"time"

	"golang.org/x/time/rate"
)

type tbReader struct {
	r io.Reader
	l *rate.Limiter
}

func (r *tbReader) Read(buf []byte) (int, error) {
	if r.l == nil {
		return r.r.Read(buf)
	}
	n, e := r.r.Read(buf)
	if n <= 0 {
		return n, e
	}
	ee := r.l.WaitN(context.Background(), n)
	if ee != nil {
		return n, ee
	}
	return n, e
}

func TokenBucketReader(r io.Reader, l *rate.Limiter) io.Reader {
	return &tbReader{
		r: r,
		l: l,
	}
}

type tbWriter struct {
	w io.Writer
	l *rate.Limiter
}

func (w *tbWriter) Write(buf []byte) (int, error) {
	if w.l != nil {
		if e := w.l.WaitN(context.Background(), len(buf)); e != nil {
			return 0, e
		}
	}
	return w.w.Write(buf)
}

func TokenBucketWriter(w io.Writer, l *rate.Limiter) io.Writer {
	return &tbWriter{
		w: w,
		l: l,
	}
}

type tbConn struct {
	conn net.Conn
	tbReader
	tbWriter
}

func (tbc *tbConn) Close() error {
	return tbc.conn.Close()
}

func (tbc *tbConn) LocalAddr() net.Addr {
	return tbc.conn.LocalAddr()
}

func (tbc *tbConn) RemoteAddr() net.Addr {
	return tbc.conn.RemoteAddr()
}

func (tbc *tbConn) SetDeadline(t time.Time) error {
	return tbc.conn.SetDeadline(t)
}

func (tbc *tbConn) SetWriteDeadline(t time.Time) error {
	return tbc.conn.SetWriteDeadline(t)
}

func (tbc *tbConn) SetReadDeadline(t time.Time) error {
	return tbc.conn.SetReadDeadline(t)
}

func TokenBucketConn(c net.Conn, readlimiter, writelimiter *rate.Limiter) net.Conn {
	return &tbConn{
		conn: c,
		tbReader: tbReader{
			r: c,
			l: readlimiter,
		},
		tbWriter: tbWriter{
			w: c,
			l: writelimiter,
		},
	}
}
