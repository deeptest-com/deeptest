import { Component,ViewEncapsulation, Pipe, OnInit, AfterViewInit, ViewChild } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Router, ActivatedRoute, Params } from '@angular/router';

import { DropdownModule} from 'ng2-bootstrap/ng2-bootstrap';

import { CONSTANT } from '../../../../utils/constant';
import { Utils } from '../../../../utils/utils';
import {ValidatorUtils, EmailValidator, DateTimeValidator} from '../../../../validator';

import { RouteService } from '../../../../service/route';
import { DatetimePickerService } from '../../../../service/datetime-picker';

import { EventService } from '../../../../service/event';

require('bootstrap-datepicker');
require('bootstrap-timepicker');

declare var jQuery;

@Component({
  selector: 'event-edit-property',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./edit.scss'),
    require('bootstrap-datepicker/dist/css/bootstrap-datepicker3.css'),
    require('bootstrap-timepicker/css/bootstrap-timepicker.css')],

  template: require('./edit.html')
})
export class EventEditProperty implements OnInit, AfterViewInit {
  eventId: number;
  model: any = { signBefore: 3};
  eventForm: any;
  tabModel: string = 'property';
  needCreate:boolean = false;

  datePickers: string[] = ['startDate', 'endDate', 'registerStartDate', 'registerEndDate'];
  timePickers: string[] = ['startTime', 'endTime','registerStartTime', 'registerEndTime'];

  constructor(private _routeService: RouteService, private _route: ActivatedRoute, private fb: FormBuilder,
              private _datetimePickerService: DatetimePickerService, private _eventService: EventService) {

    let that = this;
  }

  ngOnInit() {
    let that = this;

    that.buildForm();
    this._route.params.forEach((params: Params) => {
      that.eventId = +params['id'];
    });

    if (that.eventId) {
        that.loadData();
    }
  }

  ngAfterViewInit() {
    let that = this;

    that.initForm();
  }

  onSubmit():void {
    let that = this;

    that.model.status = undefined;
    Utils.dateCombine(that.model, 'startDate', 'startTime', 'startDatetime');
    Utils.dateCombine(that.model, 'endDate', 'endTime', 'endDatetime');
    Utils.dateCombine(that.model, 'registerStartDate', 'registerStartTime', 'registerStartDatetime');
    Utils.dateCombine(that.model, 'registerEndDate', 'registerEndTime', 'registerEndDatetime');

    that._eventService.save(that.model).subscribe((json:any) => {
        if (json.code = 1) {
          that._routeService.navTo("/pages/event/list");
        }
    });
  }

  goto($event) {
    let that = this;

    that._routeService.navTo('/pages/event/edit/' + that.eventId + '/' + $event.tabModel);
  }
  loadData() {
   let that = this;

   that._eventService.get(that.eventId).subscribe((json:any) => {
      that.model = json.event;

     Utils.dateDivide(that.model, 'startDate', 'startTime', 'startDatetime');
     Utils.dateDivide(that.model, 'endDate', 'endTime', 'endDatetime');
     Utils.dateDivide(that.model, 'registerStartDate', 'registerStartTime', 'registerStartDatetime');
     Utils.dateDivide(that.model, 'registerEndDate', 'registerEndTime', 'registerEndDatetime');

     that.initForm();
   });
  }

  initForm(): void {
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
    that.eventForm = that.fb.group(
        {
          'title': [that.model.email, [Validators.required]],

          'startDate': [that.model.startDate, [Validators.required, DateTimeValidator.validateDate()]],
          'startTime': [that.model.startTime, [Validators.required, DateTimeValidator.validateTime()]],
          'endDate': [that.model.endDate, [Validators.required, DateTimeValidator.validateDate()]],
          'endTime': [that.model.endTime, [Validators.required, DateTimeValidator.validateTime()]],

          'registerStartDate': [that.model.startDate, [Validators.required, DateTimeValidator.validateDate()]],
          'registerStartTime': [that.model.startTime, [Validators.required, DateTimeValidator.validateTime()]],
          'registerEndDate': [that.model.endDate, [Validators.required, DateTimeValidator.validateDate()]],
          'registerEndTime': [that.model.endTime, [Validators.required, DateTimeValidator.validateTime()]],

          'signBefore': [that.model.signBefore, [Validators.required]],
          'address': [that.model.address, [Validators.required]],
          'phone': [that.model.phone, [Validators.required]],
          'email': [that.model.email, [Validators.required, EmailValidator.validate]],
          'website': [that.model.website, []]
        }, {
           validator: DateTimeValidator.compareDatetime([
             ['eventTimeCompare', 'startDate','startTime','endDate','endTime'],
             ['registerTimeCompare', 'registerStartDate','registerStartTime','registerEndDate','registerEndTime']
           ])
        }
    );

    that.eventForm.valueChanges.subscribe(data => that.onValueChanged(data));
    that.onValueChanged();
  }
  onValueChanged(data?: any) {
    let that = this;
    if (!that.eventForm) { return; }

    that.formErrors = ValidatorUtils.genMsg(that.eventForm, that.validateMsg, ['eventTimeCompare', 'registerTimeCompare']);
  }

  formErrors = [];
  validateMsg = {
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

    'eventTimeCompare': '会议结束时间不能早于开始时间',
    'registerTimeCompare': '报名结束时间不能早于开始时间'
  };

}
