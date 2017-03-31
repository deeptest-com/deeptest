import {AbstractControl, FormGroup, ValidatorFn} from "@angular/forms";
import {Utils} from "../utils/utils";

export var PhoneValidator:any = {
  validate: function ():ValidatorFn {
    return (c:AbstractControl):{[key:string]:any} => {
      if (!c.value) {
        return null;
      }
      let REGEXP = /^1[0-9]{10}$/i;

      return REGEXP.test(c.value) ? null : {
        validate: {
          valid: false
        }
      };
    };
  }
};
