import {Pipe, PipeTransform} from '@angular/core';

import {CONSTANT} from '../utils/constant';

@Pipe({name: 'mapKeys'})
export class KeysPipe implements PipeTransform {
    transform(array: Array<any>, args:string[]): any {
       
        let keys:Array<any> = [];
        for (let index in array) {
            for (let key in array[index]) {
              keys.push({key: key, value: array[index][key]});
            }
        }
        return keys;
    }
}
