import { Component, Input, OnInit } from '@angular/core';

import { FieldType } from '../field-type';
import { FieldChangedEvent } from '../field-event';
import { InputTextService } from './input-text.service';


@Component({
  selector: 'input-text',
  templateUrl: './input-text.html',
  styleUrls: ['./styles.scss']
})
export class InputTextComponent implements OnInit {

    @Input() type: string = '0';
    @Input() content: string = '';

    constructor(public inputTextService: InputTextService) { }

    ngOnInit(): any {
        this.service.events.subscribe((event: InputTextEvent) => {
            if (event.type === InputTextEvent.change) {
                // this.progress = event.value;
            }
        });
    }
}
