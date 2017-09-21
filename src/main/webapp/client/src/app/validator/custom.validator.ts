import {FormGroup, AbstractControl, ValidatorFn} from "@angular/forms";

import {Utils} from "../utils/utils";

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
  },

  compareDate: function (resultKey:string, startTime:string, endTime:string): ValidatorFn {
    return (c:FormGroup) => {
      let fail = false;

      let start = c.controls[startTime];
      let end = c.controls[endTime];

      if (!start.value || !end.value) {
        start.setErrors(null);
        end.setErrors(null);
        return null;
      }

      let startDate = Utils.dateStructToDate(start.value);
      let endDate = Utils.dateStructToDate(end.value);

      let pass = startDate <= endDate;

      if (!pass) {
        console.log('compareDate fail');

        start.setErrors({});
        end.setErrors({});

        return {
            resultKey: {
              valid: false
            }
          };
      } else {
        start.setErrors(null);
        end.setErrors(null);

        return null;
      }
    }
  }
};
