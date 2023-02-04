
function arrayToList(array){
    if(array.length==0){
        return {}
    } else {
        newArray = array.slice(1,array.length);
        console.log(newArray);
        return {value:array[0],rest:arrayToList(newArray)};
    }
}

liste = arrayToList([1,2,3]);
console.log(liste);

function deepEqual(obj1, obj2){
    if(obj1==null && obj2==null){
        return true;
    } else if(obj1==null || obj2==null){
        return false;
    } else if(typeof(obj1)!="object"){
        return obj1 === obj2;
    } else {
        keys1 = Object.keys(obj1);
        keys2 = Object.keys(obj2);
        if(keys1.length!=keys2.length){
            return false
        } else if(keys1.length==0){
            return true
        } else {
            for(let i=0; i<keys1.length; i++){
                if(keys1[i]!=keys2[i]){
                    return false
                }
                if(!deepEqual(obj1[keys1[i]],obj2[keys2[i]])){
                    return false
                }
            }
        }
    }
    return true
}

