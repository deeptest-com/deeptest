import {AbstractControl, FormGroup, ValidatorFn} from "@angular/forms";
import {Utils} from "../utils/utils";

export var DateTimeValidator:any = {
   validateDate: function ():ValidatorFn {
    return (c:AbstractControl):{[key:string]:any} => {
        let DATE_REGEXP = /^[0-9]{4}-[0-9]{2}-[0-9]{2}$/i

        return DATE_REGEXP.test(c.value) ? null : {
          validateDate: {
            valid: false
          }
        };
    };
  },
  validateTime: function ():ValidatorFn {
    return (c:AbstractControl):{[key:string]:any} => {
        let TIME_REGEXP = /^[0-9]{2}\:[0-9]{2}$/i

        return TIME_REGEXP.test(c.value) ? null : {
          validateTime: {
            valid: false
          }
        };
    };
  },
  compareDatetime: function (arr0:string[][]):ValidatorFn {
    return (group:FormGroup) => {
      let result = {};
      let fail = false;
      for (var i = 0; i < arr0.length; i++) {
        let arr = arr0[i];

        let resultKey = arr[0];
        let date1 = arr[1];
        let time1 = arr[2];
        let date2 = arr[3];
        let time2 = arr[4];

        if (group.controls[date1] && group.controls[time1] && group.controls[date2] && group.controls[time2]) {
          let startValue = group.controls[date1].value + ' ' + group.controls[time1].value;
          let endValue = group.controls[date2].value + ' ' + group.controls[time2].value;

          let startTm = Utils.strToTimestamp(startValue);
          let endTm = Utils.strToTimestamp(endValue);

          if (startTm >= endTm) {
            console.log(resultKey + ' fail');
            fail = true;
            result[resultKey] = {
              valid: false
            };
          }
        }
      }

      if (fail) {
        return result;
      } else {
        return null;
      }
    }
  }
}
