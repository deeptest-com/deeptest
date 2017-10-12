import { Input, Component, OnInit, EventEmitter, Output, Inject, OnChanges, SimpleChanges } from '@angular/core';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';

import { CONSTANT } from '../../../utils/constant';
import { Utils, Deferred } from '../../../utils/utils';

import { FieldShowService } from './field-show.service';

@Component({
  selector: 'field-show',
  templateUrl: './field-show.html',
  styleUrls: ['./styles.scss'],
  providers: [FieldShowService]
})
export class FieldShowComponent implements OnInit {

  @Input()
  public model: any;
  @Input()
  public prop: string;
  @Input()
  public type: string;
  @Input()
  public list: any[];
  @Input()
  public label: string;

  @Output() onSave = new EventEmitter<any>();

  @Input()
  public form: any;

  public status: string = 'view';
  public temp: string;

  public constructor(@Inject(FieldShowService) private fieldShowService: FieldShowService) {


  }

  public ngOnInit(): void {
    let control: FormControl = new FormControl(this.prop, Validators.required);
    this.form.addControl(this.prop, control);
  }

  edit(event: any) {
    event.preventDefault();
    event.stopPropagation();

    this.status = 'edit';
    this.temp = this.model[this.prop];
  }

  save(event: any) {
    event.preventDefault();
    event.stopPropagation();

    const deferred = new Deferred();
    deferred.promise.then((data) => {
      this.status = 'view';
    }).catch((err) => {console.log('err', err);});

    if (this.model[this.prop] != this.temp) {
      this.onSave.emit({deferred: deferred, data: {prop: this.prop, value: this.model[this.prop]}});
    }
  }
  cancel(event: any) {
    event.preventDefault();
    event.stopPropagation();

    this.status = 'view';
    this.model[this.prop] = this.temp;
  }

}
