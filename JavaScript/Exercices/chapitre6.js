
//vector
class Vec{
    constructor(x,y){
        this.x = x;
        this.y = y;
    }

    plus(v2){
        return new Vec(this.x+v2.x,this.y+v2.y)
    }

    minus(v2){
        return new Vec(this.x-v2.x,this.y-v2.y)
    }

    get length(){
        return (this.x**2+this.y**2)**0.5
    }
}

