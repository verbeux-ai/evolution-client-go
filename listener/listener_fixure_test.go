package listener_test

const messageExample1 = `{"event":"messages.upsert","instance":"Ivo Teste","data":{"key":{"remoteJid":"558598437440@s.whatsapp.net","fromMe":true,"id":"3EB0742228A7FA7CB97E44"},"pushName":"Ivo","status":"SERVER_ACK","message":{"conversation":"fala doto"},"messageType":"conversation","messageTimestamp":1731071861,"instanceId":"abf7de57-cedb-4204-a9fd-8335717380a7","source":"web"},"destination":"https://webhook.site/75a58d38-80c5-4f9c-b164-7d76ac427377","date_time":"2024-11-08T13:17:41.471Z","sender":"558594138387@s.whatsapp.net","server_url":"http://localhost:8080","apikey":"xxxxxx"}`
const messageImageExample1 = `{"event":"messages.upsert","instance":"577e7717-a66d-11ef-88f4-42004e494300","data":{"key":{"remoteJid":"558594138387@s.whatsapp.net","fromMe":false,"id":"3EB025A75C7367ED422EBA"},"pushName":"Ivo","message":{"imageMessage":{"url":"https://mmg.whatsapp.net/o1/v/t62.7118-24/f1/m234/up-oil-image-100407ea-4c66-4136-9458-46ff96107c74?ccb=9-4&oh=01_Q5AaIDmDJW9myHQvuxEl7eE8J0PSY9ekG16eHysUrMXoybP6&oe=676437C1&_nc_sid=e6ed6c&mms3=true","mimetype":"image/jpeg","fileSha256":"im9CvFA5Ps86d6KSEP9I1L8qH2cvBn08pie1DqNhfUc=","fileLength":"226688","height":1600,"width":1600,"mediaKey":"cb8tqyHM3vCfn/m1TRIiAamS9ApFPKSwe3w/qdVzD5w=","fileEncSha256":"4xwfVRHnLbp8ZUg9XfM8J2KKgvTT0M942Qmw9nVXHsE=","directPath":"/o1/v/t62.7118-24/f1/m234/up-oil-image-100407ea-4c66-4136-9458-46ff96107c74?ccb=9-4&oh=01_Q5AaIDmDJW9myHQvuxEl7eE8J0PSY9ekG16eHysUrMXoybP6&oe=676437C1&_nc_sid=e6ed6c","mediaKeyTimestamp":"1732041575","jpegThumbnail":"thumblink","contextInfo":{"disappearingMode":{"initiator":"CHANGED_IN_CHAT"}},"viewOnce":false},"messageContextInfo":{"deviceListMetadata":{"senderKeyHash":"DK/EyaUp7VXSag==","senderTimestamp":"1732035332","senderAccountType":"E2EE","receiverAccountType":"E2EE","recipientKeyHash":"VEgMtMpwrVz/JA==","recipientTimestamp":"1732019837"},"deviceListMetadataVersion":2,"messageSecret":"0xIb/FC49IbKxwIvrCAKuuzdLnoGwd4/rsD8v/RKq88="},"base64":"base64img"},"contextInfo":{"disappearingMode":{"initiator":"CHANGED_IN_CHAT"}},"messageType":"imageMessage","messageTimestamp":1732041577,"instanceId":"f890a78d-5497-486d-99f7-712111aee095","source":"web"},"destination":"https://webhook.site/bfcd9054-1a18-4cd2-bb31-890c8337ecdd","date_time":"2024-11-19T18:39:38.343Z","sender":"5585920013048@s.whatsapp.net","server_url":"http://localhost:80","apikey":null}`
const messageAudioExample1 = `{"event":"messages.upsert","instance":"577e7717-a66d-11ef-88f4-42004e494300","data":{"key":{"remoteJid":"558594138387@s.whatsapp.net","fromMe":false,"id":"3EB0B5C72FAD50804DAF26"},"pushName":"Ivo","message":{"audioMessage":{"url":"https://mmg.whatsapp.net/v/t62.7117-24/19144390_561556576731467_3681810655211256607_n.enc?ccb=11-4&oh=01_Q5AaILQcNJ0HXHgIyPFQeGHFZtSf2bRP9jEcNa_L0EN_DLfz&oe=67645F65&_nc_sid=5e03e0&mms3=true","mimetype":"audio/ogg; codecs=opus","fileSha256":"SPe2uSTzO20eu3z9LNS6wN7pLTbxJued0UDbWabXzRM=","fileLength":"2881","seconds":2,"ptt":true,"mediaKey":"tDWG1TjksYj5PSphhG+SVFFAb4gCmJX4nz5ryvCLdAc=","fileEncSha256":"ATo5oVDpa6WsVkMzzDt+91PHl95Pjs/EPGwQgY2jYlM=","directPath":"/v/t62.7117-24/19144390_561556576731467_3681810655211256607_n.enc?ccb=11-4&oh=01_Q5AaILQcNJ0HXHgIyPFQeGHFZtSf2bRP9jEcNa_L0EN_DLfz&oe=67645F65&_nc_sid=5e03e0","mediaKeyTimestamp":"1732042072","contextInfo":{"disappearingMode":{"initiator":"CHANGED_IN_CHAT"}},"waveform":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAggODQoHBAEJHTAhDgAAAAAAAA4dIRIDAAAAAAAAAAAAAAAAAAAAAA==","viewOnce":false},"messageContextInfo":{"deviceListMetadata":{"senderKeyHash":"DK/EyaUp7VXSag==","senderTimestamp":"1732035332","senderAccountType":"E2EE","receiverAccountType":"E2EE","recipientKeyHash":"VEgMtMpwrVz/JA==","recipientTimestamp":"1732019837"},"deviceListMetadataVersion":2,"messageSecret":"ZGhAH4JzmRGkw30pkOS9QkO/JkeSzA+ygLFI5PjnHjA="},"base64":"base64"},"contextInfo":{"disappearingMode":{"initiator":"CHANGED_IN_CHAT"}},"messageType":"audioMessage","messageTimestamp":1732042073,"instanceId":"f890a78d-5497-486d-99f7-712111aee095","source":"web"},"destination":"https://webhook.site/bfcd9054-1a18-4cd2-bb31-890c8337ecdd","date_time":"2024-11-19T18:47:54.797Z","sender":"5585920013048@s.whatsapp.net","server_url":"http://localhost:80","apikey":null}`
const messageListExample1 = `{
  "event": "messages.upsert",
  "instance": "5d953186-2f59-11f0-b8c2-42004e494300",
  "data": {
    "key": {
      "remoteJid": "558594138387@s.whatsapp.net",
      "fromMe": false,
      "id": "3EB07AD954B951BE89C57C"
    },
    "pushName": "Ivo",
    "status": "DELIVERY_ACK",
    "message": {
      "messageContextInfo": {
        "deviceListMetadata": {
          "senderKeyHash": "UnKZt9oknJWeZQ==",
          "senderTimestamp": "1747103049",
          "senderAccountType": "E2EE",
          "receiverAccountType": "E2EE",
          "recipientKeyHash": "b8Grno02Q8WDVg==",
          "recipientTimestamp": "1747072178"
        },
        "deviceListMetadataVersion": 2
      },
      "listResponseMessage": {
        "title": "Sim",
        "listType": "SINGLE_SELECT",
        "singleSelectReply": {
          "selectedRowId": "1"
        },
        "contextInfo": {
          "stanzaId": "3EB085785EB2D70C4C75A57EF7F06CEC5BE30251",
          "participant": "554184168755@s.whatsapp.net",
          "quotedMessage": {
            "messageContextInfo": {},
            "listMessage": {
              "title": "Confirmar consulta",
              "description": "Confirme a consulta para hoje",
              "buttonText": "Confirmar",
              "listType": "SINGLE_SELECT",
              "sections": [
                {
                  "title": "Confirmação",
                  "rows": [
                    {
                      "title": "Sim",
                      "description": "Sim",
                      "rowId": "1"
                    },
                    {
                      "title": "Nao",
                      "description": "Nao",
                      "rowId": "2"
                    }
                  ]
                }
              ],
              "footerText": "Baixo texto"
            }
          }
        },
        "description": "Sim"
      }
    },
    "contextInfo": {
      "stanzaId": "3EB085785EB2D70C4C75A57EF7F06CEC5BE30251",
      "participant": "554184168755@s.whatsapp.net",
      "quotedMessage": {
        "messageContextInfo": {},
        "listMessage": {
          "title": "Confirmar consulta",
          "description": "Confirme a consulta para hoje",
          "buttonText": "Confirmar",
          "listType": "SINGLE_SELECT",
          "sections": [
            {
              "title": "Confirmação",
              "rows": [
                {
                  "title": "Sim",
                  "description": "Sim",
                  "rowId": "1"
                },
                {
                  "title": "Nao",
                  "description": "Nao",
                  "rowId": "2"
                }
              ]
            }
          ],
          "footerText": "Baixo texto"
        }
      }
    },
    "messageType": "listResponseMessage",
    "messageTimestamp": 1747254526,
    "instanceId": "243553d5-ebc5-4825-9f29-9b7b84afbeca",
    "source": "web"
  },
  "destination": "https://evolution-integration-592799294413.us-central1.run.app/webhook/5d953186-2f59-11f0-b8c2-42004e494300",
  "date_time": "2025-05-14T20:28:46.509Z",
  "sender": "554184168755@s.whatsapp.net",
  "server_url": "http://localhost:80",
  "apikey": null
}`
const contactUpsertExample1 = `{"event":"contacts.upsert","instance":"fdf3b942-1eaf-45e3-871a-0ab2c0e182b6","data":[{"remoteJid":"558586541111@s.whatsapp.net","pushName":"Beyonce Do Front end","profilePicUrl":null,"instanceId":"c44a02ad-1ba5-4177-8c57-0d359c6a9f32"}],"destination":"https://evolution-integration-592799294413.us-central1.run.app/webhook/fdf3b942-1eaf-45e3-871a-0ab2c0e182b6","date_time":"2025-05-30T16:13:20.288Z","sender":"558594138387@s.whatsapp.net","server_url":"http://localhost:80","apikey":null}`
