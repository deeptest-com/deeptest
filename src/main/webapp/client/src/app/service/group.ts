import {Injectable} from "@angular/core";

import { Cookie } from 'ng2-cookies/ng2-cookies';
import {GlobalState} from '../global.state';

import { CONSTANT } from '../utils/constant';
import { RouteService } from './route';
import {RequestService} from "./request";

@Injectable()
export class GroupService {
  constructor(private _state:GlobalState, private _reqService:RequestService, private routeService: RouteService) {
  }

  _list = 'group/list';
  _get = 'group/get';
  _save = 'group/save';


  list() {
    return this._reqService.post(this._list, {});
  }

  get(profile:any) {
    return this._reqService.post(this._get, profile);
  }
  save(model:any) {
    return this._reqService.post(this._save, model);
  }
}

