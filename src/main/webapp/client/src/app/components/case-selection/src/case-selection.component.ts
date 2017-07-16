
import { Component, Input, OnInit } from '@angular/core';

import {NgbDatepickerI18n, NgbDateStruct, NgbModal, NgbModalRef, NgbActiveModal} from '@ng-bootstrap/ng-bootstrap';

import { CaseSelectionService } from './case-selection.service';

@Component({
    selector: 'case-selection',
    templateUrl: './case-selection.html',
    styleUrls: ['./styles.scss']
})
export class CaseSelectionComponent implements OnInit {

    @Input() progress: string = '0';
    @Input() color: string = '#209e91';
    @Input() height: string = '1px';
    @Input() show: boolean = true;

    constructor(public activeModal: NgbActiveModal, public service: CaseSelectionService) { }

    ngOnInit(): any {

    }

  dismiss(): any {
    this.activeModal.dismiss('Cross click');
  }

  close(): any {
    this.activeModal.close('Close click');
  }
}
