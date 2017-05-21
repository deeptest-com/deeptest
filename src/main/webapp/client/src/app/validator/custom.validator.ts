import {FormGroup, AbstractControl, ValidatorFn} from "@angular/forms";

export var CustomValidator:any = {
  validate: function (...params: string[]):ValidatorFn {
    return (c:FormGroup):{[key:string]:any} => {

      if (!c.parent) {
        return null;
      }

      let result = {};
      let pass = true;

      let name = params[0];
      let msgKey = params[1];
      if (name === 'required_if_other_is') {
        let thisField = params[2];
        let otherFiled = params[3];
        let value = params[4];
        pass = this.required_if_other_is(c.parent, thisField, otherFiled, value);
      }

      if (!pass) {
        result[msgKey] = {
          valid: false
        };
        return result;
      } else {
        return null;
      }
    };
  },

  required_if_other_is: function (group:FormGroup, thisField:string, otherFiled: string, value: string): boolean {
    if (group.controls && group.controls[otherFiled].value === value) {
      let val = group.controls[thisField].value;
      if (!val || val === '') {
        return false;
      }
    }
    return true;
  }
};
