"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
class ArrayList {
    constructor() {
        // 定义属性集合
        this.elementData = [];
    }
    _getData() {
        return this.elementData;
    }
    size() {
        return this.elementData.length;
    }
    isEmpty() {
        return this.size() === 0;
    }
    indexOf(obj) {
        let i;
        const data = this._getData();
        const length = data.length;
        for (i = 0; i < length; i++) {
            if (obj === data[i]) {
                return i;
            }
        }
        return -1;
    }
    contains(obj) {
        return this.indexOf(obj) > -1;
    }
    get(index) {
        return this._getData()[index];
    }
    set(index, element) {
        this._getData()[index] = element;
    }
    add(arg0, arg1) {
        if (typeof arg0 === 'number' && arg1 !== undefined) {
            this.set(arg0, arg1);
        }
        else {
            if (typeof arg0 === 'object') {
                this._getData().push(arg0);
            }
        }
    }
    remove(index) {
        const oldValue = this._getData()[index];
        this._getData().splice(index, 1);
        return oldValue;
    }
    clear() {
        this._getData().length = 0;
    }
}
exports.default = ArrayList;
