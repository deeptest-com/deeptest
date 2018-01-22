import {FormGroup, ValidatorFn} from "@angular/forms";

export var PasswordsEqualValidator:any = {
  validate: function (resultKey, firstField, secondField):ValidatorFn {
    return (c:FormGroup) => {
      let result = {};
      let fail = false;

      let pass = c.controls && c.controls[firstField].value == c.controls[secondField].value;

      if (!pass) {
        // console.log('passwordsEqual fail');
        fail = true;
        result[resultKey] = {
          valid: false
        };
      }
      if (fail) {
        return result;
      } else {
        return null;
      }
    }
  }
}
