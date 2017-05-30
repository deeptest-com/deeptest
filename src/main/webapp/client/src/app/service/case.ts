import * as _ from 'lodash';

import {Injectable} from '@angular/core';

import {CONSTANT} from '../utils/constant';
import {RequestService} from './request';

import { TreeModel } from '../components/ng2-tree';

@Injectable()
export class CaseService {
    constructor(private _reqService: RequestService) { }
    _api_url = 'case/';

    query(query: TreeModel) {
        return this._reqService.post(this._api_url + 'query', query);
    }

    get(id: number) {
      let model = {id: id};
      return this._reqService.post(this._api_url + 'get', id);
    }

    save(model: number) {
      return this._reqService.post(this._api_url + 'save', model);
    }

    create(node: TreeModel) {
        let model = {id: node.id, value: node.value, type: node.type, pid: node.pid};
        return this._reqService.post(this._api_url + 'create', model);
    }

    move(target: TreeModel, src: TreeModel, options: any) {
        let model;
      if (options.mode === 'inner') {
        model = {id: src.id, newPid: target.id, prePid: src.pid};
      } else {
        model = {id: src.id, newPid: target.pid, prePid: src.pid};
      }
        _.merge(model, options);
        return this._reqService.post(this._api_url + 'move', model);
    }

    rename(node: TreeModel) {
      console.log('rename');
        let model = {id: node.id, value: node.value, type: node.type, pid: node.pid};
        return this._reqService.post(this._api_url + 'rename', model);
    }

    delete(node: TreeModel) {
        let model = {id: node.id};
        return this._reqService.post(this._api_url + 'delete', model);
    }
}


