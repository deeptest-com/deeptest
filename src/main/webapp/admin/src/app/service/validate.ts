import {AbstractControl, FormGroup, ValidatorFn} from "@angular/forms";
import {Utils} from "../utils/utils";

export var PATTERN:any = {
  Email: '^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,5}$',
  Date: '^[0-9]{4}-[0-9]{2}-[0-9]{2}$',
  Time: '^[0-9]{2}\:[0-9]{2}$'
};

export var Validate:any = {
  emailValidator: function ():ValidatorFn {
    return (control:AbstractControl):{[key:string]:any} => {

      const text = control.value;
      const reg = new RegExp(PATTERN.Email, 'i');
      let pass = reg.test(text);

      console.log('emailValidator=' + pass);

      return pass ? null : {'emailValidator': {text}};
    };
  },

  dateValidator: function ():ValidatorFn {
    return (control:AbstractControl):{[key:string]:any} => {

      const text = control.value;
      const reg = new RegExp(PATTERN.Date, 'i');
      let pass = reg.test(text);

      console.log('dateValidator=' + pass);

      return pass ? null : {'dateValidator': {text}};
    };
  },
  timeValidator: function ():ValidatorFn {
    return (control:AbstractControl):{[key:string]:any} => {

      const text = control.value;
      const reg = new RegExp(PATTERN.Time, 'i');
      let pass = reg.test(text);

      console.log('timeValidator=' + pass);

      return pass ? null : {'timeValidator': {text}};
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
            console.log(resultKey + '=true');
            fail = true;

            result[resultKey] = true;
          }
        }
      }

      if (fail) {
        return result;
      } else {
        return null;
      }
    }
  },

  genValidateInfo: function (form:any, validateMsg:any, customValidators:string[]) {
    if (!form) {
      return;
    }

    let errors = [];
    for (const field in form.controls) {
      const control = form.controls[field];

      if (control && control.dirty && !control.valid) {
        const messages = validateMsg[field];
        for (const key in control.errors) {
          errors.push(messages[key]);
        }
      }
    }
    // deal with custom validators' msg
    for (const idx in customValidators) {
      let validator = customValidators[idx];

      if (form.errors && form.errors[validator]) {
        errors.push(validateMsg[validator]);
      }
    }

    return errors;
  },

}

//@Directive({
//  selector: '[myValidate]',
//  providers: [{provide: NG_VALIDATORS, useExisting: MyValidatorDirective, multi: true}]
//})
//export class MyValidatorDirective implements Validator, OnChanges {
//  @Input() myValidate: string;
//  private valFn = Validators.nullValidator;
//
//  ngOnChanges(changes: SimpleChanges): void {
//    const change = changes['myValidate'];
//    if (change) {
//      const reg: string = change.currentValue;
//
//      this.valFn = myValidator(reg);
//    } else {
//      this.valFn = Validators.nullValidator;
//    }
//  }
//
//  validate(control: AbstractControl): {[key: string]: any} {
//
//    return this.valFn(control);
//  }
//}
