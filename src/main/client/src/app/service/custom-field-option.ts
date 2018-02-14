import * as _ from 'lodash';

import {Injectable} from "@angular/core";

import {RequestService} from "./request";

@Injectable()
export class CustomFieldOptionService {
  constructor(private _reqService: RequestService) { }
  _api_url = 'custom_field_option/';

  save(model: any, fieldId: number) {
    model.fieldId = fieldId;
    return this._reqService.post(this._api_url + 'save', {model: model});
  }

  delete(id: number, fieldId: number) {
    let model = {id: id, fieldId: fieldId};
    return this._reqService.post(this._api_url + 'delete', model);
  }

  changeOrder(id: number, act: string, fieldId: number) {
    let model = {id: id, act: act, fieldId: fieldId};
    return this._reqService.post(this._api_url + 'changeOrder', model);
  }

}

