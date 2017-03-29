import {AbstractControl, ValidatorFn} from "@angular/forms";

export var EmailValidator:any = {
  validate: function ():ValidatorFn {
    return (c:AbstractControl):{[key:string]:any} => {
      if (!c.value) {
        
        return null;
      }

      let EMAIL_REGEXP = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+.(.[a-zA-Z0-9_-])+/i;

      return EMAIL_REGEXP.test(c.value) ? null : {
        validate: {
          valid: false
        }
      };
    };
  }
};
