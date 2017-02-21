import {Component, Input, Output,OnInit, EventEmitter} from '@angular/core';
import {List,Item,Button,Icon} from 'ionic-angular';

import {ImgPathPipe} from '../../pipes/img-path';
import {CurrencyCnyPipe} from '../../pipes/currency-cny';
import {Product} from '../../models/Product';

@Component({
    selector: 'bizcard-list',
    templateUrl: 'build/components/bizcard-list/bizcard-list.html',
    directives: [List, Item,Button,Icon],
    pipes: [ImgPathPipe,CurrencyCnyPipe]
})

export class ProductListComponent {
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
