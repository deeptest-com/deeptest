import * as _ from "lodash";

import {Injectable} from "@angular/core";

import {CONSTANT} from "../utils/constant";

@Injectable()
export class PrivilegeService {
  hasPrivilege(privs: string) {
    let ret = true;

    let arr = privs.split(',');
    for (let i = 0; i < arr.length; i++) {
      if (!CONSTANT.PROFILE.projectPrivilege[arr[i]]) {
        ret = false;
      }
    }

    return ret;
  }
}


