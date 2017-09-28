package main

import (
   "net/rpc"
   "net/rpc/jsonrpc"
   "net"
   "fmt"
   "log"
)
   
type Args struct {
   Nome, Sobrenome string
}

type Cadastro string

func (t *Cadastro) Usuario(args *Args, res *string) error {
   *res = (*args).Nome + " " + (*args).Sobrenome
   fmt.Println(*res)
   return nil
}
 
func main() {
   servicoCadastro := new(Cadastro)
   rpc.Register(servicoCadastro)
   porta := ":666"

   l, e := net.Listen("tcp", porta)
   fmt.Println("Ouvindo na porta", porta);
   if e != nil {
      log.Fatal("erro no listen: ", e)
   }
   
   for {
      conn, e := l.Accept()
      if e != nil {
         log.Fatal("erro na recepcao da conexao: ", e)
      }
      jsonrpc.ServeConn(conn)
   }
}