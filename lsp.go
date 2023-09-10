package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
)

func main() {
	out, err := os.Create("changes.log")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	stream := jsonrpc2.NewStream(struct {
		io.Reader
		io.Writer
		io.Closer
	}{
		Reader: os.Stdin,
		Writer: os.Stdout,
		Closer: io.NopCloser(nil),
	})
	conn := jsonrpc2.NewConn(stream)
	ctx := context.Background()
	conn.Go(ctx, func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
		switch req.Method() {
		case protocol.MethodInitialize:
			var params protocol.InitializeParams
			if err := json.Unmarshal(req.Params(), &params); err != nil {
				return err
			}
			return reply(ctx, protocol.InitializeResult{
				Capabilities: protocol.ServerCapabilities{
					TextDocumentSync: protocol.TextDocumentSyncKindIncremental,
				},
			}, nil)
		case protocol.MethodTextDocumentDidChange:
			var params protocol.DidChangeTextDocumentParams
			if err := json.Unmarshal(req.Params(), &params); err != nil {
				return err
			}
			var buf bytes.Buffer
			if err := json.Indent(&buf, req.Params(), "", "  "); err != nil {
				return err
			}
			if _, err := fmt.Fprintln(out, req.Method(), buf.String()); err != nil {
				return err
			}
		}
		return nil
	})
	<-conn.Done()
	if err := conn.Err(); err != nil {
		log.Fatal(err)
	}
}
