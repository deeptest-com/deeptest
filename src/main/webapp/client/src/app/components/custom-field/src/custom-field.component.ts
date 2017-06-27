import { Input, Component, OnInit, EventEmitter, Output, Inject, OnChanges, SimpleChanges } from '@angular/core';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';

import { CONSTANT } from '../../../utils/constant';
import { Utils } from '../../../utils/utils';

import { Prop, Field } from './field.prop';
import { FieldChangedEvent } from './field.events';

import { CustomFieldService } from './custom-field.service';

@Component({
  selector: 'custom-field',
  templateUrl: './custom-field.html',
  styleUrls: ['./styles.scss'],
  providers: [CustomFieldService]
})
export class CustomFieldComponent implements OnInit, OnChanges {

  @Input()
  public prop: Prop;
  @Input()
  public fields: Field[];

  @Input()
  public form: any;

  @Output()
  public fieldChanged: EventEmitter<any> = new EventEmitter();

  public constructor(@Inject(CustomFieldService) private customFieldService: CustomFieldService) {


  }

  public ngOnChanges(changes: SimpleChanges): void {
    this.fields = CONSTANT.CUSTOM_FIELD_FOR_PROJECT;
    console.log(this.fields);
  }

  public ngOnInit(): void {
    let control: FormControl = new FormControl(this.prop.code, Validators.required);
    this.form.addControl(this.prop.code, control);
  }
}
