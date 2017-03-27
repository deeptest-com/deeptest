import {Pipe, PipeTransform} from "@angular/core";

@Pipe({name: 'mapKeys'})
export class KeysPipe implements PipeTransform {
  transform(array:Array<any>, ignore:string[]):any {
    if (!_.isArray(ignore)) {
      ignore = _.union([], [ignore]);
    }

    let keys:Array<any> = [];
    for (let index in array) {
      for (let key in array[index]) {
        if (!_.includes(ignore, key)) {
          keys.push({key: key, value: array[index][key]});
        }
      }
    }
    return keys;
  }
}
