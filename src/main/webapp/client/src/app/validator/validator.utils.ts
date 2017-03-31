import {AbstractControl} from '@angular/forms';

export class ValidatorUtils {
  public static genMsg(form:any, validateMsg: any, customValidators:string[]) {
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
  }

}
