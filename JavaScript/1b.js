function pow3(n){
    return n*n*n;
}

var x = require('fs').readFileSync('/dev/stdin', 'utf8');
console.log(pow3(x));