import {AbstractControl, FormGroup, ValidatorFn} from "@angular/forms";
import {Utils} from "../utils/utils";

export var EqualPasswordsValidator:any = {
  validate: function(firstField, secondField):ValidatorFn {
    return (c:FormGroup) => {
      let pass = c.controls && c.controls[firstField].value == c.controls[secondField].value;

      if (pass) {
        return null;
      } else {
        console.log('passwordsEqual fail');
        return {
          passwordsEqual: {
            valid: false
          }
        };
      }
    }
  }
};
