import { LocalSorter } from './local.sorter';
import { LocalFilter } from './local.filter';
import { LocalPager } from './local.pager';
import { DataSource } from '../data-source';
import { deepExtend } from '../../helpers';

export class LocalDataSource extends DataSource {

  protected data: Array<any> = [];
  protected sortConf: Array<any> = [];
  protected filterConf: any = {
    filters: [],
    andOperator: true,
  };
  protected pagingConf: any = {};

  constructor(data: Array<any> = []) {
    super();

    this.data = data;
  }

  load(data: Array<any>): Promise<any> {
    this.data = data;

    return super.load(data);
  }

  up(currElem: any): Promise<any> {

    let index = this.data.indexOf(currElem);
    let preIndex = index - 1;
    let preElem = this.data[preIndex];
    let currOrder = currElem.ordr;

    // 交换ordr
    currElem.ordr = preElem.ordr;
    preElem.ordr = currOrder;

    // 交换位置
    this.data.splice(index, 1);
    this.data.splice(preIndex, 0, currElem);

    return super.up(currElem);
  }
  down(currElem: any): Promise<any> {
    let index = this.data.indexOf(currElem);
    let nextIndex = index + 1;
    let nextElem = this.data[nextIndex];
    let currOrder = currElem.ordr;

    // 交换ordr
    currElem.ordr = nextElem.ordr;
    nextElem.ordr = currOrder;

    // 交换位置
    this.data.splice(nextIndex, 1);
    this.data.splice(index, 0, nextElem);

    return super.up(currElem);
  }

  create(element: any, curr: any): Promise<any> {
    let index = 0;
    if (!curr && this.data.length > 0) {
      curr = this.data[this.data.length - 1];
      index = this.data.indexOf(curr);
    } else if(!curr && this.data.length == 0) {
      index = 0;
    } else {
      index = this.data.indexOf(curr);
    }

    element.ordr = curr?curr.ordr:1;
    this.data.splice(index + 1, 0, element);

    this.data.forEach(function(elem, indx, arr) {
      if (indx > index) {
        elem.ordr += 1;
      }
    });
    return super.create(element, curr);
  }

  delete(element: any): Promise<any> {
    let index = this.data.indexOf(element);

    this.data = this.data.filter(el => el !== element);

    this.data.forEach(function(elem, indx, arr) {
      if (indx >= index) {
        elem.ordr -= 1;
      }
    });

    return super.delete(element);
  }

  save(element: any, values: any): Promise<any> {
    console.log(element, values);

    return new Promise((resolve, reject) => {
      this.find(element).then((found) => {
        found = deepExtend(found, values);
        super.save(found, values).then(resolve).catch(reject);
      }).catch(reject);
    });
  }

  find(element: any): Promise<any> {
    let found = this.data.find(el => el === element);
    if (found) {
      return Promise.resolve(found);
    }

    return Promise.reject(new Error('Element was not found in the dataset'));
  }

  getElements(): Promise<any> {
    const data = this.data.slice(0);
    return Promise.resolve(data);
  }

  getAll(): Promise<any> {
    const data = this.data.slice(0);
    return Promise.resolve(data);
  }

  empty(): Promise<any> {
    this.data = [];

    return super.empty();
  }

}
