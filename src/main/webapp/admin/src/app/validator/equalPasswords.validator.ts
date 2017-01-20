import {AbstractControl, FormGroup, ValidatorFn} from "@angular/forms";
import {Utils} from "../utils/utils";

export var EqualPasswordsValidator:any = {
  validate: function(firstField, secondField):ValidatorFn {
    return (c:FormGroup) => {
      return (c.controls && c.controls[firstField].value == c.controls[secondField].value) ? null : {
        passwordsEqual: {
          valid: false
        }
      };
    }
  }
};
