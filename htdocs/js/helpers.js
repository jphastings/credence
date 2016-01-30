var Credence = {
  detectLocalInstance: function(onInstanceFound, onNoneFound) {
    if (typeof onInstanceFound != 'function') {
      throw 'onInstanceFound must be a function';
    }
    onNoneFound = onNoneFound || function(){};

    var server = 'http://127.0.0.1:8808';
    fetch(server + '/ping').then(
      function(response) {  
        switch(response.status) {
          case 200:
            onInstanceFound(server);
            break;
          default:
            onNoneFound();
        }
      }  
    ).catch(onNoneFound);
  }
}