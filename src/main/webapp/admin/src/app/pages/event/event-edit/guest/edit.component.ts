import { Component,ViewEncapsulation, Pipe, OnInit, AfterViewInit, ViewChild } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Router, ActivatedRoute, Params } from '@angular/router';

import { DropdownModule} from 'ng2-bootstrap/ng2-bootstrap';

import {Validate} from '../../../../service/validate';
import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import { DatetimePickerService } from '../../../../service/datetime-picker';
import { EventService } from '../../../../service/event';

require('bootstrap-datepicker');
require('bootstrap-timepicker');

declare var jQuery;

@Component({
  selector: 'event-edit-guest',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./edit.scss'),
    require('bootstrap-datepicker/dist/css/bootstrap-datepicker3.css'),
    require('bootstrap-timepicker/css/bootstrap-timepicker.css')],

  template: require('./edit.html')
})
export class EventEditGuest implements OnInit, AfterViewInit {
  eventId: number;
  model: any = { signBefore: 3 };
//  active: boolean = true;
  eventForm: any;
  tabModel: string = 'guest';
  datePickers: string[] = ['startDate', 'endDate', 'registerStartDate', 'registerEndDate'];
  timePickers: string[] = ['startTime', 'endTime','registerStartTime', 'registerEndTime'];

  constructor(private _router: Router, private _route: ActivatedRoute, private fb: FormBuilder,
              private _datetimePickerService: DatetimePickerService, private _eventService: EventService) {

    let that = this;
  }

  ngOnInit() {
    let that = this;

    that.buildForm();
    console.log(this._route.params);
    this._route.params.forEach((params: Params) => {
      that.eventId = +params['id'];

      console.log(that.eventId);
    });

    if (that.eventId) {
        that.loadData();
    }
  }

  ngAfterViewInit() {
    let that = this;

    that.initForm(false);
  }

  edit(tabModel):void {
    let that = this;

    console.log(tabModel , that.eventId);
    that._router.navigateByUrl('/pages/event/edit/' + that.eventId + '/' + tabModel);
  }

  onSubmit():void {
    let that = this;

//    that.active = false;
//    setTimeout(() => that.active = true, 0);

    console.log(that.model);

    that._eventService.save(that.model).subscribe((json:any) => {
        console.log(json);
        if (json.code = 1) {
            that._router.navigateByUrl("/pages/event/list");
        }
    });
  }
  back() {
    let that = this;

    that._router.navigateByUrl("/pages/event/list");
  }
  goto(tabModel) {
    let that = this;

    that._router.navigateByUrl('/pages/event/edit/' + that.eventId + '/' + tabModel);
  }

  loadData() {
   let that = this;

   that._eventService.get(that.eventId).subscribe((json:any) => {
      that.model = json.event;

      that.initForm(true);

   });
  }

  initForm(dataLoaded): void {
    let that = this;

    jQuery.each(that.datePickers, function( index, value ) {
      that._datetimePickerService.genDatePicker(that.model, value);
    });
    jQuery.each(that.timePickers, function( index, value ) {
      that._datetimePickerService.genTimePicker(that.model, value);
    });
  }

  buildForm(): void {
    let that = this;
    this.eventForm = this.fb.group(
        {
          'subject': [that.model.email, [Validators.required]],

          'startDate': [that.model.startDate, [Validators.required, Validate.dateValidator()]],
          'startTime': [that.model.startTime, [Validators.required, Validate.timeValidator()]],
          'endDate': [that.model.endDate, [Validators.required, Validate.dateValidator()]],
          'endTime': [that.model.endTime, [Validators.required, Validate.timeValidator()]],

          'registerStartDate': [that.model.startDate, [Validators.required, Validate.dateValidator()]],
          'registerStartTime': [that.model.startTime, [Validators.required, Validate.timeValidator()]],
          'registerEndDate': [that.model.endDate, [Validators.required, Validate.dateValidator()]],
          'registerEndTime': [that.model.endTime, [Validators.required, Validate.timeValidator()]],

          'signBefore': [that.model.signBefore, [Validators.required]],
          'address': [that.model.address, [Validators.required]],
          'phone': [that.model.phone, [Validators.required]],
          'email': [that.model.email, [Validators.required, Validate.emailValidator()]],
          'website': [that.model.website, []]
        }, {}
    );

    this.eventForm.valueChanges.subscribe(data => this.onValueChanged(data));
    this.onValueChanged();
  }
  onValueChanged(data?: any) {
    let that = this;
    if (!that.eventForm) { return; }

    // that.formErrors = Validate.genEventValidateInfo(that.eventForm, that.validationMessages);

  }

  formErrors = [];
  validationMessages = {
    'title': {
      'required':      '会议名称不能为空'
    },
    'startDate': {
      'required':      '开始日期不能为空',
      'dateValidator': '开始日期格式不正确'
    },
    'startTime': {
      'required':      '开始时间不能为空',
      'timeValidator': '开始日期格式不正确'
    },
    'endDate': {
      'required':      '结束日期不能为空'
    },
    'endTime': {
      'required':      '结束时间不能为空',
      'timeValidator': '结束时间格式不正确'
    },
    'registerStartDate': {
      'required':      '注册开始日期不能为空',
      'dateValidator': '注册开始日期格式不正确'
    },
    'registerStartTime': {
      'required':      '注册开始时间不能为空',
      'timeValidator': '注册开始时间格式不正确'
    },
    'registerEndDate': {
      'required':      '注册结束日期不能为空',
      'dateValidator': '注册结束日期格式不正确'
    },
    'registerEndTime': {
      'required':      '注册结束时间不能为空',
      'timeValidator':  '注册结束时间格式不正确',
    },
    'signBefore': {
      'required':      '提前注册天数不能为空'
    },
    'address': {
      'required':      '地址不能为空'
    },
    'phone': {
      'required':      '电话不能为空'
    },
    'email': {
      'required':      '邮件不能为空',
      'emailValidator': '邮件格式错误'
    },
    'website': {'required':      '邮件不能为空'},

    'eventTimeCompareValidator': '会议结束时间不能早于开始时间',
    'registerTimeCompareValidator': '报名结束时间不能早于开始时间'
  };

}
