import { Component, Input, OnInit } from '@angular/core';

import { FieldType, CustomFieldDefinition } from '../field.definitions';
import { CustomFieldModel } from '../field.models';
import { FieldChangedEvent } from '../field.events';

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
        this.inputTextService.events.subscribe((event: FieldChangedEvent) => {
            // if (event.type === FieldChangedEvent.change) {
            //     // this.progress = event.value;
            // }
        });
    }
}
