import { Input, Component, OnInit, EventEmitter, Output, Inject, OnChanges, SimpleChanges } from '@angular/core';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';

import { CONSTANT } from '../../../utils/constant';
import { Utils } from '../../../utils/utils';

import { CustomFieldService } from './custom-field.service';

@Component({
  selector: 'custom-field',
  templateUrl: './custom-field.html',
  styleUrls: ['./styles.scss'],
  providers: [CustomFieldService]
})
export class CustomFieldComponent implements OnInit, OnChanges {

  @Input()
  public model: any;
  @Input()
  public field: any;
  @Input()
  public validateMsg: any;

  @Input()
  public form: any;

  @Output()
  public fieldChanged: EventEmitter<any> = new EventEmitter();

  public constructor(@Inject(CustomFieldService) private customFieldService: CustomFieldService) {

  }

  public ngOnChanges(changes: SimpleChanges): void {

  }

  public ngOnInit(): void {
    // console.log('===field===', this.model, this.field);

    this.validateMsg[this.field.myColumn] = {};
    let control: FormControl = new FormControl(this.field.myColumn);
    if (this.field.required) {
      control.setValidators([Validators.required]);
      this.validateMsg[this.field.myColumn]['required'] = this.field.label + '不能为空';
    }

    this.form.addControl(this.field.myColumn, control);
  }

  onEditorKeyup(event: any) {
    // console.log('===', this.model, this.field.myColumn, event);
    this.model[this.field.myColumn] = event;
  }
}
