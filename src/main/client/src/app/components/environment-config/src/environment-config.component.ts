
import { Component, Input, OnInit } from '@angular/core';

import {NgbDatepickerI18n, NgbDateStruct, NgbModal, NgbModalRef, NgbActiveModal} from '@ng-bootstrap/ng-bootstrap';

import { EnvironmentConfigService } from './environment-config.service';

@Component({
    selector: 'environment-config',
    templateUrl: './environment-config.html',
    styleUrls: ['./styles.scss']
})
export class EnvironmentConfigComponent implements OnInit {

    @Input() progress: string = '0';
    @Input() color: string = '#209e91';
    @Input() height: string = '1px';
    @Input() show: boolean = true;

    constructor(public activeModal: NgbActiveModal, public service: EnvironmentConfigService) { }

    ngOnInit(): any {

    }

  save(): any {
    this.activeModal.close('save');
  }

  dismiss(): any {
    this.activeModal.dismiss('cancel');
  }
}
