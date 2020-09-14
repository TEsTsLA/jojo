class Test {
    constructor(name) {
        this.name = name;
    }
    say(){
        console.log(this.name);
    }
}

new Test('test_name')
.say();
