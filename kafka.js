var kafka = require('kafka-node');
const { exec } = require('child_process');

var nomeDoTopico = 'arquivosNaoProcessados';
var Consumer = kafka.Consumer;
var KeyedMessage = kafka.KeyedMessage;
var km = new KeyedMessage('key', 'message');
var client = new kafka.KafkaClient();
var consumir = new Consumer(client, [{ topic: nomeDoTopico, partition: 0 }], { autoCommit: false });

consumir.on('message', function (message) {
    executarOCR(message);
});

function executarOCR(message) {
    exec(`go run main.go ${message.value}`, (erro, texto) => {
        if (erro)
            return;
        enviarParaFila(texto);
    });
}

function enviarParaFila(texto) {
    let nomeDoTopico_p = 'arquivosProcessados';
    let Producer = kafka.Producer;
    let producer = new Producer(client);
    let publicar = [{ topic: nomeDoTopico_p, messages: texto, partition: 0 }];
    producer.send(publicar, () => {
        console.log(`Enviou para fila ${nomeDoTopico_p}`)
    });
}
