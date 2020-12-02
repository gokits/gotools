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
	ctx := context.Background()
	max := len(buf)
	total := 0
	burst := r.l.Burst()
	for total < max {
		next := total + burst
		if next > max {
			next = max
		}
		n, e := r.r.Read(buf[total:next])
		if n <= 0 {
			return total, e
		}
		total += n
		ee := r.l.WaitN(ctx, n)
		if ee != nil {
			return total, ee
		}
	}
	return total, nil
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
	if w.l == nil {
		return w.w.Write(buf)
	}
	ctx := context.Background()
	max := len(buf)
	burst := w.l.Burst()
	total := 0
	for total < max {
		next := total + burst
		if next > max {
			next = max
		}
		if e := w.l.WaitN(ctx, next-total); e != nil {
			return total, e
		}
		n, e := w.w.Write(buf[total:next])
		if e != nil {
			return total + n, e
		}
		total += n
	}
	return total, nil
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
