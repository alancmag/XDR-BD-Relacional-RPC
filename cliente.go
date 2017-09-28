package main

import (
   "net/rpc/jsonrpc"
   "fmt"
   "log"
   "os"
)

type Args struct {
   Nome, Sobrenome string
}

func main() {
   args := &Args{"Warley", "Santos"}
   var resposta string

   cliente, erro := jsonrpc.Dial("tcp", "localhost:666")
   if erro != nil {
      log.Fatal("Erro dialing: ", erro);
      os.Exit(1)
   }

   erro = cliente.Call("Cadastro.Usuario", args, &resposta)
   if erro != nil {
      log.Fatal("Erro invocacao: ", erro);
      os.Exit(1)
   }
   fmt.Printf("Resposta do servidor: %s\n", resposta)
}
