const { exec } = require('child_process');

exec('go run main.go', (err, stdout) => {
  if (err) {
      return;
  }

  
  console.log(stdout);

  //enviar para a fila
});
console.log("asasdasda");