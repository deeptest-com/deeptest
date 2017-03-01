import {Injectable} from '@angular/core';

import {CONSTANT} from '../utils/constant';
import {RequestService} from './request';

import { TreeModel } from '../pages/components/ng2-tree/src/tree.types';

@Injectable()
export class TestcaseService {
    constructor(private _reqService: RequestService) { }
    _api_url = 'testcase/';

    query(query: TreeModel) {
        return this._reqService.post(this._api_url + 'query', query);
    }

    create(node: TreeModel) {
        let model = {id: node.id, value: node.value, type: node.type, pid: node.pid};
        return this._reqService.post(this._api_url + 'create', model);
    }

    move(node: TreeModel, prevParent: TreeModel, options: any) {
        let model = {id: node.id, pid: node.pid, prePid: prevParent.id};
        _.merge(model, options);
        return this._reqService.post(this._api_url + 'move', model);
    }

    rename(node: TreeModel) {
        let model = {id: node.id, value: node.value, type: node.type, pid: node.pid};
        return this._reqService.post(this._api_url + 'rename', model);
    }

    delete(node: TreeModel) {
        let model = {id: node.id};
        return this._reqService.post(this._api_url + 'delete', model);
    }
}

