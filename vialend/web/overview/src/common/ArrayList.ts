export interface List<E> {
    /**  LinkedList + ArrayList 通用接口 */
    add(element: E):void
    add(index: number, element: E):void // 添加对象
    remove(index: number): E // 指定索引删除对象
    get(index: number): E // 通过索引获取对象，for循环获取必备
    set(index:number, element:E):void
    size(): Number // 获取list长度
    isEmpty():boolean
    indexOf(obj:E):number
    contains(obj:E):boolean
    clear():void
}

export default class ArrayList<E> implements List<E> {
    // 定义属性集合
    elementData: E[] = [];

    _getData():E[] {
      return this.elementData
    }

    size():number {
      return this.elementData.length
    }

    isEmpty():boolean {
      return this.size() === 0
    }

    indexOf(obj:E):number {
      let i
      const data = this._getData()
      const length = data.length
      for (i = 0; i < length; i++) {
        if (obj === data[i]) {
          return i
        }
      }
      return -1
    }

    contains(obj:E):boolean {
      return this.indexOf(obj) > -1
    }

    get(index:number):E {
      return this._getData()[index]
    }

    set(index:number, element:E):void {
      this._getData()[index] = element
    }

    add(element:E):void
    add(index:number, element:E):void
    add(arg0?: number | E, arg1?: E):void {
      if (typeof arg0 === 'number' && arg1 !== undefined) {
        this.set(arg0, arg1)
      } else {
        if (typeof arg0 === 'object') { this._getData().push(arg0) }
      }
    }

    remove(index: number): E {
      const oldValue = this._getData()[index]
      this._getData().splice(index, 1)
      return oldValue
    }

    clear():void {
      this._getData().length = 0
    }
}
