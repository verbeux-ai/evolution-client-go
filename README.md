# Evolution Client Go

Essa biblioteca foi desenvolvida para o uso do [https://github.com/evolutionapi/evolution-api/](evolution), api do
whatsapp usando a linguagem go.


Fork Verbeux da evolution: https://github.com/verbeux-ai/evolution-api/  
Status: Em desenvolvimento

## Instalando
```go
go get github.com/verbeux-ai/evolution-client-go
```

## Enviando Mensagem

Este é um exemplo de envio de mensagem a partir do client.

```go
import evolution "github.com/verbeux-ai/evolution-client-go"
import "context"

client = evolution.NewClient(
    evolution.WithApiKey(os.Getenv("API_KEY")),
    evolution.WithBaseUrl(os.Getenv("BASE_URL")),
)
response, err := client.SendTextMessage(context.Background(), &evolution.TextMessageRequest{
    Number:   "558500000000",
    Phone: "Teste",
})
if err != nil {
    panic(err)
}
fmt.Println(response)
```

## Escutando mensagens

Este é um exemplo de como escutar mensagens no webhook

```go
import "github.com/verbeux-ai/evolution/listener"

whatsappListener := listener.NewMessageListener()
whatsappListener.HandleErrors(func (err error) {
    fmt.Println("fail", err)
})

// register listeners
whatsappListener.OnMessage(func (message *listener.MessageUpsert) error {

    // treat your text message here
    
    return nil
})

if err := whatsappListener.ReadBodyAsync(ctx.Request().Body); err != nil {
    panic(err)
}
```

## Features disponíveis

| Funcionalidade             | Implementado |
|----------------------------|--------------|
| Find Chats                 | Sim          |
| Read Messages              | Sim          |
| Create Instance            | Sim          |
| Restart Instance           | Sim          |
| Logout Instance            | Sim          |
| Delete Instance            | Sim          |
| Connect Instance           | Sim          |
| Connection State Instance  | Sim          |
| Fetch Instances            | Sim          |
| Send Text Message          | Sim          |
| Send Media Message         | Sim          |
| Get Tags                   | Sim          |
| Add Chat Tag               | Sim          |

> Você está convidado a contribuir ao repositório!