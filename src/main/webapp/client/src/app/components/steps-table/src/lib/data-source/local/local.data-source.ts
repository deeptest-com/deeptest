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

  create(element: any, curr: any): Promise<any> {
    let index = this.data.indexOf(curr);
    this.data.splice(index + 1, 0, element)

    return super.create(element, curr);
  }

  remove(element: any): Promise<any> {
    this.data = this.data.filter(el => el !== element);

    return super.remove(element);
  }

  save(element: any, values: any): Promise<any> {
    return new Promise((resolve, reject) => {
      this.find(element).then((found) => {
        found = deepExtend(found, values);
        super.save(found, values).then(resolve).catch(reject);
      }).catch(reject);
    });
  }

  find(element: any): Promise<any> {
    console.log('find', element, this.data);

    let found = this.data.find(el => el === element);
    // if (!found && !element.ordr) { // 新对象
    //   found = element;
    // }

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
