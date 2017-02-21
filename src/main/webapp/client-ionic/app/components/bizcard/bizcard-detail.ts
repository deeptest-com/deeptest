import {Component, Input, Output,OnInit, EventEmitter} from '@angular/core';
import {List,Item,Button,Icon} from 'ionic-angular';

import {ImgPathPipe} from '../../pipes/img-path';

@Component({
    selector: 'bizcard-detail',
    templateUrl: 'build/components/bizcard/bizcard-detail.html',
    directives: [List, Item,Button,Icon],
    pipes: [ImgPathPipe]
})

export class BizcardDetailComponent {
    @Input() products: any[];
    @Output() selected: any = new EventEmitter<any>();
    
    constructor() {
        
    }
    
    ngOnInit() {
        
    }
    
    onSelect(item) {
        this.selected.emit(item);
    }
}
