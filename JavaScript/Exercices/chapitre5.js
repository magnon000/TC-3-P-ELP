console.log("helo")


//flattening exercice
let arrays = [[1, 2, 3], [4, 5], [6]];
function flatten(arrayList) {
    return arrayList.reduce((a,b) => a.concat(b)) // concat pour joindre deux arrays (pas push)
}
console.log(flatten(arrays))


//decreasing function
function loop(init, testFunc, updateFunc, userFunction) {
    ////version récursive
    // if (testFunc(init)) {
    //     userFunction(init)
    //     loop(updateFunc(init), testFunc, updateFunc, userFunc)
    // }
    //version itérative (meilleure il semble)
    while(testFunc(init)) {
        userFunction(init)
        init = updateFunc(init)
    }
}
loop(3, n => n > 0, n => n - 1, console.log);

//every (juste la version avec some)
function every(array, test){
    return !array.some(element => !test(element))
}
console.log(every([1, 3, 5], n => n < 10));
console.log(every([2, 4, 16], n => n < 10));
console.log(every([], n => n < 10));


//dominant writing (en utilisant les fct définies dans la leçon)
// function dominantDirection(text) {
//     let scripts = countBy(text, char => {
//       let script = characterScript(char.codePointAt(0));
//       return script ? script.direction : "none";
//     }).filter(({direction}) => direction != "none");
  
//     let total = scripts.reduce((n, {count}) => n + count, 0);
//     if (total == 0) return "No scripts found";
  
//     return scripts.reduce((a,b) => {
//       return a.count < b.count ? b : a
//     });
// }
// console.log(dominantDirection("Hey, مساء الخير"));